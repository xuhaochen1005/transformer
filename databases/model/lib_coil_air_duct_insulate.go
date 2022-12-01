package model

type CoilAirDuctInsulate struct {
	ID                   int       `json:"id" gorm:"column:id;primary_key"`                                 // 绕线内模台账编号
	WindType             string    `json:"wind_type" gorm:"column:wind_type"`                               // 绕制类型
	CoilInnerInsulate    float32   `json:"coil_inner_insulate" gorm:"column:coil_inner_insulate"`           // 线圈内层绝缘（mm）
	CoilOuterInsulate    float32   `json:"coil_outer_insulate" gorm:"column:coil_outer_insulate"`           // 线圈外层绝缘（mm）
	AirDuctThick         string    `json:"air_duct_thick" gorm:"column:air_duct_thick"`                     // 风道厚度可选值（mm）
	AirDuctInOutInsulate float32   `json:"air_duct_in_out_insulate" gorm:"column:air_duct_in_out_insulate"` // 风道内外层绝缘（mm）
	LcadiCreator         int       `json:"lcadi_creator" gorm:"column:lcadi_creator"`                       // 创建者
	CreatorUser          *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LcadiCreator;AssociationForeignKey:ID"`
	CreatedAt            Timestamp `json:"created_at" gorm:"column:created_at"` // 绕线内模台账添加时间
	UpdatedAt            Timestamp `json:"updated_at" gorm:"column:updated_at"` // 绕线内模台账更新时间
	DeletedAt            Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 绕线内模台账删除时间
}

func (CoilAirDuctInsulate) TableName() string {
	return "lib_coil_air_duct_insulate"
}
