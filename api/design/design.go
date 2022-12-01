// Package design 产品设计流程管理
package design

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	Documents "transformer_mz/api/documents"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/httpclient"
	"transformer_mz/libs/message"
	"transformer_mz/libs/message/template"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/libs/permission/role"
	"transformer_mz/libs/third"
	"transformer_mz/libs/token"
	"transformer_mz/utils"
)

// CreateDesignTask 创建设计任务
func CreateDesignTask(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	var self *model.User
	self, err = utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	newDesignTask := &model.DesignTask{Creator: self.ID, TaskStatus: model.DesignTaskCreated}
	err = errors.New(c.BindJSON(newDesignTask), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	designProject := &model.DesignProject{ID: newDesignTask.DesignProjectID}

	err = errors.New(connector.DataSource.First(&designProject).Error, "查找任务书失败，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if designProject.Designer != self.ID {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New("你不是此任务书的设计人员")})
		return
	}
	existDesignTask := model.DesignTask{}
	err = errors.New(connector.DataSource.Where("design_project_id = ?", designProject.ID).Find(&existDesignTask).Error, "查找任务书失败，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if existDesignTask.ID != 0 {
		err = errors.New("此任务书已发起设计任务")
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	newDesignTask.Creator = self.ID
	newDesignTask.Designer = self.ID
	newDesignTask.Reviewer = designProject.Reviewer
	newDesignTask.Approve = designProject.CheckBy
	newDesignTask.TaskStatus = model.DesignTaskCreated
	newDesignTask.ApproveByCreator = designProject.NeedMasterApprove == 0
	newDesignTask.Name = newDesignTask.Name + "-" + self.UserNameZh + "-" + time.Now().Format("2006-01-02 15:04:05") //任务名为产品名称+用户名+时间戳

	db := connector.DataSource.Begin()
	defer func() {
		log.Println("错误捕捉", err)
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

	err = errors.New(db.Create(newDesignTask).Error, "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	inputData := make(map[string]interface{})
	if err = errors.New(json.Unmarshal(newDesignTask.Input, &inputData), "请求数据有误，请联系开发人员进行处理"); err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	inputData["id"] = newDesignTask.ID
	var inputJsonStr []byte
	inputJsonStr, err = json.Marshal(&inputData)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New(err, "请求数据有误，请联系开发人员进行处理")})
		return
	}
	newDesignTask.Input = inputJsonStr
	if err = errors.New(db.Model(&model.DesignTask{}).Where("id = ?", newDesignTask.ID).Select("Input").Updates(model.DesignTask{Input: newDesignTask.Input}).Error); err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New(err, "请求数据有误，请联系开发人员进行处理")})
		return
	}
	if err = StartCalculatingDesign(db, newDesignTask); err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	var approveInstanceData []byte
	approveInstanceData, err = json.Marshal(map[string]interface{}{
		"id":                newDesignTask.ID,
		"project_name":      designProject.ProjectName,
		"product_name":      designProject.ProductName,
		"product_model":     designProject.ProductModel,
		"serial_code":       designProject.SerialCode,
		"design_project_id": designProject.ID,
		"demander_id":       designProject.Creator, // 设计主管
	})
	if err != nil {
		err = errors.New(err, "设计任务添加失败，请稍后再试")
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	// 开始审批流
	client := httpclient.NewClient(c.GetHeader("Token"))
	var resp io.ReadCloser
	resp, err = client.PostSpec("/v1/approve/approve_instance", &model.ApproveInstance{
		ApproveID: model.ApproveIDDesignTask,
		Head:      model.ApproveNodeDesignTaskHeadID,
		Data:      approveInstanceData,
		Participants: []byte("[" + strings.Join([]string{
			strconv.Itoa(designProject.Designer), strconv.Itoa(designProject.Reviewer),
			strconv.Itoa(designProject.DraftingBy), strconv.Itoa(designProject.CheckBy),
		}, ",") + "]"),
	})
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	respData := &struct {
		Code int
		Err  string
		Spec string
	}{}
	err = errors.New(json.NewDecoder(resp).Decode(respData), "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if respData.Code != http.StatusOK {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New(respData.Err)})
		return
	}
	c.Set("response", &model.Response{Spec: "设计任务添加成功"})
}

