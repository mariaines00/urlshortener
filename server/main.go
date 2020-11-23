package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/mariaines00/urlshortener/config"
	"github.com/mariaines00/urlshortener/handlers"
)

func main() {

	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/favicon.ico", http.NotFound)
	r.HandleFunc("/short", handlers.Shortener).Methods("GET")
	r.HandleFunc("/short/new", handlers.RegisterShortLink).Methods("POST")
	r.HandleFunc("/{key}", handlers.Redirect)

	server := &http.Server{
		Addr:         "0.0.0.0:3000", //TODO: use env vars
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Println("Server started at port 3000") // TODO: use env vars

		err := config.Init("../database/entries.db")
		if err != nil {
			log.Fatalln(err)
		}

		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("see you later aligator")
	os.Exit(0)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/short", http.StatusSeeOther)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%v %v %v", req.Method, req.Host, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
