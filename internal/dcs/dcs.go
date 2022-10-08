package dcs

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
	"umbrella/internal/watchdog"

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
func Watcher() {
	var isRunning bool = false

	for {

		isRunning = portCheck(config.Settings.Host.Port)
		if isRunning {
			if isRunning != watchdog.IsDCSRunning {
				channels.DCSState <- isRunning
			}
		}

		//Check the Client state
		if !watchdog.IsClientRunning {
			go Client()
		}

		if !isRunning && watchdog.IsClientRunning {
			channels.ClientStop <- true
		}
		time.Sleep(config.Settings.Umbrella.Refreshrate.DcsWatcher * time.Millisecond)
	}
}

func Client() {
	var Binding Bindings
	channels.ClientState <- true

	for {
		// Check for request to stop.
		select {
		case <-channels.ClientStop:
			channels.ClientState <- false
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

		time.Sleep(config.Settings.Umbrella.Refreshrate.Client * time.Millisecond)
	}

}

func MissionState() {
	// Get the mission state from the DCS Server

	// Compare existing state to new state

	// If state has changed, IADS, Radar, Network, Power need to be stopped.

	// If state has changed, caches need to be cleared.

}

// *************
// gRPC Messages
// *************

//TODO, waiting on go-bindings to be updated:
// GetSessionId
// getDetectedTargets
// setAlarmState
// addStaticObject

// function to send a chat message to the DCS server
func SendChat(Bindings Bindings, message string) {
	Bindings.net.SendChat(context.Background(), &net.SendChatRequest{
		Message: message,
	})
}

func GetUnit(Bindings Bindings, unitName string) *unit.GetResponse {
	unit, err := Bindings.unit.Get(context.Background(), &unit.GetRequest{
		Name: unitName,
	})
	if err != nil {
		log.Fatalf("Could not get unit: %v", err)
	}
	return unit
}

func GetRadar(Bindings Bindings, unitName string) *unit.GetRadarResponse {
	radar, err := Bindings.unit.GetRadar(context.Background(), &unit.GetRadarRequest{
		Name: unitName,
	})
	if err != nil {
		log.Fatalf("Could not get radar: %v", err)
	}
	return radar
}

func SetEmmission(Bindings Bindings, unitName string, emitting bool) {
	_, err := Bindings.unit.SetEmission(context.Background(), &unit.SetEmissionRequest{
		Name:     unitName,
		Emitting: emitting,
	})
	if err != nil {
		log.Fatalf("Could not set emission: %v", err)
	}
}

func GetDescriptor(Bindings Bindings, unitName string) *unit.GetDescriptorResponse {
	descriptor, err := Bindings.unit.GetDescriptor(context.Background(), &unit.GetDescriptorRequest{
		Name: unitName,
	})
	if err != nil {
		log.Fatalf("Could not get descriptor: %v", err)
	}
	return descriptor
}

func GetAirbases(Bindings Bindings) *world.GetAirbasesResponse {
	airbases, err := Bindings.world.GetAirbases(context.Background(), &world.GetAirbasesRequest{})
	if err != nil {
		log.Fatalf("Could not get airbases: %v", err)
	}
	return airbases
}

func GetAbsoluteTime(Bindings Bindings) *timer.GetAbsoluteTimeResponse {
	time, err := Bindings.timer.GetAbsoluteTime(context.Background(), &timer.GetAbsoluteTimeRequest{})
	if err != nil {
		log.Fatalf("Could not get absolute time: %v", err)
	}
	return time
}
func GetTime(Bindings Bindings) *timer.GetTimeResponse {
	time, err := Bindings.timer.GetTime(context.Background(), &timer.GetTimeRequest{})
	if err != nil {
		log.Fatalf("Could not get time: %v", err)
	}
	return time
}

func GetTimeZero(Bindings Bindings) *timer.GetTimeZeroResponse {
	time, err := Bindings.timer.GetTimeZero(context.Background(), &timer.GetTimeZeroRequest{})
	if err != nil {
		log.Fatalf("Could not get time zero: %v", err)
	}
	return time
}

func GetUnits(Bindings Bindings) *group.GetUnitsResponse {
	units, err := Bindings.group.GetUnits(context.Background(), &group.GetUnitsRequest{})
	if err != nil {
		log.Fatalf("Could not get units: %v", err)
	}
	return units
}

func GetGroups(Bindings Bindings) *coalition.GetGroupsResponse {
	groups, err := Bindings.coalition.GetGroups(context.Background(), &coalition.GetGroupsRequest{})
	if err != nil {
		log.Fatalf("Could not get groups: %v", err)
	}
	return groups
}

func GetStaticObjects(Bindings Bindings) *coalition.GetStaticObjectsResponse {
	staticObjects, err := Bindings.coalition.GetStaticObjects(context.Background(), &coalition.GetStaticObjectsRequest{})
	if err != nil {
		log.Fatalf("Could not get static objects: %v", err)
	}

	return staticObjects
}
