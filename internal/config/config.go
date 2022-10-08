package config

import (
	"log"
	"os"
	"time"

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
	Network        time.Duration `yaml:"network"`
	Power          time.Duration `yaml:"power"`
	Radar          time.Duration `yaml:"radar"`
	Iads           time.Duration `yaml:"iads"`
	Config         time.Duration `yaml:"config"`
	Client         time.Duration `yaml:"client"`
	DcsWatcher     time.Duration `yaml:"dcsWatcher"`
	MissionWatcher time.Duration `yaml:"missionWatcher"`
}

var (
	Settings   Config
	configFile string = "config.yaml"
)

func init() {
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

}

func Watcher() {
	for {

		time.Sleep(Settings.Umbrella.Refreshrate.Config * time.Millisecond)
	}
}
