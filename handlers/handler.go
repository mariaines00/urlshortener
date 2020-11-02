package handlers

import "net/http"

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// Redirecter will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map the given short url
// paths (keys in the map) to their corresponding long URL.
// If the path is not found, then the fallback
// http.Handler will be called instead.
func Redirecter(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, req, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, req)
	}
}

// Register will ...
func Register(url string) http.HandlerFunc {

}
