package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
