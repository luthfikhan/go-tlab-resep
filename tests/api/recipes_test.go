package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/luthfikhan/go-tlab-resep/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetRecipes(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/recipes", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response model_web.WebResponse[[]model_web.RecipeResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.SuccessStatus)
}

func TestGetRecipeById(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/recipes/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response model_web.WebResponse[model_web.RecipeResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.Name, "nasi Goreng")
}

func TestDeleteRecipe(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/recipes/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response model_web.WebResponse[any]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.SuccessStatus)
}

func TestCreateRecipe(t *testing.T) {
	payload := []byte(`{"category_id":1,"ingredients":[{"amount":"1 piring","id":1}],"name":"Name"}`)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/recipes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 201 OK
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	var response model_web.WebResponse[any]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.SuccessStatus)
}

func TestCreateRecipeWithoutPayloadName(t *testing.T) {
	payload := []byte(`{"category_id":1,"ingredients":[{"amount":"1 piring","id":1}],"name":""}`)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/recipes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 400 OK
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Check the response body
	var response model_web.WebResponse[any]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.BadRequestStatus)
}

func TestUpdateRecipe(t *testing.T) {
	payload := []byte(`{"category_id":1,"ingredients":[{"amount":"1 piring","id":1}],"name":"Name"}`)
	// Create a new HTTP request
	req, err := http.NewRequest("PUT", "/api/recipes/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response model_web.WebResponse[any]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.SuccessStatus)
}

func TestUpdateRecipeNotFound(t *testing.T) {
	payload := []byte(`{"category_id":1,"ingredients":[{"amount":"1 piring","id":1}],"name":"Name"}`)
	// Create a new HTTP request
	req, err := http.NewRequest("PUT", "/api/recipes/3", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Create a new router itesting
	r := tests.SetupRouter()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code is 404 OK
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Check the response body
	var response model_web.WebResponse[any]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.NotFoundStatus)
}
