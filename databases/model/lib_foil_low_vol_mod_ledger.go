// Package model
//  箔绕低压模具台账
package model

type FoilLowVolModLedger struct {
	ID             int       `json:"id" gorm:"column:id;primary_key"`               // 箔绕低压模具台账编号
	ModSize        string    `json:"mod_size" gorm:"column:mod_size"`               // 模具尺寸
	PlatformHeight float32   `json:"platform_height" gorm:"column:platform_height"` // 平台高
	ModHeight      int       `json:"mod_height" gorm:"column:mod_height"`           // 模具高度
	ModPic         string    `json:"mod_pic" gorm:"column:mod_pic"`                 // 模具图号
	ModNum         string    `json:"mod_num" gorm:"column:mod_num"`                 // 模具编号
	ModAmount      int       `json:"mod_amount" gorm:"column:mod_amount"`           // 模具数量
	ModSuit        string    `json:"mod_suit" gorm:"column:mod_suit"`               // 适用产品型号
	Remark         string    `json:"remark" gorm:"column:remark"`                   // 备注
	LflvmlCreator  int       `json:"lflvml_creator" gorm:"column:lflvml_creator"`   // 创建者
	CreatorUser    *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LflvmlCreator;AssociationForeignKey:ID"`
	CreatedAt      Timestamp `json:"created_at" gorm:"column:created_at"` // 箔绕低压模具台账添加时间
	UpdatedAt      Timestamp `json:"updated_at" gorm:"column:updated_at"` // 箔绕低压模具台账更新时间
	DeletedAt      Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 箔绕低压模具台账删除时间
}

func (FoilLowVolModLedger) TableName() string {
	return "lib_foil_low_vol_mod_ledger"
}
