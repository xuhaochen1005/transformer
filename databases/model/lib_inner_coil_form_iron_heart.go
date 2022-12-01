// Package model
// 内线圈至铁芯距离
package model

type InnerCoilFormIronHeart struct {
	ID                        int       `json:"id" gorm:"column:id;primary_key"`                                             // 内线圈至铁芯距离编号
	RatedHighVolMin           float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"`                         // 额定高压下界
	RatedHighVolMax           float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"`                         // 额定高压上界
	InnerCoilFormIronHeartMin float32   `json:"inner_coil_form_iron_heart_min" gorm:"column:inner_coil_form_iron_heart_min"` // 内线圈至铁芯距离最小值
	LicfihCreator             int       `json:"licfih_creator" gorm:"column:licfih_creator"`                                 // 创建者
	CreatorUser               *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LicfihCreator;AssociationForeignKey:ID"`
	CreatedAt                 Timestamp `json:"created_at" gorm:"column:created_at"` // 内线圈至铁芯距离添加时间
	UpdatedAt                 Timestamp `json:"updated_at" gorm:"column:updated_at"` // 内线圈至铁芯距离更新时间
	DeletedAt                 Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 内线圈至铁芯距离删除时间

}

func (InnerCoilFormIronHeart) TableName() string {
	return "lib_inner_coil_form_iron_heart"
}