// StartCalculatingDesign 调用开始计算api
func StartCalculatingDesign(db *gorm.DB, t *model.DesignTask) error {
	designProject := model.DesignProject{}
	err := errors.New(db.Where("id = ?", t.DesignProjectID).First(&designProject).Error, "审批发起失败，请稍后再试")
	if err != nil {
		return err
	}
	client := httpclient.NewClient("", third.AI_SYSTEM_URL)
	resp, err := client.PostSpec("/transformer/task", t.Input)
	if err != nil {
		return err
	}
	respData := &struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}
	if err = errors.New(json.NewDecoder(resp).Decode(respData), "响应数据格式有误"); err != nil {
		return err
	}
	if respData.Code != http.StatusOK {
		err = errors.New(respData.Msg)
		log.Println(respData.Msg)
		return err
	}
	err = errors.New(db.Model(&model.DesignProject{}).Where("id = ?", t.DesignProjectID).Update("ProjectStatus", model.DesignProjectStarted).Error, "审批发起失败，请稍后再试")
	if err != nil {
		return err
	}
	err = errors.New(db.Model(&model.DesignTask{}).Where("id = ?", t.ID).Update("TaskStatus", model.DesignTaskStarted).Error, "审批发起失败，请稍后再试")
	if err != nil {
		return err
	}
	err = errors.New(db.Model(&model.DesignResults{}).Where("design_task_id = ?", t.ID).Update("ResultStatus", model.ResultStatusCalculating).Error, "审批发起失败，请稍后再试")
	if err != nil {
		return err
	}
	err = message.SendMessageByTemplate(template.DesignTaskSubmitted, t.Creator, []int{t.Designer},
		[]string{t.Name, designProject.SerialCode}, db)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDesignTask 取消特定的设计任务
func UpdateDesignTask(c *gin.Context) {
	var (
		designTask model.DesignTask
		self       *model.User
		id         int
		err        error
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
	id, err = utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&designTask).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&designTask), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}

	switch self.ID {
	case designTask.Creator:
		if designTask.TaskStatus > model.DesignTaskStarted {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError,
				Err: errors.New("设计任务已开始，不能修改此任务")})
			return
		}
	default:
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("您不是这个设计任务的设计人")})
		return
	}
	// 清空设计结果 重新开始计算 调用API成功才能设置状态为开始
	// 创建下一个审批节点
	// 查询此设计任务的当前节点
	var (
		approveInstance     model.ApproveInstance
		approveNodeInstance model.ApproveNodeInstance
	)
	err = errors.New(connector.DataSource.Where("JSON_EXTRACT(data,'$.id') = ?", designTask.ID).First(&approveInstance).Error, "设计任务修改，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("approve_instance_id = ? and next = 0", approveInstance.ID).First(&approveNodeInstance).Error, "设计任务修改，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	// 重新创建头节点
	client := httpclient.NewClient(c.GetHeader("Token"))
	var resp io.ReadCloser
	resp, err = client.PatchSpec("/v1/approve/approve_node_instance/"+strconv.Itoa(approveNodeInstance.ID), &approveNodeInstance)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	respData := &struct {
		Code int    `json:"code"`
		Err  string `json:"err"`
		Spec string `json:"spec"`
	}{}
	err = errors.New(json.NewDecoder(resp).Decode(respData), "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if respData.Code != http.StatusOK {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New(respData.Err)})
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
	designTask.DesignResults = nil
	designTask.TaskStatus = model.DesignTaskStarted
	err = errors.New(db.Model(&model.DesignTask{}).Where("id=?", id).Select("*").
		Updates(&designTask).Error, "设计任务修改失败，请稍后再试")

	//调用API接口进行计算,发送开始计算的信息
	err = errors.New(StartCalculatingDesign(db, &designTask), "调用计算接口失败，请联系开发人员")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "设计任务修改成功"})
}

