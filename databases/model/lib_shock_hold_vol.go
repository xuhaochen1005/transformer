// Package model
//  冲击电压
package model

type ShockHoldVol struct {
	ID              int       `json:"id" gorm:"column:id;primary_key"`                     // 冲击电压编号
	RatedHighVolMin float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"` // 额定高压下界
	RatedHighVolMax float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"` // 额定高压上界
	ShockHoldVol    float32   `json:"shock_hold_vol" gorm:"column:shock_hold_vol"`         // 冲击电压
	LshvCreator     int       `json:"lshv_creator" gorm:"column:lshv_creator"`             // 创建者
	CreatorUser     *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LshvCreator;AssociationForeignKey:ID"`
	CreatedAt       Timestamp `json:"created_at" gorm:"column:created_at"` // 冲击电压添加时间
	UpdatedAt       Timestamp `json:"updated_at" gorm:"column:updated_at"` // 冲击电压更新时间
	DeletedAt       Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 冲击电压删除时间
}

func (ShockHoldVol) TableName() string {
	return "lib_shock_hold_vol"
}
