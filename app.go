package main

import (
	"fmt"
	"log"

	"./pkg/config"
	model "./pkg/model"
	"./pkg/mongo"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

//WishApp exposes references to the router and the database that the application uses
type WishApp struct {
	session *mongo.MongoSession
	Router *mux.Router
	DB     *mgo.Collection
}

//Initialize take in details required to connect to db
//and create db and wire up the routes to respond according
//to requirements
func (a *WishApp) Initialize(config *config.MongoConfig) {

	//var err error
	mongoSession, err := mongo.NewMongoSession(config)
	if err != nil {
		log.Println("No mongodb session")
		return
	}

	wishService := mongo.NewWishService(mongoSession, config)

	wish := model.Wish{
		Title: "wish titles",
		Story: "wish storysssswww",
	}

	wishService.CreateWish(&wish)
	fmt.Println("Inserting wish ......", wish)
}

//Start the application
func (a *WishApp) Start(addr string) {

}
