package main

import (
	"fmt"
	"log"

	"./pkg/config"
	"./pkg/mongo"
	"./pkg/server"
)

//WishApp exposes references to the router and the database that the application uses
type WishApp struct {
	httpServer *server.HTTPServer
	session    *mongo.MongoSession
}

//Initialize take in details required to connect to db
//and create db and wire up the routes to respond according
//to requirements
func (wa *WishApp) Initialize(mongoConfig *config.MongoConfig) {


	//var err error
	mongoSession, err := mongo.NewMongoSession(mongoConfig)
	if err != nil {
		log.Println("No mongodb session")
		return
	}


	//u := mongo.NewUserService(a.session.Copy(), a.config.Mongo)

	wishService := mongo.NewWishService(mongoSession, mongoConfig)

	wa.httpServer = server.NewServer(wishService)

}

//Start the application
func (wa *WishApp) Start() {
	fmt.Println("Starting the http httpServer")
	wa.httpServer.Start()

}
