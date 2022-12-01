// Package model
//  扁导线规格
package model

type FlatWireSize struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`         // 扁导线规格编号
	StdLength   float32   `json:"std_length" gorm:"column:std_length"`     // 标称长（mm）
	StdWidth    float32   `json:"std_width" gorm:"column:std_width"`       // 标称宽（mm）
	SectionArea float32   `json:"section_area" gorm:"column:section_area"` // 截面积（mm^2）
	RoundCorner float32   `json:"round_corner" gorm:"column:round_corner"` // 圆角半径（mm）
	Remark      string    `json:"remark" gorm:"column:remark"`             // 备注
	LfwsCreator int       `json:"lfws_creator" gorm:"column:lfws_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LfwsCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 扁导线规格添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 扁导线规格更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 扁导线规格删除时间
}

func (FlatWireSize) TableName() string {
	return "lib_flat_wire_size"
}
