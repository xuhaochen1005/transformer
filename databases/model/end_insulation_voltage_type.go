// Package model
// 端绝缘电压类型

package model

type EndInsulationVType struct {
	ID          int    `json:"id" gorm:"column:id;primary_key"`         // 编号
	VoltageType string `json:"voltage_type" gorm:"column:voltage_type"` // 电压类型
}

func (EndInsulationVType) TableName() string {
	return "end_insulation_voltage_type"
}
