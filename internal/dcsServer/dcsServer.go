package dcsServer

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/DCS-gRPC/go-bindings/dcs/v0/controller"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/mission"
	"github.com/DCS-gRPC/go-bindings/dcs/v0/net"
)

type bindings struct {
	conn       *grpc.ClientConn
	mission    mission.MissionServiceClient
	net        net.NetServiceClient
	controller controller.ControllerServiceClient
}

func NewBindings(addr string, port int) *bindings {

	conn, _ := connect(addr, port)

	return &bindings{
		conn:       conn,
		mission:    mission.NewMissionServiceClient(conn),
		net:        net.NewNetServiceClient(conn),
		controller: controller.NewControllerServiceClient(conn),
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
func SendChat(bindings bindings, message string) {
	bindings.net.SendChat(context.Background(), &net.SendChatRequest{
		Message: message,
	})
}
