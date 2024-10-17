package user

import (
	"database/sql"
	"fmt"

	"github.com/jiayishen21/resume-comp-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	if err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) UserExists(id string, email string) bool {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ? OR email = ?", id, email)
	if err != nil {
		return false
	}

	defer rows.Close()
	exists := rows.Next()

	return exists
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	found := false
	for rows.Next() {
		found = true
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if !found {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserById(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	found := false
	for rows.Next() {
		found = true
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if !found {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (id, email) VALUES (?, ?)", user.ID, user.Email)
	return err
}
