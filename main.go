package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func handleConfig() error {
	viper.SetConfigName("http2d")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("/etc")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return nil
}

func setupLogging() error {
	log.Info("Starting")
	return nil
}

func main() {
	var srv Server
	var err error

	// _ = "breakpoint"

	setupLogging()

	if err = handleConfig(); err != nil {
		log.Fatal(err)
	}

	if err = srv.Init(); err != nil {
		log.Fatal(err)
	}

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}

	srv.Quit()
}
