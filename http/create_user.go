package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hex_ddd_cqs_example/user"
	"net/http"
)

func CreteUser(userRepo user.UserRepository) gin.HandlerFunc { //NEW injecting user repository
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := uuid.Parse(req.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		if err := user.CreateUser(c.Request.Context(), userRepo, user.User{
			ID:       id,
			UserName: req.Username,
			Phone:    req.Phone,
		}); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	}
}

type CreateUserRequest struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Phone    *string `json:"phone,omitempty"`
}
