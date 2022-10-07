package dcsServer

import (
	"context"
	"fmt"
	"log"
	network "net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"umbrella/internal/channels"
	"umbrella/internal/config"

	"github.com/DCS-gRPC/go-bindings/dcs/v0/atmosphere"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/coalition"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/controller"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/custom"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/group"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/hook"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/mission"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/net"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/timer"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/trigger"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/unit"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/world"
)

type Bindings struct {
	ready      bool
	conn       *grpc.ClientConn
	atmosphere atmosphere.AtmosphereServiceClient
	coalition  coalition.CoalitionServiceClient
	controller controller.ControllerServiceClient
	custom     custom.CustomServiceClient
	group      group.GroupServiceClient
	hook       hook.HookServiceClient
	mission    mission.MissionServiceClient
	net        net.NetServiceClient
	timer      timer.TimerServiceClient
	trigger    trigger.TriggerServiceClient
	unit       unit.UnitServiceClient
	world      world.WorldServiceClient
}

func NewBindings(addr string, port int) *Bindings {

	conn, _ := connect(addr, port)

	return &Bindings{
		ready:      true,
		conn:       conn,
		atmosphere: atmosphere.NewAtmosphereServiceClient(conn),
		coalition:  coalition.NewCoalitionServiceClient(conn),
		controller: controller.NewControllerServiceClient(conn),
		custom:     custom.NewCustomServiceClient(conn),
		group:      group.NewGroupServiceClient(conn),
		hook:       hook.NewHookServiceClient(conn),
		mission:    mission.NewMissionServiceClient(conn),
		net:        net.NewNetServiceClient(conn),
		timer:      timer.NewTimerServiceClient(conn),
		trigger:    trigger.NewTriggerServiceClient(conn),
		unit:       unit.NewUnitServiceClient(conn),
		world:      world.NewWorldServiceClient(conn),
	}
}

func connect(addr string, port int) (*grpc.ClientConn, error) {
	// Concate the address and port
	addr = fmt.Sprint(addr, ":", port)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTimeout(5 * time.Second),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Panicf("Failed to connect to server: %v", err)
	}
	return conn, err
}

func portCheck(port int) bool {
	var isRunning bool = false
	_, err := network.Listen("tcp", ":"+strconv.FormatInt(int64(port), 10))
	if err == nil {
		isRunning = true
	}
	return isRunning
}

/*
	Watcher runs at main execution and checks to see if gRPC port is listening.

When server is tected, Watcher will spawn the Client exactly once.
*/
func ServerWatcher() {
	var isRunning bool = false
	var isClient bool = false

	//Check the Client state

	for {
		isRunning = portCheck(config.Settings.Host.Port)
		// Check to see if DCS is running
		if isRunning {
			go Client()
			time.Sleep(1 * time.Second)

			select {
			case msg := <-channels.ClientState:
				if msg == "Client Running" {
					isClient = true
				}
				if msg == "Client Stopped" {
					isClient = false
				}
			default:
			}
		}

		if isRunning && !isClient {
			go Client()
		}

		if !isRunning && isClient {
			channels.ClientStop <- true
		}
		time.Sleep(1 * time.Second)
	}
}

func Client() {
	var Binding Bindings
	channels.ClientState <- "Client Running"

	for {
		// Check for request to stop.
		select {
		case <-channels.ClientStop:
			channels.ClientState <- "Client Stopped"
			return
		default:
		}

		// Create binding if it does not exist.
		if !Binding.ready {
			Binding := NewBindings(config.Settings.Host.Address, config.Settings.Host.Port)
			var message string = fmt.Sprintf("%v Client has started", config.Settings.Iads.Name)
			SendChat(*Binding, message)
		}

		//Process Client Call Queue
		select {
		case request := <-channels.ClientCallQueue:
			println(request)
			//TODO Handle Request
		}

		time.Sleep(1 * time.Microsecond)
	}

}

// function to send a chat message to the DCS server
func SendChat(Bindings Bindings, message string) {
	Bindings.net.SendChat(context.Background(), &net.SendChatRequest{
		Message: message,
	})
}
