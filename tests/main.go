package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luthfikhan/go-tlab-resep/controller"
	"github.com/luthfikhan/go-tlab-resep/middleware"
	model_db "github.com/luthfikhan/go-tlab-resep/models/db"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/service"
	"github.com/luthfikhan/go-tlab-resep/tests/mocks/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	dsn := "postgres://user:password@localhost:5432/resep?sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn))

	db.AutoMigrate(
		&model_db.Category{},
		&model_db.Ingredient{},
		&model_db.Recipe{},
		&model_db.RecipeIngredient{},
	)

	router := gin.New()
	router.Use(middleware.Log, middleware.Recover())
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, model_web.NotFound)
	})

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository)
	categoryCtrl := controller.NewCategoryController(categoryService)

	categoryRouter := router.Group("/api/categories")
	categoryRouter.POST("", categoryCtrl.AddCategory)
	categoryRouter.GET("", categoryCtrl.GetCategories)
	categoryRouter.PUT("/:id", categoryCtrl.UpdateCategory)
	categoryRouter.GET("/:id", categoryCtrl.GetCategoryById)
	categoryRouter.DELETE("/:id", categoryCtrl.DeleteCategory)

	ingredientRepository := repository.NewIngredientRepository()
	ingredientService := service.NewIngredientService(db, ingredientRepository)
	ingredientCtrl := controller.NewIngredientController(ingredientService)

	ingredientRouter := router.Group("/api/ingredients")
	ingredientRouter.POST("", ingredientCtrl.AddIngredient)
	ingredientRouter.GET("", ingredientCtrl.GetIngredients)
	ingredientRouter.PUT("/:id", ingredientCtrl.UpdateIngredient)
	ingredientRouter.DELETE("/:id", ingredientCtrl.DeleteIngredient)
	ingredientRouter.GET("/:id", ingredientCtrl.GetIngredientById)

	recipeRepository := repository.NewRecipeRepository()
	recipeIngredientRepository := repository.NewRecipeIngredientRepository()
	recipeService := service.NewRecipeService(db, recipeRepository, categoryRepository, ingredientRepository, recipeIngredientRepository)
	recipeCtrl := controller.NewRecipeController(recipeService)

	recipeRouter := router.Group("/api/recipes")
	recipeRouter.POST("", recipeCtrl.AddRecipe)
	recipeRouter.GET("", recipeCtrl.GetRecipes)
	recipeRouter.PUT("/:id", recipeCtrl.UpdateRecipe)
	recipeRouter.DELETE("/:id", recipeCtrl.DeleteRecipe)
	recipeRouter.GET("/:id", recipeCtrl.GetRecipeById)

	return router
}
