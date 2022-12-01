// Package model
//  涡流损耗系数
package model

type EddyLossFactor struct {
	ID             int       `json:"id" gorm:"column:id;primary_key"`                 // 涡流损耗系数编号
	RatedCapacity  float32   `json:"rated_capacity" gorm:"column:rated_capacity"`     // 额定容量
	EddyLossFactor float32   `json:"eddy_loss_factor" gorm:"column:eddy_loss_factor"` // 涡流损耗系数
	LelfCreator    int       `json:"lelf_creator" gorm:"column:lelf_creator"`         // 创建者
	CreatorUser    *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LelfCreator;AssociationForeignKey:ID"`
	CreatedAt      Timestamp `json:"created_at" gorm:"column:created_at"` // 涡流损耗系数添加时间
	UpdatedAt      Timestamp `json:"updated_at" gorm:"column:updated_at"` // 涡流损耗系数更新时间
	DeletedAt      Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 涡流损耗系数删除时间
}

func (EddyLossFactor) TableName() string {
	return "lib_eddy_loss_factor"
}
