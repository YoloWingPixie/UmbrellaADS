package dcsServer

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

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

// function to send a chat message to the DCS server
func SendChat(bindings Bindings, message string) {
	bindings.net.SendChat(context.Background(), &net.SendChatRequest{
		Message: message,
	})
}
