package controller

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes(r *mux.Router) {
	r.HandleFunc("/ng/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	h.homeTemplate.Execute(w, nil)
}
