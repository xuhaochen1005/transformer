// Package model
//  绕线内模台账
package model

type WindInModLedger struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`           // 绕线内模台账编号
	ModDiameter  int       `json:"mod_diameter" gorm:"column:mod_diameter"`   // 模具直径
	ModSize      string    `json:"mod_size" gorm:"column:mod_size"`           // 模具规格
	ModAmount    int       `json:"mod_amount" gorm:"column:mod_amount"`       // 模具数量
	ModPic       string    `json:"mod_pic" gorm:"column:mod_pic"`             // 模具图号
	ModNum       string    `json:"mod_num" gorm:"column:mod_num"`             // 模具编号
	Remark       string    `json:"remark" gorm:"column:remark"`               // 备注
	ModSuit      string    `json:"mod_suit" gorm:"column:mod_suit"`           // 适用产品型号
	LwimlCreator int       `json:"lwiml_creator" gorm:"column:lwiml_creator"` // 创建者
	CreatorUser  *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LwimlCreator;AssociationForeignKey:ID"`
	CreatedAt    Timestamp `json:"created_at" gorm:"column:created_at"` // 绕线内模台账添加时间
	UpdatedAt    Timestamp `json:"updated_at" gorm:"column:updated_at"` // 绕线内模台账更新时间
	DeletedAt    Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 绕线内模台账删除时间
}

func (WindInModLedger) TableName() string {
	return "lib_wind_in_mod_ledger"
}
