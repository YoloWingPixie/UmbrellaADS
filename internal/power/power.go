package power

import (
	"time"

	"umbrella/internal/channels"
)

type Grid struct {
	Name          string
	PowerStations []PowerStation
	SubStations   []SubStation
	Coalition     string
}

type PowerStation struct {
	Name string
	Grid Grid
}

type SubStation struct {
	Name          string
	Grid          Grid
	PowerStations []PowerStation
}

type PowerGenerator struct {
	Name                 string
	ConnectedSubStations []SubStation
}

// interface for power stations
type PowerStationInterface interface {
}

// interface for sub stations
type SubStationInterface interface {
}

// interface for grids
type GridInterface interface {
}

// interface for power
type PowerInterface interface {
}

func NewGrid(Name string) *Grid {
	return &Grid{}
}

func NewPowerStation(Name string) *PowerStation {

	return &PowerStation{}
}

func NewSubStation(Name string) *SubStation {

	return &SubStation{}
}

func Run() {
	channels.PowerState <- true
	for {
		//check for stop signal
		select {
		case <-channels.PowerStop:
			channels.PowerState <- false
			return
		default:
		}

		time.Sleep(5 * time.Millisecond)
	}
}
