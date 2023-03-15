package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/service"
)

func NewIngredientController(ingredientService service.IngredientService) IngredientController {
	return &ingredientController{
		ingredientService: ingredientService,
	}
}

type ingredientController struct {
	ingredientService service.IngredientService
}

// GetIngredientById implements IngredientController
func (ctrl *ingredientController) GetIngredientById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}
	response, code := ctrl.ingredientService.GetIngredientById(uint(id))
	ctx.JSON(int(code), response)
}

// AddIngredient implements IngredientController
func (ctrl *ingredientController) AddIngredient(ctx *gin.Context) {
	payload := model_web.IngredientPayload{}
	if err := ctx.ShouldBindJSON(&payload); err == nil {
		response, code := ctrl.ingredientService.AddIngredient(&payload)
		ctx.JSON(int(code), response)
	} else {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: err.Error(),
			Status:  model_web.BadRequestStatus,
		})
	}
}

// DeleteIngredient implements IngredientController
func (ctrl *ingredientController) DeleteIngredient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}
	response, code := ctrl.ingredientService.DeleteIngredient(uint(id))
	ctx.JSON(int(code), response)
}

// GetIngredients implements IngredientController
func (ctrl *ingredientController) GetIngredients(ctx *gin.Context) {
	q := ctx.Query("q")

	response, code := ctrl.ingredientService.GetIngredients(q)
	ctx.JSON(int(code), response)
}

// UpdateIngredient implements IngredientController
func (ctrl *ingredientController) UpdateIngredient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}

	payload := model_web.IngredientPayload{}
	if err := ctx.ShouldBindJSON(&payload); err == nil {
		payload.ID = uint(id)
		response, code := ctrl.ingredientService.UpdateIngredient(&payload)
		ctx.JSON(int(code), response)
	} else {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: err.Error(),
			Status:  model_web.BadRequestStatus,
		})
	}
}
