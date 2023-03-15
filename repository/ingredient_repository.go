package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

type IngredientRepository interface {
	AddIngredient(db *gorm.DB, ingredient *model_db.Ingredient) error
	GetIngredients(db *gorm.DB, query string) (*[]model_db.Ingredient, error)
	GetIngredientById(db *gorm.DB, ingredientId uint) (*model_db.Ingredient, error)
	UpdateIngredient(db *gorm.DB, newIngredient *model_db.Ingredient) error
	DeleteIngredient(db *gorm.DB, ingredientId uint) error
}
