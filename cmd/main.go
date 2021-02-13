package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irgalieri/item-compatibility-checker/handlers"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("ping", handlers.PingGetHandler())

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
