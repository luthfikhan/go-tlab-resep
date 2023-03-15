package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

type categoryRepository struct {
}

// GetCategoryById implements CategoryRepository
func (*categoryRepository) GetCategoryById(db *gorm.DB, categoryId uint) (*model_db.Category, error) {
	category := &model_db.Category{}
	result := db.First(category, categoryId)

	return category, result.Error
}

// AddCategory implements CategoryRepository
func (*categoryRepository) AddCategory(db *gorm.DB, category *model_db.Category) error {
	return db.Create(category).Error
}

// DeleteCategory implements CategoryRepository
func (*categoryRepository) DeleteCategory(db *gorm.DB, categoryId uint) error {
	return db.Delete(&model_db.Category{}, categoryId).Error
}

// GetCategories implements CategoryRepository
func (*categoryRepository) GetCategories(db *gorm.DB, query string) (*[]model_db.Category, error) {
	categories := &[]model_db.Category{}

	result := db.Where("name ILIKE ?", "%"+query+"%").Find(categories)

	return categories, result.Error
}

// UpdateCategory implements CategoryRepository
func (*categoryRepository) UpdateCategory(db *gorm.DB, newCategory *model_db.Category) error {
	return db.Save(newCategory).Error
}
