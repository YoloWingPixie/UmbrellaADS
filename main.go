package main

import (
	"umbrella/internal/channels"
	"umbrella/internal/config"
	"umbrella/internal/dcs"
	"umbrella/internal/logger"
)

func main() {
	go logger.Run()
	channels.Logs <- "Starting Umbrella"

	//Start the watcher
	go config.Watcher()
	go dcs.Watcher()

	<-channels.ProcessStop
	channels.Logs <- "Process stop signal received. Exiting."
}
