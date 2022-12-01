package model

// ApproveNodeStatus 审批节点状态
type ApproveNodeStatus int

const (
	ApproveNodeInAction ApproveNodeStatus = iota + 1 // 审批节点进行中
	ApproveNodeRejected                              // 审批节点被拒绝
	ApproveNodeAccepted                              // 审批节点被通过
	ApproveNodeClosed                                // 审批节点被关闭
)

// ApproveNodeInstance 审批节点实例
type ApproveNodeInstance struct {
	ID                int               `json:"id" gorm:"column:id;primary_key"`               // 审批节点实例编号
	ApproveNodeID     int               `json:"approve_node_id" gorm:"column:approve_node_id"` // 审批节点编号
	ApproveNode       *ApproveNode      `json:"approve_node,omitempty" gorm:"ForeignKey:ApproveNodeID"`
	ApproveInstanceID int               `json:"approve_instance_id" gorm:"column:approve_instance_id"` // 审批实例编号
	Approval          int               `json:"approval" gorm:"column:approval"`                       // 审批人
	ApprovalUser      *User             `json:"approval_user,omitempty" gorm:"ForeignKey:Approval"`
	Status            ApproveNodeStatus `json:"status" gorm:"column:status"`         // 审批节点实例状态
	Data              []byte            `json:"data" gorm:"column:data"`             // 审批节点实例相关数据
	Comment           string            `json:"comment" gorm:"column:comment"`       // 审批意见
	Pre               int               `json:"pre" gorm:"pre"`                      // 上一个审批节点实例
	Next              int               `json:"next" gorm:"next"`                    // 下一个审批节点实例
	CreatedAt         Timestamp         `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt         Timestamp         `json:"updated_at" gorm:"column:updated_at"` // 更新时间
	DeletedAt         Timestamp         `json:"-" gorm:"column:deleted_at"`          // 删除时间
}

func (ApproveNodeInstance) TableName() string {
	return "approve_node_instance"
}
