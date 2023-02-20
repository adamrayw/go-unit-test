package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Creating a new instance of the `CategoryRepositoryMock` struct.
var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}

// Creating a new instance of the `CategoryService` struct.
var categoryService = CategoryService{Repository: &categoryRepository}

// It tests the CategoryService
func TestCategoryService(t *testing.T) {
	// program mock
	// Mocking the `FindById` method of the `CategoryRepository` struct.
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// Calling the `Get` method of the `CategoryService` struct.
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

// `TestCategoryService_GetSuccess`
func TestCategoryService_GetSuccess(t *testing.T) {
	// Creating a new instance of the `Category` struct.
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	// Mocking the `FindById` method of the `CategoryRepository` struct.
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// Calling the `Get` method of the `CategoryService` struct.
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
