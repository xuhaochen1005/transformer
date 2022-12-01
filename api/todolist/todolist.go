package todolist

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// FindTodoLists 查找文档信息
func FindTodoLists(c *gin.Context) {
	var todoLists []model.TodoList

	err := permission.Check(c, resource.Dashboard, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Where("created_by = ?", self.ID), &todoLists)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: todoLists})
}

// UpdateTodoList 更新待完成列表
func UpdateTodoList(c *gin.Context) {
	var todoList model.TodoList
	err := permission.Check(c, resource.Dashboard, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&todoList), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	db := connector.DataSource.Model(&model.TodoList{}).Where("id=?", id).Updates(&todoList)
	if todoList.CheckedAt != nil {
		todoList.CheckedAt = &model.Timestamp{Time: time.Now()}
	} else {
		db = db.Update("checked_at", gorm.Expr("null"))
	}
	err = errors.New(db.Error, "待完成列表更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "待完成列表更新成功"})
}

// CreateTodoList 创建待更新完成列表
func CreateTodoList(c *gin.Context) {
	var todoList model.TodoList
	err := permission.Check(c, resource.Dashboard, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&todoList), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	todoList.CreatedBy = self.ID
	err = errors.New(connector.DataSource.Model(&model.TodoList{}).Create(&todoList).Error,
		"待完成列表创建失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "待完成列表创建成功"})

}

// DeleteTodoList 删除待完成列表
func DeleteTodoList(c *gin.Context) {
	err := permission.Check(c, resource.Dashboard, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.TodoList{}).Error,
		"待完成列表删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "待完成列表删除成功"})
}
