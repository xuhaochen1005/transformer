// Package model
//  硅钢片性能数据
package model

type StackPerformData struct {
	ID                     int       `json:"id" gorm:"column:id;primary_key"`                                   // 硅钢片性能数据编号
	StackType              string    `json:"stack_type" gorm:"column:stack_type"`                               // 硅钢片型号
	CoreFluxDensity        float32   `json:"core_flux_density" gorm:"column:core_flux_density"`                 // 铁芯磁密
	UnitIronLoss           float32   `json:"unit_iron_loss" gorm:"column:unit_iron_loss"`                       // 单位铁损
	UnitMassMagnetCapacity float32   `json:"unit_mass_magnet_capacity" gorm:"column:unit_mass_magnet_capacity"` // 单位质量磁化容量
	UnitAreaSeamVa         float32   `json:"unit_area_seam_va" gorm:"column:unit_area_seam_va"`                 // 单位面积接缝伏安值
	LspdCreator            int       `json:"lspd_creator" gorm:"column:lspd_creator"`                           // 创建者
	CreatorUser            *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LspdCreator;AssociationForeignKey:ID"`
	CreatedAt              Timestamp `json:"created_at" gorm:"column:created_at"` // 硅钢片性能数据添加时间
	UpdatedAt              Timestamp `json:"updated_at" gorm:"column:updated_at"` // 硅钢片性能数据更新时间
	DeletedAt              Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 硅钢片性能数据删除时间
}

func (StackPerformData) TableName() string {
	return "lib_stack_perform_data"
}
