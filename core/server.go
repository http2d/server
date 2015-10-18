package core

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
)

// Server implements the HTTP2d Server class
type Server struct {
	http.Server
	mux *mux.Router
}

func testResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hi tester %q\n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)
}

// Init initializes the server
func (srv *Server) Init() error {
	// Settings
	srv.Addr = ":8080"
	http2.VerboseLogs = true

	http2.ConfigureServer(&srv.Server, nil)

	// Gorilla's mux router
	srv.mux = mux.NewRouter()

	// Handlers
	/// http.HandleFunc("/", testResponse)
	return nil
}

// Run runs the server (forever)
func (srv *Server) Run() error {
	var err error
	var tlsKey string
	var tlsCrt string

	if tlsKey != "" && tlsCrt != "" {
		err = srv.ListenAndServeTLS(tlsCrt, tlsKey)
	} else {
		err = srv.ListenAndServe()
	}

	return err
}

// Quit cleans up before destroying the object
func (srv *Server) Quit() {
	if srv == nil {
		return
	}
}
