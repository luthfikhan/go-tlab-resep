package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	AddCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	GetCategoryById(ctx *gin.Context)
}
