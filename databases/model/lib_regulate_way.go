// Package model
//  调压方式
package model

type RegulateWay struct {
	ID              int       `json:"id" gorm:"column:id;primary_key"`                   // 调压方式编号
	RegulateWay     string    `json:"regulate_way" gorm:"column:regulate_way"`           // 调压方式
	RegulateWaySign string    `json:"regulate_way_sign" gorm:"column:regulate_way_sign"` // 调压方式字母表示
	LrwCreator      int       `json:"lrw_creator" gorm:"column:lrw_creator"`             // 创建者
	CreatorUser     *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LrwCreator;AssociationForeignKey:ID"`
	CreatedAt       Timestamp `json:"created_at" gorm:"column:created_at"` // 调压方式添加时间
	UpdatedAt       Timestamp `json:"updated_at" gorm:"column:updated_at"` // 调压方式更新时间
	DeletedAt       Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 调压方式删除时间
}

func (RegulateWay) TableName() string {
	return "lib_regulate_way"
}
