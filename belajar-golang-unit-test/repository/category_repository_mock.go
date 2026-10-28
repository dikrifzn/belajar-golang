package repository

import (
"github.com/stretchr/testify/mock"
"belajar_golang/entity"
)

type CategoryRepositoryMock struct{
	mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil{
		return nil
	}else{
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}