// Package model
//  电流密度
package model

type CurrentDensity struct {
	ID                int       `json:"id" gorm:"column:id;primary_key"`                       // 电流密度编号
	WindingMaterial   string    `json:"winding_material" gorm:"column:winding_material"`       // 导线材料，铜/铝
	CurrentDensityMin float32   `json:"current_density_min" gorm:"column:current_density_min"` // 电流密度下界
	CurrentDensityMax float32   `json:"current_density_max" gorm:"column:current_density_max"` // 电流密度上界
	LcdCreator        int       `json:"lcd_creator" gorm:"column:lcd_creator"`                 // 创建者
	CreatorUser       *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LcdCreator;AssociationForeignKey:ID"`
	CreatedAt         Timestamp `json:"created_at" gorm:"column:created_at"` // 电流密度添加时间
	UpdatedAt         Timestamp `json:"updated_at" gorm:"column:updated_at"` // 电流密度更新时间
	DeletedAt         Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 电流密度删除时间
}

func (CurrentDensity) TableName() string {
	return "lib_current_density"
}
