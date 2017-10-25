package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/", handler.Index)
	r := NewRouter()
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
