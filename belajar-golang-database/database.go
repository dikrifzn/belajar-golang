package belajar_golang_database

import (
	"database/sql"
	"time"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5) // Set the maximum number of idle connections in the pool
	db.SetMaxOpenConns(20) // Set the maximum number of open connections to the database
	db.SetConnMaxIdleTime(5 * time.Minute) // Set the maximum amount of time a connection may be idle
	db.SetConnMaxLifetime(60 * time.Minute) // Set the maximum amount of time a connection may be reused
	return db
}