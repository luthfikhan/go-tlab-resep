package repository

import (
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	"gorm.io/gorm"
)

func NewIngredientRepository() IngredientRepository {
	return &ingredientRepository{}
}

type ingredientRepository struct {
}

// AddIngredient implements IngredientRepository
func (*ingredientRepository) AddIngredient(db *gorm.DB, ingredient *model_db.Ingredient) error {
	return db.Create(ingredient).Error
}

// DeleteIngredient implements IngredientRepository
func (*ingredientRepository) DeleteIngredient(db *gorm.DB, ingredientId uint) error {
	return db.Delete(&model_db.Ingredient{}, ingredientId).Error
}

// GetIngredientById implements IngredientRepository
func (*ingredientRepository) GetIngredientById(db *gorm.DB, ingredientId uint) (*model_db.Ingredient, error) {
	ingredient := &model_db.Ingredient{}
	result := db.First(ingredient, ingredientId)

	return ingredient, result.Error
}

// GetIngredients implements IngredientRepository
func (*ingredientRepository) GetIngredients(db *gorm.DB, query string) (*[]model_db.Ingredient, error) {
	ingredients := &[]model_db.Ingredient{}

	result := db.Where("name ILIKE ?", "%"+query+"%").Find(ingredients)

	return ingredients, result.Error
}

// UpdateIngredient implements IngredientRepository
func (*ingredientRepository) UpdateIngredient(db *gorm.DB, newIngredient *model_db.Ingredient) error {
	return db.Save(newIngredient).Error
}
