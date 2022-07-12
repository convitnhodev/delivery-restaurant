package restaurantstorage

import "gorm.io/gorm"

// create struct store db
type sqlStore struct {
	db *gorm.DB
}

// declare function NewSQLStore return sqlStore
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}
