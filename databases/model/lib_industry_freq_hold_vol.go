// Package model
//  工频耐压
package model

type IndustryFreqHoldVol struct {
	ID                  int       `json:"id" gorm:"column:id;primary_key"`                             // 工频耐压编号
	RatedHighVolMin     float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"`         // 额定高压下界
	RatedHighVolMax     float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"`         // 额定高压上界
	IndustryFreqHoldVol float32   `json:"industry_freq_hold_vol" gorm:"column:industry_freq_hold_vol"` // 工频耐压
	LifhvCreator        int       `json:"lifhv_creator" gorm:"column:lifhv_creator"`                   // 创建者
	CreatorUser         *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LifhvCreator;AssociationForeignKey:ID"`
	CreatedAt           Timestamp `json:"created_at" gorm:"column:created_at"` // 层间绝缘距离添加时间
	UpdatedAt           Timestamp `json:"updated_at" gorm:"column:updated_at"` // 层间绝缘距离更新时间
	DeletedAt           Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 层间绝缘距离删除时间

}

func (IndustryFreqHoldVol) TableName() string {
	return "lib_industry_freq_hold_vol"
}
