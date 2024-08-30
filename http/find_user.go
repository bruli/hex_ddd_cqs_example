package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hex_ddd_cqs_example/internal/domain/user"
	"net/http"
)

func FindUser(userRepo user.UserRepository) gin.HandlerFunc { //NEW injecting user repository
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		found, err := user.FindUser(c.Request.Context(), userRepo, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, found)
	}
}
