package belajar_errors

import (
	"errors"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Dikri" {
		return NotFoundError
	}

	return nil
}