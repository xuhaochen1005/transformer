// Package message 消息通知发送接口
package message

import (
	"gorm.io/gorm"
	"strconv"

	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	template2 "transformer_mz/libs/message/template"
	"transformer_mz/utils"
)

var templates = make(map[string]*model.MessageTemplate)

func SendMessage(sendBy, sendTo int, title, content string, db ...*gorm.DB) error {
	ct := connector.DataSource
	if len(db) != 0 {
		ct = db[0]
	}
	return errors.New(ct.Create(&model.Message{
		Title:   title,
		Content: content,
		SendBy:  sendBy,
		SendTo:  sendTo,
		Status:  model.UnRead,
	}).Error, "消息通知发送失败，请稍后再试")
}

func SendMessageByTemplate(templateName string, sendBy int, sendToList []int, params []string, db ...*gorm.DB) error {
	template, ok := templates[templateName]
	if !ok {
		return errors.New("名称为" + templateName + "的模板不存在")
	}
	sendByUser, err := utils.GetUserByID(sendBy)
	if err != nil {
		return err
	}
	data := make(map[string]string)
	data["send_by"] = sendByUser.UserNameZh
	for i, param := range params {
		data[strconv.Itoa(i+1)] = param
	}
	for _, sendTo := range sendToList {
		err = SendMessage(sendBy, sendTo, template.Title, template2.Process(template.Content, data), db...)
		if err != nil {
			return err
		}
	}
	return nil
}

func Init() error {
	var messageTemplates []*model.MessageTemplate
	err := errors.New(connector.DataSource.Find(&messageTemplates).Error)
	if err != nil {
		return err
	}
	for _, messageTemplate := range messageTemplates {
		templates[messageTemplate.Name] = messageTemplate
	}
	return nil
}
