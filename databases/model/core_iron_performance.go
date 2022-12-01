// Package model
//铁心性能表
package model

type CoreIronPerformance struct {
	ID                    int     `json:"id" gorm:"column:id;primary_key"`                             //
	MagneticDensity       float32 `json:"magnetic_density" gorm:"column:magnetic_density"`             // 铁心磁密
	IronLoss              float32 `json:"iron_loss" gorm:"column:iron_loss"`                           // 单位铁损(W/kg)
	MagnetizationCapacity float32 `json:"magnetization_capacity" gorm:"column:magnetization_capacity"` //单位质量磁化容量
	SeamVoltammetric      float32 `json:"seam_voltammetric" gorm:"column:seam_voltammetric"`           // 单位面积接缝伏安值VA/cm2
	SiliconSteel          string  `json:"silicon_steel" gorm:"column:silicon_steel"`                   // 硅钢片类型
}

func (CoreIronPerformance) TableName() string {
	return "core_iron_performance"
}
