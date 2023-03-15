package model_db

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null" json:"name"`
}

type Ingredient struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null" json:"name"`
	Amount    string         `gorm:"not null" json:"amount"`
}

type Recipe struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name       string         `gorm:"not null" json:"name"`
	CategoryID uint           `gorm:"not null" json:"category_id"`
	Category   string         `json:"category"`
}

type RecipeIngredient struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	RecipeID     uint           `gorm:"not null" json:"recipe_id"`
	Name         string         `gorm:"<-:false" json:"name"`
	IngredientID uint           `gorm:"not null" json:"ingredient_id"`
	Amount       string         `gorm:"not null" json:"amount"`
	Recipe       Recipe         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"recipe"`
	Ingredient   Ingredient     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"ingredient"`
}
