package repository

import (
	"gorm.io/gorm"
)

// repository struct
type Repository struct {
	DB      *gorm.DB
}

// new repository
func New(db *gorm.DB) *Repository {
	return &Repository{db}
}



