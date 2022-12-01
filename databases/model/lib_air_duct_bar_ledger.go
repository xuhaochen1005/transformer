// Package model
//  风道条台账
package model

type AirDuctBarLedger struct {
	ID                  int       `json:"id" gorm:"column:id;primary_key"`                             // 风道条台账编号
	AirDuctBarThickness int       `json:"air_duct_bar_thickness" gorm:"column:air_duct_bar_thickness"` // 风道条厚度
	AirDuctBarWidth     int       `json:"air_duct_bar_width" gorm:"column:air_duct_bar_width"`         // 风道条宽度
	AirDuctBarLength    int       `json:"air_duct_bar_length" gorm:"column:air_duct_bar_length"`       // 风道条长度
	AirDuctBarNum       int       `json:"air_duct_bar_num" gorm:"column:air_duct_bar_num"`             // 风道条数量
	LadblCreator        int       `json:"ladbl_creator" gorm:"column:ladbl_creator"`                   // 创建者
	CreatorUser         *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LadblCreator;AssociationForeignKey:ID"`
	CreatedAt           Timestamp `json:"created_at" gorm:"column:created_at"` // 风道条台账添加时间
	UpdatedAt           Timestamp `json:"updated_at" gorm:"column:updated_at"` // 风道条台账更新时间
	DeletedAt           Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 风道条台账删除时间
}

func (AirDuctBarLedger) TableName() string {
	return "lib_air_duct_bar_ledger"
}
