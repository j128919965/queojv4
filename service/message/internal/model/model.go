package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id        uint64         `gorm:"primaryKey,autoIncrement"`
	Receiver  uint64         `gorm:"index:idx_msg"`
	Read      bool           `gorm:"index:idx_msg,default:0"`
	Type      byte           `gorm:"default:1"`
	Title     string         `gorm:"size:255"`
	Content   *string        `gorm:"size:512"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_msg"`
	CreatedAt time.Time
}
