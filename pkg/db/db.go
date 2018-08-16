/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/
package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// globally defined, because the db object/connection should generally be in go
// FIXME: Should this be exported?
var db *sql.DB

// MetalScanner is ... FIXME
type MetalScanner struct {
	valid bool
	value interface{}
}

// ConnString contains the connection details for the database
type ConnString struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

// Scan for any unsigned 8-bit integers or bytes
func (scanner *MetalScanner) getBytes(src interface{}) []byte {
	if a, ok := src.([]uint8); ok {
		return a
	}
	return nil
}

// Scan the column data type and send back its value.
func (scanner *MetalScanner) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		if value, ok := src.(int64); ok {
			scanner.value = value
			scanner.valid = true
		}
	case float64:
		if value, ok := src.(float64); ok {
			scanner.value = value
			scanner.valid = true
		}
	case bool:
		if value, ok := src.(bool); ok {
			scanner.value = value
			scanner.valid = true
		}
	case string:
		value := scanner.getBytes(src)
		scanner.value = value
		scanner.valid = true
	case []byte:
		value := scanner.getBytes(src)
		scanner.value = string(value)
		scanner.valid = true
	case time.Time:
		if value, ok := src.(time.Time); ok {
			scanner.value = value
			scanner.valid = true
		}
	case nil:
		scanner.value = nil
		scanner.valid = true
	}
	return nil
}

// Establish a connection to the database.
func (connString *ConnString) establishConnection() {
	var err error
	uri := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		connString.Username, connString.Password, connString.Hostname, connString.Port, connString.Database)
	log.WithField("uri", uri).Debug("Connecting to the database using the connection string: " + uri)
	db, err = sql.Open("postgres", uri)
	if err != nil {
		log.Error(err)
		log.Panic("Failed to establish a connection to database")
	}
	err = db.Ping()
	if err != nil {
		log.Panic("Failed to ping database")
	}
}

// Close the database connection once the query has been executed
func (connString *ConnString) closeConnection() {
	err := db.Close()
	if err != nil {
		log.Warn("Failed to close database connection")
		log.Warn(err)
	}
}

// Execute the query that was supplied. We will send the error back to the user
// so that they can decide if they want to keep the error or exit the code, so
// no err will be handled here.
func (connString *ConnString) ExecuteQuery(query string) ([]map[string]interface{}, error) {

	// Make a connection to the database
	connString.establishConnection()

	// close database connection once done
	defer connString.closeConnection()

	// Initialize a data array map, which we will return
	var data []map[string]interface{}

	// Execute the query to the database.
	log.Debug("Executing the statement: " + query)
	rows, err := db.Query(query)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	// Get all the column names from the query provided by the users.
	columns, err := rows.Columns()
	if err != nil {
		return data, err
	}

	// scan row by row for the queries result
	for rows.Next() {

		// Create a interface map based on the number of columns
		// so that we can store the scanned row
		row := make([]interface{}, len(columns))
		for idx := range columns {
			row[idx] = new(MetalScanner)
		}

		// Scan for the rows and placed it on the interface map (row)
		err := rows.Scan(row...)
		if err != nil {
			return data, err
		}

		// A temp placeholder
		temp := make(map[string]interface{})

		// Now lets create a JSON for that output
		for idx, column := range columns {
			var scanner = row[idx].(*MetalScanner)
			temp[column] = scanner.value
		}

		// Store that on the data array map
		data = append(data, temp)

	}

	// Send the data back to the user for further
	// manipulation or to what their code depends
	return data, nil
}
