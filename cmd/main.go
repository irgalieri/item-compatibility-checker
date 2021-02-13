package main

import (
	"github.com/irgalieri/item-compatibility-checker/server"
	"log"
)

func main() {
	err := server.NewServer().Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
