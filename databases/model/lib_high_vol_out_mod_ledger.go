// Package model
//  高压外模台账
package model

type HighVolOutModLedger struct {
	ID              int       `json:"id" gorm:"column:id;primary_key"`                   // 高压外模台账编号
	ModSize         string    `json:"mod_size" gorm:"column:mod_size"`                   // 模具尺寸
	ModType         string    `json:"mod_type" gorm:"column:mod_type"`                   // 模具类型
	ModHeight       int       `json:"mod_height" gorm:"column:mod_height"`               // 模具高度
	ModAmount       int       `json:"mod_amount" gorm:"column:mod_amount"`               // 模具数量
	ModPic          string    `json:"mod_pic" gorm:"column:mod_pic"`                     // 模具图号
	ModNum          string    `json:"mod_num" gorm:"column:mod_num"`                     // 模具编号
	ModSuit         string    `json:"mod_suit" gorm:"column:mod_suit"`                   // 适用产品型号
	NutSize         int       `json:"nut_size" gorm:"column:nut_size"`                   // 螺母尺寸
	BossWidth       string    `json:"boss_width" gorm:"column:boss_width"`               // 凸台高度（封板尺寸）
	Groove          string    `json:"groove" gorm:"column:groove"`                       // 是否开槽
	TapHoleDistance float32   `json:"tap_hole_distance" gorm:"column:tap_hole_distance"` // 抽头孔距
	ClosurePic      string    `json:"closure_pic" gorm:"column:closure_pic"`             // 封板图号
	ASize           int       `json:"a_size" gorm:"column:a_size"`                       // A头尺寸
	Remark          string    `json:"remark" gorm:"column:remark"`                       // 备注
	LhvomlCreator   int       `json:"lhvoml_creator" gorm:"column:lhvoml_creator"`       // 创建者
	CreatorUser     *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LhvomlCreator;AssociationForeignKey:ID"`
	CreatedAt       Timestamp `json:"created_at" gorm:"column:created_at"` // 高压外模台账添加时间
	UpdatedAt       Timestamp `json:"updated_at" gorm:"column:updated_at"` // 高压外模台账更新时间
	DeletedAt       Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 高压外模台账删除时间
}

func (HighVolOutModLedger) TableName() string {
	return "lib_high_vol_out_mod_ledger"
}
