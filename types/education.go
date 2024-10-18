package types

import (
	"time"
)

type Education struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`

	Field string  `json:"field"`
	Minor string  `json:"minor"`
	GPA   float64 `json:"gpa"`

	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`

	Current   bool      `json:"current"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

// store

type EducationStore interface {
	GetEducationByUserId(userId string) ([]*Education, error)
	AddEducation(education *Education) error
	UpdateEducation(education *Education) error
	DeleteEducation(id int, userId string) error
}

// payloads

type AddEducationPayload struct {
	Institution string  `json:"institution" validate:"required,max=255"`
	Degree      string  `json:"degree" validate:"required,max=255"`
	Field       string  `json:"field" validate:"max=255"`
	Minor       string  `json:"minor" validate:"max=255"`
	GPA         float64 `json:"gpa" validate:"gte=0,lte=100"`
	Country     string  `json:"country" validate:"max=255"`
	State       string  `json:"state" validate:"max=255"`
	City        string  `json:"city" validate:"max=255"`
	Current     bool    `json:"current"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
}

type UpdateEducationPayload struct {
	ID          int     `json:"id" validate:"required"`
	Institution string  `json:"institution" validate:"required,max=255"`
	Degree      string  `json:"degree" validate:"required,max=255"`
	Field       string  `json:"field" validate:"max=255"`
	Minor       string  `json:"minor" validate:"max=255"`
	GPA         float64 `json:"gpa" validate:"gte=0,lte=100"`
	Country     string  `json:"country" validate:"max=255"`
	State       string  `json:"state" validate:"max=255"`
	City        string  `json:"city" validate:"max=255"`
	Current     bool    `json:"current"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
}

type DeleteEducationPayload struct {
	ID int `json:"id" validate:"required"`
}
