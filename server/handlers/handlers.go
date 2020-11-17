package handlers

import (
	"log"
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
	remoteURL := req.FormValue("url")
	e, err := models.RegisterShortLink(req.Host, remoteURL)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "success.gohtml", e)
}

// Redirect sends it woosh
func Redirect(w http.ResponseWriter, req *http.Request) {
	path := req.RequestURI[1:]
	e, err := models.GetLongLink(path)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	err = models.IncreaseHits(path)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, e.OutsideAddr, http.StatusTemporaryRedirect)
}
