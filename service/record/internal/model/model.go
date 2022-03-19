package model

type Record struct {
	Id uint64  `gorm:"primaryKey,autoIncrement"`
	Uid uint64
	Pid int32
	Time uint64
	Language uint32
	Status uint32
	TimeUsed uint64
	SpaceUsed uint64
	Code string
}