// GetSpecDesignTask 查询特定的设计任务
func GetSpecDesignTask(c *gin.Context) {
	designTask := &model.DesignTask{}
	err := permission.Check(c, resource.DesignTask, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	db := connector.DataSource
	// 判断总工
	hasChiefEngineerRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if !hasChiefEngineerRole {
		db = db.Where("(design_task.creator=? or design_task.designer=? or design_task.reviewer=? or design_task.approve=?) and design_task.deleted_at IS NULL",
			self.ID, self.ID, self.ID, self.ID)
	}
	err = errors.New(db.Model(&model.DesignTask{}).Where("id = ?", id).
		Preload("DesignProject").
		Preload("FinalDesignResults").
		Preload("FinalDesignResults.DesignProject").
		Preload("FinalDesignResults.DesignProject.CreatorUser").
		Preload("FinalDesignResults.DesignProject.DraftingUser").
		Preload("FinalDesignResults.DesignProject.CreatorUser").
		Preload("FinalDesignResults.DesignProject.ReviewerUser").
		Preload("FinalDesignResults.DesignProject.DesignerUser").
		Preload("FinalDesignResults.DesignProject.CheckUser").
		Preload("FinalDesignResults.DesignTask").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.DraftingUser").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.ReviewerUser").
		Preload("DesignProject.DesignerUser").
		Preload("DesignProject.CheckUser").First(designTask).Error, "内部服务出错，请联系开发人员处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: designTask})
}

// FindDesignTasks 查找设计任务信息
func FindDesignTasks(c *gin.Context) {
	var designTasks []model.DesignTask
	err := permission.Check(c, resource.DesignTask, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	db := connector.DataSource
	// 判断总工
	hasChiefEngineerRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if !hasChiefEngineerRole {
		db = db.Where("(design_task.creator=? or design_task.designer=? or design_task.reviewer=? or design_task.approve=?) and design_task.deleted_at IS NULL",
			self.ID, self.ID, self.ID, self.ID)
	}
	total, err := connector.GetRecords(c, db.
		Clauses(clause.Select{
			Expression: clause.Expr{SQL: "design_task.*", WithoutParentheses: true}}).
		Joins("LEFT JOIN design_project as p on p.id = design_task.design_project_id").
		Preload("DesignProject").
		Set("skipCallback", true), &designTasks)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: designTasks})
}

// FindTaskDesignResults 查找设计结果
func FindTaskDesignResults(c *gin.Context) {
	var designTasks []model.DesignTask
	err := permission.Check(c, resource.DesignTask, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	// 判断总工，总工可以看见所有设计书
	db := connector.DataSource
	hasRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if !hasRole {
		db = db.Where("d.creator = ? or d.designer = ? or d.reviewer = ? or d.drafting_by = ? or d.check_by = ?",
			self.ID, self.ID, self.ID, self.ID, self.ID)
	}
	preloadDB := db.
		Clauses(clause.Select{
			Expression: clause.Expr{SQL: "design_task.*", WithoutParentheses: true}}).
		Joins("LEFT JOIN design_project as p on p.id = design_task.design_project_id").
		Joins("LEFT JOIN design_results on design_results.design_task_id = design_task.id").
		Where("design_results.deleted_at IS NULL").
		Preload("DesignProject").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.DraftingUser").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.ReviewerUser").
		Preload("DesignProject.DesignerUser").
		Preload("DesignProject.CheckUser").
		Preload("FinalDesignResults").
		Preload("FinalDesignResults.DesignProject").
		Preload("FinalDesignResults.DesignProject.CreatorUser").
		Preload("FinalDesignResults.DesignProject.DraftingUser").
		Preload("FinalDesignResults.DesignProject.CreatorUser").
		Preload("FinalDesignResults.DesignProject.ReviewerUser").
		Preload("FinalDesignResults.DesignProject.DesignerUser").
		Preload("FinalDesignResults.DesignProject.CheckUser").
		Preload("FinalDesignResults.DesignTask").
		Group("design_task.id").
		Set("skipCallback", true)

	total, err := connector.GetRecords(c, preloadDB, &designTasks)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: designTasks})
}

// RecommendDesignTasks 获取推荐的设计结果
func RecommendDesignTasks(c *gin.Context) {
	// step1 解析前端请求参数，获取推荐必要项
	// step2 依据业务方推荐规则，构造查询sql 语句
	// step3 返回结果给前端
	var designTasks []model.DesignTask
	err := permission.Check(c, resource.DesignTask, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("DesignProject").
		Preload("FinalDesignResults").
		Preload("FinalDesignResults.DesignProject").
		Preload("FinalDesignResults.DesignProject.CreatorUser").
		Preload("FinalDesignResults.DesignProject.DraftingUser").
		Preload("FinalDesignResults.DesignProject.CreatorUser").
		Preload("FinalDesignResults.DesignProject.ReviewerUser").
		Preload("FinalDesignResults.DesignProject.DesignerUser").
		Preload("FinalDesignResults.DesignProject.CheckUser").
		Preload("FinalDesignResults.DesignTask"), &designTasks)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: designTasks})
}

