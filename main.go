package main

import (
	"io/ioutil"
	"log"

	"umbrella/internal/dcsServer"

	"gopkg.in/yaml.v3"
)

func main() {

	// Read the config file
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Panicf("Failed to read config file: %v", err)
	}
	config := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(configFile, &config)

	if err2 != nil {
		log.Fatal(err2)
	}

	//Create new bindings
	bindings := dcsServer.NewBindings(config["address"].(string), config["port"].(int))

	//Send a chat message
	dcsServer.SendChat(*bindings, "Umbrella ADS is running")
}
