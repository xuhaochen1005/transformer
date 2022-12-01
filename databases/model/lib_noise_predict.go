// Package model
//  噪声预测
package model

type NoisePredict struct {
	ID               int       `json:"id" gorm:"column:id;primary_key"`                     // 噪声预测编号
	RatedCapacityMin float32   `json:"rated_capacity_min" gorm:"column:rated_capacity_min"` // 额定容量下界
	RatedCapacityMax float32   `json:"rated_capacity_max" gorm:"column:rated_capacity_max"` // 额定容量上界
	RatedHighVolMin  float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"` // 额定高压下界
	RatedHighVolMax  float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"` // 额定高压上界
	CoolType         string    `json:"cool_type" gorm:"column:cool_type"`                   // 冷却方式
	NoisePredict     int       `json:"noise_predict" gorm:"column:noise_predict"`           // 噪声预测
	LnpCreator       int       `json:"lnp_creator" gorm:"column:lnp_creator"`               // 创建者
	CreatorUser      *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LnpCreator;AssociationForeignKey:ID"`
	CreatedAt        Timestamp `json:"created_at" gorm:"column:created_at"` // 噪声预测添加时间
	UpdatedAt        Timestamp `json:"updated_at" gorm:"column:updated_at"` // 噪声预测更新时间
	DeletedAt        Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 噪声预测删除时间
}

func (NoisePredict) TableName() string {
	return "lib_noise_predict"
}
