package service

import (
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
)

type IngredientService interface {
	AddIngredient(ingredient *model_web.IngredientPayload) (*model_web.WebResponse[any], uint)
	GetIngredients(q string) (*model_web.WebResponse[any], uint)
	UpdateIngredient(ingredient *model_web.IngredientPayload) (*model_web.WebResponse[any], uint)
	DeleteIngredient(ingredientId uint) (*model_web.WebResponse[any], uint)
	GetIngredientById(ingredientId uint) (*model_web.WebResponse[any], uint)
}
