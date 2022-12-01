package model

// Approve 审批
type Approve struct {
	ID        int       `json:"id" gorm:"column:id;primary_key"`             // 审批编号
	Name      string    `json:"name" validate:"required" gorm:"column:name"` // 审批名称
	Head      int       `json:"head" gorm:"column:head"`                     // 审批头节点
	CreatedBy string    `json:"created_by" gorm:"column:created_by"`         // 创建者
	DeletedBy string    `json:"deleted_by" gorm:"column:deleted_by"`         // 删除者
	CreatedAt Timestamp `json:"created_at" gorm:"column:created_at"`         // 创建时间
	UpdatedAt Timestamp `json:"updated_at" gorm:"column:updated_at"`         // 更新时间
	DeletedAt Timestamp `json:"deleted_at" gorm:"column:deleted_at"`         // 删除时间
}

func (Approve) TableName() string {
	return "approve"
}
