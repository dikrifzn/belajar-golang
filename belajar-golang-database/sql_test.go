package belajar_golang_database

import (
	"context"
	"testing"
	"fmt"
	"time"
	"database/sql"
	"strconv"
)

func TestExecuteSql(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('dikri', 'Dikri', 'dikri@example.com', 10000, 5.0, '2023-01-01', true)"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQueryGetSql(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()
	
	ctx := context.Background()

	script := "SELECT id, name FROM customer WHERE id = 'dikri'"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
	defer rows.Close()
}

func TestQueryGetComplex(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()
	
	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString // Use sql.NullString to handle nullable string values
		var balance int
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime // Use sql.NullTime to handle nullable time values
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		} // Check if email is valid before printing
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date :", birthDate.Time)
		} // Check if birth date is valid before printing
		fmt.Println("Married :", married)
		fmt.Println("Created At :", createdAt)
	}
	defer rows.Close()
}


func TestQueryGetSqlInjection(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin' ; #"
	password := "sala"

	// script := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	scrtipt := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"

	// rows, err := db.QueryContext(ctx, script)
	rows, err := db.QueryContext(ctx, scrtipt, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Login successful:", username)
	
	}else {
		fmt.Println("Username not found")
	}
	defer rows.Close()
}

func TestExecuteSqlInjection(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()

	ctx := context.Background()

	// script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('dikri', 'Dikri', 'dikri@example.com', 10000, 5.0, '2023-01-01', true)"
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, script, "fikri", "Fikri", "fikri@example.com", 10000, 5.0, "2023-01-01", true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, "fikri@example.com", "This is a test comment")
	if err != nil {
		panic(err)
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Insert ID:", insertedId)

	fmt.Println("Success new comment")
}

func TestPrepareStatement(t *testing.T) {
	db , _ := OpenConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "dikri" + strconv.Itoa(i) + "@example.com"
		comment := "Comment ke-" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new comment ke-", id)
	}
}

func TestTransaction(t *testing.T) {
	db, _ := OpenConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do something with transaction
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i := 0; i < 10; i++ {
		email := "fikri" + strconv.Itoa(i) + "@example.com"
		comment := "Comment ke-" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new comment ke-", id)
	}

	err = tx.Commit() // commit transaction
	// er = tx.Rollback() // rollback transaction if error occurs
	if err != nil {
		panic(err)
	}
}