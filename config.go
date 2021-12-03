package main

import (
	"log"

	"github.com/spf13/viper"
)

type (
	TunnelAddress struct {
		Interface string
		Gateway   string
		Mask      string
	}

	Tunnel struct {
		Name    string
		Address TunnelAddress
	}

	Proxy struct {
		Port    uint64
		Address string
	}

	Configuration struct {
		Proxy  Proxy
		Tunnel Tunnel
	}
)

func Config() Configuration {

	/**
	* Find the correct configuration file from the command
	* argument or /etc location and verify.
	 */
	var config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("/etc/ttos")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load configuration file")
	}

	var configuration Configuration
	config.Unmarshal(&configuration)

	return configuration

}
