// Package model
//  干式变压器铁心直径经验系数K
package model

type CoreK struct {
	ID               int       `json:"id" gorm:"column:id;primary_key"`                     // 系数K编号
	WindingMaterial  string    `json:"winding_material" gorm:"column:winding_material"`     // 绕组材质
	RatedCapacityMin float32   `json:"rated_capacity_min" gorm:"column:rated_capacity_min"` // 额定容量下限
	RatedCapacityMax float32   `json:"rated_capacity_max" gorm:"column:rated_capacity_max"` // 额定容量上限
	KMin             float32   `json:"k_min" gorm:"column:k_min"`                           // k值下限
	KMax             float32   `json:"k_max" gorm:"column:k_max"`                           // k值上限
	LckCreator       int       `json:"lck_creator" gorm:"column:lck_creator"`               // 创建者
	CreatorUser      *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LckCreator;AssociationForeignKey:ID"`
	CreatedAt        Timestamp `json:"created_at" gorm:"column:created_at"` // 系数K添加时间
	UpdatedAt        Timestamp `json:"updated_at" gorm:"column:updated_at"` // 系数K更新时间
	DeletedAt        Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 系数K删除时间
}

func (CoreK) TableName() string {
	return "lib_core_k"
}
