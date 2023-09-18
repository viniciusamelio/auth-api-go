package database

import "gorm.io/gorm"

type GoOrmDatabase struct {
	*gorm.DB
}
