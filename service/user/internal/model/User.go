package model

import (
	"database/sql"
)

type User struct {
	ID       uint64         `gorm:"primarykey"`
	Role     uint8          `gorm:"default:1"`
	Password sql.NullString `gorm:"size:255"`
	Email    string         `gorm:"size:200;not null;unique"`
}

type UserInfo struct {
	ID           uint64
	Nickname     sql.NullString `gorm:"size:45"`
	Introduction sql.NullString `gorm:"size:255"`
	Github       sql.NullString `gorm:"size:255"`
	Website      sql.NullString `gorm:"size:255"`
	Wechat       sql.NullString `gorm:"size:255"`
	Favicon      sql.NullString `gorm:"size:255"`
	Phone        sql.NullString `gorm:"size:45"`
	Email        string         `gorm:"size:200;not null;unique"`
}

type UserAccount struct {
	ID    uint64 `gorm:"primarykey"`
	Coins int32 `gorm:"default:0;not null"`
	Point int32 `gorm:"default:0;not null;index:idx_point"`
}

type PeReq struct {
	ID uint64 `gorm:"primarykey,autoIncrement"`
	UserId uint64
	Role uint32
	Reason string
	Approval int32 `gorm:"default:0"`
}