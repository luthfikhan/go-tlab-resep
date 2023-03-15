package service

import (
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
)

type RecipeService interface {
	AddRecipe(recipe *model_web.RecipePayload) (*model_web.WebResponse[any], uint)
	GetRecipes(q string) (*model_web.WebResponse[any], uint)
	UpdateRecipe(recipe *model_web.RecipePayload) (*model_web.WebResponse[any], uint)
	DeleteRecipe(recipeId uint) (*model_web.WebResponse[any], uint)
	GetRecipeById(recipeId uint) (*model_web.WebResponse[any], uint)
}
