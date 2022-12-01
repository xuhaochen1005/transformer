// Package model
// 端绝缘表
package model

type EndInsulation struct {
	ID                          int    `json:"id" gorm:"column:id;primary_key"`                                             // 部门编号
	EndInsulationVoltage        string `json:"end_insulation_voltage" gorm:"column:end_insulation_voltage"`                 //电压               // 部门名称
	VoltageTypeID               int    `json:"voltage_type_id" gorm:"column:voltage_type_id"`                               //电压类型id                     // 部门状态
	CoilStructurePreliminmaryID int    `json:"coil_structure_preliminmary_id" gorm:"column:coil_structure_preliminmary_id"` // 部门信息添加时间

}

func (EndInsulation) TableName() string {
	return "end_insulation"
}
