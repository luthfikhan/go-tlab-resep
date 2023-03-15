package service

import (
	"net/http"

	"github.com/luthfikhan/go-tlab-resep/helper"
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"gorm.io/gorm"
)

func NewIngredientService(db *gorm.DB, ingredientRepository repository.IngredientRepository) IngredientService {
	return &ingredientService{
		db:                   db,
		ingredientRepository: ingredientRepository,
	}
}

type ingredientService struct {
	db                   *gorm.DB
	ingredientRepository repository.IngredientRepository
}

func (service *ingredientService) GetIngredientById(ingredientId uint) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	_, err := service.ingredientRepository.GetIngredientById(tx, ingredientId)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Ingredient Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	ingredientDB, err := service.ingredientRepository.GetIngredientById(tx, ingredientId)
	helper.PanifIfError(err)
	ingredient, err := helper.TypeConverter[model_web.IngredientResponse](ingredientDB)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
		Data:   ingredient,
	}, http.StatusOK
}

func (service *ingredientService) AddIngredient(ingredient *model_web.IngredientPayload) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	err := service.ingredientRepository.AddIngredient(tx, &model_db.Ingredient{Name: ingredient.Name})
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusCreated
}

func (service *ingredientService) DeleteIngredient(ingredientId uint) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	_, err := service.ingredientRepository.GetIngredientById(tx, ingredientId)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Ingredient Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	err = service.ingredientRepository.DeleteIngredient(tx, ingredientId)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusOK
}

func (service *ingredientService) GetIngredients(q string) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	ingredientDB, err := service.ingredientRepository.GetIngredients(tx, q)
	helper.PanifIfError(err)
	ingredients, err := helper.TypeConverter[[]model_web.IngredientResponse](ingredientDB)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
		Data:   ingredients,
	}, http.StatusOK
}

func (service *ingredientService) UpdateIngredient(ingredient *model_web.IngredientPayload) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	c, err := service.ingredientRepository.GetIngredientById(tx, ingredient.ID)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Ingredient Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	c = helper.UpdateStructFromAnother(*c, *ingredient)
	err = service.ingredientRepository.UpdateIngredient(tx, c)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusOK
}
