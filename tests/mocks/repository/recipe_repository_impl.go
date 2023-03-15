package repository

import (
	"strings"
	"time"

	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"gorm.io/gorm"
)

var recipeMock = []model_db.Recipe{
	{
		ID:         1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Name:       "nasi Goreng",
		CategoryID: 1,
		Category:   "Category1",
	},
}

func NewRecipeRepository() repository.RecipeRepository {
	return &recipeRepository{}
}

type recipeRepository struct {
}

// AddRecipe implements RecipeRepository
func (*recipeRepository) AddRecipe(db *gorm.DB, recipe *model_db.Recipe) error {
	return nil
}

// DeleteRecipe implements RecipeRepository
func (*recipeRepository) DeleteRecipe(db *gorm.DB, recipeId uint) error {
	return nil
}

// GetRecipeById implements RecipeRepository
func (*recipeRepository) GetRecipeById(db *gorm.DB, recipeId uint) (*model_db.Recipe, error) {
	for _, v := range recipeMock {
		if v.ID == recipeId {
			return &v, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

// GetRecipes implements RecipeRepository
func (*recipeRepository) GetRecipes(db *gorm.DB, query string) (*[]model_db.Recipe, error) {
	recipes := []model_db.Recipe{}
	if query == "" {
		return &recipeMock, nil
	}

	for _, v := range recipeMock {
		if strings.Contains(v.Name, query) {
			recipes = append(recipes, v)
		}
	}

	return &recipes, nil
}

// UpdateRecipe implements RecipeRepository
func (*recipeRepository) UpdateRecipe(db *gorm.DB, newRecipe *model_db.Recipe) error {
	return nil
}
