package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/service"
)

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

type categoryController struct {
	categoryService service.CategoryService
}

// GetCategoryById implements CategoryController
func (ctrl *categoryController) GetCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}
	response, code := ctrl.categoryService.GetCategoryById(uint(id))
	ctx.JSON(int(code), response)
}

// AddCategory implements CategoryController
func (ctrl *categoryController) AddCategory(ctx *gin.Context) {
	payload := model_web.CategoryPayload{}
	if err := ctx.ShouldBindJSON(&payload); err == nil {
		response, code := ctrl.categoryService.AddCategory(&payload)
		ctx.JSON(int(code), response)
	} else {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: err.Error(),
			Status:  model_web.BadRequestStatus,
		})
	}
}

// DeleteCategory implements CategoryController
func (ctrl *categoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}
	response, code := ctrl.categoryService.DeleteCategory(uint(id))
	ctx.JSON(int(code), response)
}

// GetCategories implements CategoryController
func (ctrl *categoryController) GetCategories(ctx *gin.Context) {
	q := ctx.Query("q")

	response, code := ctrl.categoryService.GetCategories(q)
	ctx.JSON(int(code), response)
}

// UpdateCategory implements CategoryController
func (ctrl *categoryController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: "Invalid id",
			Status:  model_web.BadRequestStatus,
		})
		return
	}

	payload := model_web.CategoryPayload{}
	if err := ctx.ShouldBindJSON(&payload); err == nil {
		payload.ID = uint(id)
		response, code := ctrl.categoryService.UpdateCategory(&payload)
		ctx.JSON(int(code), response)
	} else {
		ctx.JSON(400, model_web.WebResponse[any]{
			Message: err.Error(),
			Status:  model_web.BadRequestStatus,
		})
	}
}
