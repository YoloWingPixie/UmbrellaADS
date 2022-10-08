package config

import (
	"log"
	"os"
	"time"

	"umbrella/internal/channels"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     Host     `yaml:"host"`
	Iads     Iads     `yaml:"iads"`
	Umbrella Umbrella `yaml:"umbrella"`
}

type Host struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Iads struct {
	Name      string `yaml:"name"`
	Coalition string `yaml:"coalition"`
}

type Umbrella struct {
	Refreshrate RefreshRate `yaml:"refreshrate"`
}

type RefreshRate struct {
	Network        int `yaml:"network"`
	Power          int `yaml:"power"`
	Radar          int `yaml:"radar"`
	Iads           int `yaml:"iads"`
	Config         int `yaml:"config"`
	Client         int `yaml:"client"`
	DcsWatcher     int `yaml:"dcsWatcher"`
	MissionWatcher int `yaml:"missionWatcher"`
}

var (
	Settings   Config
	configFile string = "config.yaml"
)

func init() {
	channels.Logs <- "Loading config file: " + configFile

	// Read the config file
	configFile, err := os.Open(configFile)
	if err != nil {
		log.Panicf("Failed to read config file: %v", err)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err2 := decoder.Decode(&Settings)

	if err2 != nil {
		log.Fatal(err2)
	}

	channels.Logs <- "Config file loaded."
}

func Watcher() {
	var NewSettings Config
	for {
		//Read the config file
		configFile, err := os.Open(configFile)
		if err != nil {
			log.Panicf("Failed to read config file: %v", err)
		}
		defer configFile.Close()

		decoder := yaml.NewDecoder(configFile)
		err2 := decoder.Decode(&NewSettings)

		if err2 != nil {
			log.Fatal(err2)
		}

		//Check if the config file has changed
		if NewSettings != Settings {
			Settings = NewSettings
			channels.Logs <- "Config file changed."
		}

		time.Sleep(time.Duration(Settings.Umbrella.Refreshrate.Config) * time.Millisecond)
	}
}
