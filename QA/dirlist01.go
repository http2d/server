package main

import (
	"net/http"

	"github.com/http2d/server/core"
	"github.com/http2d/server/handlers"
	"golang.org/x/net/http2"
)

func main() {
	var srv core.Server

	srv.Addr = ":8080"
	http2.ConfigureServer(&srv.Server, nil)

	http.HandleFunc("/", handlers.NewDirList().Reply)

	srv.Init()
	srv.Run()
	srv.Quit()
}
