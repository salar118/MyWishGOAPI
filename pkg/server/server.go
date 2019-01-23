package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../config"
	"../model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//HTTPServer contains router
type HTTPServer struct {
	router *mux.Router
	config *config.HttpServerConfig
}

//NewServer Create a new HttpServer
func NewServer(userServicer model.WishServicer) *HTTPServer {
	s := HTTPServer{
		router: mux.NewRouter(),
	}

	s.config = config.GetConfig()
	s.router = NewWishServer(userServicer, s.getSubrouter("/wish"))

	if s.router != nil {
		fmt.Println("Mux router is not nil")
	}
	return &s
}

func (s *HTTPServer) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}

//Start http server
func (s *HTTPServer) Start() {
	fmt.Printf("%+v\n", s)
	log.Println("Listening on port ", s.config)
	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
