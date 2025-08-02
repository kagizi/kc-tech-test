package entity

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
}

func NewUser(name, phone string) (*User, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	now := time.Now()

	user := &User{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Phone:     phone,
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name must not be empty")
	}

	if u.Phone == "" {
		return errors.New("phone must not be empty")
	}

	if len(u.Phone) < 10 {
		return errors.New("phone must be at least 10 characters")
	}

	return nil
}
