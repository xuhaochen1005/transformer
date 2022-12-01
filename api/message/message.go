// Package notify 通知服务
package message

import (
	"net/http"
	"transformer_mz/libs/permission/role"

	"github.com/gin-gonic/gin"
	"strconv"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// FindMessages 查询接收到的通知消息
func FindMessages(c *gin.Context) {
	var messages []model.Message
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	// 判断是否只查询自己的消息
	selfMsg, err := strconv.Atoi(c.DefaultQuery("self", "1"))
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	dataSource := connector.DataSource.Preload("SendByUser").Preload("SendToUser").Preload("SendByUser.Department")
	if selfMsg == 2 {
		err := permission.Check(c, resource.Message, action.Read)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
			return
		}
	} else {
		dataSource = dataSource.Where("send_to=?", self.ID)
	}
	total, err := connector.GetRecords(c, dataSource, &messages)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: messages})
}

// GetSpecMessage 查询特定的通知消息
func GetSpecMessage(c *gin.Context) {
	var message model.Message
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&message).Error,
		"消息查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if message.SendTo != self.ID {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("不能查询他人接收到的消息", "请求有误，请联系开发人员进行处理")})
		return
	}
	c.Set("response", &model.Response{Spec: message})
}

// UpdateMessageStatus 更新消息状态为已读或未读
func UpdateMessageStatus(c *gin.Context) {
	var message model.Message
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
	err = errors.New(connector.DataSource.Where("id=?", id).Find(&message).Error, "消息查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&message), "消息查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	permCheckRes := false
	err = permission.Check(c, resource.Message, action.Read)
	if err == nil {
		permCheckRes = true
	}
	hasRootRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.Root))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
		return
	}
	if message.SendTo != self.ID && !hasRootRole {
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError,
				Err: errors.New("不能更新他人接收到的消息的状态")})
			return
		}
	}
	if message.Status == model.HasRead && permCheckRes {
		err = errors.New(connector.DataSource.Model(&message).Update("status", model.HasRead).Error,
			"消息已读设置失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	} else {
		err = errors.New(connector.DataSource.Model(&message).Updates(&message).Error,
			"消息未读设置失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	c.Set("response", &model.Response{Spec: "消息状态更新成功"})
}

func UpdateAllMessageRead(c *gin.Context) {
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if err = connector.DataSource.Model(&model.Message{}).Where("send_to = ? ", self.ID).Update("status", model.HasRead).Error; err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "更新成功"})
}

// DeleteMessage 删除消息
func DeleteMessage(c *gin.Context) {
	var message model.Message
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
	err = errors.New(connector.DataSource.Where("id=?", id).Find(&message).Error,
		"消息查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	hasRootRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.Root))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
		return
	}
	if message.SendTo != self.ID && !hasRootRole {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("您没有权限删除他人的消息")})
		return
	}
	err = errors.New(connector.DataSource.Delete(&message).Error, "消息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "消息删除成功"})
}
