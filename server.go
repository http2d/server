package main

import (
	"fmt"
	"html"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

func main() {
	var srv http.Server

	log.Info("Starting")
	handleConfig()

	srv.Addr = ":8080"
	http2.VerboseLogs = true

	http2.ConfigureServer(&srv, nil)

	// Handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi tester %q\n", html.EscapeString(r.URL.Path))
		showRequestInfoHandler(w, r)
	})

	// Listen loop
	var err error

	tlsKey := viper.GetString("server.http2.tls.key")
	tlsCrt := viper.GetString("server.http2.tls.cert")

	if tlsKey != "" && tlsCrt != "" {
		err = srv.ListenAndServeTLS(tlsCrt, tlsKey)
	} else {
		err = srv.ListenAndServe()
	}

	log.Fatal(err)
}

func handleConfig() {
	_ = "breakpoint"

	viper.SetConfigName("http2d")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("/etc")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func showRequestInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
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
	fmt.Fprintf(w, "hola\n")
	r.Header.Write(w)
}
