package controller

import "github.com/gin-gonic/gin"

type IngredientController interface {
	AddIngredient(ctx *gin.Context)
	GetIngredients(ctx *gin.Context)
	UpdateIngredient(ctx *gin.Context)
	DeleteIngredient(ctx *gin.Context)
	GetIngredientById(ctx *gin.Context)
}
