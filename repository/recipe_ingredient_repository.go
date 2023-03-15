package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

type RecipeIngredientRepository interface {
	AddRecipeIngredient(db *gorm.DB, recipeIngredient *model_db.RecipeIngredient) error
	GetRecipeIngredientsByRecipeId(db *gorm.DB, recipeId uint) (*[]model_db.RecipeIngredient, error)
	GetRecipeIngredientById(db *gorm.DB, recipeIngredientId uint) (*model_db.RecipeIngredient, error)
	UpdateRecipeIngredient(db *gorm.DB, newRecipeIngredient *model_db.RecipeIngredient) error
	DeleteRecipeIngredient(db *gorm.DB, recipeIngredientId uint) error
}
