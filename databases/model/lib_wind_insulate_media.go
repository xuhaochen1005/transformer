// Package model
//  绕组外绝缘介质
package model

type WindInsulateMedia struct {
	ID                    int       `json:"id" gorm:"column:id;primary_key"`                                 // 绕组外绝缘介质编号
	WindInsulateMedia     string    `json:"wind_insulate_media" gorm:"column:wind_insulate_media"`           // 模绕组外绝缘介质
	WindInsulateMediaSign string    `json:"wind_insulate_media_sign" gorm:"column:wind_insulate_media_sign"` // 绕组外绝缘介质字母代号
	LwimCreator           int       `json:"lwim_creator" gorm:"column:lwim_creator"`                         // 创建者
	CreatorUser           *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LwimCreator;AssociationForeignKey:ID"`
	CreatedAt             Timestamp `json:"created_at" gorm:"column:created_at"` // 绕组外绝缘介质添加时间
	UpdatedAt             Timestamp `json:"updated_at" gorm:"column:updated_at"` // 绕组外绝缘介质更新时间
	DeletedAt             Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 绕组外绝缘介质删除时间
}

func (WindInsulateMedia) TableName() string {
	return "lib_wind_insulate_media"
}
