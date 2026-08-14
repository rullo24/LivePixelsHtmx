package main

import (
	"html/template"
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

	// parse HTML templates for static init
	var pageTmpl = template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/square.html",
	))

	var mux *http.ServeMux = http.NewServeMux()

	// render base HTML, using parsed parameters

	// showcase HTML on network

	// use HTMX to real-time capture clicks on squares and update, live
}
