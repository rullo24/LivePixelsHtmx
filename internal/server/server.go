package server

import (
	"html/template"
	"livepixelshtmx/internal/cli"
	"livepixelshtmx/internal/game"
	"net/http"
)

const (
	HOSTED_NETWORK_PORT int = 8080
)

// parse HTML templates for static init
var pageTmpl = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/square.html",
))

type Server struct {
	tmpl    *template.Template
	squares []game.Square
}

// DESCRIPTION
// Init server artifact
func NewServerWithTemplate(cfg *cli.CliConfig, tmpl *template.Template) *Server {

	// creating squares with random start states
	var totalNumSquares int = cfg.NumColumns * cfg.NumRows
	var squares []game.Square = make([]game.Square, totalNumSquares)
	for i := range squares {
		var currSquare game.Square = game.Square{}
		currSquare.ID = i              // setting index
		currSquare.SetColourToRandom() // rand colour before assignment
		squares[i] = currSquare
	}

	// send back new instance with randomised starting squares
	return &Server{
		tmpl:    tmpl,
		squares: squares,
	}
}

// DESCRIPTION
// Handler func that runs when POST GET is called on root (/)
func (s *Server) IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Squares []game.Square
	}{
		Squares: s.squares,
	}

	if err := s.tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DESCRIPTION
// Handler func that runs when POST PATCH is called on square id (/square/{id})
func (s *Server) SquarePatchHandler(w http.ResponseWriter, r *http.Request) {

}
