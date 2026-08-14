package main

import (
	"livepixelshtmx/internal/cli"
	"log"
)

func main() {

	// capturing flags from user
	var cfg *cli.CliConfig = &cli.CliConfig{}
	if err := cfg.ParseArgs(); err != nil {
		log.Fatalf("Failed to parse arguments: (%s)", err.Error())
	}

	// render base HTML, using parsed parameters

	// showcase HTML on network

	// use HTMX to real-time capture clicks on squares and update, live
}
