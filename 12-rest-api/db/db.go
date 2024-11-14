package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connectionString := os.Getenv("DB")
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	  id SERIAL PRIMARY KEY,
	  email TEXT NOT NULL UNIQUE,
	  password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	  id SERIAL PRIMARY KEY,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime TIMESTAMP NOT NULL,
	  user_id INTEGER,
	  FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	  event_id INTEGER,
	  user_id INTEGER,
	  FOREIGN KEY (event_id) REFERENCES events(id),
	  FOREIGN KEY (user_id) REFERENCES users(id),
	  PRIMARY KEY (event_id, user_id)
	)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
