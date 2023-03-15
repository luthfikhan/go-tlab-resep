package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

func NewRecipeRepository() RecipeRepository {
	return &recipeRepository{}
}

type recipeRepository struct {
}

// AddRecipe implements RecipeRepository
func (*recipeRepository) AddRecipe(db *gorm.DB, recipe *model_db.Recipe) error {
	return db.Create(recipe).Error
}

// DeleteRecipe implements RecipeRepository
func (*recipeRepository) DeleteRecipe(db *gorm.DB, recipeId uint) error {
	return db.Delete(&model_db.Recipe{}, recipeId).Error
}

// GetRecipeById implements RecipeRepository
func (*recipeRepository) GetRecipeById(db *gorm.DB, recipeId uint) (*model_db.Recipe, error) {
	recipe := &model_db.Recipe{}
	result := db.
		Select("recipes.*, categories.id as category_id, categories.name as category").
		Joins("JOIN categories ON recipes.category_id = categories.id").
		First(recipe, recipeId)

	return recipe, result.Error
}

// GetRecipes implements RecipeRepository
func (*recipeRepository) GetRecipes(db *gorm.DB, query string) (*[]model_db.Recipe, error) {
	recipes := &[]model_db.Recipe{}

	result := db.Where("recipes.name ILIKE ?", "%"+query+"%").
		Select("recipes.*, categories.id as category_id, categories.name as category").
		Joins("JOIN categories ON recipes.category_id = categories.id").
		Find(recipes)

	return recipes, result.Error
}

// UpdateRecipe implements RecipeRepository
func (*recipeRepository) UpdateRecipe(db *gorm.DB, newRecipe *model_db.Recipe) error {
	return db.Save(newRecipe).Error
}
