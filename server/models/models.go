package models

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"time"

	"../config"
	"../shared"
)

// RegisterShortLink does things alright
func RegisterShortLink(req *http.Request) (shared.Entry, error) {
	e := shared.Entry{}
	url := req.FormValue("url")
	if url == "" {
		return e, errors.New("400. Bad Request")
	}

	//mac := hmac.New(sha512.New, TODO:
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	id := fmt.Sprintf("%x\n", bs)

	e.Path = fmt.Sprintf("%s/%s", req.Host, id)
	e.OutsideAddr = url
	e.CreatedAt = time.Now()

	err := config.CreateEntry(id, e)
	if err != nil {
		return e, err
	}

	return e, nil
}

// RemoveShortLink also does things
func RemoveShortLink(req *http.Request) (shared.Entry, error) {
	e := shared.Entry{}
	return e, nil
}
