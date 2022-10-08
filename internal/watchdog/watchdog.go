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
	channels.Logs <- "Watchdog thread started."
	for {
		// Check if DCS state changed
		select {
		case msg := <-channels.DCSState:
			if msg {
				IsDCSRunning = true
				channels.Logs <- "Watchdog: DCS-gRPC Server has been detected."
			}
			if !msg {
				IsDCSRunning = false
				channels.Logs <- "Watchdog: DCS-gRPC Server has been lost."
			}
		default:
		}

		// Check if Mission state changed
		select {
		case msg := <-channels.MissionState:
			if msg == 1 {
				IsMissionRunning = true
				channels.Logs <- "Watchdog: Mission has been detected."
			}
			if msg == 0 {
				IsMissionRunning = false
				channels.Logs <- "Watchdog: Mission has been lost."
			}
			if msg == 2 {
				HasMissionChanged = true
				channels.Logs <- "Watchdog: Mission has been changed."
			}
		default:
		}

		// Check if Client state changed
		select {
		case msg := <-channels.ClientState:
			if msg {
				IsClientRunning = true
				channels.Logs <- "Watchdog: Client has been detected."
			}
			if !msg {
				IsClientRunning = false
				channels.Logs <- "Watchdog: Client has been lost."
			}
		default:
		}

		// Check if IADS state changed
		select {
		case msg := <-channels.IADSState:
			if msg {
				IsIADSRunning = true
				channels.Logs <- "Watchdog: IADS has been detected."
			}
			if !msg {
				IsIADSRunning = false
				channels.Logs <- "Watchdog: IADS has been lost."
			}
		default:
		}

		// Check if Power state changed
		select {
		case msg := <-channels.PowerState:
			if msg {
				IsPowerRunning = true
				channels.Logs <- "Watchdog: Power has been detected."
			}
			if !msg {
				IsPowerRunning = false
				channels.Logs <- "Watchdog: Power has been lost."
			}
		default:
		}

		// Check if Network state changed
		select {
		case msg := <-channels.NetworkState:
			if msg {
				IsNetworkRunning = true
				channels.Logs <- "Watchdog: Network has been detected."
			}
			if !msg {
				IsNetworkRunning = false
				channels.Logs <- "Watchdog: Network has been lost."
			}
		default:
		}

		// Check if Radar state changed
		select {
		case msg := <-channels.RadarState:
			if msg {
				IsTargetRunning = true
				channels.Logs <- "Watchdog: Radar has been detected."
			}
			if !msg {
				IsTargetRunning = false
				channels.Logs <- "Watchdog: Radar has been lost."
			}
		default:
		}

		time.Sleep(1 * time.Millisecond)
	}
}
