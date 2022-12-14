package network

import (
	"time"

	"umbrella/internal/channels"
	"umbrella/internal/config"
	"umbrella/internal/power"
	"umbrella/internal/radar"
)

type Network struct {
	Name         string
	PowerGrids   []power.Grid
	EWRadars     []radar.EWRadar
	SearchRadars []radar.SearchRadar
	Elements     []Element // Elements connected directly to the network root
	Sections     []Section
	CoreCommand  []CommandUnit
}

type Section struct {
	Name        string
	Network     Network
	Elements    []Element
	CommandUnit CommandUnit
}

type Element struct {
	Name                string
	SearchRadar         radar.SearchRadar
	TrackRadar          radar.TrackRadar
	PowerGenerator      power.PowerGenerator
	ConnectedSubStation power.SubStation
}

type CommandUnit struct {
	Name                string
	PowerGenerator      power.PowerGenerator
	ConnectedSubStation power.SubStation
}

func Run() {
	channels.NetworkState <- true
	for {
		//check for stop signal
		select {
		case <-channels.NetworkStop:
			channels.NetworkState <- false
			return
		default:
		}

		time.Sleep(time.Duration(config.Settings.Umbrella.Refreshrate.Network) * time.Millisecond)
	}
}
