package handlers

import (
	"net/http"

	"../config"
)

// Shortener is the entrypoint
func Shortener(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "shorten.gohtml", nil)
}

// RegisterNewShortener is responsible for adding new entries to the db
func RegisterNewShortener(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "success.gohtml", nil)
}

// Redirect sends it woosh
func Redirect(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	//http.Redirect(w, req, string(url), http.StatusMovedPermanently)
}
