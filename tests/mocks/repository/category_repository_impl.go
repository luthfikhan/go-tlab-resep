package repository

import (
	"strings"
	"time"

	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"gorm.io/gorm"
)

var categoriesMock = []model_db.Category{
	{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Category 1",
	},
	{
		ID:        2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Category 2",
	},
}

func NewCategoryRepository() repository.CategoryRepository {
	return &categoryRepository{}
}

type categoryRepository struct {
}

// GetCategoryById implements CategoryRepository
func (*categoryRepository) GetCategoryById(db *gorm.DB, categoryId uint) (*model_db.Category, error) {
	for _, v := range categoriesMock {
		if v.ID == categoryId {
			return &v, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

// AddCategory implements CategoryRepository
func (*categoryRepository) AddCategory(db *gorm.DB, category *model_db.Category) error {
	return nil
}

// DeleteCategory implements CategoryRepository
func (*categoryRepository) DeleteCategory(db *gorm.DB, categoryId uint) error {
	return nil
}

// GetCategories implements CategoryRepository
func (*categoryRepository) GetCategories(db *gorm.DB, query string) (*[]model_db.Category, error) {
	if query == "" {
		return &categoriesMock, nil
	}

	categories := []model_db.Category{}
	for _, v := range categoriesMock {
		if strings.Contains(v.Name, query) {
			categories = append(categories, v)
		}
	}

	return &categories, nil
}

// UpdateCategory implements CategoryRepository
func (*categoryRepository) UpdateCategory(db *gorm.DB, newCategory *model_db.Category) error {
	return nil
}
