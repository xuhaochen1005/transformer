// Package model
// 干变低压箔绕模具
package model

type DryChangeLowvFoilWindingMold struct {
	ID           int     `json:"id" gorm:"column:id;primary_key"`           // 部门编号
	PlatHeight   float32 `json:"plat_height" gorm:"column:plat_height"`     //平台高度
	MoldSize     float32 `json:"mold_size" gorm:"column:mold_size"`         //模具尺寸
	MoldHeight   float32 `json:"mold_height" gorm:"column:mold_height"`     //模具高度
	MoldDrawings string  `json:"mold_drawings" gorm:"column:mold_drawings"` //模具图号
	MoldNumber   int     `json:"mold_number" gorm:"column:mold_number"`     //模具编号
	Number       int     `json:"number" gorm:"column:number"`               //数量
	Info         string  `json:"info" gorm:"column:info"`                   // 备注

}

func (DryChangeLowvFoilWindingMold) TableName() string {
	return "dry_change_low_voltage_foil_winding_mold"
}
