// Package model
//  线圈形状
package model

type CoilShape struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`       // 线圈形状编号
	CoilShape   string    `json:"coil_shape" gorm:"column:coil_shape"`   // 线圈形状
	LcsCreator  int       `json:"lcs_creator" gorm:"column:lcs_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LcsCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 线圈形状添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 线圈形状更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 线圈形状删除时间
}

func (CoilShape) TableName() string {
	return "lib_coil_shape"
}
