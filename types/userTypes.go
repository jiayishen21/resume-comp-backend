package types

import (
	"time"
)

type User struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`

	Private bool `json:"private"`

	Company  string `json:"company"`
	Position string `json:"position"`

	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`

	CreatedAt time.Time `json:"createdAt"`
}

type Auth0UserInfo struct {
	Sub           string `json:"sub" validate:"required"` // of format auth0|ab3sdjas134
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	UpdatedAt     string `json:"updated_at"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

// store

type UserStore interface {
	UserExists(id string, email string) bool
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(user *User) error
}

// payloads

// type RegisterUserPayload struct {
// 	FirstName string `json:"firstName" validate:"required"`
// 	LastName  string `json:"lastName" validate:"required"`
// 	Email     string `json:"email" validate:"required,email"`
// }
