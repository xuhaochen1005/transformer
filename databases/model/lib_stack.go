// Package model
//  硅钢片
package model

type Stack struct {
	ID            int       `json:"id" gorm:"column:id;primary_key"`             // 硅钢片编号
	Type          string    `json:"type" gorm:"column:type"`                     // 硅钢片型号
	Density       float32   `json:"density" gorm:"column:density"`               // 密度(kg/cm^3)
	Price         float32   `json:"price" gorm:"column:price"`                   // 单价
	LstackCreator int       `json:"lstack_creator" gorm:"column:lstack_creator"` // 创建者
	CreatorUser   *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LstackCreator;AssociationForeignKey:ID"`
	CreatedAt     Timestamp `json:"created_at" gorm:"column:created_at"` // 硅钢片添加时间
	UpdatedAt     Timestamp `json:"updated_at" gorm:"column:updated_at"` // 硅钢片更新时间
	DeletedAt     Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 硅钢片删除时间
}

func (Stack) TableName() string {
	return "lib_stack"
}
