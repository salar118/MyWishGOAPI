package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//HttpServer contains router
type HttpServer struct {
	router *mux.Router
	port   string
}

//NewServer Create a new HttpServer
func NewServer() *HttpServer {
	s := HttpServer{
		router: mux.NewRouter(),
	}

	return &s
}

//Start  the http server
func (a *HttpServer) Start() {
	log.Println("Linsting on port", a.port)
	err := http.ListenAndServe(a.port, nil)

	if err == nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
