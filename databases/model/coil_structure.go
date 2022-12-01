// Package model
//线圈结构表
package model

type CoilStructure struct {
	ID                int       `json:"id" gorm:"column:id;primary_key"`                       // 编号
	StructureTypeName string    `json:"structure_type_name" gorm:"column:structure_type_name"` // 线圈结构类型名称
	CreatedAt         Timestamp `json:"created_at" gorm:"column:created_at"`                   // 添加时间
	UpdatedAt         Timestamp `json:"updated_at" gorm:"column:updated_at"`                   // 更新时间
}

func (CoilStructure) TableName() string {
	return "coil_structure"
}
