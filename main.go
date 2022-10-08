package main

import (
	"umbrella/internal/channels"
	"umbrella/internal/config"
	"umbrella/internal/dcs"
)

func main() {

	//Start the watcher
	go config.Watcher()
	go dcs.Watcher()

	//exitProgram()
	<-channels.ProcessStop
}
