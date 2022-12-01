// Package model
//  圆铜（铝）线规格
package model

type RoundWireSize struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`         // 圆铜（铝）线规格编号
	StdDiameter float32   `json:"std_diameter" gorm:"column:std_diameter"` // 标称直径（mm）
	SectionArea float32   `json:"section_area" gorm:"column:section_area"` // 截面积（mm^2）
	MaxDiameter float32   `json:"max_diameter" gorm:"column:max_diameter"` // 最大外径（mm）
	LrwsCreator int       `json:"lrws_creator" gorm:"column:lrws_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LrwsCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 圆铜（铝）线规格添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 圆铜（铝）线规格更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 圆铜（铝）线规格删除时间
}

func (RoundWireSize) TableName() string {
	return "lib_round_wire_size"
}
