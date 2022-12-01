package model

type Boss struct {
	ID       int    `json:"id" gorm:"column:id;primary_key"`
	BossSpec string `json:"boss_spec" gorm:"column:boss_spec"`
	C        int    `json:"c" gorm:"column:c"`
}

func (Boss) TableName() string {
	return "lib_boss"
}
