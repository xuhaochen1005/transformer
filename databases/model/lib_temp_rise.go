// Package model
//  绕组温升
package model

type TempRise struct {
	ID              int       `json:"id" gorm:"column:id;primary_key"`                     // 绕组温升编号
	Temp            float32   `json:"temp" gorm:"column:temp"`                             // 绝缘系统温度
	TempSign        string    `json:"temp_sign" gorm:"column:temp_sign"`                   // 绝缘系统温度字母代号
	LowVolTempRise  float32   `json:"low_vol_temp_rise" gorm:"column:low_vol_temp_rise"`   // 低压温升
	HighVolTempRise float32   `json:"high_vol_temp_rise" gorm:"column:high_vol_temp_rise"` // 高压温升
	LtrCreator      int       `json:"ltr_creator" gorm:"column:ltr_creator"`               // 创建者
	CreatorUser     *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LtrCreator;AssociationForeignKey:ID"`
	CreatedAt       Timestamp `json:"created_at" gorm:"column:created_at"` // 绕组温升添加时间
	UpdatedAt       Timestamp `json:"updated_at" gorm:"column:updated_at"` // 绕组温升更新时间
	DeletedAt       Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 绕组温升删除时间
}

func (TempRise) TableName() string {
	return "lib_temp_rise"
}
