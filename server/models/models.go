package models

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mariaines00/urlshortener/config"
	"github.com/mariaines00/urlshortener/shared"
)

const (
	alphabet    = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
	alphabetLen = len(alphabet)
)

// RegisterShortLink does things alright
func RegisterShortLink(host string, url string) (shared.Entry, error) {
	e := shared.Entry{}

	if url == "" || !isValidURL(url) {
		return e, shared.NewHTTPError(nil, http.StatusBadRequest, "Invalid parameters")
	}

	index := dbNextSequence()
	id := encode(index)

	e.Path = fmt.Sprintf("%s/%s", host, id)
	e.OutsideAddr = url
	e.CreatedAt = time.Now()

	err := config.CreateEntry(id, e)
	if err != nil {
		return e, err
	}

	return e, nil
}

// GetLongLink returns the entry corresponding to the long URL
func GetLongLink(path string) (shared.Entry, error) {
	e, err := config.GetEntryByID(path)
	return *e, err
}

// IncreaseHits calls the db function to update the hits counter
func IncreaseHits(path string) error {
	return config.IncreaseHits(path)
}

/* Helpers */

// isValidURL returns false if the provided input is not a url
func isValidURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// dbNextSequence returns the next db index(self incremental)
func dbNextSequence() int {
	n, _ := config.GetSequence()
	n++
	return n
}

//encode takes an ID and turns it into a short string
// based on https://stackoverflow.com/questions/742013/how-do-i-create-a-url-shortener#742047
func encode(n int) string {
	sb := strings.Builder{}
	for n > 0 {
		sb.WriteByte(alphabet[n%alphabetLen])
		n = n / alphabetLen
	}
	return sb.String()
}

//decode takes a string and turns it into an ID
func decode(s string) (n int) {
	for _, r := range s {
		n = n*alphabetLen + strings.IndexRune(alphabet, r)
	}
	return
}
