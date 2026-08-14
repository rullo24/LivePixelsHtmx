package main

import (
	"html/template"
	"livepixelshtmx/internal/cli"
	"livepixelshtmx/internal/server"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// capturing flags from user
	var cfg *cli.CliConfig = &cli.CliConfig{}
	if err := cfg.ParseArgs(); err != nil {
		log.Fatalf("Failed to parse arguments: (%s)", err.Error())
	}

	// creating new server instances
	var mux *http.ServeMux = http.NewServeMux()
	var tmpl *template.Template = template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/square.html",
	))
	var srv *server.Server = server.NewServerWithTemplate(cfg, tmpl)

	// set callbacks against POST
	mux.HandleFunc("GET /", srv.IndexGetHandler)
	mux.HandleFunc("PATCH /square/{id}", srv.SquarePatchHandler)

	// allow server mux to access local assets
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// start listening server on selected network port
	log.Printf("Listening on :%d\n", server.HOSTED_NETWORK_PORT)
	var portStr string = ":" + strconv.Itoa(server.HOSTED_NETWORK_PORT)
	log.Fatal(http.ListenAndServe(portStr, mux))
}
