package main

import (
	"umbrella/internal/channels"
	"umbrella/internal/config"
	"umbrella/internal/dcsServer"
)

func main() {

	//Start the watcher
	go config.ConfigWatcher()
	go dcsServer.ServerWatcher()

	//exitProgram()
	<-channels.ProcessStop
}
