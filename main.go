package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the application ......")
	a := WishApp{}
	a.Initialize()
	a.Start()
}
