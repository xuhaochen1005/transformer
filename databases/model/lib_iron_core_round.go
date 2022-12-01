// Package model
// 铁芯（圆形）
package model

type IronCoreRound struct {
	ID                      int       `json:"id" gorm:"column:id;primary_key"`                                       // 铁芯（圆形）编号
	CoreDiameter            float32   `json:"core_diameter" gorm:"column:core_diameter"`                             // 铁芯直径
	PullPlate               int       `json:"pull_plate" gorm:"column:pull_plate"`                                   // 有无拉板，有1无0
	SectionArea             float32   `json:"section_area" gorm:"column:section_area"`                               // 截面积
	IronCornerWeight        float32   `json:"iron_corner_weight" gorm:"column:iron_corner_weight"`                   // 角重
	StackWidth1             float32   `json:"stack_width_1" gorm:"column:stack_width_1"`                             // 片宽（1级）
	StackWidth2             float32   `json:"stack_width_2" gorm:"column:stack_width_2"`                             // 片宽（2级）
	StackWidth3             float32   `json:"stack_width_3" gorm:"column:stack_width_3"`                             // 片宽（3级）
	StackWidth4             float32   `json:"stack_width_4" gorm:"column:stack_width_4"`                             // 片宽（4级）
	StackWidth5             float32   `json:"stack_width_5" gorm:"column:stack_width_5"`                             // 片宽（5级）
	StackWidth6             float32   `json:"stack_width_6" gorm:"column:stack_width_6"`                             // 片宽（6级）
	StackWidth7             float32   `json:"stack_width_7" gorm:"column:stack_width_7"`                             // 片宽（7级）
	StackWidth8             float32   `json:"stack_width_8" gorm:"column:stack_width_8"`                             // 片宽（8级）
	StackThick1             float32   `json:"stack_thick_1" gorm:"column:stack_thick_1"`                             // 迭厚（1级）
	StackThick2             float32   `json:"stack_thick_2" gorm:"column:stack_thick_2"`                             // 迭厚（2级）
	StackThick3             float32   `json:"stack_thick_3" gorm:"column:stack_thick_3"`                             // 迭厚（3级）
	StackThick4             float32   `json:"stack_thick_4" gorm:"column:stack_thick_4"`                             // 迭厚（4级）
	StackThick5             float32   `json:"stack_thick_5" gorm:"column:stack_thick_5"`                             // 迭厚（5级）
	StackThick6             float32   `json:"stack_thick_6" gorm:"column:stack_thick_6"`                             // 迭厚（6级）
	StackThick7             float32   `json:"stack_thick_7" gorm:"column:stack_thick_7"`                             // 迭厚（7级）
	StackThick8             float32   `json:"stack_thick_8" gorm:"column:stack_thick_8"`                             // 迭厚（8级）
	StackThickSum           float32   `json:"stack_thick_sum" gorm:"column:stack_thick_sum"`                         // 总迭厚
	MainLevelRealStackThick float32   `json:"main_level_real_stack_thick" gorm:"column:main_level_real_stack_thick"` // 主级实迭厚
	LicrCreator             int       `json:"licr_creator" gorm:"column:licr_creator"`                               // 创建者
	CreatorUser             *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LicrCreator;AssociationForeignKey:ID"`
	CreatedAt               Timestamp `json:"created_at" gorm:"column:created_at"` // 铁芯（圆形）添加时间
	UpdatedAt               Timestamp `json:"updated_at" gorm:"column:updated_at"` // 铁芯（圆形）更新时间
	DeletedAt               Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 铁芯（圆形）删除时间

}

func (IronCoreRound) TableName() string {
	return "lib_iron_core_round"
}
