package repository

import "go-unit-test/entity"

// // Go
// type CategoryRepository interface {
// 	FindById(id string) *entity.Category
// }
// @property FindById - This is the method that will be used to find a category by its id.
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
