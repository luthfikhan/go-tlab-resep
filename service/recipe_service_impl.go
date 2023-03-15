package service

import (
	"net/http"
	"strconv"

	"github.com/luthfikhan/go-tlab-resep/helper"
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewRecipeService(
	db *gorm.DB,
	recipeRepository repository.RecipeRepository,
	catRepository repository.CategoryRepository,
	ingredientRepository repository.IngredientRepository,
	recipeIngredientRepository repository.RecipeIngredientRepository,
) RecipeService {
	return &recipeService{
		db:                         db,
		recipeRepository:           recipeRepository,
		categoryRepository:         catRepository,
		ingredientRepository:       ingredientRepository,
		recipeIngredientRepository: recipeIngredientRepository,
	}
}

type recipeService struct {
	db                         *gorm.DB
	recipeRepository           repository.RecipeRepository
	categoryRepository         repository.CategoryRepository
	ingredientRepository       repository.IngredientRepository
	recipeIngredientRepository repository.RecipeIngredientRepository
}

func (service *recipeService) AddRecipe(recipe *model_web.RecipePayload) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	_, err := service.categoryRepository.GetCategoryById(tx, uint(recipe.CategoryID))
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Category Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	for i := 0; i < len(recipe.Ingredients); i++ {
		ingredient := recipe.Ingredients[i]

		_, err = service.ingredientRepository.GetIngredientById(tx, ingredient.ID)
		if err == gorm.ErrRecordNotFound {
			return &model_web.WebResponse[any]{
				Status:  model_web.NotFoundStatus,
				Message: "Ingredient (" + strconv.Itoa(int(ingredient.ID)) + ") Not Found",
			}, http.StatusNotFound
		}
	}
	helper.PanifIfError(err)

	dbRecipe := &model_db.Recipe{
		Name:       recipe.Name,
		CategoryID: uint(recipe.CategoryID),
	}
	err = service.recipeRepository.AddRecipe(tx, dbRecipe)
	helper.PanifIfError(err)

	for i := 0; i < len(recipe.Ingredients); i++ {
		ingredient := recipe.Ingredients[i]

		dbRecipeIngredient := model_db.RecipeIngredient{
			RecipeID:     dbRecipe.ID,
			Amount:       ingredient.Amount,
			IngredientID: ingredient.ID,
		}
		err = service.recipeIngredientRepository.AddRecipeIngredient(tx, &dbRecipeIngredient)
		helper.PanifIfError(err)
	}

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusCreated
}

func (service *recipeService) DeleteRecipe(recipeId uint) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	_, err := service.recipeRepository.GetRecipeById(tx, recipeId)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Recipe Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	err = service.recipeRepository.DeleteRecipe(tx, recipeId)
	helper.PanifIfError(err)

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusOK
}

func (service *recipeService) GetRecipeById(recipeId uint) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	recipeDB, err := service.recipeRepository.GetRecipeById(tx, recipeId)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Recipe Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)
	recipe, err := helper.TypeConverter[model_web.RecipeResponse](recipeDB)
	helper.PanifIfError(err)

	ingredients, err := service.recipeIngredientRepository.GetRecipeIngredientsByRecipeId(tx, recipe.ID)
	helper.PanifIfError(err)

	for _, v := range *ingredients {
		logrus.Info(v)
		recipe.Ingredients = append(recipe.Ingredients, model_web.Ingredient{
			Amount: v.Amount,
			ID:     v.ID,
			Name:   v.Name,
		})
	}

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
		Data:   recipe,
	}, http.StatusOK
}

func (service *recipeService) GetRecipes(q string) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	recipeDB, err := service.recipeRepository.GetRecipes(tx, q)
	helper.PanifIfError(err)

	recipes, err := helper.TypeConverter[[]model_web.RecipeResponse](recipeDB)
	helper.PanifIfError(err)

	copyRecipes := *recipes
	for i := 0; i < len(copyRecipes); i++ {
		ingredients, err := service.recipeIngredientRepository.GetRecipeIngredientsByRecipeId(tx, copyRecipes[i].ID)
		helper.PanifIfError(err)

		for _, v := range *ingredients {
			logrus.Info(v)
			copyRecipes[i].Ingredients = append(copyRecipes[i].Ingredients, model_web.Ingredient{
				Amount: v.Amount,
				ID:     v.ID,
				Name:   v.Name,
			})
		}
	}
	recipes = &copyRecipes

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
		Data:   recipes,
	}, http.StatusOK
}

func (service *recipeService) UpdateRecipe(recipe *model_web.RecipePayload) (*model_web.WebResponse[any], uint) {
	tx := service.db.Begin()
	defer helper.CheckErrorToCommitOrRollback(tx)

	recipeDb, err := service.recipeRepository.GetRecipeById(tx, recipe.ID)
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Recipe Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	_, err = service.categoryRepository.GetCategoryById(tx, uint(recipe.CategoryID))
	if err == gorm.ErrRecordNotFound {
		return &model_web.WebResponse[any]{
			Status:  model_web.NotFoundStatus,
			Message: "Category Not Found",
		}, http.StatusNotFound
	}
	helper.PanifIfError(err)

	for i := 0; i < len(recipe.Ingredients); i++ {
		ingredient := recipe.Ingredients[i]

		_, err = service.ingredientRepository.GetIngredientById(tx, ingredient.ID)
		if err == gorm.ErrRecordNotFound {
			return &model_web.WebResponse[any]{
				Status:  model_web.NotFoundStatus,
				Message: "Ingredient (" + strconv.Itoa(int(ingredient.ID)) + ") Not Found",
			}, http.StatusNotFound
		}
	}
	helper.PanifIfError(err)

	recipeDb = helper.UpdateStructFromAnother(*recipeDb, *recipe)

	err = service.recipeRepository.UpdateRecipe(tx, recipeDb)
	helper.PanifIfError(err)

	currentIngredient, err := service.recipeIngredientRepository.GetRecipeIngredientsByRecipeId(tx, recipeDb.ID)
	helper.PanifIfError(err)
	var ingredientToDelete []model_web.Ingredient
	var ingredientToAdd []model_web.Ingredient

	for _, ingredient := range *currentIngredient {
		var notSameLen int
		for _, v := range recipe.Ingredients {
			if ingredient.ID != v.ID {
				notSameLen++

				if notSameLen == len(recipe.Ingredients) {
					ingredientToDelete = append(ingredientToDelete, model_web.Ingredient{
						ID:     ingredient.ID,
						Amount: ingredient.Amount,
					})
				}
			}
		}
	}

	for _, ingredient := range recipe.Ingredients {
		var notSameLen int
		for _, v := range *currentIngredient {
			if ingredient.ID != v.ID {
				notSameLen++

				if notSameLen == len(*currentIngredient) {
					ingredientToAdd = append(ingredientToAdd, model_web.Ingredient{
						ID:     ingredient.ID,
						Amount: ingredient.Amount,
					})
				}
			}
		}
	}

	for _, ingredient := range ingredientToAdd {
		dbRecipeIngredient := model_db.RecipeIngredient{
			RecipeID:     recipeDb.ID,
			Amount:       ingredient.Amount,
			IngredientID: ingredient.ID,
		}
		err = service.recipeIngredientRepository.AddRecipeIngredient(tx, &dbRecipeIngredient)
		helper.PanifIfError(err)
	}

	for _, ingredient := range ingredientToDelete {
		err = service.recipeIngredientRepository.DeleteRecipeIngredient(tx, ingredient.ID)
		helper.PanifIfError(err)
	}

	return &model_web.WebResponse[any]{
		Status: model_web.SuccessStatus,
	}, http.StatusOK
}
