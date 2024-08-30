package user

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrInvalidUserName = errors.New("invalid user name")
)

type User struct {
	id       uuid.UUID
	userName string
	phone    *string
}

func (u User) Id() uuid.UUID {
	return u.id
}

func (u User) UserName() string {
	return u.userName
}

func (u User) Phone() *string {
	return u.phone
}

func (u User) validate() error {
	switch {
	case u.id == uuid.Nil:
		return ErrInvalidUserID
	case u.userName == "":
		return ErrInvalidUserName
	default:
		return nil
	}
}

func NewUser(id uuid.UUID, userName string, phone *string) (*User, error) {
	u := User{id: id, userName: userName, phone: phone}
	if err := u.validate(); err != nil {
		return nil, err
	}
	return &u, nil
}
