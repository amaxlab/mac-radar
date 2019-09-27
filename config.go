package main

import (
	"github.com/amaxlab/go-lib/config"
	"github.com/amaxlab/mac-radar/mikrotik"
)

const (
	DefaultPort             = 8080
	DefaultMikrotikAddress  = "192.168.88.1:8728"
	DefaultMikrotikUsername = "admin"
	DefaultMikrotikPassword = "admin"
)

type WatchMacAddress struct {
	Id    string
	Alias string
}

type Configuration struct {
	Debug                 bool
	Port                  int
	WatchMacAddress       []WatchMacAddress
	MikrotikConfiguration mikrotik.Configuration
}

func NewConfiguration() *Configuration {
	loader := config.NewConfigLoader()

	configuration := Configuration{}
	configuration.Debug = loader.Bool("debug", false)
	configuration.Port = loader.Int("port", DefaultPort)

	configuration.MikrotikConfiguration = mikrotik.Configuration{}
	configuration.MikrotikConfiguration.Address = loader.String("mikrotik-address", DefaultMikrotikAddress)
	configuration.MikrotikConfiguration.Username = loader.String("mikrotik-username", DefaultMikrotikUsername)
	configuration.MikrotikConfiguration.Password = loader.String("mikrotik-password", DefaultMikrotikPassword)

	return &configuration
}
