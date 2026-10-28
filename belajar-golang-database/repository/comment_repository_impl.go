package repository

import (
	"database/sql"
	"context"
	"belajar_golang_database/entity"
	"errors"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	
	if err != nil {
		return comment, err
	}
	resultId, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(resultId)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ?"
	
	rows, err := repository.DB.QueryContext(ctx, script, id)

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		//tidak ada
		return comment, errors.New("comment not found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		//ada
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}