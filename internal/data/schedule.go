package data

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

// Schedule struct is a data provider for appointment scheduling
type Schedule struct {
	db *sql.DB
}

// NewSchedule creates and returns a new data handler
func NewSchedule(db *sql.DB) *Schedule {
	return &Schedule{
		db: db,
	}
}

// Insert inserts a new appointment into the database
func (d *Schedule) Insert(start time.Time, end time.Time) error {

	q := `INSERT INTO appointment (start_date, end_date) VALUES ($1, $2) RETURNING id`

	var id int

	err := d.db.QueryRow(q, start, end).Scan(&id)

	log.Printf("Inserted row with id %d", id)

	return err
}

// IsAvailable checks ithe database to see if start and end dates are free to book
func (d *Schedule) IsAvailable(start time.Time, end time.Time) (bool, error) {

	q := `
SELECT id
FROM appointment
WHERE $1 >= start_date AND $1 < end_date AND $2 > start_date AND $2 <= end_date
`
	row := d.db.QueryRow(q, start.Format(time.RFC3339), end.Format(time.RFC3339))

	var data int
	err := row.Scan(&data)

	if err != nil {
		// if we have no rows, sql returns an error
		// check for this error and return true.
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
	}

	log.Printf("Found existing appointment row with id %d", data)

	return false, err
}

// Delete removes an appointment from the database
func (d *Schedule) Delete(start time.Time, end time.Time) (bool, error) {

	q := `DELETE FROM appointment WHERE start_date = $1 AND end_date = $2`

	res, err := d.db.Exec(q, start.Format(time.RFC3339), end.Format(time.RFC3339))

	if err != nil {
		log.Print(err)
		return false, err
	}

	count, err := res.RowsAffected()
	log.Printf("Deleted %d rows", count)

	return count > 0, err
}
