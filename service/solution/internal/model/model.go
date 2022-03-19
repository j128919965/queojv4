package model

type Solution struct {
	Id uint64  `gorm:"primaryKey,autoIncrement"`
	Uid uint64
	Pid int32
	Time uint64
	Nickname string `gorm:"size:50"`
	Title string `gorm:"size:50"`
	Summary string `gorm:"size:255"`
	Content string
}
