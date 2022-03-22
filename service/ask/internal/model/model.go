package model

type Ask struct {
	Id       uint64 `gorm:"primaryKey,autoIncrement"`
	Uid      uint64
	Time     int64
	Nickname string `gorm:"size:50"`
	Title    string `gorm:"size:50"`
	Content  string `gorm:"size:511"`
}

type Reply struct {
	Id       uint64 `gorm:"primaryKey,autoIncrement"`
	Uid      uint64
	AskId      uint64
	Time     int64
	Nickname string `gorm:"size:50"`
	Content  string `gorm:"size:511"`
	IsTeacher bool
}