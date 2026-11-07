package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang_rest_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
