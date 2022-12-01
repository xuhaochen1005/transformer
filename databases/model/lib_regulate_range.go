// Package model
//  调压范围
package model

type RegulateRange struct {
	ID                int       `json:"id" gorm:"column:id;primary_key"`                       // 调压范围编号
	RegulateRangeMin  float32   `json:"regulate_range_min" gorm:"column:regulate_range_min"`   // 调压范围下界
	RegulateRangeMax  float32   `json:"regulate_range_max" gorm:"column:regulate_range_max"`   // 调压范围上界
	RegulateRangeStep float32   `json:"regulate_range_step" gorm:"column:regulate_range_step"` // 调压范围步长
	LrrCreator        int       `json:"lrr_creator" gorm:"column:lrr_creator"`                 // 创建者
	CreatorUser       *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LrrCreator;AssociationForeignKey:ID"`
	CreatedAt         Timestamp `json:"created_at" gorm:"column:created_at"` // 调压范围添加时间
	UpdatedAt         Timestamp `json:"updated_at" gorm:"column:updated_at"` // 调压范围更新时间
	DeletedAt         Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 调压范围删除时间
}

func (RegulateRange) TableName() string {
	return "lib_regulate_range"
}
