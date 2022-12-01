// Package check 复核管理
package check

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"transformer_mz/libs/message"
	"transformer_mz/libs/message/template"

	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/role"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
)

// UpdateDesignTaskStatus 复核设计任务的结果，状态修改只能为复核通过或者不接受
func UpdateDesignTaskStatus(c *gin.Context) {

	var (
		designTask    model.DesignTask
		designProject model.DesignProject
	)

	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	// 判断总工
	hasRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
	}
	if !hasRole {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	req := &struct {
		ApproveInstanceStatus model.ApproveStatus
		Data                  []byte
		Status                model.ApproveNodeStatus
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

	switch req.Status {
	case model.ApproveNodeAccepted:
		err = errors.New(db.Model(&designTask).Where("id=?", id.ID).
			Update("TaskStatus", model.DesignTaskChecked).Error, "审核操作失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = errors.New(db.Model(&model.DesignProject{}).Where("id = ?", designTask.DesignProjectID).
			Update("ProjectStatus", model.DesignProjectChecked).Error, "内部服务出错，请联系开发人员处理")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = errors.New(db.Model(&model.DesignResults{}).Where("design_task_id = ?", id.ID).
			Update("ResultStatus", model.ResultStatusFinished).Error, "内部服务出错，请联系开发人员处理")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = message.SendMessageByTemplate(template.DesignTaskChecked, self.ID,
			[]int{designProject.Creator, designTask.Designer},
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
		approval := designTask.Approve
		if designTask.ApproveByCreator {
			approval = designTask.Creator
		}
		c.Set("response", &model.Response{Spec: struct {
			ApproveInstanceStatus model.ApproveStatus
			Data                  []byte
			Approval              int
		}{
			ApproveInstanceStatus: model.ApproveInAction,
			Data:                  req.Data,
			Approval:              approval,
		}})
		return
	case model.ApproveNodeRejected:
		err = errors.New(db.Model(&designTask).Where("id=?", designTask.ID).
			Select([]string{"TaskStatus", "ResultProgress", "Comment"}).
			Updates(&model.DesignTask{
				TaskStatus:     model.DesignTaskCanceled, // 审批失败会回到头节点重新计算
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
		err = message.SendMessageByTemplate(template.DesignTaskCheckUnaccepted, self.ID,
			[]int{designProject.Creator, designTask.Designer},
			[]string{designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		c.Set("response", &model.Response{Spec: struct {
			ApproveInstanceStatus model.ApproveStatus
			Data                  []byte
			Approval              int
		}{
			ApproveInstanceStatus: model.ApproveInAction,
			Data:                  req.Data,
			Approval:              designTask.Designer,
		}})
	default:
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("内部服务出错，请联系开发人员处理")})
		return
	}

}
