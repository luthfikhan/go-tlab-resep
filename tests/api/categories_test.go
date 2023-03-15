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

func TestGetCategories(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/categories", nil)
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
	var response model_web.WebResponse[[]model_web.CategoryResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.SuccessStatus)
}

func TestGetCategoriesWithQuery(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/categories?q=1", nil)
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
	var response model_web.WebResponse[[]model_web.CategoryResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Status, model_web.SuccessStatus)
	assert.Equal(t, len(response.Data), 1)
}

func TestGetCategoryById(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/categories/1", nil)
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
	var response model_web.WebResponse[model_web.CategoryResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.Name, "Category 1")
}

func TestDeleteCategory(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/categories/1", nil)
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

func TestCreateCategory(t *testing.T) {
	payload := []byte(`{"name":"test"}`)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(payload))
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

func TestCreateCategoryWithoutPayloadName(t *testing.T) {
	payload := []byte(`{"name":""}`)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(payload))
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

func TestUpdateCategory(t *testing.T) {
	payload := []byte(`{"name":"test"}`)
	// Create a new HTTP request
	req, err := http.NewRequest("PUT", "/api/categories/1", bytes.NewBuffer(payload))
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

func TestUpdateCategoryNotFound(t *testing.T) {
	payload := []byte(`{"name":"test"}`)
	// Create a new HTTP request
	req, err := http.NewRequest("PUT", "/api/categories/3", bytes.NewBuffer(payload))
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
