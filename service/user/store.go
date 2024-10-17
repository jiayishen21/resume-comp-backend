package user

import (
	"database/sql"
	"fmt"

	"github.com/jiayishen21/resume-comp-backend/types"
	"github.com/jiayishen21/resume-comp-backend/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
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
		user, err = utils.ScanRowIntoUser(rows)
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
		user, err = utils.ScanRowIntoUser(rows)
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
	_, err := s.db.Exec("INSERT INTO users (id, email, display_name) VALUES (?, ?, ?)", user.ID, user.Email, user.DisplayName)
	return err
}
