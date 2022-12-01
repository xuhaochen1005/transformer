// Package model
//  引线损耗
package model

type LeadLoss struct {
	ID            int       `json:"id" gorm:"column:id;primary_key"`             // 引线损耗编号
	RatedCapacity float32   `json:"rated_capacity" gorm:"column:rated_capacity"` // 额定容量
	LeadLoss      float32   `json:"lead_loss" gorm:"column:lead_loss"`           // 引线损耗
	LlsCreator    int       `json:"lls_creator" gorm:"column:lls_creator"`       // 创建者
	CreatorUser   *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LlsCreator;AssociationForeignKey:ID"`
	CreatedAt     Timestamp `json:"created_at" gorm:"column:created_at"` // 引线损耗添加时间
	UpdatedAt     Timestamp `json:"updated_at" gorm:"column:updated_at"` // 引线损耗更新时间
	DeletedAt     Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 引线损耗更新时间

}

func (LeadLoss) TableName() string {
	return "lib_lead_loss"
}
