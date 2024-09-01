package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is a pointer to a sql.DB struct. sql.DB is a database handle representing a pool of zero or more underlying connections.
// It's safe for concurrent use by multiple goroutines.
var DB *sql.DB

func InitDB() {
	var err error

	// sql.Open is used to open a database specified by its database driver name and a driver-specific data source name, 
	// usually consisting of at least a database name and connection information.
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database. 
	//It means that the database can have up to 10 open connections at the same time.
	DB.SetMaxOpenConns(10)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// It means that the database can have up to 5 idle connections at the same time. What are idle connections?
	// Idle connections are connections that are not being used at the moment.
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)`
	
	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create Users table!")
	}

	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`

	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create Events table!")
	}
}