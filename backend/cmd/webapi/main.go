package main

import (
	"github.com/gorilla/mux"
	week "github.com/ubahwin/week-of-learn/internal/handlers"
	"github.com/ubahwin/week-of-learn/internal/middleware"
	"log"
	"net/http"
)

func handleRequests() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.HeadersMiddleware)
	router.HandleFunc("/api/week.what-a-week", week.WhatWeek)

	http.Handle("/", router)

	return router
}

func main() {
	err := http.ListenAndServe(":7171", handleRequests())
	if err != nil {
		log.Fatal("Internal error!")
	}
}
