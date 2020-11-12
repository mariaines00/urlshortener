package models

import (
	"errors"
	"net/http"

	"../shared"
)

func RegisterShortLink(req *http.Request) (shared.Entry, error) {
	e := shared.Entry{}

	url := req.FormValue("url")
	if url == "" {
		return e, errors.New("400. Bad Request")
	}

	//TODO: boltdb operations

	return e, nil
}

func RemoveShortLink(req *http.Request) (shared.Entry, error) {
	e := shared.Entry{}
	return e, nil
}
