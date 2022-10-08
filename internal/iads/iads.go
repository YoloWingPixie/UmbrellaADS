package iads

import (
	"time"

	"umbrella/internal/channels"
	"umbrella/internal/config"
	"umbrella/internal/network"
)

type IADS struct {
	Name    string
	Network network.Network
}

func NewIads(name string) *IADS {
	var iads *IADS
	iads.Name = name

	return iads
}

func Run() {
	channels.IADSState <- true
	channels.Logs <- "IADS thread started."

	for {
		//check for stop signal
		select {
		case <-channels.IADSStop:
			channels.IADSState <- false
			channels.Logs <- "IADS thread stopped."
			return
		default:
		}

		time.Sleep(time.Duration(config.Settings.Umbrella.Refreshrate.Iads) * time.Millisecond)
	}
}
