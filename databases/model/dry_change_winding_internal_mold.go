// Package model
// 干变低压内膜模台账
package model

type DryChangeWindingInternalrMold struct {
	ID           int     `json:"id" gorm:"column:id;primary_key"`           // 编号
	MoldDiameter string  `json:"mold_diameter" gorm:"column:mold_diameter"` // 模具直径
	MoldLength   float32 `json:"mold_length" gorm:"column:mold_length"`     // 模具长
	MoldWidth    float32 `json:"mold_width" gorm:"column:mold_width"`       // 模具宽
	Number       int     `json:"number" gorm:"column:number"`               //数量
	MoldDrawings string  `json:"mold_drawings" gorm:"column:mold_drawings"` //模具图号
	MoldNumber   int     `json:"mold_number" gorm:"column:mold_number"`     //模具编号
	Info         string  `json:"info" gorm:"column:info"`                   //备注
	MoleSize     string  `json:"mold_size" gorm:"column:mold_size"`         //模具规格

}

func (DryChangeWindingInternalrMold) TableName() string {
	return "dry_change_winding_internal_mold"
}
