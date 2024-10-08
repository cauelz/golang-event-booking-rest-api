package models

import (
	"time"

	"github.com/cauelz/golang-event-booking-rest-api/db"
)

type Event struct {
	ID		  int64    `json:"id"`
	Name 	string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location string `json:"location" binding:"required"`
	DateTime time.Time `json:"dateTime" binding:"required"`
	UserId int64 `json:"user_id"`
}

func (e *Event) Save() error {

	// query is a SQL query that inserts a new event into the events table.
	// The query uses placeholders (?) to represent the values that will be inserted into the table. 
	//It's a good practice to use placeholders to prevent SQL injection attacks.
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	row, error := db.DB.Query(query)
 
	if(error != nil) {
		return nil, error
	}
	// defer is used to ensure that the row is closed after the function returns.
	defer row.Close()

	events := []Event{}

	// row.Next() is used to iterate over the rows returned by the query.
	for row.Next() {
		var e Event
		// row.Scan() is used to scan the values from the current row into the Event struct.
		// The values are scanned in the same order as they appear in the SELECT statement.
		// The values are passed as pointers to the fields of the Event struct.
		err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		if(err != nil) {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(query, id)

	var e Event
	error := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

	if error != nil {
		return nil, error
	}

	return &e, nil

}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.ID)

	return err
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (e *Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func Unregister(userId int64, eventId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(eventId, userId)

	return err
}