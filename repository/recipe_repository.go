package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	AddRecipe(db *gorm.DB, recipe *model_db.Recipe) error
	GetRecipes(db *gorm.DB, query string) (*[]model_db.Recipe, error)
	GetRecipeById(db *gorm.DB, recipeId uint) (*model_db.Recipe, error)
	UpdateRecipe(db *gorm.DB, newRecipe *model_db.Recipe) error
	DeleteRecipe(db *gorm.DB, recipeId uint) error
}
