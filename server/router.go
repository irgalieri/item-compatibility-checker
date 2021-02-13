package server

import (
	"github.com/gin-gonic/gin"
	"github.com/irgalieri/item-compatibility-checker/handlers"
)

func SetRoutes(r *gin.Engine) {
	r.GET("ping", handlers.PingGetHandler())
}
