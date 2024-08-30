package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hex_ddd_cqs_example/internal/app"
	"hex_ddd_cqs_example/internal/domain/user"
	"net/http"
)

func CreteUser(ch app.CommandHandler) gin.HandlerFunc { //NEW injecting user repository
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
		if _, err := ch.Handle(c.Request.Context(), app.CreateUserCommand{
			ID:       id,
			UserName: req.Username,
			Phone:    req.Phone,
		}); err != nil {
			switch {
			case errors.As(err, &user.CreateUserError{}):
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			default:
				c.Status(http.StatusInternalServerError)
			}
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