// TasksExcelExport 导出设计任务
func TasksExcelExport(c *gin.Context) {
	designTask := &model.DesignTask{}
	err := permission.Check(c, resource.DesignTask, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.DesignTask{}).Where("id = ?", id).
		Preload("DesignProject").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.DraftingUser").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.ReviewerUser").
		Preload("DesignProject.DesignerUser").
		Preload("DesignProject.CheckUser").First(designTask).Error, "内部服务出错，请联系开发人员处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	basePath := Documents.BasePath
	f, err := excelize.OpenFile(basePath + "DesignTaskTemplate.xlsx")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New(err, "文件打开失败，请稍后再试")})
		return
	}
	s := "Sheet1"
	f.SetActiveSheet(f.GetSheetIndex(s))
	if err = setDesignProjectSheet(f, designTask.DesignProject); err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if err = setDesignTaskSheet(f, designTask); err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	fileName := "DesignTask-" + designTask.DesignProject.SerialCode + fmt.Sprintf("_%v", designTask.ID) + ".xlsx"
	fullPath := Documents.BasePath + fileName
	if err = errors.New(f.SaveAs(fullPath), "导出失败，请稍后重试"); err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	stat, err := os.Stat(fullPath)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New("文件不存在")})
		return
	}
	size, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(stat.Size())/1024/1024), 64)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New("文件不存在")})
		return
	}

	err = connector.DataSource.Create(&model.Documents{
		Name:            "输入参数：" + strings.TrimSpace(designTask.DesignProject.ProjectName) + "-" + time.Now().Format("2006-01-02 15:04:05"),
		UUID:            uuid.NewV4().String(),
		Location:        fileName,
		DocsCreator:     self.ID,
		DocumentType:    ".xlsx",
		FileSize:        size,
		Category:        model.DocumentCategoryDesign,
		DesignProjectID: designTask.DesignProject.ID,
	}).Error
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New(err, "导出失败，请稍后重试")})
		return
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8")
	_ = f.Write(c.Writer)

}

func ApproveByDesignTask(t *model.DesignTask, approveStatus model.ApproveNodeStatus) error {
	requestToken, err := token.Create(t.Creator)
	if err != nil {
		return errors.New(err, "计算结果对接：审批流创建token失败")
	}
	client := httpclient.NewClient(requestToken)
	approveInstance := model.ApproveInstance{}
	approveNodeInstance := model.ApproveNodeInstance{}
	err = connector.DataSource.Where("JSON_EXTRACT(data,'$.id') = ?", t.ID).First(&approveInstance).Error
	if err != nil {
		return errors.New(err, "计算结果对接：查找审批实例失败")
	}
	err = connector.DataSource.Where("approve_instance_id = ? and next = 0", approveInstance.ID).Preload("ApproveNode").First(&approveNodeInstance).Error
	if err != nil {
		return errors.New(err, "计算结果对接：查找审批节点实例失败")
	}
	approveNodeInstance.Status = approveStatus
	response, err := client.PatchSpec("/v1/approve/approve_node_instance/"+strconv.Itoa(approveNodeInstance.ID), approveNodeInstance)
	if err != nil {
		return errors.New(err, "计算结果对接：审批流审批失败")
	}
	responseData := &struct {
		Code int    `json:"code"`
		Err  string `json:"error"`
		Spec string `json:"spec"`
	}{}
	if err = errors.New(json.NewDecoder(response).Decode(responseData), "计算结果对接：审批流审批响应格式有误"); err != nil {
		return err
	}
	if responseData.Code != http.StatusOK {
		return errors.New("计算结果对接：审批流审批响应结果错误")
	}
	return nil
}

