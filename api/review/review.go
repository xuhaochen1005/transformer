// Package review 复核管理
package review

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"transformer_mz/libs/message"
	"transformer_mz/libs/message/template"

	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
)

// UpdateDesignTaskStatus 复核设计任务的结果，状态修改只能为复核通过或者不接受
func UpdateDesignTaskStatus(c *gin.Context) {
	var (
		designTask    model.DesignTask
		designProject model.DesignProject
	)
	err := permission.Check(c, resource.DesignTask, action.Review)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
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
	if self.ID == designTask.Reviewer || self.ID == designProject.Creator {
		switch req.Status {
		case model.ApproveNodeAccepted:
			err = errors.New(db.Model(model.DesignTask{}).Where("id=?", id.ID).
				Updates(model.DesignTask{TaskStatus: model.DesignTaskReviewed}).Error, "审核操作失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = errors.New(db.Model(&model.DesignProject{}).Where("id = ?", designTask.DesignProjectID).
				Update("ProjectStatus", model.DesignProjectReviewed).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = message.SendMessageByTemplate(template.DesignTaskReviewedForDesigner, self.ID,
				[]int{designTask.Designer},
				[]string{designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			designer := &model.User{}
			err = errors.New(db.Where("id=?", designTask.Designer).First(designer).Error,
				"设计人员信息查询失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = message.SendMessageByTemplate(template.DesignTaskReviewedForApproval, self.ID,
				[]int{designTask.Approve},
				[]string{designer.UserNameZh, designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			err = errors.New(db.Create(&model.TodoList{
				Title:     fmt.Sprintf("审核设计任务 %s ", designProject.ProjectName),
				ReferType: "project",
				ReferID:   designProject.ID,
				CreatedBy: designTask.Approve,
			}).Error, "内部服务出错，请联系开发人员处理")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
			approval := designTask.Approve
			c.Set("response", &model.Response{Spec: struct {
				Data     []byte
				Approval int
			}{
				Data:     req.Data,
				Approval: approval,
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
			err = message.SendMessageByTemplate(template.DesignTaskReviewUnaccepted, self.ID,
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
		Err: errors.New("您不是指定的审批人员或设计主管")})
}
