package model_web

import "time"

type RecipeResponse struct {
	ID          uint         `json:"id,omitempty"`
	Name        string       `json:"name"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Category    string       `json:"category"`
	CategoryId  uint         `json:"category_id"`
	Ingredients []Ingredient `json:"ingredients" binding:"required"`
}

type Ingredient struct {
	ID     uint   `json:"id,omitempty" binding:"required"`
	Amount string `json:"amount" binding:"required"`
	Name   string `json:"name"`
}

type RecipePayload struct {
	ID          uint         `json:"id,omitempty"`
	Name        string       `json:"name" binding:"required"`
	CategoryID  uint         `json:"category_id" binding:"required"`
	Ingredients []Ingredient `json:"ingredients" binding:"required"`
}
