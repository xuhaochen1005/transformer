// Package approve 审批管理
package approve

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"transformer_mz/libs/permission/role"

	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/message"
	"transformer_mz/libs/message/template"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
)

// UpdateDesignTaskStatus 审核设计任务的结果，状态修改只能为审核通过或者不接受
func UpdateDesignTaskStatus(c *gin.Context) {
	var (
		designTask    model.DesignTask
		designProject model.DesignProject
		err           error
		self          *model.User
	)

	err = permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}

	self, err = utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	req := &struct {
		ApproveInstanceID     int                     `json:"approveInstanceID"`
		ApproveNodeInstanceID int                     `json:"approveNodeInstanceID"`
		Data                  []byte                  `json:"data"`
		Status                model.ApproveNodeStatus `json:"status"`
	}{}
	err = errors.New(c.BindJSON(&req), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id := &struct {
		ID      int    `json:"id"`
		Comment string `json:"comment"`
	}{}
	err = errors.New(json.Unmarshal(req.Data, id), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id.ID).First(&designTask).Error, "设计任务查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", designTask.DesignProjectID).First(&designProject).Error, "设计任务查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	db := connector.DataSource.Begin()
	defer func() {
		if err != nil {
			e := errors.New(db.Rollback().Error)
			if e != nil {
				log.Println(e)
			}
		} else {
			e := errors.New(db.Commit().Error)
			if e != nil {
				log.Println(e)
			}
		}
	}()
	if self.ID == designTask.Approve || self.ID == designProject.Creator {
		switch req.Status {
		case model.ApproveNodeAccepted:
			taskStatus := model.DesignTaskApproved
			projectStatus := model.DesignProjectApproved
			// 如果是发起者代为审批则直接通过
			if designTask.ApproveByCreator {
				taskStatus = model.DesignTaskChecked
				projectStatus = model.DesignProjectChecked
				// 审核结果通过
				err = errors.New(db.Model(&model.DesignResults{}).Where("design_task_id = ?", id.ID).
					Update("ResultStatus", model.ResultStatusFinished).Error, "内部服务出错，请联系开发人员处理")
				if err != nil {
					c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
					return
				}
			}
			err = errors.New(db.Model(&designTask).Where("id=?", id.ID).
				Update("TaskStatus", taskStatus).Error, "审核操作失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = errors.New(db.Model(&model.DesignProject{}).Where("id = ?", designTask.DesignProjectID).
				Update("ProjectStatus", projectStatus).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			approveStatus := model.ApproveInAction
			if designTask.ApproveByCreator {
				approveStatus = model.ApproveAccepted
				// 直接将审核单状态改为通过
				err = errors.New(db.Model(&model.ApproveInstance{}).Where("id = ?", req.ApproveInstanceID).
					Update("Status", approveStatus).Error, "内部服务出错，请联系开发人员处理")
				if err != nil {
					c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
					return
				}
			}
			err = message.SendMessageByTemplate(template.DesignTaskApproved, self.ID,
				[]int{designProject.Creator, designTask.Designer},
				[]string{designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			designer := &model.User{}
			err = errors.New(db.Where("id=?", designTask.Designer).First(designer).Error,
				"内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			// 找出总工的用户id
			var chiefEngineerUserIDs []int
			err = errors.New(connector.DataSource.Model(&model.UserRole{}).
				Where("role_name = ? ", role.ChiefEngineer).
				Joins("left join user as u on u.name = user_role.username").Pluck("u.id", &chiefEngineerUserIDs).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			if len(chiefEngineerUserIDs) == 0 {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New(nil, "必须有总工角色用户才能进行才一步操作")})
				return
			}
			for _, uid := range chiefEngineerUserIDs {
				err = errors.New(db.Create(&model.TodoList{
					Title:     fmt.Sprintf("审批 %s 设计任务", designProject.ProjectName),
					ReferType: "project",
					ReferID:   designProject.ID,
					CreatedBy: uid,
				}).Error, "内部服务出错，请联系开发人员处理")
				if err != nil {
					c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
					return
				}
			}
			err = message.SendMessageByTemplate(template.DesignTaskCheckForChiefEngineer, self.ID,
				chiefEngineerUserIDs,
				[]string{designer.UserNameZh, designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}

			c.Set("response", &model.Response{Spec: struct {
				ApproveInstanceStatus model.ApproveStatus `json:"approve_instance_status"`
				Data                  []byte              `json:"data"`
				Approval              int                 `json:"approval"`
			}{
				ApproveInstanceStatus: approveStatus,
				Data:                  req.Data,
				Approval:              chiefEngineerUserIDs[0],
			}})
			return
		case model.ApproveNodeRejected:
			err = errors.New(db.Model(&designTask).Where("id=?", designTask.ID).
				Select([]string{"TaskStatus", "ResultProgress", "Comment"}).
				Updates(&model.DesignTask{
					TaskStatus:     model.DesignTaskCanceled,
					ResultProgress: 0,
					Comment:        id.Comment,
				}).Error, "审核操作失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = errors.New(db.Model(&model.DesignProject{}).Where("id = ?", designTask.DesignProjectID).
				Update("ProjectStatus", model.DesignProjectCreated).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = errors.New(db.Model(&model.DesignResults{}).Where("design_task_id = ?", designTask.ID).
				Update("ResultStatus", model.ResultStatusUnStarted).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = message.SendMessageByTemplate(template.DesignTaskApproveUnaccepted, self.ID,
				[]int{designProject.Creator, designTask.Designer},
				[]string{designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = errors.New(db.Create(&model.TodoList{
				Title:     fmt.Sprintf("修改设计任务 %s ", designProject.ProjectName),
				ReferType: "project",
				ReferID:   designProject.ID,
				CreatedBy: designTask.Designer,
			}).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			c.Set("response", &model.Response{Spec: struct {
				ApproveInstanceStatus model.ApproveStatus `json:"approve_instance_status"`
				Data                  []byte              `json:"data"`
				Approval              int                 `json:"approval"`
			}{
				ApproveInstanceStatus: model.ApproveInAction,
				Data:                  req.Data,
				Approval:              designTask.Designer,
			}})
			return
		default:
			c.Set("response", &model.Response{Code: http.StatusInternalServerError,
				Err: errors.New("内部服务出错，请联系开发人员处理")})
			return
		}
	}
	c.Set("response", &model.Response{Code: http.StatusInternalServerError,
		Err: errors.New("内部服务出错，请联系开发人员处理")})
}

// CreateApprove 添加审批
func CreateApprove(c *gin.Context) {
	var (
		err        error
		newApprove *model.Approve
	)
	err = permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	err = errors.New(c.BindJSON(newApprove), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newApprove).Error, "审批添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "审批添加成功"})
}

// UpdateApprove 更新审批
func UpdateApprove(c *gin.Context) {
	var (
		approve model.Approve
		err     error
		id      int
	)
	err = permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err = utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&approve), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.Approve{}).Where("id=?", id).Updates(&approve).Error,
		"审批更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "审批更新成功"})
}

// DeleteApprove 删除审批
func DeleteApprove(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.Approve{}).Error,
		"审批删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "审批删除成功"})
}

// GetSpecApprove 查找特定的审批
func GetSpecApprove(c *gin.Context) {
	var approve model.Approve
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&approve).Error,
		"审批查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: approve})
}

// FindApproves 查找所有的审批
func FindApproves(c *gin.Context) {
	var approves []model.Approve
	total, err := connector.GetRecords(c, connector.DataSource, &approves)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: approves})
}
