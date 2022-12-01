// Package notify 消息模板管理服务
package message

import (
	"net/http"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"

	"github.com/gin-gonic/gin"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/utils"
)

// CreateMessageTemplate 添加消息模板
func CreateMessageTemplate(c *gin.Context) {
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = permission.Check(c, resource.Message, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	newMessageTemplate := &model.MessageTemplate{}
	err = errors.New(c.BindJSON(newMessageTemplate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	newMessageTemplate.CreateBy = self.ID
	err = errors.New(connector.DataSource.Create(newMessageTemplate).Error, "消息模板添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "消息模板添加成功"})
}

// UpdateMessageTemplate 更新消息模板
func UpdateMessageTemplate(c *gin.Context) {
	var messageTemplate model.MessageTemplate
	err := permission.Check(c, resource.Message, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&messageTemplate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.MessageTemplate{}).Where("id=?", id).
		Updates(&messageTemplate).Error, "消息模板更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "消息模板更新成功"})
}

// DeleteMessageTemplate 删除消息模板
func DeleteMessageTemplate(c *gin.Context) {
	err := permission.Check(c, resource.Message, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.MessageTemplate{}).Error,
		"消息模板删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "消息模板删除成功"})
}

// GetSpecMessageTemplate 查找特定消息模板
func GetSpecMessageTemplate(c *gin.Context) {
	var messageTemplate model.MessageTemplate
	err := permission.Check(c, resource.Message, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&messageTemplate).Error,
		"消息模板查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: messageTemplate})
}

// FindMessageTemplates 查找消息模板
func FindMessageTemplates(c *gin.Context) {
	var messageTemplates []model.MessageTemplate
	err := permission.Check(c, resource.Message, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreateByUser"), &messageTemplates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: messageTemplates})
}
