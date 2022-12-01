package role

type Role string

const (
	Root          Role = "root"           // 超级用户
	Demander      Role = "demander"       // 任务需求方
	Designer      Role = "designer"       // 设计者
	Reviewer      Role = "reviewer"       // 复核者
	ChiefEngineer Role = "chief_engineer" // 审批人
)
