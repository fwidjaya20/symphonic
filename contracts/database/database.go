package database

import "gorm.io/gorm"

type Database interface {
	GetSession() *gorm.DB
}
