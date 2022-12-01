// Package model
//  电阻率
package model

type Resistivity struct {
	ID             int       `json:"id" gorm:"column:id;primary_key"`               // 电阻率编号
	WireMaterial   string    `json:"wire_material" gorm:"column:wire_material"`     // 导线材质
	Temp           int       `json:"temp" gorm:"column:temp"`                       // 温度（℃）
	Resistivity    float32   `json:"resistivity" gorm:"column:resistivity"`         // 电阻率(Ω·m)
	LresistCreator int       `json:"lresist_creator" gorm:"column:lresist_creator"` // 创建者
	CreatorUser    *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LresistCreator;AssociationForeignKey:ID"`
	CreatedAt      Timestamp `json:"created_at" gorm:"column:created_at"` // 电阻率添加时间
	UpdatedAt      Timestamp `json:"updated_at" gorm:"column:updated_at"` // 电阻率更新时间
	DeletedAt      Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 电阻率删除时间
}

func (Resistivity) TableName() string {
	return "lib_resistivity"
}
