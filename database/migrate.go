package database

import (
	"gorm.io/gorm"
)

// Migrate :
func Migrate(db *gorm.DB) {
	// Master data dan polymorphiz table
	db.AutoMigrate(&User{})
}
