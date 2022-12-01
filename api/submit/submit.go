package submit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/message"
	"transformer_mz/libs/message/template"

	"github.com/gin-gonic/gin"
)

// UpdateDesignTaskStatus 提交设计任务的结果，状态修改只能为通过
func UpdateDesignTaskStatus(c *gin.Context) {
	var (
		designTask    model.DesignTask
		designProject model.DesignProject
	)
	req := &struct {
		ApproveInstanceID     int                     `json:"approveInstanceID"`
		ApproveNodeInstanceID int                     `json:"approveNodeInstanceID"`
		Data                  []byte                  `json:"data"`
		Status                model.ApproveNodeStatus `json:"status"`
	}{}
	err := errors.New(c.BindJSON(&req), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id := &struct {
		ID int `json:"id"`
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
	err = errors.New(connector.DataSource.Where("id = ?", designTask.DesignProjectID).First(&designProject).Error, "设计任务查询失败，请稍后再试")
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
		err = message.SendMessageByTemplate(template.DesignTaskSuccess, designTask.Designer,
			[]int{designTask.Designer},
			[]string{designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = errors.New(db.Create(&model.TodoList{
			Title:     fmt.Sprintf("选择 %s 设计结果", designProject.ProjectName),
			ReferType: "project",
			ReferID:   designProject.ID,
			CreatedBy: designProject.Designer,
		}).Error, "内部服务出错，请联系开发人员处理")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = message.SendMessageByTemplate(template.DesignTaskNeedReview, designTask.Designer,
			[]int{designTask.Reviewer},
			[]string{designTask.Name, designProject.SerialCode, strconv.Itoa(designTask.ID)}, db)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = errors.New(db.Create(&model.TodoList{
			Title:     fmt.Sprintf("复核设计任务 %s ", designProject.ProjectName),
			ReferType: "project",
			ReferID:   designProject.ID,
			CreatedBy: designTask.Reviewer,
		}).Error, "内部服务出错，请联系开发人员处理")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = errors.New(db.Model(&model.DesignTask{}).Where("id = ?", designTask.ID).
			Update("TaskStatus", model.DesignTaskFinished).Error, "内部服务出错，请联系开发人员处理")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		err = errors.New(db.Model(&model.DesignProject{}).Where("id = ?", designTask.DesignProjectID).
			Update("ProjectStatus", model.DesignProjectFinished).Error, "内部服务出错，请联系开发人员处理")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		c.Set("response", &model.Response{Spec: struct {
			Data     []byte
			Approval int
		}{
			Data:     req.Data,
			Approval: designTask.Reviewer,
		}})
		return
	case model.ApproveNodeRejected:
		err = errors.New(message.SendMessageByTemplate(template.DesignTaskFailed, designTask.Designer, []int{designTask.Designer},
			[]string{designTask.Name, designProject.SerialCode}, db), "发送设计任务计算失败消息失败")
		if err != nil {
			log.Println(err)
			return
		}
		err = errors.New(db.Create(&model.TodoList{
			Title:     fmt.Sprintf("修改设计任务 %s ", designProject.ProjectName),
			ReferType: "project",
			ReferID:   designProject.ID,
			CreatedBy: designTask.Designer,
		}).Error, "内部服务出错，请联系开发人员处理")
		err = errors.New(db.Model(&designTask).Where("id=?", designTask.ID).
			Updates(&model.DesignTask{
				TaskStatus:     model.DesignTaskCanceled,
				ResultProgress: 0,
				Comment:        "",
			}).Error, "审核操作失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		if err != nil {
			log.Println(err)
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

	default:
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("内部服务出错，请联系开发人员处理")})
		return
	}
}