type inputForm struct {
	RatedCapacity     float32        `json:"rated_capacity"`
	RatedHighVol      float32        `json:"rated_high_vol"`
	RatedLowVol       float32        `json:"rated_low_vol"`
	VolRegulationMax  float32        `json:"vol_regulation_max"`
	VolRegulationMin  float32        `json:"vol_regulation_min"`
	VolRegulationStep float32        `json:"vol_regulation_step"`
	PhaseNumber       int            `json:"phase_number"`
	ConnectType       string         `json:"connect_type"`
	RatedFrequency    float32        `json:"rated_frequency"`
	InsulationClass   string         `json:"insulation_class"`
	HasPullPlate      string         `json:"has_pull_plate"`
	InsulationTemp    int            `json:"insulation_temp"`
	WireMaterial      string         `json:"wire_material"`
	InsulationTube1   string         `json:"insulation_tube1"`
	InsulationTube2   string         `json:"insulation_tube2"`
	BossSpec          string         `json:"boss_spec"`
	ResinType         [2]interface{} `json:"resin_type"`
	LvWireType        struct {
		Round [2]interface{} `json:"round"`
		Flat  [2]interface{} `json:"flat"`
		Foil  [2]interface{} `json:"foil"`
	} `json:"lv_wire_type"`
	HvWireType struct {
		Round [2]interface{} `json:"round"`
		Flat  [2]interface{} `json:"flat"`
	} `json:"hv_wire_type"`
	SteelType                [][3]interface{} `json:"steel_type"`
	LimitTotalLoss           float32          `json:"limit_total_loss"`
	LimitLoadLoss            float32          `json:"limit_load_loss"`
	LimitNoLoadLoss          float32          `json:"limit_no_load_loss"`
	LimitNoLoadCurrent       float32          `json:"limit_no_load_current"`
	LimitImpedanceVol        float32          `json:"limit_impedance_vol"`
	LimitLvTempRise          float32          `json:"limit_lv_temp_rise"`
	LimitHvTempRise          float32          `json:"limit_hv_temp_rise"`
	LimitNoisePrediction     float32          `json:"limit_noise_prediction"`
	DeviationTotalLoss       float32          `json:"deviation_total_loss"`
	DeviationLoadLoss        float32          `json:"deviation_load_loss"`
	DeviationNoLoadLoss      float32          `json:"deviation_no_load_loss"`
	DeviationNoLoadCurrent   float32          `json:"deviation_no_load_current"`
	DeviationImpedanceVol    float32          `json:"deviation_impedance_vol"`
	DeviationLvTempRise      float32          `json:"deviation_lv_temp_rise"`
	DeviationHvTempRise      float32          `json:"deviation_hv_temp_rise"`
	DeviationNoisePrediction float32          `json:"deviation_noise_prediction"`
	EddyCurrentLossRate      float32          `json:"eddy_current_loss_rate"`
	LeadLoss                 float32          `json:"lead_loss"`
	Ks                       float32          `json:"ks"`
	TempRiseCoeff            struct {
		Hv          [4]float32 `json:"hv"`
		Lv          [4]float32 `json:"lv"`
		InsulaTube1 float32    `json:"insula_tube1"`
		InsulaTube2 float32    `json:"insula_tube2"`
	} `json:"temp_rise_coeff"`
	RadiatingSurfaceCoeff struct {
		Hv          [4][2]float32 `json:"hv"`
		Lv          [4][2]float32 `json:"lv"`
		InsulaTube1 [2]float32    `json:"insula_tube1"`
		InsulaTube2 [2]float32    `json:"insula_tube2"`
	} `json:"radiating_surface_coeff"`
}

