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

	err = db.AutoMigrate(&Problem{})
	err = db.AutoMigrate(&ProblemIO{})
	err = db.AutoMigrate(&ProblemTag{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
