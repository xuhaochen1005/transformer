package model

// ApproveStatus 审批状态
type ApproveStatus int

const (
	ApproveInAction ApproveStatus = 1 // 审批进行中
	ApproveRejected ApproveStatus = 2 // 审批被拒绝
	ApproveAccepted ApproveStatus = 3 // 审批被通过
	ApproveClosed   ApproveStatus = 4 // 审批被关闭
)

const (
	ApproveIDDesignTask int = 1
)

// ApproveInstance 审批实例
type ApproveInstance struct {
	ID                   int                    `json:"id" gorm:"column:id;primary_key"`                    // 审批实例编号
	ApproveID            int                    `json:"approve_id" gorm:"column:approve_id"`                // 审批编号
	Head                 int                    `json:"head" gorm:"column:head"`                            // 审批实例头节点
	Data                 []byte                 `json:"data" gorm:"column:data"`                            // 审批实例关联的数据
	Status               ApproveStatus          `json:"status" validate:"gte=1,lte=4" gorm:"column:status"` // 审批实例状态
	Applicant            int                    `json:"applicant" gorm:"column:applicant"`                  // 审批实例发起者
	Participants         []byte                 `json:"participants" gorm:"column:participants"`            // 审批实例参与者
	Approve              *Approve               `json:"approve" gorm:"ForeignKey:ApproveID"`
	ApplicantUser        *User                  `json:"applicant_user" gorm:"ForeignKey:Applicant"`
	ApproveNodeInstances *[]ApproveNodeInstance `json:"approve_node_instances" gorm:"AssociationForeignKey:ApproveInstanceID"`
	CreatedAt            Timestamp              `json:"created_at" gorm:"column:created_at"` // 创建时间
}

func (ApproveInstance) TableName() string {
	return "approve_instance"
}
