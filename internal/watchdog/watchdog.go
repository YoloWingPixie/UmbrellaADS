package watchdog

import (
	"time"
	"umbrella/internal/channels"
)

var (
	IsDCSRunning      bool
	IsMissionRunning  bool
	IsClientRunning   bool
	IsIADSRunning     bool
	IsPowerRunning    bool
	IsTargetRunning   bool
	IsNetworkRunning  bool
	IsConfigRunning   bool
	HasMissionChanged bool
)

func Run() {
	for {
		// Check if DCS state changed
		select {
		case msg := <-channels.DCSState:
			if msg {
				IsDCSRunning = true
			}
			if !msg {
				IsDCSRunning = false
			}
		default:
		}

		// Check if Mission state changed
		select {
		case msg := <-channels.MissionState:
			if msg == 1 {
				IsMissionRunning = true
			}
			if msg == 0 {
				IsMissionRunning = false
			}
			if msg == 2 {
				HasMissionChanged = true
			}
		default:
		}

		// Check if Client state changed
		select {
		case msg := <-channels.ClientState:
			if msg {
				IsClientRunning = true
			}
			if !msg {
				IsClientRunning = false
			}
		default:
		}

		// Check if IADS state changed
		select {
		case msg := <-channels.IADSState:
			if msg {
				IsIADSRunning = true
			}
			if !msg {
				IsIADSRunning = false
			}
		default:
		}

		// Check if Power state changed
		select {
		case msg := <-channels.PowerState:
			if msg {
				IsPowerRunning = true
			}
			if !msg {
				IsPowerRunning = false
			}
		default:
		}

		// Check if Network state changed
		select {
		case msg := <-channels.NetworkState:
			if msg {
				IsNetworkRunning = true
			}
			if !msg {
				IsNetworkRunning = false
			}
		default:
		}

		// Check if Radar state changed
		select {
		case msg := <-channels.RadarState:
			if msg {
				IsTargetRunning = true
			}
			if !msg {
				IsTargetRunning = false
			}
		default:
		}

		time.Sleep(1 * time.Millisecond)
	}
}
