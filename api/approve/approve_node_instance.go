// Package approve 审批管理
package approve

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/httpclient"
	"transformer_mz/libs/log"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// UpdateApproveNodeInstance 更新审批节点实例
func UpdateApproveNodeInstance(c *gin.Context) {
	var (
		resp                       io.ReadCloser
		approveInstance            model.ApproveInstance
		approveNode                model.ApproveNode
		approveNodeInstance        model.ApproveNodeInstance
		currentApproveNodeInstance model.ApproveNodeInstance
		self                       *model.User
		err                        error
	)
	err = permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err = utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&approveNodeInstance), "请求数据有误，请联系开发人员进行处理")
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
	err = errors.New(db.Where("id=?", approveNodeInstance.ApproveInstanceID).
		First(&approveInstance).Error, "审批节点实例更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(db.Where("id=?", approveNodeInstance.ApproveNodeID).
		First(&approveNode).Error, "审批节点实例更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(db.Where("id=?", approveNodeInstance.ID).
		First(&currentApproveNodeInstance).Error, "审批节点实例更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}

	// 如果原节点已经被拒绝了，那么此时是恢复审批流操作
	if currentApproveNodeInstance.Status == model.ApproveNodeRejected &&
		approveInstance.Status != model.ApproveClosed {
		// 头节点
		headApproveNode := model.ApproveNode{}
		err = errors.New(db.Where("approve_id = ? and pre = 0", approveInstance.ApproveID).
			First(&headApproveNode).Error, "审批节点实例更新失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		nextApproveNodeInstance := &model.ApproveNodeInstance{
			ApproveNodeID:     headApproveNode.ID,
			ApproveInstanceID: approveNodeInstance.ApproveInstanceID,
			Approval:          self.ID,
			Status:            model.ApproveNodeInAction,
			Data:              approveNodeInstance.Data,
			Pre:               approveNodeInstance.ID,
			Next:              0,
		}
		err = errors.New(db.Create(nextApproveNodeInstance).Error, "审批节点创建失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		approveNodeInstance.Next = nextApproveNodeInstance.ID
		err = errors.New(db.Model(&approveNodeInstance).Where("id = ?", approveNodeInstance.ID).
			Select("Status", "Next", "Comment").Updates(&approveNodeInstance).Error, "审批节点保存失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		c.Set("response", &model.Response{Spec: "审批节点实例更新成功"})
		return
	}

	switch approveNodeInstance.Status {
	case model.ApproveNodeAccepted, model.ApproveNodeRejected:
		client := httpclient.NewClient(c.GetHeader("Token"))
		resp, err = client.PatchSpec(approveNode.API, struct {
			ID                    int                     `json:"id"`
			ApproveInstanceID     int                     `json:"approveInstanceID"`
			ApproveNodeInstanceID int                     `json:"approveNodeInstanceID"`
			Data                  []byte                  `json:"data"`
			Status                model.ApproveNodeStatus `json:"status"`
		}{
			ID:                    approveNodeInstance.ID,
			ApproveInstanceID:     approveInstance.ID,
			ApproveNodeInstanceID: approveNodeInstance.ID,
			Data:                  approveNodeInstance.Data,
			Status:                approveNodeInstance.Status,
		})
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		respData := &struct {
			Code int    `json:"code"`
			Err  string `json:"error"`
			Spec struct {
				ApproveInstanceStatus model.ApproveStatus `json:"approve_instance_status"`
				Data                  []byte              `json:"data"`
				Approval              int                 `json:"approval"`
			} `json:"spec"`
		}{}
		err = errors.New(json.NewDecoder(resp).Decode(respData), "响应数据格式有误")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		if respData.Code != http.StatusOK {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New(respData.Err)})
			return
		}
		err = errors.New(db.Where("id=?", approveNodeInstance.ApproveInstanceID).
			First(&approveInstance).Error, "审批节点实例更新失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		nextApproveNodeInstance := &model.ApproveNodeInstance{
			ApproveNodeID:     approveNode.Next,
			ApproveInstanceID: approveNodeInstance.ApproveInstanceID,
			Approval:          respData.Spec.Approval,
			Status:            model.ApproveNodeInAction,
			Data:              respData.Spec.Data,
			Pre:               approveNodeInstance.ID,
			Next:              0,
		}
		if approveNodeInstance.Status == model.ApproveNodeRejected {
			headApproveNode := &model.ApproveNode{
				ApproveID: approveNode.ApproveID,
			}
			err = errors.New(db.Where("pre = ?", 0).First(&headApproveNode).Error, "审批节点创建失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
				return
			}
			nextApproveNodeInstance.ApproveNodeID = headApproveNode.ID
		}
		// 审批单状态没通过
		if respData.Spec.ApproveInstanceStatus != model.ApproveAccepted {
			// 如果状态是通过且不是底部节点了或审核单状态非审批中，则不创建下一个节点
			if approveNode.Next != 0 && approveNodeInstance.Status == model.ApproveNodeAccepted {
				err = errors.New(db.Create(nextApproveNodeInstance).Error, "审批节点创建失败，请稍后再试")
				if err != nil {
					c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
					return
				}
				approveNodeInstance.Next = nextApproveNodeInstance.ID
			}
		}

		// 最底部节点
		if approveNode.Next == 0 {
			if approveNodeInstance.Status == model.ApproveNodeAccepted {
				err = errors.New(db.Model(&approveInstance).Where("id = ?", approveNodeInstance.ApproveInstanceID).Update("Status", model.ApproveAccepted).Error, "审批节点失败，请稍后再试")
				if err != nil {
					c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
					return
				}
			}
		}
		err = errors.New(db.Model(&approveNodeInstance).Where("id = ?", approveNodeInstance.ID).
			Select("Status", "Next", "Comment").Updates(&approveNodeInstance).Error, "审批节点保存失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		c.Set("response", &model.Response{Spec: "审批节点实例更新成功"})
		return
	case model.ApproveNodeClosed:
		if approveNodeInstance.Status == model.ApproveNodeClosed {
			err = errors.New(db.Model(&model.ApproveInstance{}).
				Where("id=?", approveNodeInstance.ApproveInstanceID).
				Update("status", model.ApproveClosed).Error,
				"审批实例更新失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
		}
		err = errors.New(db.Model(&model.ApproveNodeInstance{}).Where("id=?", id).
			Updates(&approveNodeInstance).Error,
			"审批节点实例更新失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		err = errors.New(db.Model(&approveInstance).Where("id = ?", approveNodeInstance.ApproveInstanceID).Update("Status", model.ApproveClosed).Error, "审批节点失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
			return
		}
		c.Set("response", &model.Response{Spec: "审批节点实例更新成功"})
		return
	default:
		err = errors.New("请求数据有误，请联系开发人员进行处理")
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
}

// GetSpecApproveNodeInstance 查找特定的审批节点实例
func GetSpecApproveNodeInstance(c *gin.Context) {
	var approveNodeInstance model.ApproveNodeInstance
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&approveNodeInstance).Error,
		"审批节点查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: approveNodeInstance})
}

// FindApproveNodeInstances 查找所有的审批节点实例
func FindApproveNodeInstances(c *gin.Context) {
	var approveNodeInstances []model.ApproveNodeInstance
	total, err := connector.GetRecords(c, connector.DataSource, &approveNodeInstances)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: approveNodeInstances})
}
