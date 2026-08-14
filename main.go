package main

import (
	"livepixelshtmx/internal/cli"
	// "livepixelshtmx/internal/server"
	"log"
	"net/http"
)

func main() {
	// capturing flags from user
	var cfg *cli.CliConfig = &cli.CliConfig{}
	if err := cfg.ParseArgs(); err != nil {
		log.Fatalf("Failed to parse arguments: (%s)", err.Error())
	}

	var mux *http.ServeMux = http.NewServeMux()

	// render base HTML, using parsed parameters

	// showcase HTML on network

	// use HTMX to real-time capture clicks on squares and update, live
}
