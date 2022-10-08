package main

import (
	"umbrella/internal/channels"
	"umbrella/internal/config"
	"umbrella/internal/dcs"
	"umbrella/internal/logger"
	"umbrella/internal/watchdog"
)

func main() {
	go logger.Run()
	channels.Logs <- "Starting Umbrella"
	channels.Logs <- "Umbrella will idle until DCS is running."
	go watchdog.Run()

	//Start the watcher
	go config.Watcher()
	go dcs.Watcher()

	<-channels.ProcessStop
	channels.Logs <- "Process stop signal received. Exiting."
}
