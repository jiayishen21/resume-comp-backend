package education

import (
	"database/sql"

	"github.com/jiayishen21/resume-comp-backend/types"
	"github.com/jiayishen21/resume-comp-backend/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetEducationByUserId(userId string) ([]*types.Education, error) {
	rows, err := s.db.Query("SELECT * FROM education WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educations []*types.Education
	for rows.Next() {
		education, err := utils.ScanRowIntoEducation(rows)
		if err != nil {
			return nil, err
		}
		educations = append(educations, education)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return educations, nil
}

func (s *Store) AddEducation(education *types.Education) error {
	query := `
        INSERT INTO education (user_id, institution, degree, field, minor, gpa, country, state, city, current, start_date, end_date)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	startDate := utils.TimeToNullTime(education.StartDate)
	endDate := utils.TimeToNullTime(education.EndDate)

	// Make sure to pass the actual startDate
	_, err := s.db.Exec(query,
		education.UserID,
		education.Institution,
		education.Degree,
		education.Field,
		education.Minor,
		education.GPA,
		education.Country,
		education.State,
		education.City,
		education.Current,
		startDate,
		endDate,
	)

	return err
}
