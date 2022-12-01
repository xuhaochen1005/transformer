// Package approve 审批管理
package approve

import (
	"net/http"
	"strings"
	"transformer_mz/libs/log"
	"transformer_mz/libs/permission/role"

	"gorm.io/gorm/clause"

	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
)

// CreateApproveInstance 添加审批实例
func CreateApproveInstance(c *gin.Context) {
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	approveInstance := &model.ApproveInstance{}
	err = errors.New(c.BindJSON(approveInstance), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	approve := &model.Approve{}
	err = errors.New(connector.DataSource.Where("id=?", approveInstance.ApproveID).First(approve).Error,
		"审批查询失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	approveNode := &model.ApproveNode{}
	err = errors.New(connector.DataSource.Where("id=?", approve.Head).First(approveNode).Error,
		"审批节点查询失败，请稍后再试")
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
	approveInstance.Applicant = self.ID
	// approveInstance.Participants = []byte("[" + strconv.Itoa(self.ID) + "]")
	approveInstance.Status = model.ApproveInAction
	err = errors.New(db.Create(approveInstance).Error, "审批发起失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	approveNodeInstance := &model.ApproveNodeInstance{
		ApproveNodeID:     approveNode.ID,
		ApproveInstanceID: approveInstance.ID,
		Approval:          self.ID,
		Data:              approveInstance.Data,
		Status:            model.ApproveNodeInAction,
	}
	err = errors.New(db.Create(approveNodeInstance).Error, "审批发起失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(db.Model(approveInstance).Update("head", approveNodeInstance.ID).Error, "审批发起失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "审批发起成功"})
}

// GetSpecApproveInstance 查找特定的审批实例
func GetSpecApproveInstance(c *gin.Context) {
	var approveInstance model.ApproveInstance
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&approveInstance).Error,
		"审批实例查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: approveInstance})
}

// FindApproveInstances 查找所有的审批实例
func FindApproveInstances(c *gin.Context) {
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	var approveInstances []model.ApproveInstance
	dbClause := clause.Select{
		Expression: clause.Expr{SQL: strings.Join([]string{
			"approve_instance.*",
			"JSON_EXTRACT(approve_instance.data,'$.project_name') as project_name",
			"JSON_EXTRACT(approve_instance.data,'$.serial_code') as serial_code"}, ","),
			WithoutParentheses: true},
	}
	db := connector.DataSource.Clauses(dbClause)
	// 判断总工，总工可以看见所有复核
	hasRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
	}
	if !hasRole {
		db = db.Where("id in (?)", connector.DataSource.Model(&model.ApproveNodeInstance{}).Select("approve_instance_id").Where("approval = ?", self.ID))
	}
	projectName := c.DefaultQuery("project_name", "")
	if projectName != "" {
		db = db.Where("JSON_EXTRACT(approve_instance.data,'$.project_name') like ?", "%"+projectName+"%")
	}
	serialCode := c.DefaultQuery("serial_code", "")
	if serialCode != "" {
		db = db.Where("JSON_EXTRACT(approve_instance.data,'$.serial_code') = ?", serialCode)
	}
	db = db.Preload("Approve").
		Preload("ApplicantUser").
		Preload("ApproveNodeInstances").
		Preload("ApproveNodeInstances.ApprovalUser").
		Preload("ApproveNodeInstances.ApproveNode")

	total, err := connector.GetRecords(c, db, &approveInstances)

	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: approveInstances})
}
