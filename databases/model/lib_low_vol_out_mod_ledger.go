// Package model
//  低压外模台账
package model

type LowVolOutModLedger struct {
	ID            int       `json:"id" gorm:"column:id;primary_key"`             // 低压外模台账编号
	ModDiameter   int       `json:"mod_diameter" gorm:"column:mod_diameter"`     // 低压外模外径
	ModType       string    `json:"mod_type" gorm:"column:mod_type"`             // 模具类型
	ModHeight     int       `json:"mod_height" gorm:"column:mod_height"`         // 模具高度
	ModAmount     int       `json:"mod_amount" gorm:"column:mod_amount"`         // 模具数量
	ModNum        string    `json:"mod_num" gorm:"column:mod_num"`               // 模具编号
	ModSuit       string    `json:"mod_suit" gorm:"column:mod_suit"`             // 适用产品
	Remark        string    `json:"remark" gorm:"column:remark"`                 // 备注
	LlvomlCreator int       `json:"llvoml_creator" gorm:"column:llvoml_creator"` // 创建者
	CreatorUser   *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LlvomlCreator;AssociationForeignKey:ID"`
	CreatedAt     Timestamp `json:"created_at" gorm:"column:created_at"` // 损低压外模台账添加时间
	UpdatedAt     Timestamp `json:"updated_at" gorm:"column:updated_at"` // 低压外模台账更新时间
	DeletedAt     Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 低压外模台账删除时间
}

func (LowVolOutModLedger) TableName() string {
	return "lib_low_vol_out_mod_ledger"
}
