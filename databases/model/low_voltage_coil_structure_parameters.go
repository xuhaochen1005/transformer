// Package model
// 低压线圈参数
package model

type LowVCoilStructureParam struct {
	ID                int    `json:"id" gorm:"column:id;primary_key"`                   // 编号
	InitialTurns      int    `json:"initial_turns" gorm:"column:initial_turns"`         //初选匝数         // 部门名称
	InitialConductors string `json:"intial_conductors" gorm:"column:intial_conductors"` // 初选导线
	LowVstructure     string `json:"low_v_structure" gorm:"column:low_v_structure"`     // 低压结构
	Duct1Thickness    int    `json:"duct1_thickness" gorm:"column:duct1_thickness"`     // 风道1厚度
	Duct2Thickness    int    `json:"duct2_thickness" gorm:"column:duct2_thickness"`     // 风道2厚度

	CoilStructureID             int `json:"coil_structure_id" gorm:"column:coil_structure_id"`                           //线圈结构id                        // 部门信息更新时间
	CoilStructurePreliminmaryID int `json:"coil_structure_preliminmary_id" gorm:"column:coil_structure_preliminmary_id"` // 线圈结构初选id

}

func (LowVCoilStructureParam) TableName() string {
	return "low_voltage_coil_structure_parameters"
}
