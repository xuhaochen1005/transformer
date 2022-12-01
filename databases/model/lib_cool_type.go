// Package model
//  冷却方式
package model

type CoolType struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`       // 冷却方式编号
	CoolType    string    `json:"cool_type" gorm:"column:cool_type"`     // 冷却方式
	LctCreator  int       `json:"lct_creator" gorm:"column:lct_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LctCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 冷却方式添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 冷却方式更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 冷却方式删除时间
}

func (CoolType) TableName() string {
	return "lib_cool_type"
}
