package server

import (
	"fmt"
	"net/http"

	"../model"
	"github.com/gorilla/mux"
)

type wishRouter struct {
	wishService model.WishServicer
}

//NewWishServer returns a router for processing wish requests
func NewWishServer(u model.WishServicer, router *mux.Router) *mux.Router {
	wishRouter := wishRouter{
		wishService: u,
	}

	router.HandleFunc("/", wishRouter.createWishHandler).Methods("GET")
	return router

}

func (wishRouter *wishRouter) createWishHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("zit in  createWishHandler")
	//Remove later
	wish := model.Wish{
		Title: "wish titles",
		Story: "wish storysssswww",
	}

	wishRouter.wishService.CreateWish(&wish)

	//
	Json(w, http.StatusOK, wish)
}
