package model

const (
	ApproveNodeDesignTaskHeadID int = 1
)

// ApproveNode 审批节点
type ApproveNode struct {
	ID        int       `json:"id" gorm:"column:id;primary_key"`                         // 审批节点编号
	ApproveID int       `json:"approve_id" validate:"required" gorm:"column:approve_id"` // 审批编号
	Approve   *Approve  `json:"approve" gorm:"ForeignKey:ApproveID"`
	Name      string    `json:"name" validate:"required" gorm:"column:name"` // 审批节点名称
	API       string    `json:"api" validate:"required" gorm:"column:api"`   // 审批节点操作API
	Pre       int       `json:"pre" gorm:"column:pre"`                       // 上一个审批节点
	Next      int       `json:"next" gorm:"column:next"`                     // 下一个审批节点
	CreatedBy string    `json:"created_by" gorm:"column:created_by"`         // 创建者
	DeletedBy string    `json:"deleted_by" gorm:"column:deleted_by"`         // 删除者
	CreatedAt Timestamp `json:"created_at" gorm:"column:created_at"`         // 创建时间
	UpdatedAt Timestamp `json:"updated_at" gorm:"column:updated_at"`         // 更新时间
	DeletedAt Timestamp `json:"deleted_at" gorm:"column:deleted_at"`         // 删除时间
}

func (ApproveNode) TableName() string {
	return "approve_node"
}
