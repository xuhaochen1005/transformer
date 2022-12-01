package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
)

func GetIdFromRequest(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	return id, errors.New(err, "请求地址有误，请联系开发人员进行处理")
}

func GetUserFromRequest(c *gin.Context) (*model.User, error) {
	userIf, ok := c.Get("user")
	if !ok {
		return nil, errors.New("无法从请求中获取到用户信息", "内部服务出错，请联系开发人员进行处理")
	}
	return userIf.(*model.User), nil
}

func GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := errors.New(connector.DataSource.Where("id=?", id).First(&user).Error, "用户信息查询失败，请稍后再试")
	return &user, err
}
