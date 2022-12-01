// Package model
// 扁铜(铝)线标称尺寸表
package model

type FlatCopperAluminumWireSizeCrossSection struct {
	ID           int     `json:"id" gorm:"column:id;primary_key"`           // 编号
	NominalSize  float32 `json:"nominal_size" gorm:"column:nominal_size"`   //标称尺寸
	CrossSection float32 `json:"cross_section" gorm:"column:cross_section"` //截面
	FilletRadius float32 `json:"fillet_radius" gorm:"column:fillet_radius"` //圆角半径
}

func (FlatCopperAluminumWireSizeCrossSection) TableName() string {
	return "flat_copper_aluminum_ wire_size_cross_section"
}
