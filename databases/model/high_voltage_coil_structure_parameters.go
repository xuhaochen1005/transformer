// Package model
//y 高压线圈参数
package model

type HighVcoilStructureParam struct {
	ID                          int    `json:"id" gorm:"column:id;primary_key"`
	PackageNumber               string `json:"packages_number" gorm:"column:packages_number"`                               //包封数
	Package1SegmentLayer        string `json:"package1_segment_layer" gorm:"column:package1_segment_layer"`                 //包封1段数/层数
	Package2SegmentLayer        string `json:"package2_segment_layer" gorm:"column:package2_segment_layer"`                 //包封2段数/层数
	InitialWire                 string `json:"initial_wire" gorm:"column:initial_wire"`                                     //初选导线
	WindwayThickness            int    `json:"windway_thickness" gorm:"column:windway_thickness"`                           //风道厚
	CoilStructureID             int    `json:"coil_structure_id" gorm:"column:coil_structure_id"`                           //线圈结构id
	CoilStructurePreliminmaryID int    `json:"coil_structure_preliminmary_id" gorm:"column:coil_structure_preliminmary_id"` //线圈初选结构id
}

func (HighVcoilStructureParam) TableName() string {
	return "high_voltage_coil_structure_parameters"
}
