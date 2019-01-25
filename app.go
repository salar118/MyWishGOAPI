package main

import (
	"fmt"
	"log"

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
func (wa *WishApp) Initialize() {

	//var err error
	mongoSession, err := mongo.NewMongoSession()
	if err != nil {
		log.Println("No mongodb session")
		return
	}

	wishService := mongo.NewWishService(mongoSession)
	wa.httpServer = server.NewServer(wishService)
}

//Start the application
func (wa *WishApp) Start() {
	fmt.Println("Starting the http httpServer")
	wa.httpServer.Start()

}
