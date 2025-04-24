package main

import (
	"net/http"
	"github.com/gorilla/mux"
)


type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()
	router.HandleFunc("/", DisplayHome).Methods("GET")
	router.HandleFunc("/world", UpdateWorld).Methods("GET")

  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return &Server{Router: router}
}

func DisplayHome(http.ResponseWriter, *http.Request){}
func UpdateWorld(http.ResponseWriter, *http.Request){}
