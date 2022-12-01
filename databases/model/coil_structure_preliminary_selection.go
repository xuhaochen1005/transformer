// Package model
// 线圈结构初选表
package model

type CoilStructurePreliminarySelection struct {
	ID                            int     `json:"id" gorm:"column:id;primary_key"`                                                   // 编号
	Capacity                      int     `json:"capacity" gorm:"column:capacity"`                                                   //  容量
	CoreDiameter                  float32 `json:"core_diameter" gorm:"column:core_diameter"`                                         // 铁心直径
	CoreLowvInnerDiameterDistance float32 `json:"core_lowv_inner_diameter_distance" gorm:"column:core_lowv_inner_diameter_distance"` //低压内径距铁心
	CoilHeight                    float32 `json:"coil_height" gorm:"column:coil_height"`                                             //	线圈高
	IronYokeDistance              float32 `json:"iron_yoke_distance" gorm:"column:iron_yoke_distance"`                               //	距铁轭
	InitialWindowHeight           float32 `json:"initial_window_height" gorm:"column:initial_window_height"`                         //窗高（初选）
	StructureType                 int     `json:"structure_type" gorm:"column:structure_type"`                                       //结构(1:端绝缘；2:低压；3：高压)
}

func (CoilStructurePreliminarySelection) TableName() string {
	return "coil_structure_preliminary_selection"
}
