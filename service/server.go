package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/matches/{shortname}", getMatchHandler(formatter)).Methods("GET")
	mx.HandleFunc("/matches", matchesCollectionHandler(formatter)).Methods("GET")
	mx.HandleFunc("/matches", createMatchHandler(formatter)).Methods("POST")
}
