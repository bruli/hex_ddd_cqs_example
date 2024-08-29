package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Ok"})
	}
}
