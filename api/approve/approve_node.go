// Package approve 审批管理
package approve

import (
	"net/http"
	"transformer_mz/libs/log"

	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
)

// CreateApproveNode 添加审批节点
func CreateApproveNode(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	newApproveNode := &model.ApproveNode{}
	err = errors.New(c.BindJSON(newApproveNode), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
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
	err = errors.New(db.Create(newApproveNode).Error, "审批节点添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if newApproveNode.Pre == 0 {
		err = errors.New(db.Model(&model.Approve{}).Where("id=?", newApproveNode.ApproveID).
			Update("head", newApproveNode.ID).Error)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	} else {
		err = errors.New(db.Model(&model.ApproveNode{}).Where("id=?", newApproveNode.Pre).
			Update("next", newApproveNode.ID).Error)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	if newApproveNode.Next != 0 {
		err = errors.New(db.Model(&model.ApproveNode{}).Where("id=?", newApproveNode.Next).
			Update("pre", newApproveNode.ID).Error)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	c.Set("response", &model.Response{Spec: "审批节点添加成功"})
}

// UpdateApproveNode 更新审批节点
func UpdateApproveNode(c *gin.Context) {
	var approveNode model.ApproveNode
	err := permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&approveNode), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.ApproveNode{}).Where("id=?", id).Updates(&approveNode).Error,
		"审批节点更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "审批节点更新成功"})
}

// DeleteApproveNode 删除审批节点
func DeleteApproveNode(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Approve)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	approveNode := &model.ApproveNode{}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&approveNode).Error,
		"审批节点查找失败，请稍后再试")
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
	err = errors.New(db.Where("id=?", id).Delete(&model.ApproveNode{}).Error,
		"审批节点删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if approveNode.Pre != 0 {
		err = errors.New(db.Model(&model.ApproveNode{}).Where("id=?", approveNode.Pre).
			Update("next", approveNode.Next).Error)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		approveNode = &model.ApproveNode{}
		err = errors.New(connector.DataSource.Where("id=?", approveNode.Pre).First(&approveNode).Error,
			"审批节点查找失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		if approveNode.Pre == 0 {
			err = errors.New(db.Model(&model.Approve{}).Where("id=?", approveNode.ApproveID).
				Update("head", approveNode.ID).Error)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
		}
	}
	if approveNode.Next != 0 {
		err = errors.New(db.Model(&model.ApproveNode{}).Where("id=?", approveNode.Next).
			Update("Pre", approveNode.Pre).Error)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	c.Set("response", &model.Response{Spec: "审批节点删除成功"})
}

// GetSpecApproveNode 查找特定的审批节点
func GetSpecApproveNode(c *gin.Context) {
	var approveNode model.ApproveNode
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&approveNode).Error,
		"审批节点查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: approveNode})
}

// FindApproveNodes 查找所有的审批节点
func FindApproveNodes(c *gin.Context) {
	var approveNodes []model.ApproveNode
	total, err := connector.GetRecords(c, connector.DataSource, &approveNodes)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: approveNodes})
}
