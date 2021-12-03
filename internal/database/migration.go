package database

import (
	"gorm.io/gorm"

	"github.com/wfen/go-rest-api-course/internal/comment"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&comment.Comment{}); err != nil {
		return err
	}
	return nil
}
