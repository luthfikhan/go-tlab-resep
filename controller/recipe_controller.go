package controller

import "github.com/gin-gonic/gin"

type RecipeController interface {
	AddRecipe(ctx *gin.Context)
	GetRecipes(ctx *gin.Context)
	UpdateRecipe(ctx *gin.Context)
	DeleteRecipe(ctx *gin.Context)
	GetRecipeById(ctx *gin.Context)
}
