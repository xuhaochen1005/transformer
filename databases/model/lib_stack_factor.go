// Package model
//  迭片系数
package model

type StackFactor struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`         // 迭片系数编号
	StackThick  float32   `json:"stack_thick" gorm:"column:stack_thick"`   // 片厚
	StackFactor float32   `json:"stack_factor" gorm:"column:stack_factor"` // 迭片系数
	LsfCreator  int       `json:"lsf_creator" gorm:"column:lsf_creator"`   // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LsfCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 迭片系数添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 迭片系数更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 迭片系数删除时间
}

func (StackFactor) TableName() string {
	return "lib_stack_factor"
}
