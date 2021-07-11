package main

import (
	"gowebapp1/middlewares"
	"gowebapp1/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	Host = "0.0.0.0"
	Port = "8000"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Index)
	r.HandleFunc("/hello", routes.Hello)
	r.HandleFunc("/login", routes.Login)
	r.HandleFunc("/logout", routes.Logout)
	r.HandleFunc("/register", routes.Register)
	r.HandleFunc("/poll/create", routes.CreatePoll)
	r.HandleFunc("/poll/{ID}", routes.Vote)
	r.HandleFunc("/poll/{ID}/results", routes.Results)
	WrappedRouter := middlewares.RequestLogger(r)
	log.Fatal(http.ListenAndServe(Host+":"+Port, WrappedRouter))
}
