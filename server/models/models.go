package models

import (
	"errors"
	"net/http"
)

type Link struct {
}

func RegisterShortLink(req *http.Request) (Link, error) {
	l := Link{}

	url := req.FormValue("url")
	if url == "" {
		return l, errors.New("400. Bad Request")
	}

	//TODO: boltdb operations

	return l, nil
}

func RemoveShortLink(req *http.Request) (Link, error) {
	l := Link{}
	return l, nil
}
