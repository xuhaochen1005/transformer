// Package notify 通知服务
package message

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	client "transformer_mz/libs/message"
	"transformer_mz/utils"
)

// Notify 发送消息
func Notify(c *gin.Context) {
	var users []*model.User
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	messageData := &struct {
		SendTo  []int  `json:"send_to"`      // 消息的接收者
		Title   string `validate:"required"` // 消息的标题
		Content string `validate:"required"` // 消息的内容
	}{}
	err = errors.New(c.BindJSON(messageData), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if len(messageData.SendTo) > 0 {
		err := connector.DataSource.Transaction(func(tx *gorm.DB) error {
			for _, sendTo := range messageData.SendTo {
				if sendTo != 0 {
					err = client.SendMessage(self.ID, sendTo, messageData.Title, messageData.Content)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	} else {
		err := connector.DataSource.Transaction(func(tx *gorm.DB) error {
			err = errors.New(connector.DataSource.Find(&users).Error, "消息通知发送失败，请稍后再试")
			if err != nil {
				return err
			}
			for _, user := range users {
				if user.ID != self.ID {
					err = client.SendMessage(self.ID, user.ID, messageData.Title, messageData.Content)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}

	c.Set("response", &model.Response{Spec: "消息通知发送成功"})
}
