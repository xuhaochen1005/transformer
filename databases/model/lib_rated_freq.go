// Package model
//  额定频率
package model

type RatedFreq struct {
	ID          int       `json:"id" gorm:"column:id;primary_key"`       // 额定频率编号
	RatedFreq   int       `json:"rated_freq" gorm:"column:rated_freq"`   // 额定频率
	LrfCreator  int       `json:"lrf_creator" gorm:"column:lrf_creator"` // 创建者
	CreatorUser *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LrfCreator;AssociationForeignKey:ID"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"` // 额定频率添加时间
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"` // 额定频率更新时间
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 额定频率删除时间
}

func (RatedFreq) TableName() string {
	return "lib_rated_freq"
}
