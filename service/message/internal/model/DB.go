package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(&Message{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
