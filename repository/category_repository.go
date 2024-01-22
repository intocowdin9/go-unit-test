package repository

import "kelas-golang-pzn/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
