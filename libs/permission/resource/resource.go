package resource

import "transformer_mz/libs/permission/action"

const (
	User           = "user"            // 用户
	Approve        = "approve"         // 审批
	Department     = "department"      // 部门
	Documents      = "documents"       // 文档
	PermissionRule = "permission_rule" // 权限配置
	DesignProject  = "design_project"  // 任务书管理
	DesignTask     = "design_task"     // 设计任务
	Message        = "message"         // 消息管理
	Dashboard      = "dashboard"       // 首页
	Libraries      = "libraries"       // 标准库
)

var resourceList = []string{
	User,
	Approve,
	Department,
	PermissionRule,
	DesignProject,
	DesignTask,
	Message,
	Message,
	Libraries,
	Libraries,
}

var ActionInResource = map[string][]string{
	DesignTask: {action.Cancel, action.Design, action.Review, action.Approve},
}

func init() {
	for _, resource := range resourceList {
		ActionInResource[resource] = append([]string{action.Read, action.Write}, ActionInResource[resource]...)
	}
}
