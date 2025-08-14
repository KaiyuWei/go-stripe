package models

import (
	"database/sql"
)

// type for database connection values
type DBModel struct {
	DB *sql.DB
}

// wapper for all models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Widget struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	InventoryLevel int    `json:"inventory_level"`
	Price          int    `json:"price"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
}
