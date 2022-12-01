// Package model
//  导线表
package model

type Wire struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`           // 导线表编号
	WireType     string    `json:"wire_type" gorm:"column:wire_type"`         // 导线型号
	WireMaterial string    `json:"wire_material" gorm:"column:wire_material"` // 导线材质(铜/铝/铜箔/铝箔)
	WireShape    string    `json:"wire_shape" gorm:"column:wire_shape"`       // 导线形状(圆/扁)
	WirePrice    float32   `json:"wire_price" gorm:"column:wire_price"`       // 导线价格
	WireDensity  float32   `json:"wire_density" gorm:"column:wire_density"`   // 导线密度(kg/m^3)
	LwireCreator int       `json:"lwire_creator" gorm:"column:lwire_creator"` // 创建者
	CreatorUser  *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LwireCreator;AssociationForeignKey:ID"`
	CreatedAt    Timestamp `json:"created_at" gorm:"column:created_at"` // 导线表添加时间
	UpdatedAt    Timestamp `json:"updated_at" gorm:"column:updated_at"` // 导线表更新时间
	DeletedAt    Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 导线表删除时间
}

func (Wire) TableName() string {
	return "lib_wire"
}
