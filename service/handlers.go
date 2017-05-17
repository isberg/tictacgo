package service

import (
	"github.com/unrolled/render"
	"net/http"
)

type match struct {
	ShortName string `json:"shortname"`
}

func getMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, match { ShortName: "james-vs-joe"})
	}
}

func matchesCollectionHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, []match{match { ShortName: "james-vs-joe"}})
	}
}

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Location", "/matches/adam-vs-nils")
		formatter.JSON(w, http.StatusCreated, match { ShortName: "adam-vs-nils"})
	}
}
