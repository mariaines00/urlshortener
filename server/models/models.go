package models

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"../config"
	"../shared"
)

// RegisterShortLink does things alright
func RegisterShortLink(req *http.Request) (shared.Entry, error) {
	e := shared.Entry{}
	url := req.FormValue("url")
	if url == "" || !isValidURL(url) {
		return e, errors.New("400. Bad Request")
	}

	//mac := hmac.New(sha512.New, TODO:
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	id := fmt.Sprintf("%x", bs)

	e.Path = fmt.Sprintf("%s/%s", req.Host, id)
	e.OutsideAddr = url
	e.CreatedAt = time.Now()

	err := config.CreateEntry(id, e)
	if err != nil {
		return e, err
	}

	return e, nil
}

// GetLongLink returns the entry corresponding to the long URL
func GetLongLink(id string) (shared.Entry, error) {
	e, err := config.GetEntryByID(id)
	if err != nil {
		return *e, err
	}

	return *e, nil
}

// IncreaseHits calls the db function to update the hits counter
func IncreaseHits(id string) error {
	return config.IncreaseHits(id)
}

/* Helpers */

// isValidURL returns false if the provided input is not a url
func isValidURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}
