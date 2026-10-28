package repository

import "belajar_golang/entity"

type CategoryRepository interface{
	FindById(id string) *entity.Category
}