// Package model
//  线圈导线类型
package model

type CoilWireType struct {
	ID               int       `json:"id" gorm:"column:id;primary_key"`                       // 线圈导线类型编号
	CoilWireType     string    `json:"coil_wire_type" gorm:"column:coil_wire_type"`           // 线圈导线类型
	CoilWireTypeSign string    `json:"coil_wire_type_sign" gorm:"column:coil_wire_type_sign"` // 线圈导线类型字母代号
	LcwtCreator      int       `json:"lcwt_creator" gorm:"column:lcwt_creator"`               // 创建者
	CreatorUser      *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LcwtCreator;AssociationForeignKey:ID"`
	CreatedAt        Timestamp `json:"created_at" gorm:"column:created_at"` // 线圈导线类型添加时间
	UpdatedAt        Timestamp `json:"updated_at" gorm:"column:updated_at"` // 线圈导线类型更新时间
	DeletedAt        Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 线圈导线类型删除时间
}

func (CoilWireType) TableName() string {
	return "lib_coil_wire_type"
}
