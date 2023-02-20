package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

// CategoryService is a struct with a field named Repository of type repository.CategoryRepository.
// @property Repository - This is the repository that will be used to perform the database operations.
type CategoryService struct {
	Repository repository.CategoryRepository
}

// A method that belongs to the CategoryService struct.
func (service CategoryService) Get(id string) (*entity.Category, error) {
	// Calling the FindById method of the CategoryRepository struct.
	category := service.Repository.FindById(id)

	// Checking if the category is nil and if it is, it returns an error.
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
