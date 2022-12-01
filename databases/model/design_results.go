// Package model
// 设计结果

package model

type ResultStatus int

const (
	ResultStatusUnStarted ResultStatus = iota
	ResultStatusCalculating
	ResultStatusFinished
)

type DesignResults struct {
	ID          int    `json:"id" gorm:"column:id;type:integer;primary_key;comment:消息模板编号"`            // 消息模板编号
	ProductName string `json:"product_name" gorm:"column:product_name;type:varchar(255);comment:产品型号"` // 产品型号
	ProjectName string `json:"project_name" gorm:"column:project_name;type:varchar(255);comment:项目名称"` // 项目名称
	ProjectCode string `json:"project_code" gorm:"column:project_code;type:varchar(255);comment:项目型号"` // 项目型号

	//endregion
	DesignTaskID      int            `json:"design_task_id" gorm:"column:design_task_id;type:integer"`
	DesignProjectID   int            `json:"design_project_id" gorm:"column:design_project_id;type:integer"`
	DesignProject     *DesignProject `json:"design_project" gorm:"foreignKey:DesignProjectID"`
	DesignTask        *DesignTask    `json:"design_task" gorm:"foreignKey:DesignTaskID"`
	ResultStatus      ResultStatus   `json:"result_status" gorm:"column:result_status;type:integer"`
	ContinueVars      []byte         `json:"continue_vars" gorm:"column:continue_vars"`
	CoreResult        []byte         `json:"core_result" gorm:"column:core_result"`
	HvResult          []byte         `json:"hv_result" gorm:"column:hv_result"`
	LvResult          []byte         `json:"lv_result" gorm:"column:lv_result"`
	ImpedanceResult   []byte         `json:"impedance_result" gorm:"column:impedance_result"`
	PerformanceResult []byte         `json:"performance_result" gorm:"column:performance_result"`
	TempResult        []byte         `json:"temp_result" gorm:"column:temp_result"`
	CostResult        []byte         `json:"cost_result" gorm:"column:cost_result"`
	TotalCost         float32        `json:"total_cost" gorm:"column:total_cost"`
	CreatedAt         Timestamp      `json:"created_at" gorm:"column:created_at;type:bigint"` // 中间变量值添加时间
	UpdatedAt         Timestamp      `json:"updated_at" gorm:"column:updated_at;type:bigint"` // 中间变量值更新时间
	DeletedAt         Timestamp      `json:"deleted_at" gorm:"column:deleted_at;type:bigint"` // 中间变量值删除时间
}

func (DesignResults) TableName() string {
	return "design_results"
}
