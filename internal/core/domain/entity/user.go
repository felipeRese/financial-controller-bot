package entity

import (
	"errors"
	"time"
)

type User struct {
	ID             string
	Name           string
	IsActive       bool
	ExpirationDate time.Time
}

func NewUser(id, name string) *User {
	return &User{
		ID:             id,
		Name:           name,
		IsActive:       false,
		ExpirationDate: time.Time{},
	}
}

func (u *User) SetPlan(duration time.Duration) error {
	if duration <= 0 {
		return errors.New("invalid duration: must be greater than zero")
	}

	u.IsActive = true
	u.ExpirationDate = time.Now().Add(duration)
	return nil
}
