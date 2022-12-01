// Package model
//  损耗
package model

type Loss struct {
	ID                int       `json:"id" gorm:"column:id;primary_key"`                       // 损耗编号
	LossLevel         int       `json:"loss_level" gorm:"column:loss_level"`                   // 损耗水平代号
	RegulateWay       string    `json:"regulate_way" gorm:"column:regulate_way"`               // 调压方式
	RatedCapacityMin  float32   `json:"rated_capacity_min" gorm:"column:rated_capacity_min"`   // 额定容量下界
	RatedCapacityMax  float32   `json:"rated_capacity_max" gorm:"column:rated_capacity_max"`   // 额定容量上界
	RatedHighVolMin   float32   `json:"rated_high_vol_min" gorm:"column:rated_high_vol_min"`   // 额定高压下界
	RatedHighVolMax   float32   `json:"rated_high_vol_max" gorm:"column:rated_high_vol_max"`   // 额定高压上界
	RatedLowVolMin    float32   `json:"rated_low_vol_min" gorm:"column:rated_low_vol_min"`     // 额定低压下界
	RatedLowVolMax    float32   `json:"rated_low_vol_max" gorm:"column:rated_low_vol_max"`     // 额定低压上界
	RegulateRangeMin  float32   `json:"regulate_range_min" gorm:"column:regulate_range_min"`   // 调压范围下界
	RegulateRangeMax  float32   `json:"regulate_range_max" gorm:"column:regulate_range_max"`   // 调压范围上界
	RegulateRangeStep float32   `json:"regulate_range_step" gorm:"column:regulate_range_step"` // 调压范围步长
	Temperature       float32   `json:"temperature" gorm:"column:temperature"`                 // 绝缘温度
	LoadLoss          float32   `json:"load_loss" gorm:"column:load_loss"`                     // 负载损耗
	NoLoadLoss        float32   `json:"no_load_loss" gorm:"column:no_load_loss"`               // 空载损耗
	NoLoadCurrent     float32   `json:"no_load_current" gorm:"column:no_load_current"`         // 空载电流
	ShortCircuitImped float32   `json:"short_circuit_imped" gorm:"column:short_circuit_imped"` // 短路阻抗
	LlCreator         int       `json:"ladbl_creator" gorm:"column:ll_creator"`                // 创建者
	CreatorUser       *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LlCreator;AssociationForeignKey:ID"`
	CreatedAt         Timestamp `json:"created_at" gorm:"column:created_at"` // 损耗添加时间
	UpdatedAt         Timestamp `json:"updated_at" gorm:"column:updated_at"` // 损耗更新时间
	DeletedAt         Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 损耗删除时间
}

func (Loss) TableName() string {
	return "lib_loss"
}
