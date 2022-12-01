// Package model
//  层间绝缘距离
package model

type LayerInsulate struct {
	ID             int       `json:"id" gorm:"column:id;primary_key"`                   // 层间绝缘距离编号
	WindingType    string    `json:"winding_type" gorm:"column:winding_type"`           // 绕制类型，箔绕/线绕
	RatedLowVolMin float32   `json:"rated_low_vol_min" gorm:"column:rated_low_vol_min"` // 额定低压下界
	RatedLowVolMax float32   `json:"rated_low_vol_max" gorm:"column:rated_low_vol_max"` // 额定低压上界
	LayerVolMin    float32   `json:"layer_vol_min" gorm:"column:layer_vol_min"`         // 层间电压下界
	LayerVolMax    float32   `json:"layer_vol_max" gorm:"column:layer_vol_max"`         // 层间电压上界
	LayerInsulate  float32   `json:"layer_insulate" gorm:"column:layer_insulate"`       // 层间绝缘距离
	LliCreator     int       `json:"lli_creator" gorm:"column:lli_creator"`             // 创建者
	CreatorUser    *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LliCreator;AssociationForeignKey:ID"`
	CreatedAt      Timestamp `json:"created_at" gorm:"column:created_at"` // 层间绝缘距离添加时间
	UpdatedAt      Timestamp `json:"updated_at" gorm:"column:updated_at"` // 层间绝缘距离更新时间
	DeletedAt      Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 层间绝缘距离删除时间

}

func (LayerInsulate) TableName() string {
	return "lib_layer_insulate"
}
