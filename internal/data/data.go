package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tariqc80/appointments-exercise/internal/config"
)

// Connect opens a connection to the database using the given config
func Connect(c *config.Config) (*sql.DB, error) {
	str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName)
	db, err := sql.Open("postgres", str)

	log.Print(str)
	if err != nil {
		log.Print(err)
	}

	return db, err
}
