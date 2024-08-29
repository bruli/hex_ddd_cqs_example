package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
	"hex_ddd_cqs_example/user"
	"net/http"
)

func FindUser(sess db.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		found, err := user.FindUser(c.Request.Context(), sess, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, found)
	}
}
