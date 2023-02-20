package service

import (
	"go-unit-test/repository"
)

// CategoryService is a struct with a field named Repository of type repository.CategoryRepository.
// @property Repository - This is the repository that will be used to perform the database operations.
type CategoryService struct {
	Repository repository.CategoryRepository
}
