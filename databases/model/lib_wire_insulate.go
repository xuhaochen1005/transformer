// Package model
//  导线绝缘厚度表
package model

type WireInsulate struct {
	ID               int       `json:"id" gorm:"column:id;primary_key"`                   // 导线绝缘厚度表编号
	WireID           int       `json:"wire_id" gorm:"column:wire_id"`                     // 导线编号
	WidthMin         float32   `json:"width_min" gorm:"column:width_min"`                 // 线厚区间下界(mm)
	WidthMax         float32   `json:"width_max" gorm:"column:width_max"`                 // 线厚区间上界(mm)
	AxialInsulation  float32   `json:"axial_insulation" gorm:"column:axial_insulation"`   // 轴向绝缘厚(mm)
	RadialInsulation float32   `json:"radial_insulation" gorm:"column:radial_insulation"` // 辐向绝缘厚(mm)
	LwiCreator       int       `json:"lwi_creator" gorm:"column:lwi_creator"`             // 创建者
	CreatorUser      *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LwiCreator;AssociationForeignKey:ID"`
	CreatedAt        Timestamp `json:"created_at" gorm:"column:created_at"` // 导线绝缘厚度表添加时间
	UpdatedAt        Timestamp `json:"updated_at" gorm:"column:updated_at"` // 导线绝缘厚度表更新时间
	DeletedAt        Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 导线绝缘厚度表删除时间
}

func (WireInsulate) TableName() string {
	return "lib_wire_insulate"
}
