package utils

import (
	"database/sql"
	"time"

	"github.com/jiayishen21/resume-comp-backend/types"
)

type SqlUser struct {
	ID          string
	Email       string
	DisplayName string

	Private bool

	Company  sql.NullString
	Position sql.NullString

	Country sql.NullString
	State   sql.NullString
	City    sql.NullString

	CreatedAt time.Time
}

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	sqlUser := new(SqlUser)
	if err := rows.Scan(
		&sqlUser.ID,
		&sqlUser.Email,
		&sqlUser.DisplayName,

		&sqlUser.Private,

		&sqlUser.Company,
		&sqlUser.Position,

		&sqlUser.Country,
		&sqlUser.State,
		&sqlUser.City,

		&sqlUser.CreatedAt,
	); err != nil {
		return nil, err
	}

	user := &types.User{
		ID:          sqlUser.ID,
		Email:       sqlUser.Email,
		DisplayName: sqlUser.DisplayName,

		Private: sqlUser.Private,

		Company:  sqlUser.Company.String,
		Position: sqlUser.Position.String,

		Country: sqlUser.Country.String,
		State:   sqlUser.State.String,
		City:    sqlUser.City.String,

		CreatedAt: sqlUser.CreatedAt,
	}

	return user, nil
}

type SqlEducation struct {
	ID          int
	UserID      string
	Institution string
	Degree      string

	Field sql.NullString
	Minor sql.NullString
	GPA   sql.NullFloat64

	Country sql.NullString
	State   sql.NullString
	City    sql.NullString

	Current   bool
	StartDate sql.NullTime
	EndDate   sql.NullTime
}

func ScanRowIntoEducation(rows *sql.Rows) (*types.Education, error) {
	sqlEducation := new(SqlEducation)
	if err := rows.Scan(
		&sqlEducation.ID,
		&sqlEducation.UserID,
		&sqlEducation.Institution,
		&sqlEducation.Degree,

		&sqlEducation.Field,
		&sqlEducation.Minor,
		&sqlEducation.GPA,

		&sqlEducation.Country,
		&sqlEducation.State,
		&sqlEducation.City,

		&sqlEducation.Current,
		&sqlEducation.StartDate,
		&sqlEducation.EndDate,
	); err != nil {
		return nil, err
	}

	education := &types.Education{
		ID:          sqlEducation.ID,
		UserID:      sqlEducation.UserID,
		Institution: sqlEducation.Institution,
		Degree:      sqlEducation.Degree,

		Field: sqlEducation.Field.String,
		Minor: sqlEducation.Minor.String,
		GPA:   sqlEducation.GPA.Float64,

		Country: sqlEducation.Country.String,
		State:   sqlEducation.State.String,
		City:    sqlEducation.City.String,

		Current:   sqlEducation.Current,
		StartDate: sqlEducation.StartDate.Time,
		EndDate:   sqlEducation.EndDate.Time,
	}

	return education, nil
}
