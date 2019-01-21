package main

import (
	"fmt"

	"./pkg/config"
)

func main() {
	fmt.Println("Starting the application ......")
	a := WishApp{}
	mongoconfig := config.MongoConfig{
		IP:     "127.0.0.1:27017",
		Port:   ":27017",
		DbName: "WishDB",
	}
	a.Initialize(&mongoconfig)
	a.Start(mongoconfig.Port)
}
