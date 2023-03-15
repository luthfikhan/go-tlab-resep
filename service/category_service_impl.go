package service

import (
	"net/http"

	"github.com/luthfikhan/go-tlab-resep/helper"
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"gorm.io/gorm"
)

func NewCategoryService(db *gorm.DB, categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		db:                 db,
		categoryRepository: categoryRepository,
	}
}

type categoryService struct {
	db                 *gorm.DB
	categoryRepository repository.CategoryRepository
}

// GetCategoryById implements CategoryService
func (service *categoryService) GetCategoryById(categoryId uint) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	categoryDB, err := service.categoryRepository.GetCategoryById(tx, categoryId)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Category Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	category, err := helper.TypeConverter[model_web.CategoryResponse](categoryDB)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
		Data:   category,
	}, http.StatusOK
}

// AddCategory implements CategoryService
func (service *categoryService) AddCategory(category *model_web.CategoryPayload) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	err := service.categoryRepository.AddCategory(tx, &model_db.Category{Name: category.Name})
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusCreated
}

// DeleteCategory implements CategoryService
func (service *categoryService) DeleteCategory(categoryId uint) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	_, err := service.categoryRepository.GetCategoryById(tx, categoryId)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Category Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	err = service.categoryRepository.DeleteCategory(tx, categoryId)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusOK
}

// GetCategories implements CategoryService
func (service *categoryService) GetCategories(q string) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	categoriesDB, err := service.categoryRepository.GetCategories(tx, q)
	helper.PanifIfError(err)
	categories, err := helper.TypeConverter[[]model_web.CategoryResponse](categoriesDB)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
		Data:   categories,
	}, http.StatusOK
}

// UpdateCategory implements CategoryService
func (service *categoryService) UpdateCategory(category *model_web.CategoryPayload) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	c, err := service.categoryRepository.GetCategoryById(tx, category.ID)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Category Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	c = helper.UpdateStructFromAnother(*c, *category)
	err = service.categoryRepository.UpdateCategory(tx, c)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusOK
}
