// db/db.go

package db

import (
	"chat-app-microservice/calling-service/model"
	"database/sql"
	"fmt"
)

// Database struct to hold the connection
type Database struct {
	Conn *sql.DB
}

// Initialize initializes the database connection
func Initialize() (*Database, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "mychatapp"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

// CreateCall inserts a new call into the database
func (db *Database) CreateCall(call model.Call) error {
	// Implement the logic to insert a new call into the database
	// You should use prepared statements to prevent SQL injection
	// Example: _, err := db.Conn.Exec("INSERT INTO calls (call_id, channel_id, call_type, status) VALUES ($1, $2, $3, $4)", call.CallID, call.ChannelID, call.Type, call.Status)
	// Handle the error appropriately
	return nil
}

// GetCall retrieves call details from the database based on the call ID
func (db *Database) GetCall(callID string) (*model.Call, error) {
	// Implement the logic to retrieve call details from the database based on the call ID
	// Example: row := db.Conn.QueryRow("SELECT id, call_id, channel_id, call_type, status FROM calls WHERE call_id = $1", callID)
	// Scan the row into a Call struct and return it
	return nil, nil
}

// UpdateCall updates the status of an existing call in the database
func (db *Database) UpdateCall(call model.Call) error {
	// Implement the logic to update the status of an existing call in the database
	// Example: _, err := db.Conn.Exec("UPDATE calls SET status = $1 WHERE call_id = $2", call.Status, call.CallID)
	// Handle the error appropriately
	return nil
}

// CloseCallIfEmpty closes a call if both users have left
func (db *Database) CloseCallIfEmpty(channelID int) error {
	// Implement the logic to close a call if both users have left
	// Example: _, err := db.Conn.Exec("DELETE FROM calls WHERE channel_id = $1", channelID)
	// Handle the error appropriately
	return nil
}
