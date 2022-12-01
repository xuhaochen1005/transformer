// Package model 部门信息数据库模型
package model

type Department struct {
	ID        int       `json:"id" gorm:"column:id;primary_key"`                             // 部门编号
	Name      string    `json:"name" validate:"required,min=4,max=24" gorm:"column:name"`    // 部门名称
	Status    int       `json:"status" validate:"required,gte=1,lte=2" gorm:"column:status"` // 部门状态 1:正常 2:停用等
	CreatedAt Timestamp `json:"created_at" gorm:"column:created_at"`                         // 部门信息添加时间
	UpdatedAt Timestamp `json:"updated_at" gorm:"column:updated_at"`                         // 部门信息更新时间
}

func (Department) TableName() string {
	return "department"
}
