package lib

import "gorm.io/gorm"

// DatabaseInstance holds all database instances and it's methods.
type DatabaseInstance struct {
	OrderDB *gorm.DB
}

// NewDatabaseInstance returns constructed database instance.
func NewDatabaseInstance() *DatabaseInstance {
	return &DatabaseInstance{}
}
