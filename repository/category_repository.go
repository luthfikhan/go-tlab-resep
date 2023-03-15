package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	AddCategory(db *gorm.DB, category *model_db.Category) error
	GetCategories(db *gorm.DB, query string) (*[]model_db.Category, error)
	GetCategoryById(db *gorm.DB, categoryId uint) (*model_db.Category, error)
	UpdateCategory(db *gorm.DB, newCategory *model_db.Category) error
	DeleteCategory(db *gorm.DB, categoryId uint) error
}
