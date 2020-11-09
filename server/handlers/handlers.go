package handlers

import (
	"net/http"

	"../config"
)

// Shortener is ...
func Shortener(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "shorten.gohtml", nil)
}

// RegisterNewShortener is ...
func RegisterNewShortener(w http.ResponseWriter, req *http.Request) {

}
