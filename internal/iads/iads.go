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

	for {
		//check for stop signal
		select {
		case <-channels.IADSStop:
			channels.IADSState <- false
			return
		default:
		}

		time.Sleep(config.Settings.Umbrella.Refreshrate.Iads * time.Millisecond)
	}
}
