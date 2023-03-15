package repository

import (
	"strings"
	"time"

	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"gorm.io/gorm"
)

var ingredientMock = []model_db.Ingredient{
	{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "ingredient 1",
	},
	{
		ID:        2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "ingredient 2",
	},
}

func NewIngredientRepository() repository.IngredientRepository {
	return &ingredientRepository{}
}

type ingredientRepository struct {
}

// AddIngredient implements IngredientRepository
func (*ingredientRepository) AddIngredient(db *gorm.DB, ingredient *model_db.Ingredient) error {
	return nil
}

// DeleteIngredient implements IngredientRepository
func (*ingredientRepository) DeleteIngredient(db *gorm.DB, ingredientId uint) error {
	return nil
}

// GetIngredientById implements IngredientRepository
func (*ingredientRepository) GetIngredientById(db *gorm.DB, ingredientId uint) (*model_db.Ingredient, error) {
	for _, v := range ingredientMock {
		if v.ID == ingredientId {
			return &v, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

// GetIngredients implements IngredientRepository
func (*ingredientRepository) GetIngredients(db *gorm.DB, query string) (*[]model_db.Ingredient, error) {
	if query == "" {
		return &ingredientMock, nil
	}

	ingredients := []model_db.Ingredient{}
	for _, v := range ingredientMock {
		if strings.Contains(v.Name, query) {
			ingredients = append(ingredients, v)
		}
	}

	return &ingredients, nil
}

// UpdateIngredient implements IngredientRepository
func (*ingredientRepository) UpdateIngredient(db *gorm.DB, newIngredient *model_db.Ingredient) error {
	return db.Save(newIngredient).Error
}
