// Package design 产品设计流程管理
package design

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"transformer_mz/api/documents"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/log"
	"transformer_mz/libs/message"
	"transformer_mz/libs/message/template"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/libs/permission/role"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// CreateDesignProjects 创建设计任务
func CreateDesignProjects(c *gin.Context) {
	var (
		err     error
		project model.DesignProject
		self    *model.User
	)
	err = permission.Check(c, resource.DesignProject, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err = utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	//	newDesignTask := &model.DesignTask{Creator: self.ID, TaskStatus: model.DesignTaskCreated}
	err = errors.New(c.BindJSON(&project), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	project.Creator = self.ID
	project.DraftingBy = self.ID
	project.CheckBy = self.ID
	project.CreatedAt = model.Timestamp{Time: time.Now()}
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
	var projectNum int64
	now := time.Now()
	monthlyStartTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	err = errors.New(db.Model(&model.DesignProject{}).Where("created_at >= ?", monthlyStartTime.Unix()).Set("skipCallback", true).Count(&projectNum).Error, "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	project.SerialCode = time.Now().Format("200601") + fmt.Sprintf("%03d", projectNum+1) //日期加产品代码
	project.ProjectStatus = model.DesignProjectImported
	/*	project.DrawingAt = int(time.Now().Unix())
		project.DraftingBy = self.ID*/
	err = errors.New(db.Create(&project).Error, "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = message.SendMessageByTemplate(template.DesignTaskCreated, self.ID, []int{project.Designer},
		[]string{project.ProjectName, project.SerialCode}, db)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(db.Create(&model.TodoList{
		Title:     fmt.Sprintf("设计 %s", project.ProjectName),
		ReferType: "project",
		ReferID:   project.ID,
		CreatedBy: self.ID,
	}).Error, "内部服务出错，请联系开发人员处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	/*newDesignTask := &model.DesignTask{Creator: self.ID, TaskStatus: model.DesignTaskCreated}
	err = errors.New(c.BindJSON(&project), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newDesignTask).Error, "设计任务添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = message.SendMessageByTemplate(template.DesignTask, self.ID, []int{newDesignTask.Designer},
		[]string{newDesignTask.Name, strconv.Itoa(newDesignTask.ID)})
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}*/
	c.Set("response", &model.Response{Spec: "设计任务添加成功"})
}

// CancelDesignProjectsStatus 取消特定的设计任务
func CancelDesignProjectsStatus(c *gin.Context) {
	var designTask model.DesignTask
	err := permission.Check(c, resource.DesignTask, action.Cancel)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&designTask).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	switch self.ID {
	case designTask.Creator:
		if designTask.TaskStatus > model.DesignTaskStarted {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError,
				Err: errors.New("设计任务已开始，不能取消此任务")})
			return
		}
	default:
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New("内部服务出错，请联系开发人员处理")})
		return
	}
	err = errors.New(connector.DataSource.Model(&designTask).Where("id=?", id).
		Update("status", model.DesignTaskCanceled).Error, "设计任务取消失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "设计任务取消成功"})
}

// GetSpecDesignProjects 查询特定的设计任务
func GetSpecDesignProjects(c *gin.Context) {
	var designProject model.DesignProject
	err := permission.Check(c, resource.DesignProject, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	/*self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}*/
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	err = errors.New(connector.DataSource.Where("id=?", id).First(&designProject).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	c.Set("response", &model.Response{Spec: designProject})
}

// UpdateSpecDesignProjects 更新项目任务书
func UpdateSpecDesignProjects(c *gin.Context) {
	var designProject model.DesignProject
	err := permission.Check(c, resource.DesignProject, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	/*self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}*/
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&designProject), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id = ?", id).Select("*").Updates(&designProject).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "更新设计任务书成功"})
}

// DeleteSpecDesignProjects 删除项目任务书
func DeleteSpecDesignProjects(c *gin.Context) {
	var (
		designProject model.DesignProject
		err           error
		id            int
	)
	err = permission.Check(c, resource.DesignProject, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err = utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&designProject).Error,
		"设计任务信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	c.Set("response", &model.Response{Spec: nil})
}

// FindDesignProjects 查询所有项目任务书
func FindDesignProjects(c *gin.Context) {

	var (
		designProjects []model.DesignProject
		response       model.Response
		err            error
	)
	err = permission.Check(c, resource.DesignProject, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}

	/*
		db := connector.DataSource.Joins("JOIN department ON department.id=user.department_id").
			Select([]string{"user.id as id", "user.name as name", "department.name as department", "department_id",
				"email", "telephone", "user.status as status", "info", "user.created_at as created_at",
				"user.updated_at as updated_at"}).Where("user.id=?", self.ID)
	*/

	/*connector.DataSource.Select("design_project.*,user.*").Joins("JOIN user on user.id=design_project.creator").
	Find(&designProjects)*/
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	//fmt.Print(self.Name)
	// 总工可以看见所有设计书
	db := connector.DataSource
	// 判断总工
	hasChiefEngineerRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	hasRootRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.Root))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
	}

	if hasRootRole == false && hasChiefEngineerRole == false {
		db = db.Where("creator = ? or designer = ? or reviewer = ? or drafting_by = ? or check_by = ?", self.ID, self.ID,
			self.ID, self.ID, self.ID)
	}
	response.Total, err = connector.GetRecords(c, db.
		Preload("CreatorUser").
		//Preload("ApproveUser").
		Preload("DraftingUser").
		Preload("CreatorUser").
		Preload("ReviewerUser").
		Preload("DesignerUser").
		Preload("CheckUser"), &designProjects)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	response.Spec = designProjects
	c.Set("response", &response)
}

