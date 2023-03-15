package service

import (
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
)

type CategoryService interface {
	AddCategory(category *model_web.CategoryPayload) (*model_web.WebResponse[any], uint)
	GetCategories(q string) (*model_web.WebResponse[any], uint)
	UpdateCategory(category *model_web.CategoryPayload) (*model_web.WebResponse[any], uint)
	DeleteCategory(categoryId uint) (*model_web.WebResponse[any], uint)
	GetCategoryById(categoryId uint) (*model_web.WebResponse[any], uint)
}
