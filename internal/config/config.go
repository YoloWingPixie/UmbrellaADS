package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host Host `yaml:"host"`
	Iads Iads `yaml:"iads"`
}

type Host struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Iads struct {
	Name      string
	Coalition string
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

		time.Sleep(5 * time.Second)
	}
}