// ProjectsExport 导出项目任务书
func ProjectsExport(c *gin.Context) {
	var DesignProject model.DesignProject
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
	err = errors.New(connector.DataSource.Where("id=?", id).
		Preload("DesignerUser").
		Preload("CreatorUser").
		Preload("DraftingUser").
		Preload("CreatorUser").
		Preload("ReviewerUser").
		Preload("DesignerUser").
		Preload("CheckUser").
		First(&DesignProject).Error,
		"设计任务信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	basePath := Documents.BasePath
	f, err := excelize.OpenFile(basePath + "DesignProjectTemplate.xlsx")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New(err, "文件打开失败，请稍后再试")})
		return
	}
	if err = errors.New(setDesignProjectSheet(f, &DesignProject), "写入文件失败，请稍后重试"); err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	fileName := "DesignProject-" + DesignProject.SerialCode + ".xlsx"
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
		Name:            "设计任务书：" + strings.TrimSpace(DesignProject.ProjectName) + "-" + time.Now().Format("2006-01-02 15:04:05"),
		UUID:            uuid.NewV4().String(),
		Location:        fileName,
		DocsCreator:     self.ID,
		DocumentType:    ".xlsx",
		FileSize:        size,
		Category:        model.DocumentCategoryDesign,
		DesignProjectID: DesignProject.ID,
	}).Error
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New(err, "导出失败，请稍后重试")})
		return
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8")
	_ = f.Write(c.Writer)
}

func setDesignProjectSheet(f *excelize.File, DesignProject *model.DesignProject) error {
	errMsg := "写入文件失败，请稍后重试"
	f.SetActiveSheet(f.GetSheetIndex("Sheet1"))
	designFinishAt := time.Unix(int64(DesignProject.DesignFinishAt), 0).Format("2006-01-02")
	deliverAt := time.Unix(int64(DesignProject.DeliverAt), 0).Format("2006-01-02")
	designAt := time.Unix(int64(DesignProject.DesignAt), 0).Format("2006-01-02")
	reviewAt := time.Unix(int64(DesignProject.ReviewAt), 0).Format("2006-01-02")
	draftingAt := time.Unix(int64(DesignProject.DraftingAt), 0).Format("2006-01-02")
	checkedAt := time.Unix(int64(DesignProject.CheckedAt), 0).Format("2006-01-02")
	var needMaterApprove string
	if DesignProject.NeedMasterApprove == 0 {
		needMaterApprove = "无需总工审核"
	} else {
		needMaterApprove = "需要总工审核"
	}
	var designerSigned string
	if DesignProject.DesignerSigned == 0 {
		designerSigned = "否"
	} else {
		designerSigned = "是"
	}
	var reviewerSigned string
	if DesignProject.ReviewerSigned == 0 {
		reviewerSigned = "否"
	} else {
		reviewerSigned = "是"
	}
	var cellMap = make(map[string]interface{})
	cellMap = map[string]interface{}{
		"J1":  DesignProject.SerialCode,
		"C3":  DesignProject.ProjectName,
		"C4":  DesignProject.ProductModel,
		"C5":  DesignProject.ProductCode,
		"C6":  DesignProject.ProductUsage,
		"H5":  designFinishAt,
		"H6":  deliverAt,
		"C9":  DesignProject.ProductPhrase,
		"C10": DesignProject.ProductFrequency,
		"E12": DesignProject.ProductRatedVHigh,
		"I12": DesignProject.ProductRatedVLow,
		"C13": DesignProject.ProductTapRange,
		"C14": DesignProject.ProductIndustryFreqHoldVol,
		"C15": DesignProject.ProductInductHighVol,
		"C16": DesignProject.ProductInsulationLevel,
		"C17": DesignProject.ProductTempLimit,
		"C18": DesignProject.ProductConnectionGroup,
		"C19": DesignProject.ProductIPLevel,
		"C20": DesignProject.ProductShortCircuitResistance,
		"C21": DesignProject.ProductAltitude,
		"C22": DesignProject.ProductCoolingMethod,
		"F23": DesignProject.ProductLoadLoss,
		"F24": DesignProject.ProductNoloadLoss,
		"F25": DesignProject.ProductTotalLoss,
		"C26": DesignProject.ProductWireMaterial,
		"C27": DesignProject.TechRequirments,
		"C31": DesignProject.DesignerUser.UserNameZh,
		"C32": DesignProject.ReviewerUser.UserNameZh,
		"C33": DesignProject.DesignComment,
		"C36": DesignProject.DraftingUser.UserNameZh,
		"C37": DesignProject.CheckUser.UserNameZh,
		"C38": needMaterApprove,
		"G31": designAt,
		"G32": reviewAt,
		"K31": designerSigned,
		"K32": reviewerSigned,
		"H36": draftingAt,
		"H37": checkedAt,
	}
	for key, v := range cellMap {
		if err := errors.New(f.SetCellValue("Sheet1", key, v), errMsg); err != nil {
			return err
		}
	}
	return nil
}
