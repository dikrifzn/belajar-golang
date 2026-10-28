package repository

import (
	"testing"
	"context"
	"belajar_golang_database"
	"belajar_golang_database/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentRepositoryInsert(t *testing.T) {
	comentRepository := NewCommentRepository(belajar_golang_database.OpenConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test repository",
	}

	result, err := comentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new comment with id", result)
}

func TestCommentRepositoryFindById(t *testing.T) {
	comentRepository := NewCommentRepository(belajar_golang_database.OpenConnection())

	ctx := context.Background()

	comment, err := comentRepository.FindById(ctx, 55)
	if err != nil {
		panic(err)
	}

	fmt.Println("success Find new comment with id", comment)
}

func TestCommentRepositoryFindAll(t *testing.T) {
		comentRepository := NewCommentRepository(belajar_golang_database.OpenConnection())

	ctx := context.Background()

	comments, err := comentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println("Comment", comment)
	}

	fmt.Println("success Find all comments")
}