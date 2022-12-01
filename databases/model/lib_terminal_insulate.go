// Package model
//  端绝缘距离
package model

type TerminalInsulate struct {
	ID               int       `json:"id" gorm:"column:id;primary_key"`                     // 端绝缘距离编号
	RatedHighVolMin  float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"` // 额定高压下界
	RatedHighVolMax  float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"` // 额定高压上界
	RatedLowVolMin   float32   `json:"rated_low_vol_min" gorm:"column:rated_low_vol_min"`   // 额定低压下界
	RatedLowVolMax   float32   `json:"rated_low_vol_max" gorm:"column:rated_low_vol_max"`   // 额定低压上界
	TerminalInsulate float32   `json:"terminal_insulate" gorm:"column:terminal_insulate"`   // 端绝缘距离
	LtiCreator       int       `json:"lti_creator" gorm:"column:lti_creator"`               // 创建者
	CreatorUser      *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LtiCreator;AssociationForeignKey:ID"`
	CreatedAt        Timestamp `json:"created_at" gorm:"column:created_at"` // 端绝缘距离添加时间
	UpdatedAt        Timestamp `json:"updated_at" gorm:"column:updated_at"` // 端绝缘距离更新时间
	DeletedAt        Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 端绝缘距离删除时间
}

func (TerminalInsulate) TableName() string {
	return "lib_terminal_insulate"
}
