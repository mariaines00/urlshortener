package handlers

import (
	"net/http"

	"../config"
	"../models"
)

// I can try the pattern to return the handler/controler with the db connection

// Shortener is the entrypoint
func Shortener(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "shorten.gohtml", nil)
}

// RegisterShortLink is responsible for adding new entries to the db
func RegisterShortLink(w http.ResponseWriter, req *http.Request) {
	e, err := models.RegisterShortLink(req)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "success.gohtml", e)
}

// Redirect sends it woosh
func Redirect(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	//http.Redirect(w, req, string(url), http.StatusMovedPermanently)
}
