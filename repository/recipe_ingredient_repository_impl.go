package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

func NewRecipeIngredientRepository() RecipeIngredientRepository {
	return &recipeIngredientRepository{}
}

type recipeIngredientRepository struct {
}

// AddRecipeIngredient implements RecipeIngredientRepository
func (*recipeIngredientRepository) AddRecipeIngredient(db *gorm.DB, recipeIngredient *model_db.RecipeIngredient) error {
	return db.Create(recipeIngredient).Error
}

// DeleteRecipeIngredient implements RecipeIngredientRepository
func (*recipeIngredientRepository) DeleteRecipeIngredient(db *gorm.DB, recipeIngredientId uint) error {
	return db.Delete(&model_db.RecipeIngredient{}, recipeIngredientId).Error
}

// GetRecipeIngredientById implements RecipeIngredientRepository
func (*recipeIngredientRepository) GetRecipeIngredientById(db *gorm.DB, recipeIngredientId uint) (*model_db.RecipeIngredient, error) {
	recipeIngredient := &model_db.RecipeIngredient{}
	result := db.First(recipeIngredient, recipeIngredientId)

	return recipeIngredient, result.Error
}

// GetRecipeIngredientsByRecipeId implements RecipeIngredientRepository
func (*recipeIngredientRepository) GetRecipeIngredientsByRecipeId(db *gorm.DB, recipeId uint) (*[]model_db.RecipeIngredient, error) {
	recipeIngredients := &[]model_db.RecipeIngredient{}

	result := db.Where("recipe_id = ?", recipeId).
		Select("recipe_ingredients.*, ingredients.name").
		Joins("JOIN ingredients ON recipe_ingredients.ingredient_id = ingredients.id").
		Find(recipeIngredients)

	return recipeIngredients, result.Error
}

// UpdateRecipeIngredient implements RecipeIngredientRepository
func (*recipeIngredientRepository) UpdateRecipeIngredient(db *gorm.DB, newRecipeIngredient *model_db.RecipeIngredient) error {
	return db.Save(newRecipeIngredient).Error
}
