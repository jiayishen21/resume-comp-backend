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
