package model_web

import "time"

type CategoryPayload struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
