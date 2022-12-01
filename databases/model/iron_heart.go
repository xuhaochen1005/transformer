// Package model
// 铁心表
package model

type IronHeart struct {
	ID        int       `json:"id" gorm:"column:id;primary_key"`     // 部门编号
	Name      string    `json:"name" gorm:"column:name"`             // 部门名称
	Status    int       `json:"status" gorm:"column:status"`         // 部门状态
	CreatedAt Timestamp `json:"created_at" gorm:"column:created_at"` // 部门信息添加时间
	UpdatedAt Timestamp `json:"updated_at" gorm:"column:updated_at"` // 部门信息更新时间
}

func (IronHeart) TableName() string {
	return "iron_heart"
}
