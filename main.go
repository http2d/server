package main

import log "github.com/Sirupsen/logrus"

func setupLogging() error {
	log.Info("Starting")
	return nil
}

func main() {
	var srv Server
	var err error

	// _ = "breakpoint"

	setupLogging()

	if err = srv.Init(); err != nil {
		log.Fatal(err)
	}

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}

	log.Info("Exiting")
	srv.Quit()
}
