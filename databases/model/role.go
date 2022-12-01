package model

type Role struct {
	ID          int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Name        string `json:"name" gorm:"column:name"`       //角色名
	CnName      string `json:"cn_name" gorm:"column:cn_name"` //角色中文名
	Description string `json:"description" gorm:"column:description"`
}

func (Role) TableName() string {
	return "role"
}
