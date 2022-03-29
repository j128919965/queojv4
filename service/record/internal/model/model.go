package model

type Record struct {
	Id uint64  `gorm:"primaryKey,autoIncrement" json:"id"`
	Uid uint64 `json:"uid"`
	Pid int32 `json:"pid"`
	Time uint64 `json:"time"`
	Language uint32 `json:"language"`
	Status uint32 `json:"status"`
	TimeUsed uint64 `json:"time_used"`
	SpaceUsed uint64 `json:"space_used"`
	Code string `json:"code"`
}

type UserSuccessStatistic struct {
	Uid uint64 `gorm:"primaryKey"`
	Easy int32 `gorm:"default:0"`
	Medium int32 `gorm:"default:0"`
	Hard int32 `gorm:"default:0"`
}