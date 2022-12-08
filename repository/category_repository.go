package repository

import "github.com/RandyWiratamaa/app-golang/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
