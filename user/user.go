package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	UserName string
	Phone    *string
}
