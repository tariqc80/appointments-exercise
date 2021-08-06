package data

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)

// Insert tests

func TestInsertSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()

	defer db.Close()

	q := `INSERT INTO appointment \(start_date, end_date\) .*`

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery(q).WillReturnRows(rows)

	s := NewSchedule(db)
	start, _ := time.Parse(time.RFC3339, "2022-10-10T09:00:00")
	end, _ := time.Parse(time.RFC3339, "2022-10-10T09:30:00")
	err := s.Insert(start, end)

	if err != nil {
		t.Errorf("error %s, want no error", err)
	}
}

func TestInsertEndBeforeStartError(t *testing.T) {

}

func TestInsertStartInPastError(t *testing.T) {

}

func TestInsertDatabaseError(t *testing.T) {

}

// IsAvailable tests

func TestIsAvailableTrueSuccess(t *testing.T) {

}

func TestIsAvailableFalseSuccess(t *testing.T) {

}

func TestIsAvailableInvalidDatesError(t *testing.T) {

}

func TestIsAvailableDatabaseError(t *testing.T) {

}

// Delete tests

func TestDeleteSuccess(t *testing.T) {

}

func TestDeleteNotFoundError(t *testing.T) {

}

func TestDeleteInvalidDatesError(t *testing.T) {

}

func TestDeleteDatabaseError(t *testing.T) {

}
