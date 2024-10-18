package utils

import (
	"database/sql"
	"fmt"
	"time"
)

// ParseDate parses a date string in the format "2006-01-02" and returns a time.Time object.
func ParseDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}
	return parsedDate, nil
}

func TimeToNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: t, Valid: true}
}
