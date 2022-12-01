// Package model
// 线圈接法
package model

type CoilConnect struct {
	ID                 int       `json:"id" gorm:"column:id;primary_key"`                           // 线圈接法编号
	LowVolCoilConnect  string    `json:"low_vol_coil_connect" gorm:"column:low_vol_coil_connect"`   // 高压线圈接法
	HighVolCoilConnect string    `json:"high_vol_coil_connect" gorm:"column:high_vol_coil_connect"` // 低压线圈接法
	LccCreator         int       `json:"lcc_creator" gorm:"column:lcc_creator"`                     // 创建者
	CreatorUser        *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LccCreator;AssociationForeignKey:ID"`
	CreatedAt          Timestamp `json:"created_at" gorm:"column:created_at"` // 线圈接法添加时间
	UpdatedAt          Timestamp `json:"updated_at" gorm:"column:updated_at"` // 线圈接法更新时间
	DeletedAt          Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 线圈接法删除时间

}

func (CoilConnect) TableName() string {
	return "lib_coil_connect"
}
