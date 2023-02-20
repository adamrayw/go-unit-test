package repository

import (
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock is a struct with a field named Mock of type mock.Mock.
// @property Mock - This is the mock object that will be used to mock the methods of the
// CategoryRepository interface.
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// A method that is being defined on the CategoryRepositoryMock struct.
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {

	// Calling the `Called` method of the `Mock` object.
	arguments := repository.Mock.Called(id)

	// This is a way to return a pointer to a struct.
	if arguments.Get(0) == nil {
		return nil
	} else {
		// Casting the value returned by the `Get` method to the `entity.Category` type.
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
