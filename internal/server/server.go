package server

import (
	"html/template"
	"livepixelshtmx/internal/cli"
	"livepixelshtmx/internal/game"
	"net/http"
	"strconv"
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
	cfg     *cli.CliConfig
}

// DESCRIPTION
// Init server artifact
func NewServerWithTemplate(cfg *cli.CliConfig, tmpl *template.Template) *Server {

	// creating squares with random start states
	var totalNumSquares int = cfg.NumCols * cfg.NumRows
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
		cfg:     cfg,
	}
}

// DESCRIPTION
// Handler func that runs when POST GET is called on root (/)
func (s *Server) IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Squares []game.Square
		Cfg     *cli.CliConfig
	}{
		Squares: s.squares,
		Cfg:     s.cfg,
	}

	if err := s.tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DESCRIPTION
// Handler func that runs when POST PATCH is called on square id (/square/{id})
func (s *Server) SquarePatchHandler(w http.ResponseWriter, r *http.Request) {
	var idStr string = r.PathValue("id")
	id, strConvToIntErr := strconv.Atoi(idStr)
	if strConvToIntErr != nil {
		http.Error(w, "Invalid ID parsed", http.StatusBadRequest)
		return
	}

	// check bounds of returned square ID
	if id < 0 || id >= len(s.squares) {
		http.Error(w, "Square not found", http.StatusNotFound)
		return
	}

	// change colour of square that was pressed
	s.squares[id].SetColourToRandom()

	// update template to render new coloured square
	if err := s.tmpl.ExecuteTemplate(w, "square", s.squares[id]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
