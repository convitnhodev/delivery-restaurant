package storage

import "gorm.io/gorm"

// create struct store db
type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}
