package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(dsn string) (*gorm.DB,error)  {
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		return db,err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&UserAccount{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&PeReq{})
	if err != nil {
		return nil, err
	}
	return db,nil
}