package model

type DesignTaskStatus int

const (
	DesignTaskCanceled DesignTaskStatus = iota + 1
	DesignTaskCreated
	DesignTaskStarted
	DesignTaskFinished
	DesignTaskResultSelected
	DesignTaskReviewUnaccepted
	DesignTaskReviewed
	DesignTaskApproveUnaccepted
	DesignTaskApproved
	DesignTaskCheckUnaccepted
	DesignTaskChecked
)

type DesignTask struct {
	ID                  int                  `json:"id" gorm:"column:id;primary_key"`                                            // 设计任务编号
	DesignProjectID     int                  `json:"design_project_id" validate:"required,gt=0" gorm:"column:design_project_id"` // 关联的任务书编号
	DesignProject       *DesignProject       `json:"design_project" gorm:"<-:false;foreignKey:DesignProjectID"`                  // 关联的任务书
	FinalDesignResults  []*DesignResults     `json:"final_design_results" gorm:"<-:false;foreignKey:DesignTaskID"`               // 设计结果
	ApproveNodeInstance *ApproveNodeInstance `json:"approve_node_instance" gorm:"-"`
	Name                string               `json:"name" validate:"required,min=8,max=64" gorm:"column:name"` // 设计任务的名称
	Creator             int                  `json:"creator" validate:"required" gorm:"column:creator"`        // 设计任务发起者
	Designer            int                  `json:"designer" validate:"required" gorm:"column:designer"`      // 设计任务执行者
	DesignerUser        *User                `json:"designer_user" gorm:"<-:false;foreignKey:Designer"`
	Reviewer            int                  `json:"reviewer" validate:"required" gorm:"column:reviewer"`           // 设计任务复核者
	Approve             int                  `json:"approve" validate:"required" gorm:"column:approve"`             // 设计任务审核者
	ApproveByCreator    bool                 `json:"approve_by_creator" gorm:"approve_by_creator"`                  // 设计任务是否由发起者代为审批
	TaskStatus          DesignTaskStatus     `json:"task_status" validate:"gte=1,lte=10" gorm:"column:task_status"` // 设计任务执行状态
	Comment             string               `json:"comment" gorm:"column:comment"`                                 // 设计任务取消，复核不通过或者审批不通过的原因
	Input               []byte               `json:"input" gorm:"column:input"`
	DesignResults       []byte               `json:"design_results" gorm:"column:design_results"`   //设计结果
	ResultProgress      float32              `json:"result_progress" gorm:"column:result_progress"` //设计结果进度
	CreatedAt           Timestamp            `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           Timestamp            `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt           Timestamp            `json:"deleted_at" gorm:"column:deleted_at"`
}

func (DesignTask) TableName() string {
	return "design_task"
}
