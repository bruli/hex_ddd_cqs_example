package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hex_ddd_cqs_example/internal/app"
	"hex_ddd_cqs_example/internal/domain/user"
	"net/http"
)

func FindUser(qh app.QueryHandler) gin.HandlerFunc { //NEW injecting user repository
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		found, err := qh.Handle(c.Request.Context(), app.FindUserQuery{ID: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		us, ok := found.(*user.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user type"})
			return
		}
		c.JSON(http.StatusOK, UserResponse{
			ID:          us.Id(),
			UserName:    us.UserName(),
			PhoneNumber: us.Phone(),
		})
	}
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	UserName    string    `json:"user_name"`
	PhoneNumber *string   `json:"phone_number,omitempty"`
}
