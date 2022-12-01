// Package model
// 相间距离最小值
package model

type PhaseDist struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`             // 相间距离最小值编号
	PhaseDistMin int       `json:"phase_dist_min" gorm:"column:phase_dist_min"` // 相间距离最小值（mm）
	VolMin       float32   `json:"vol_min" gorm:"column:vol_min"`               // 电压下限（kV）
	VolMax       float32   `json:"vol_max" gorm:"column:vol_max"`               // 电压上限（kV）
	LpdCreator   int       `json:"lpd_creator" gorm:"column:lpd_creator"`       // 创建者
	CreatorUser  *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LpdCreator;AssociationForeignKey:ID"`
	CreatedAt    Timestamp `json:"created_at" gorm:"column:created_at"` // 相间距离最小值添加时间
	UpdatedAt    Timestamp `json:"updated_at" gorm:"column:updated_at"` // 相间距离最小值更新时间
	DeletedAt    Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 相间距离最小值删除时间
}

func (PhaseDist) TableName() string {
	return "lib_phase_dist"
}
