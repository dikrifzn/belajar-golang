package belajar_golang_database

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}