func setDesignTaskSheet(f *excelize.File, designTask *model.DesignTask) error {
	errMsg := "写入文件失败，请稍后重试"
	s := "Sheet1"
	f.SetActiveSheet(f.GetSheetIndex(s))
	in := &inputForm{}
	err := errors.New(json.Unmarshal(designTask.Input, &in), errMsg)
	if err != nil {
		return err
	}
	var regulationArr []float32
	regulationArrLen := (in.VolRegulationMax-in.VolRegulationMin)/in.VolRegulationStep - 1
	regulationArr = append(regulationArr, in.VolRegulationMin)
	for i := 0; i < int(regulationArrLen); i++ {
		regulationArr = append(regulationArr, in.VolRegulationMin+float32(i+1)*in.VolRegulationStep)
	}
	regulationArr = append(regulationArr, in.VolRegulationMax)
	phaseNumber := "单相"
	if in.PhaseNumber == 3 {
		phaseNumber = "三相"
	}
	cellMap := map[string]interface{}{
		"O4": in.RatedCapacity, "S4": in.RatedHighVol, "W4": in.RatedLowVol, "AA4": regulationArr,
		"O5": phaseNumber, "S5": in.ConnectType, "W5": in.RatedFrequency, "AA5": in.InsulationClass,
		"O7": in.HasPullPlate, "S7": in.InsulationTemp, "W7": in.WireMaterial, "AA7": in.InsulationTube1,
		"O8": in.InsulationTube2, "S8": in.BossSpec, "W8": in.ResinType[0], "AA8": in.ResinType[1],
		"Q10": in.LvWireType.Round[0], "Y10": in.LvWireType.Round[1],
		"Q11": in.LvWireType.Flat[0], "Y11": in.LvWireType.Flat[1],
		"Q12": in.LvWireType.Foil[0], "Y12": in.LvWireType.Foil[1],
		"Q14": in.HvWireType.Round[0], "Y14": in.HvWireType.Round[1],
		"Q15": in.HvWireType.Flat[0], "Y15": in.HvWireType.Flat[1],
		//"M17": in.SteelType[0][0], "S17": in.SteelType[0][1], "Y17": in.SteelType[0][2],
		"O24": in.LimitTotalLoss, "S24": in.LimitLoadLoss, "W24": in.LimitNoLoadLoss, "AA24": in.LimitNoLoadCurrent,
		"O25": in.LimitImpedanceVol, "S25": in.LimitLvTempRise, "W25": in.LimitHvTempRise, "AA25": in.LimitNoisePrediction,
		"O27": in.DeviationTotalLoss, "S27": in.DeviationLoadLoss, "W27": in.DeviationNoLoadLoss, "AA27": in.DeviationNoLoadCurrent,
		"O28": in.DeviationImpedanceVol, "S28": in.DeviationLvTempRise, "W28": in.DeviationHvTempRise, "AA28": in.DeviationNoisePrediction,
		"O29": in.EddyCurrentLossRate, "S29": in.LeadLoss, "W29": in.Ks,
		"S31": in.TempRiseCoeff.InsulaTube1, "Y31": in.TempRiseCoeff.InsulaTube2,
		"S32": in.TempRiseCoeff.Lv[0], "T32": in.TempRiseCoeff.Lv[1], "U32": in.TempRiseCoeff.Lv[2], "Y32": in.TempRiseCoeff.Lv[3],
		"S33": in.TempRiseCoeff.Hv[0], "T33": in.TempRiseCoeff.Hv[1], "U33": in.TempRiseCoeff.Hv[2], "Y33": in.TempRiseCoeff.Hv[3],
		"S35": in.RadiatingSurfaceCoeff.InsulaTube1[0], "U35": in.RadiatingSurfaceCoeff.InsulaTube1[1], "Y35": in.RadiatingSurfaceCoeff.InsulaTube2[0], "AA35": in.RadiatingSurfaceCoeff.InsulaTube2[1],
		"S36": in.RadiatingSurfaceCoeff.Lv[0][0], "T36": in.RadiatingSurfaceCoeff.Lv[0][1], "U36": in.RadiatingSurfaceCoeff.Lv[1][0], "V36": in.RadiatingSurfaceCoeff.Lv[1][1], "W36": in.RadiatingSurfaceCoeff.Lv[2][0], "X36": in.RadiatingSurfaceCoeff.Lv[2][1], "Y36": in.RadiatingSurfaceCoeff.Lv[3][0], "Z36": in.RadiatingSurfaceCoeff.Lv[3][1],
		"S37": in.RadiatingSurfaceCoeff.Hv[0][0], "T37": in.RadiatingSurfaceCoeff.Hv[0][1], "U37": in.RadiatingSurfaceCoeff.Hv[1][0], "V37": in.RadiatingSurfaceCoeff.Hv[1][1], "W37": in.RadiatingSurfaceCoeff.Hv[2][0], "X37": in.RadiatingSurfaceCoeff.Hv[2][1], "Y37": in.RadiatingSurfaceCoeff.Hv[3][0], "Z37": in.RadiatingSurfaceCoeff.Hv[3][1],
	}
	// 不定长钢片牌号
	for i, item := range in.SteelType {
		rowIndex := fmt.Sprintf("%d", i+17)
		cellMap["M"+rowIndex] = item[0]
		cellMap["S"+rowIndex] = item[1]
		cellMap["Y"+rowIndex] = item[2]
	}
	for key, v := range cellMap {
		if err = errors.New(f.SetCellValue(s, key, v), "写入文件失败"); err != nil {
			return err
		}
	}
	return nil
}
