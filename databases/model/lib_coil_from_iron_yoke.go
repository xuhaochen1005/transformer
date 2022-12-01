// Package model
//  线圈端部距铁轭距离
package model

type CoilFromIronYoke struct {
	ID                      int       `json:"id" gorm:"column:id;primary_key"`                                         // 线圈端部距铁轭距离编号
	RatedHighVolMin         float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"`                     // 额定高压下界
	RatedHighVolMax         float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"`                     // 额定高压上界
	RatedLowVolMin          float32   `json:"rated_low_vol_min" gorm:"column:rated_low_vol_min"`                       // 额定低压下界
	RatedLowVolMax          float32   `json:"rated_low_vol_max" gorm:"column:rated_low_vol_max"`                       // 额定低压上界
	LowVolCoilFromIronYoke  float32   `json:"low_vol_coil_from_iron_yoke" gorm:"column:low_vol_coil_from_iron_yoke"`   // 低压线圈端部距铁轭距离
	HighVolCoilFromIronYoke float32   `json:"high_vol_coil_from_iron_yoke" gorm:"column:high_vol_coil_from_iron_yoke"` // 高压线圈端部距铁轭距离
	LcfiyCreator            int       `json:"lcfiy_creator" gorm:"column:lcfiy_creator"`                               // 创建者
	CreatorUser             *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LcfiyCreator;AssociationForeignKey:ID"`
	CreatedAt               Timestamp `json:"created_at" gorm:"column:created_at"` // 线圈端部距铁轭距离添加时间
	UpdatedAt               Timestamp `json:"updated_at" gorm:"column:updated_at"` // 线圈端部距铁轭距离更新时间
	DeletedAt               Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 线圈端部距铁轭距离删除时间
}

func (CoilFromIronYoke) TableName() string {
	return "lib_coil_from_iron_yoke"
}
