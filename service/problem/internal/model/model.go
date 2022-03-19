package model

type Problem struct {
	Id          int32 `gorm:"primaryKey,autoIncrement"`
	Name        string `gorm:"size:50"`
	Point       int32
	Level       int32
	Description string `gorm:"type:text"`
	TimeLimit   uint64
	SpaceLimit  uint64
}

type ProblemTag struct {
	Id  int32 `gorm:"primaryKey,autoIncrement"`
	Pid int32 `gorm:"index:idx_pid"`
	Tid int32 `gorm:"index:idx_tid"`
}

type ProblemIO struct {
	Id     int32  `gorm:"primaryKey,autoIncrement"`
	Pid    int32  `gorm:"index:idx_pid,unique"`
	InTxt  string `gorm:"type:text"`
	OutTxt string `gorm:"type:text"`
}
