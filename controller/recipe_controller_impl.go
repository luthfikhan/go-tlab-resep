package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/service"
)

func NewRecipeController(recipeService service.RecipeService) RecipeController {
	return &recipeController{
		recipeService: recipeService,
	}
}

type recipeController struct {
	recipeService service.RecipeService
}

// AddRecipe implements RecipeController
func (ctrl *recipeController) AddRecipe(ctx *gin.Context) {
	payload := model_web.RecipePayload{}
	if err := ctx.ShouldBindJSON(&payload); err == nil {
		response, code := ctrl.recipeService.AddRecipe(&payload)
		ctx.JSON(int(code), response)
	} else {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: err.Error(),
			Status:  model_web.BadRequestStatus,
		})
	}
}

// GetRecipeById implements RecipeController
func (ctrl *recipeController) GetRecipeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}
	response, code := ctrl.recipeService.GetRecipeById(uint(id))
	ctx.JSON(int(code), response)
}

// DeleteRecipe implements RecipeController
func (ctrl *recipeController) DeleteRecipe(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}
	response, code := ctrl.recipeService.DeleteRecipe(uint(id))
	ctx.JSON(int(code), response)
}

// GetRecipes implements RecipeController
func (ctrl *recipeController) GetRecipes(ctx *gin.Context) {
	q := ctx.Query("q")

	response, code := ctrl.recipeService.GetRecipes(q)
	ctx.JSON(int(code), response)
}

// UpdateRecipe implements RecipeController
func (ctrl *recipeController) UpdateRecipe(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}

	payload := model_web.RecipePayload{}
	if err := ctx.ShouldBindJSON(&payload); err == nil {
		payload.ID = uint(id)
		response, code := ctrl.recipeService.UpdateRecipe(&payload)
		ctx.JSON(int(code), response)
	} else {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: err.Error(),
			Status:  model_web.BadRequestStatus,
		})
	}
}
