package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	// Connection string
	db, err := sql.Open("postgres", "host=database user=pedro dbname=bankapidb password=1234567 sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Check the connection to the database
	/*err = db.Ping()
	if err != nil {
		return nil, err
	}
	*/
	return db, nil
}
