package feedback

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/libs/permission/role"
	"transformer_mz/utils"
)

// FindFeedback 查询接收到的通知消息
func FindFeedback(c *gin.Context) {
	var feedbacks []model.Feedback
	dataSource := connector.DataSource.Preload("CreatorUser")
	err := permission.Check(c, resource.Dashboard, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, dataSource, &feedbacks)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: feedbacks})
}

// GetSpecFeedback 查询特定的反馈信息
func GetSpecFeedback(c *gin.Context) {
	var feedback model.Feedback
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&feedback).Preload("CreatorUser").Error,
		"消息查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: feedback})
}

// CreateFeedback 创建待更新完成列表
func CreateFeedback(c *gin.Context) {
	var feedback model.Feedback
	err := permission.Check(c, resource.Documents, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&feedback), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	feedback.Creator = self.ID
	err = errors.New(connector.DataSource.Model(&model.Feedback{}).Create(&feedback).Error,
		"反馈消息创建失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "反馈消息创建成功"})

}

// DeleteFeedback 删除消息
func DeleteFeedback(c *gin.Context) {
	var feedback model.Feedback
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
	err = errors.New(connector.DataSource.Where("id=?", id).Find(&feedback).Error,
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
	if feedback.Creator != self.ID && !hasRootRole {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("您没有权限删除他人的消息")})
		return
	}
	err = errors.New(connector.DataSource.Delete(&feedback).Error, "消息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "消息删除成功"})
}
