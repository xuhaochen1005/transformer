// Package model
//  主风道
package model

type MainAirDuct struct {
	ID              int       `json:"id" gorm:"column:id;primary_key"`                     // 主风道编号
	RatedHighVolMin float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"` // 额定高压最小值
	RatedHighVolMax float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"` // 额定高压最大值
	MainAirDuctMin  float32   `json:"main_air_duct_min" gorm:"column:main_air_duct_min"`   // 主风道初选最小值
	LmadCreator     int       `json:"lmad_creator" gorm:"column:lmad_creator"`             // 创建者
	CreatorUser     *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LmadCreator;AssociationForeignKey:ID"`
	CreatedAt       Timestamp `json:"created_at" gorm:"column:created_at"` // 主风道添加时间
	UpdatedAt       Timestamp `json:"updated_at" gorm:"column:updated_at"` // 主风道更新时间
	DeletedAt       Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 主风道删除时间
}

func (MainAirDuct) TableName() string {
	return "lib_main_air_duct"
}
