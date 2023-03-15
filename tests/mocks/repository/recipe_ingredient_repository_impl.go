package repository

import (
	"time"

	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"gorm.io/gorm"
)

var recipeIngredientMock = []model_db.RecipeIngredient{
	{
		ID:           1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Name:         "nasi",
		Amount:       "1 piring",
		RecipeID:     1,
		IngredientID: 2,
	},
	{
		ID:           2,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Name:         "gula",
		Amount:       "1 sendok teh",
		RecipeID:     1,
		IngredientID: 1,
	},
}

func NewRecipeIngredientRepository() repository.RecipeIngredientRepository {
	return &recipeIngredientRepository{}
}

type recipeIngredientRepository struct {
}

// AddRecipeIngredient implements RecipeIngredientRepository
func (*recipeIngredientRepository) AddRecipeIngredient(db *gorm.DB, recipeIngredient *model_db.RecipeIngredient) error {
	return nil
}

// DeleteRecipeIngredient implements RecipeIngredientRepository
func (*recipeIngredientRepository) DeleteRecipeIngredient(db *gorm.DB, recipeIngredientId uint) error {
	return nil
}

// GetRecipeIngredientById implements RecipeIngredientRepository
func (*recipeIngredientRepository) GetRecipeIngredientById(db *gorm.DB, recipeIngredientId uint) (*model_db.RecipeIngredient, error) {
	for _, v := range recipeIngredientMock {
		if v.ID == recipeIngredientId {
			return &v, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

// GetRecipeIngredientsByRecipeId implements RecipeIngredientRepository
func (*recipeIngredientRepository) GetRecipeIngredientsByRecipeId(db *gorm.DB, recipeId uint) (*[]model_db.RecipeIngredient, error) {
	recipeIngredients := []model_db.RecipeIngredient{}
	for _, v := range recipeIngredientMock {
		if v.RecipeID == recipeId {
			recipeIngredients = append(recipeIngredients, v)
		}
	}

	if len(recipeIngredients) > 0 {
		return &recipeIngredients, nil
	}

	return nil, gorm.ErrRecordNotFound
}

// UpdateRecipeIngredient implements RecipeIngredientRepository
func (*recipeIngredientRepository) UpdateRecipeIngredient(db *gorm.DB, newRecipeIngredient *model_db.RecipeIngredient) error {
	return nil
}
