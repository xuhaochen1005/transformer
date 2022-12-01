package template

import "strings"

const (
	DesignTaskCreated               = "design_task_created"
	DesignTaskCanceled              = "design_task_canceled"
	DesignTaskSubmitted             = "design_task_submitted"
	DesignTaskFailed                = "design_task_failed"
	DesignTaskSuccess               = "design_task_success"
	DesignTaskNeedReview            = "design_task_need_review"
	DesignTaskReviewUnaccepted      = "design_task_review_unaccepted"
	DesignTaskReviewedForDesigner   = "design_task_reviewed_for_designer"
	DesignTaskReviewedForApproval   = "design_task_reviewed_for_approval"
	DesignTaskApproveUnaccepted     = "design_task_approve_unaccepted"
	DesignTaskApproved              = "design_task_approved"
	DesignTaskCheckForChiefEngineer = "design_task_check_for_chief_engineer"
	DesignTaskCheckUnaccepted       = "design_task_check_unaccepted"
	DesignTaskChecked               = "design_task_checked"
)

const (
	DesignTaskCreatedContent = `
<div>{{send_by}}给您分配了新的设计任务{{1}}，请及时点击<a href="/#/design/design_projects?code={{2}}">任务地址</a>进行设计处理</div>
`
	DesignTaskCanceledContent = `
<div>{{send_by}}取消了给您分配的设计任务{{1}}，请知悉。可点击<a href="/#/design/design_projects?code={{2}}">任务地址</a>查看详情</div>
`
	DesignTaskSubmittedContent = `
<div>设计任务{{1}}已提交计算，请点击<a href="/#/design/design_check?code={{2}}">任务地址</a>查看详情</div>
`
	DesignTaskFailedContent = `
<div>设计任务{{1}}计算失败，请点击<a href="/#/user/profile?task_code={{2}}">任务地址</a>进行修改</div>`
	DesignTaskSuccessContent = `
<div>设计任务{{1}}计算成功，请点击<a href="/#/design/design_check?code={{2}}">任务地址</a>查看详情以及选择计算结果</div>
`
	DesignTaskNeedReviewContent = `
<div>{{send_by}}完成了设计任务{{1}}，请点击<a href="/#/design/design_check?code={{2}}">任务地址</a>进行设计的复核</div>
`
	DesignTaskReviewUnacceptedContent = `
<div>{{send_by}}没有复核通过您的设计任务{{1}}，请点击<a href="/#/user/profile?task_code=?code={{2}}">任务地址</a>进行修改</div>
`
	DesignTaskReviewedForDesignerContent = `
<div>{{send_by}}复核通过了您的设计任务{{1}}，请点击<a href="/#/design/design_check?code={{2}}">任务地址</a>查看详情</div>
`
	DesignTaskReviewedForApprovalContent = `
<div>{{send_by}}复核通过了{{1}}完成的设计任务{{2}}，请点击<a href="/#/design/design_check?code={{3}}">任务地址</a>进行审批</div>
`
	DesignTaskApproveUnacceptedContent = `
<div>{{send_by}}没有审批通过设计任务{{1}}的复核，请点击<a href="/#/user/profile?task_code=?code={{2}}">任务地址</a>进行修改</div>
`
	DesignTaskApprovedContent = `
<div>{{send_by}}审批通过了设计任务{{1}}，请点击<a href="/#/design/design_check?code={{2}}">任务地址</a>查看详情</div>
`
	DesignTaskCheckForChiefEngineerContent = `
<div>{{send_by}}审批通过了{{1}}完成的设计任务{{2}}，请点击<a href="/#/design/design_check?code={{3}}">任务地址</a>进行审批</div>	`
	DesignTaskCheckUnacceptedContent = `
<div>{{send_by}}没有审批通过设计任务{{1}}的复核，请点击<a href="/#/user/profile?task_code=?code={{2}}">任务地址</a>进行修改</div>
	`
	DesignTaskCheckedContent = `
<div>{{send_by}}审批通过了{{1}}完成的设计任务{{2}}，请点击<a href="/#/design/design_check?code={{2}}">任务地址</a>进行审批</div>
	`
)

func Process(template string, data map[string]string) string {
	state := 0
	key := ""
	content := ""
	start := 0
	end := 0
	pos := 0
	for i, c := range template {
		switch state {
		case 0:
			if c == '{' {
				state = 1
			}
		case 1:
			if c == '{' {
				end = i - 1
				pos = i + 1
				state = 2
			} else {
				state = 0
			}
		case 2:
			switch c {
			case '{':
				end = i - 1
				pos = i + 1
			case '}':
				key = template[pos:i]
				state = 3
			}
		case 3:
			if c == '}' {
				key = strings.TrimSpace(key)
				value, found := data[key]
				if found {
					content += template[start:end] + value
					start = i + 1
				}
			}
			state = 0
		}
	}
	if start < len(template) {
		return content + template[start:]
	}
	return content
}
