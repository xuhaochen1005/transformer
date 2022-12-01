// Package model
// 树脂规格表
package model

type Resin struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`     // 涡流损耗系数编号
	Type        string    `json:"type" gorm:"column:type"`             // 树脂型号
	Density     float32   `json:"density" gorm:"column:density"`       // 树脂密度（kg/m^3）
	Price       float32   `json:"price" gorm:"column:price"`           // 价格
	LrCreator   int       `json:"lr_creator" gorm:"column:lr_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LrCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 涡流损耗系数添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 涡流损耗系数更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 涡流损耗系数删除时间
}

func (Resin) TableName() string {
	return "lib_resin"
}
