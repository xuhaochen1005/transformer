// Package design ai 调用远程API 进行智能设计与计算优化
package design

import (
	"net/http"
	"strconv"

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

// 调用算法接口，发起AI计算任务
/**todo
1 、调用接口，发起任务，
2、更新任务状态
3、调用通知模块
*/
func StartAIDesign(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	newDesignTask := &model.DesignTask{Creator: self.ID, TaskStatus: model.DesignTaskCreated}
	err = errors.New(c.BindJSON(newDesignTask), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newDesignTask).Error, "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = message.SendMessageByTemplate(template.DesignTaskSubmitted, self.ID, []int{newDesignTask.Designer},
		[]string{newDesignTask.Name, strconv.Itoa(newDesignTask.ID)})
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "设计任务添加成功"})
}

// CancelDesignTaskStatus 取消特定的设计任务
/***
todo
1 、调用接口，取消任务，
2、更新任务状态
3、调用通知模块
*/
func CanceAIDesign(c *gin.Context) {
	var designTask model.DesignTask
	err := permission.Check(c, resource.DesignTask, action.Cancel)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&designTask).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	switch self.ID {
	case designTask.Creator:
		if designTask.TaskStatus > model.DesignTaskStarted {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError,
				Err: errors.New("设计任务已开始，不能取消此任务")})
			return
		}
	default:
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("内部服务出错，请联系开发人员处理")})
		return
	}
	err = errors.New(connector.DataSource.Model(&designTask).Where("id=?", id).
		Update("status", model.DesignTaskCanceled).Error, "设计任务取消失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "设计任务取消成功"})
}

// AI设计算法完成后调用hook函数，更新状态，发消息
/***
todo
1、更新任务状态
2、触发消息
*/
func CreateAIDesignHook(c *gin.Context) {
	var designTask model.DesignTask
	err := permission.Check(c, resource.DesignTask, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&designTask).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	switch self.ID {
	case designTask.Creator:
	case designTask.Designer:
		if designTask.TaskStatus < model.DesignTaskCreated {
			c.Set("response", &model.Response{Spec: nil})
			return
		}
	case designTask.Reviewer:
		if designTask.TaskStatus < model.DesignTaskFinished {
			c.Set("response", &model.Response{Spec: nil})
			return
		}
	case designTask.Approve:
		if designTask.TaskStatus < model.DesignTaskReviewed {
			c.Set("response", &model.Response{Spec: nil})
			return
		}
	default:
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("内部服务出错，请联系开发人员处理")})
		return
	}
	c.Set("response", &model.Response{Spec: designTask})
}
