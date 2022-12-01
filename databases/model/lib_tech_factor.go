// Package model
//  工艺系数
package model

type TechFactor struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`         // 工艺系数编号
	StackAmount int       `json:"stack_amount" gorm:"column:stack_amount"` // 片数，硅钢片型号的后三位
	TechFactor  float32   `json:"tech_factor" gorm:"column:tech_factor"`   // 工艺系数
	LtfCreator  int       `json:"ltf_creator" gorm:"column:ltf_creator"`   // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LtfCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 工艺系数添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 工艺系数更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 工艺系数删除时间
}

func (TechFactor) TableName() string {
	return "lib_tech_factor"
}
