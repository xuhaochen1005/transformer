// Package model
//  铁心磁密Bm初选值
package model

type MagnetDensity struct {
	ID               int       `json:"id" gorm:"column:id;primary_key"`                     // 铁心磁密Bm初选值编号
	RatedCapacityMin float32   `json:"rated_capacity_min" gorm:"column:rated_capacity_min"` // 额定容量下界
	RatedCapacityMax float32   `json:"rated_capacity_max" gorm:"column:rated_capacity_max"` // 额定容量上界
	MagnetDensityMin float32   `json:"magnet_density_min" gorm:"column:magnet_density_min"` // 铁心磁密Bm下界
	MagnetDensityMax float32   `json:"magnet_density_max" gorm:"column:magnet_density_max"` // 铁心磁密Bm上界
	LmdCreator       int       `json:"lmd_creator" gorm:"column:lmd_creator"`               // 创建者
	CreatorUser      *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LmdCreator;AssociationForeignKey:ID"`
	CreatedAt        Timestamp `json:"created_at" gorm:"column:created_at"` // 铁心磁密Bm初选值添加时间
	UpdatedAt        Timestamp `json:"updated_at" gorm:"column:updated_at"` // 铁心磁密Bm初选值更新时间
	DeletedAt        Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 铁心磁密Bm初选值删除时间
}

func (MagnetDensity) TableName() string {
	return "lib_magnet_density"
}
