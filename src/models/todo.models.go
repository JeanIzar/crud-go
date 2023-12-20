package models

import (
	"gorm.io/gorm"
)

// Defines todo table for database communications
type Todo struct {
	gorm.Model
	Name        string
	Description string
}