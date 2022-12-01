package statistics

import (
	"gorm.io/gorm/clause"
	"net/http"
	"time"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"

	"github.com/gin-gonic/gin"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/role"
	"transformer_mz/utils"
)

func GetStatistics(c *gin.Context) {
	type DesignProjectStatusCount struct {
		// 统计不同状态设计项目数(design_project)
		DesignProjectImportedCount          int `json:"design_project_imported_count"`
		DesignProjectCreatedCount           int `json:"design_project_created_count"`
		DesignProjectStartedCount           int `json:"design_project_started_count"`
		DesignProjectFinishedCount          int `json:"design_project_finished_count"`
		DesignProjectReviewUnacceptedCount  int `json:"design_project_review_unaccepted_count"`
		DesignProjectReviewedCount          int `json:"design_project_reviewed_count"`
		DesignProjectApproveUnacceptedCount int `json:"design_project_approve_unaccepted_count"`
		DesignProjectApprovedCount          int `json:"design_project_approved_count"`
		DesignProjectCheckUnacceptedCount   int `json:"design_project_check_unaccepted_count"`
		DesignProjectCheckedCount           int `json:"design_project_checked_count"`
		DesignProjectOverdueCount           int `json:"design_project_overdue_count"`
	}
	type resultStruct struct {
		DesignProjectCount          int64 `json:"design_project_count"`           // 设计任务数
		LibrariesCount              int64 `json:"libraries_count"`                //标注库数量
		DesignResultsCount          int64 `json:"design_results_count"`           //设计结果数
		UserCount                   int64 `json:"user_count"`                     //系统使用人数
		DesignProjectDeliveredCount int64 `json:"design_project_delivered_count"` //已交货数量
		DesignProjectExceededCount  int64 `json:"design_project_exceeded_count"`  //设计延期数
		DesignProjectStatusCount
		//统计一周时间下各状态设计任务数(design_project)
		DesignProjectWeeklyStatusData [7]struct {
			DesignProjectCount int          `json:"design_project_count"`
			Date               string       `json:"date"`
			Weekday            time.Weekday `json:"-"`
			Time               time.Time    `json:"-"`
			//DesignProjectStatusCount
		} `json:"design_project_weekly_status_count"`
		//  列出设计项目对应产品要交货日期等信息 (design_project)
	}
	result := resultStruct{}
	// 获取用户
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	// 判断总工
	hasChiefEngineerRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.ChiefEngineer))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
	}
	// 判断管理员
	haRootRole, err := permission.Permission.HasRoleForUser(self.Name, string(role.Root))
	if err != nil {
		err = errors.New(err, "内部服务出错，请联系开发人员处理")
	}
	//设计任务数
	if hasChiefEngineerRole == true || haRootRole == true {
		err = errors.New(connector.DataSource.Model(model.DesignProject{}).Count(&result.DesignProjectCount).Error, "数据查询失败，请稍后重试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	} else {
		err = errors.New(connector.DataSource.Model(model.DesignProject{}).Where("designer = ? or reviewer = ? or drafting_by = ? or check_by = ?", self.ID, self.ID, self.ID, self.ID).Count(&result.DesignProjectCount).Error, "数据查询失败，请稍后重试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	//标准库数
	err = errors.New(connector.DataSource.Model(&model.Libraries{}).Count(&result.LibrariesCount).Error, "数据查询失败，请稍后重试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	//设计结果数
	if hasChiefEngineerRole == true || haRootRole == true {
		err = errors.New(connector.DataSource.Model(&model.DesignResults{}).Where("result_status = ?", model.ResultStatusFinished).Count(&result.DesignResultsCount).Error, "数据查询失败，请稍后重试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	} else {
		err = errors.New(connector.DataSource.Model(&model.DesignResults{}).
			Clauses(clause.Select{
				Expression: clause.Expr{SQL: "design_results.id", WithoutParentheses: true}}).
			Joins("LEFT JOIN design_project as d on design_results.design_project_id = d.id").
			Where("d.creator = ? or d.designer = ? or d.reviewer = ? or d.drafting_by = ? or d.check_by = ?",
				self.ID, self.ID, self.ID, self.ID, self.ID).
			Where("design_results.deleted_at IS NULL").
			Where("result_status = ?", model.ResultStatusFinished).
			Set("skipCallback", true).
			Count(&result.DesignResultsCount).Error, "数据查询失败，请稍后重试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	//使用人数
	err = errors.New(connector.DataSource.Model(&model.User{}).Count(&result.UserCount).Error, "数据查询失败，请稍后重试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	//已交货产品数
	if hasChiefEngineerRole == true || haRootRole == true {
		err = errors.New(connector.DataSource.Model(&model.DesignProject{}).Where("deliver_at not null and project_status = ?", model.DesignProjectChecked).
			Count(&result.DesignProjectDeliveredCount).Error, "数据查询失败，请稍后重试")
	} else {
		err = errors.New(connector.DataSource.Model(&model.DesignProject{}).
			Where("d.creator = ? or d.designer = ? or d.reviewer = ? or d.drafting_by = ? or d.check_by = ?",
				self.ID, self.ID, self.ID, self.ID, self.ID).
			Where("deliver_at not null and project_status = ?", model.DesignProjectChecked).
			Count(&result.DesignProjectDeliveredCount).Error, "数据查询失败，请稍后重试")
	}
	//设计延期数
	if hasChiefEngineerRole == true || haRootRole == true {
		err = errors.New(connector.DataSource.Model(&model.DesignProject{}).Where("deliver_at < ? and project_status != ?", time.Now().Unix(), model.DesignProjectChecked).
			Count(&result.DesignProjectExceededCount).Error, "数据查询失败，请稍后重试")
	} else {
		err = errors.New(connector.DataSource.Model(&model.DesignProject{}).
			Where("d.creator = ? or d.designer = ? or d.reviewer = ? or d.drafting_by = ? or d.check_by = ?",
				self.ID, self.ID, self.ID, self.ID, self.ID).
			Where("deliver_at < ? and project_status != ?", time.Now().Unix(), model.DesignProjectChecked).
			Count(&result.DesignProjectExceededCount).Error, "数据查询失败，请稍后重试")
	}

	type projectStatusStatisticsStruct struct {
		ProjectStatus model.DesignProjectStatus `json:"project_status" gorm:"column:project_status"`
		Count         int                       `json:"count" gorm:"column:count"`
		OverdueCount  int
	}
	// 查找项目状态
	var projectStatusStatistics []*projectStatusStatisticsStruct
	rows, err := connector.DataSource.Select([]string{"project_status", "count(project_status) as count"}).
		Group("project_status").Find(&model.DesignProject{}).Rows()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	// 统计不同状态设计项目数(design_project)
	for rows.Next() {
		projectStatusStatisticsRow := projectStatusStatisticsStruct{}
		if hasChiefEngineerRole == true || haRootRole == true {
			err := connector.DataSource.ScanRows(rows, &projectStatusStatisticsRow)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
		} else {
			err := connector.DataSource.Where("designer = ? or reviewer = ? or drafting_by = ? or check_by = ?", self.ID, self.ID, self.ID, self.ID).ScanRows(rows, &projectStatusStatisticsRow)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
		}
		projectStatusStatistics = append(projectStatusStatistics, &projectStatusStatisticsRow)
	}

	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for _, projectStatus := range projectStatusStatistics {
		if projectStatus.ProjectStatus != model.DesignProjectChecked {
			result.DesignProjectOverdueCount += 1
		}
		switch projectStatus.ProjectStatus {
		case model.DesignProjectImported:
			result.DesignProjectImportedCount = projectStatus.Count
		case model.DesignProjectCreated:
			result.DesignProjectCreatedCount = projectStatus.Count
		case model.DesignProjectStarted:
			result.DesignProjectStartedCount = projectStatus.Count
		case model.DesignProjectFinished:
			result.DesignProjectStartedCount = projectStatus.Count
		case model.DesignProjectReviewUnaccepted:
			result.DesignProjectReviewUnacceptedCount = projectStatus.Count
		case model.DesignProjectReviewed:
			result.DesignProjectReviewedCount = projectStatus.Count
		case model.DesignProjectApproveUnaccepted:
			result.DesignProjectApproveUnacceptedCount = projectStatus.Count
		case model.DesignProjectApproved:
			result.DesignProjectApprovedCount = projectStatus.Count
		case model.DesignProjectCheckUnaccepted:
			result.DesignProjectCheckUnacceptedCount = projectStatus.Count
		case model.DesignProjectChecked:
			result.DesignProjectCheckedCount = projectStatus.Count
		}
	}

	// 展示一周内各状态的设计任务
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weeklyDesignProjects := []model.DesignProject{}
	mondayDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	sundayDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, offset+6)
	err = errors.New(connector.DataSource.Where("created_at >= ? and created_at <= ?", mondayDate.Unix(), sundayDate.Unix()).
		Select([]string{"project_status", "created_at"}).Find(&weeklyDesignProjects).Error, "数据查询失败，请稍后重试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for i, _ := range &result.DesignProjectWeeklyStatusData {
		weekDate := mondayDate.AddDate(0, 0, i)
		result.DesignProjectWeeklyStatusData[i].Date = weekDate.Format("2006-01-02")
		result.DesignProjectWeeklyStatusData[i].Time = weekDate
		result.DesignProjectWeeklyStatusData[i].Weekday = weekDate.Weekday()
	}
	for _, designProject := range weeklyDesignProjects {
		for i, _ := range &result.DesignProjectWeeklyStatusData {
			//	if int64(designProject.DeliverAt) < result.DesignProjectWeeklyStatusData[i].Time.Unix() {
			//		result.DesignProjectWeeklyStatusData[i].DesignProjectOverdueCount += 1
			//	}
			if result.DesignProjectWeeklyStatusData[i].Weekday == designProject.CreatedAt.Weekday() {
				result.DesignProjectWeeklyStatusData[i].DesignProjectCount += 1
				//		switch designProject.ProjectStatus {
				//		case model.DesignProjectImported:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectImportedCount += 1
				//		case model.DesignProjectCreated:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectCreatedCount += 1
				//		case model.DesignProjectStarted:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectStartedCount += 1
				//		case model.DesignProjectFinished:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectStartedCount += 1
				//		case model.DesignProjectReviewUnaccepted:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectReviewUnacceptedCount += 1
				//		case model.DesignProjectReviewed:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectReviewedCount += 1
				//		case model.DesignProjectApproveUnaccepted:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectApproveUnacceptedCount += 1
				//		case model.DesignProjectApproved:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectApprovedCount += 1
				//		case model.DesignProjectCheckUnaccepted:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectCheckUnacceptedCount += 1
				//		case model.DesignProjectChecked:
				//			result.DesignProjectWeeklyStatusData[i].DesignProjectCheckedCount += 1
				//		}
			}
		}
	}
	c.Set("response", &model.Response{Spec: &result})

}
