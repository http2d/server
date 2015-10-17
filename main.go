package http2d

import (
	log "github.com/Sirupsen/logrus"
	"github.com/http2d/server/core"
)

func setupLogging() error {
	log.Info("Starting")
	return nil
}

func main() {
	var srv core.Server
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
