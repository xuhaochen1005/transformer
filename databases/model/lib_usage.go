// Package model
//  变压器用途
package model

type Usage struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`       // 变压器用途编号
	Usage       string    `json:"usage" gorm:"column:usage"`             // 变压器用途
	UsageSign   string    `json:"usage_sign" gorm:"column:usage_sign"`   // 变压器用途字母代号
	LusCreator  int       `json:"lus_creator" gorm:"column:lus_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LusCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 变压器用途添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 变压器用途更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 变压器用途删除时间
}

func (Usage) TableName() string {
	return "lib_usage"
}
