package controller

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

type index struct {
	indexTemplate *template.Template
}

func (h index) registerRoutes(r *mux.Router) {

	r.HandleFunc("/web/index", h.handleIndex)
	r.HandleFunc("/web/home", h.handleIndex)
	r.HandleFunc("/web/topology", h.handleIndex)
	r.HandleFunc("/web/devices", h.handleIndex)
	r.HandleFunc("/web/", h.handleIndex)
}

func (h index) handleIndex(w http.ResponseWriter, r *http.Request) {
	h.indexTemplate.Execute(w, nil)
}
