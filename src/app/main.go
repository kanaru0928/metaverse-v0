package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kanaru0928/metaverse-v0/src/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)

	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir(""))))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
