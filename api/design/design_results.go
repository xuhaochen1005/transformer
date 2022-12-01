// Package design_results 设计结果API
package design

import (
	"encoding/json"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/httpclient"
	"transformer_mz/libs/log"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/libs/permission/role"
	"transformer_mz/libs/third"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
)

// CreateDesignResults 添加设计结果规则
func CreateDesignResult(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	newDesignResults := &model.DesignResults{}
	err = errors.New(c.BindJSON(newDesignResults), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newDesignResults).Error, "设计结果记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	c.Set("response", &model.Response{Spec: "设计结果记录信息添加成功"})
}

func CreateDesignResults(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	var designResults []model.DesignResults
	err = errors.New(c.BindJSON(&designResults), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	if len(designResults) == 0 {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New("请求数据有误，请联系开发人员进行处理")})
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
	for _, designResult := range designResults {
		designResult.ResultStatus = model.ResultStatusCalculating
		err = errors.New(connector.DataSource.Create(&designResult).Error, "设计结果记录信息添加失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	err = errors.New(db.Model(&model.DesignTask{}).Where("id=?", designResults[0].DesignTaskID).Update("TaskStatus", model.DesignTaskResultSelected).Error, "设计结果记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	c.Set("response", &model.Response{Spec: "设计结果记录信息添加成功"})
}

// UpdateSpecDesignResults 更新设计结果规则
func UpdateSpecDesignResults(c *gin.Context) {
	var (
		updateResults *model.DesignResults
		err           error
		id            int
	)
	err = permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err = utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&updateResults), "请求数据有误，请联系开发人员进行处理")
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
	designResults := model.DesignResults{}
	err = errors.New(db.Model(&model.DesignResults{}).Where("id=?", id).First(&designResults).Error, "设计结果记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	updateResults.ResultStatus = model.ResultStatusFinished
	err = errors.New(db.Model(&model.DesignResults{}).Where("id=?", id).Select("*").Updates(&updateResults).Error, "设计结果记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(db.Model(&model.DesignTask{}).Where("id=?", designResults.DesignTaskID).Update("TaskStatus", model.DesignTaskResultSelected).Error, "设计结果记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "设计结果记录信息更新成功"})
}

// DeleteSpecDesignResults 删除设计结果规则
func DeleteSpecDesignResults(c *gin.Context) {
	err := permission.Check(c, resource.DesignTask, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.DesignResults{}).Error,
		"设计结果记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "设计结果记录信息删除成功"})
}

// GetSpecDesignResults 查找设计结果规则
func GetSpecDesignResults(c *gin.Context) {
	var DesignResults model.DesignResults
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&DesignResults).Error,
		"设计结果记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: DesignResults})
}

// FindDesignResults 查找设计结果规则
func FindDesignResults(c *gin.Context) {
	var DesignResults []model.DesignResults
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
			Expression: clause.Expr{SQL: "design_results.*", WithoutParentheses: true}}).
		Joins("LEFT JOIN design_project as d on design_results.design_project_id = d.id").
		Where("design_results.deleted_at IS NULL").
		Preload("DesignProject").
		Preload("DesignTask").
		Preload("DesignProject.CreatorUser").
		//Preload("ApproveUser").
		Preload("DesignProject.DraftingUser").
		Preload("DesignProject.CreatorUser").
		Preload("DesignProject.ReviewerUser").
		Preload("DesignProject.DesignerUser").
		Preload("DesignProject.CheckUser").
		Set("skipCallback", true)

	total, err := connector.GetRecords(c, preloadDB, &DesignResults)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: DesignResults})
}

type TaskProgressStructData struct {
	Msg      string  `json:"msg"`
	Progress float32 `json:"progress"`
}
type TaskProgressStruct struct {
	Code int                               `json:"code"`
	Msg  string                            `json:"msg"`
	Data map[string]TaskProgressStructData `json:"data"`
}

// CheckDesignResults 定时任务，检测计算是否已完成
func CheckDesignResults() {
	// 检测是否出结果： 状态是开始计算 但是DesignResults结果不为空，此时可以判断为已经出结果了，更新状态并且发送消息
	var designTasks []model.DesignTask
	err := connector.DataSource.Where("task_status = ?", model.DesignTaskStarted).
		Preload("DesignProject").
		Find(&designTasks).Error
	if err != nil {
		log.Println(err)
	}
	// 查询结果
	client := httpclient.NewClient("", third.AI_SYSTEM_URL)
	resp, err := client.GetSpec("/transformer/task/progress", nil)
	if err != nil {
		log.Println(errors.New(err, "计算接口对接：调用获取计算进度接口失败"))
		return
	}
	respData := &TaskProgressStruct{}
	if err = errors.New(json.NewDecoder(resp).Decode(respData), "响应数据格式有误"); err != nil {
		log.Println(errors.New(err, "计算接口对接：处理计算进度返回数据失败"))
		return
	}
	if respData.Code != http.StatusOK {
		err = errors.New("计算接口对接：调用获取计算进度接口失败")
		log.Println(respData.Msg)
		return
	}
	for _, designTask := range designTasks {
		func() {
			// 计算结果为空，检查进度
			if strings.Trim(string(designTask.DesignResults), " ") == "" {
				mapKey := strconv.Itoa(designTask.ID)
				progressData, ok := respData.Data[mapKey]
				// 进度接口中不存在这个ID 判断为计算失败 TODO: 失败原因呢？
				if !ok {
					// 拒绝审批节点
					if err = errors.New(ApproveByDesignTask(&designTask, model.ApproveNodeRejected), "拒绝审批节点失败"); err != nil {
						log.Println(err)
						return
					}
				}
				// TODO：超时处理
				if progressData.Progress < 1 {
					if err = connector.DataSource.Model(&model.DesignTask{}).Where("id = ?", designTask.ID).
						Update("result_progress", progressData.Progress).Error; err != nil {
						log.Println(err)
						return
					}
					return
				}
			}
			if err = connector.DataSource.Model(&model.DesignTask{}).Where("id = ?", designTask.ID).
				Update("result_progress", 1).Error; err != nil {
				log.Println(err)
				return
			}
			//进度 >= 1，计算成功
			if err = errors.New(ApproveByDesignTask(&designTask, model.ApproveNodeAccepted), "批准审批节点失败"); err != nil {
				log.Println(err)
				return
			}

		}()
	}
}

// DesignResultsExport 设计结果规则导出
func DesignResultsExport(c *gin.Context) {
	//var DesignResults model.DesignResults
	//err := permission.Check(c, resource.DesignTask, action.Read)
	//if err != nil {
	//	c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
	//	return
	//}
	////id, err := utils.GetIdFromRequest(c)
	////if err != nil {
	////	c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	////	return
	////}
	////err = errors.New(connector.DataSource.Where("id=?", id).First(&DesignResults).Error,
	////	"设计结果记录信息查找失败，请稍后再试")
	////if err != nil {
	////	c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	////	return
	////}
	//lowVolEachLayerNum := utils.StringToArray(DesignResults.LowVolEachLayerNum)
	//lowVolEachTurnsNum := utils.StringToArray(DesignResults.LowVolEachTurnsNum)
	//lowVolWireThick := utils.StringToArray(DesignResults.LowVolWireThick)
	//lowVolWireWidth := utils.StringToArray(DesignResults.LowVolWireWidth)
	//lowVolParalNum	:= utils.StringToArray(DesignResults.LowVolParalNum)
	//lowVolFoldNum := utils.StringToArray(DesignResults.LowVolFoldNum)
	//lowVolSectionFactor := utils.StringToArray(DesignResults.LowVolSectionFactor)
	//lowVolWireSquare := utils.StringToArray(DesignResults.LowVolWireSquare)
	//lowVolLayerInsulate := utils.StringToArray(DesignResults.LowVolLayerInsulate)
	//lowVolCurrentDensity := utils.StringToArray(DesignResults.LowVolCurrentDensity)
	//lowVolRefPackThick := utils.StringToArray(DesignResults.LowVolRefPackThick)
	//lowVolRadialMargin := utils.StringToArray(DesignResults.LowVolRadialMargin)
	//lowVolRealPackThick := utils.StringToArray(DesignResults.LowVolRealPackThick)
	//lowVolAvgTurnsLen := utils.StringToArray(DesignResults.LowVolAvgTurnsLen)
	//lowVolWireLen := utils.StringToArray(DesignResults.LowVolWireLen)
	//lowVolResist := utils.StringToArray(DesignResults.LowVolResist)
	//lowVolResistLoss := utils.StringToArray(DesignResults.LowVolResistLoss)
	//lowVolEddyLoss := utils.StringToArray(DesignResults.LowVolEddyLoss)
	//lowVolWireWeight := utils.StringToArray(DesignResults.LowVolWireWeight)
	//lowVolSegsNum := utils.StringToArray(DesignResults.LowVolSegsNum)
	//		lowVolFirstLayerTurnsNum := utils.StringToArray(DesignResults.LowVolFirstLayerTurnsNum)
	//		lowVolRefSegsHeight := utils.StringToArray(DesignResults.LowVolRefSegsHeight)
	//		lowVolRealSegsHeight := utils.StringToArray(DesignResults.LowVolRealSegsHeight)
	//		lowVolAxialMargin := utils.StringToArray(DesignResults.LowVolAxialMargin)
	//		lowVolSegsInter := utils.StringToArray(DesignResults.LowVolSegsInter)
	//		lowVolTerminalInsulate := utils.StringToArray(DesignResults.LowVolTerminalInsulate)
	//		coilFromIronYoke := utils.StringToArray(DesignResults.CoilFromIronYoke)
	//basePath := Documents.BasePath
	//f, err := excelize.OpenFile(basePath + "DesignResultsTemplate.xlsx")
	//if err != nil {
	//	return
	//}
	//f.SetActiveSheet(f.GetSheetIndex("Sheet1"))
	//// 项目信息
	//_ = f.SetCellStr("Sheet1", "B3", DesignResults.ProjectName)
	//for idx, value := range []interface{}{DesignResults.ProductName,
	//	DesignResults.ProjectCode,
	//	DesignResults.RatedCapacity,
	//	DesignResults.RatedHighVol,
	//	"短路阻抗",
	//	"DesignResults.CoolType",
	//	"DesignResults.IndustryFreqHoldVol",
	//	"DesignResults.InductHighVol",
	//	"调压范围"} {
	//	err = f.SetCellValue("Sheet1", "B"+strconv.Itoa(5+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//for idx, value := range []interface{}{DesignResults.PhaseNum,
	//	DesignResults.RatedLowVol,
	//	"联结组别",
	//	DesignResults.IPLevel,
	//	DesignResults.ShockHoldVol,
	//	DesignResults.InsulateLevel,
	//} {
	//	err = f.SetCellValue("Sheet1", "D"+strconv.Itoa(6+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 设计值
	//for idx, value := range []interface{}{DesignResults.LowVolSumEddyLoss,
	//	DesignResults.HighVolSumEddy,
	//	DesignResults.LeadLoss,
	//	DesignResults.LoadLoss,
	//	DesignResults.NoLoadLoss,
	//	DesignResults.SumLoss,
	//	DesignResults.NoLoadCurrent,
	//	DesignResults.ImpedanceVol,
	//	DesignResults.IronCoreTempRise,
	//	DesignResults.LowVolTempRise,
	//	DesignResults.HighVolTempRise,
	//} {
	//	err = f.SetCellValue("Sheet1", "B"+strconv.Itoa(15+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 标准值
	//for idx, value := range []interface{}{
	//	DesignResults.LoadLossStd,
	//	DesignResults.NoLoadLossStd,
	//	DesignResults.SumLossStd,
	//	DesignResults.NoLoadCurrentStd,
	//	DesignResults.ImpedanceVolStd,
	//	DesignResults.IronCoreTempRiseStd,
	//	DesignResults.LowVolTempRiseStd,
	//	DesignResults.HighVolTempRiseStd,
	//} {
	//	err = f.SetCellValue("Sheet1", "C"+strconv.Itoa(18+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 偏差
	//for idx, value := range []interface{}{
	//	DesignResults.LoadLossDev,
	//	DesignResults.NoLoadLossDev,
	//	DesignResults.SumLossDev,
	//	DesignResults.NoLoadCurrentDev,
	//	DesignResults.ImpedanceVolDev,
	//	"铁心温升",
	//	"低压温升",
	//	"高压温升",
	//} {
	//	err = f.SetCellValue("Sheet1", "C"+strconv.Itoa(14+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 铁心
	//// 铁心参数
	//for idx, value := range []interface{}{"钢片牌号",
	//	DesignResults.IronCoreDiameter,
	//	DesignResults.StackFactor,
	//	DesignResults.SeamType,
	//	DesignResults.PullPlate,
	//	DesignResults.IronCoreSquare,
	//	DesignResults.MagnetDensity,
	//	DesignResults.UnitIronLoss,
	//	DesignResults.UnitMassMagnetCapacity,
	//	DesignResults.UnitAreaSeamVa,
	//	DesignResults.WindingHeight,
	//	DesignResults.CenterDistance,
	//	DesignResults.IronPillarWeight,
	//	DesignResults.IronYokeWeight,
	//	DesignResults.IronCornerWeight,
	//	DesignResults.IronSumWeight,
	//	DesignResults.TechFactor,
	//} {
	//	err = f.SetCellValue("Sheet1", "D"+strconv.Itoa(28+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 片宽
	//_ = f.SetCellValue("Sheet1", "A28", DesignResults.StackWidth)
	//// 叠厚
	//_ = f.SetCellValue("Sheet1", "B28", DesignResults.StackThick)
	//// 低压线圈
	//for idx, value := range []interface{}{DesignResults.RatedCapacity,
	//	DesignResults.LowVolCoilConnect,
	//	DesignResults.LowVolTurnsNum,
	//	DesignResults.TurnsVol,
	//	DesignResults.LowVolLineVol,
	//	DesignResults.LowVolPhaseVol,
	//	DesignResults.LowVolLineCurrent,
	//	DesignResults.LowVolPhaseCurrent,
	//} {
	//	err = f.SetCellValue("Sheet1", "G"+strconv.Itoa(3+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 低压线圈涡耗率
	//_ = f.SetCellStr("Sheet1", "K4", "涡耗率")
	//// 高压线圈涡耗率
	//_ = f.SetCellStr("Sheet1", "U4", DesignResults.HighVolEddyLossFactor)
	//for idx, value := range []interface{}{DesignResults.HighVolLineVol,
	//	DesignResults.HighVolPhaseVol,
	//	DesignResults.HighVolTurnsNumMax,
	//	DesignResults.HighVolTurnsNumRated,
	//} {
	//	err = f.SetCellValue("Sheet1", "T"+strconv.Itoa(7+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//for i=0;i<DesignResults.LowVolPackNum;i++{
	//	for idx, value := range []interface{}{lowVolEachLayerNum[i],
	//	lowVolEachTurnsNum[i],
	//	lowVolWireThick[i],
	//	lowVolWireWidth[i],
	//	lowVolParalNum[i],
	//	lowVolFoldNum[i],
	//	lowVolSectionFactor[i],
	//	lowVolWireSquare[i],
	//	lowVolLayerInsulate[i],
	//	lowVolCurrentDensity[i],
	//	lowVolRefPackThick[i],
	//	lowVolRadialMargin[i],
	//	lowVolRealPackThick[i],
	//	lowVolAvgTurnsLen[i],
	//	lowVolWireLen[i],
	//	lowVolResist[i],
	//	lowVolResistLoss[i],
	//	lowVolEddyLoss[i],
	//	lowVolWireWeight[i],
	//	} {
	//		err = f.SetCellValue("Sheet1", "G"+strconv.Itoa(12+idx), value)
	//		if err != nil {
	//			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//			return
	//		}
	//	}
	//}
	//_ = f.SetCellValue("Sheet1", "G31", DesignResults.LowVolSumWireWeight)
	//for i=0;i<DesignResults.LowVolPackNum;i++{
	//	for idx, value := range []interface{}{lowVolSegsNum,
	//			lowVolFirstLayerTurnsNum,
	//			lowVolRefSegsHeight,
	//			lowVolRealSegsHeight,
	//			lowVolAxialMargin,
	//			lowVolSegsInter,
	//			lowVolTerminalInsulate,
	//			coilFromIronYoke,
	//	} {
	//		err = f.SetCellValue("Sheet1", "G"+strconv.Itoa(12+idx), value)
	//		if err != nil {
	//			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//			return
	//		}
	//	}
	//}
	//_ = f.SetCellValue("Sheet1", "G40", DesignResults.WindowHeight)
	//_ = f.SetCellValue("Sheet1", "G41", DesignResults.LowVolResinWeight)
	//////成本校验
	//for idx, value := range []interface{}{DesignResults.HighVolWireType,
	//	DesignResults.LowVolWireType,
	//	DesignResults.LoadLossStd,
	//	DesignResults.ResinType,
	//	DesignResults.IronCopperRatio,
	//	DesignResults.SumCost,
	//	DesignResults.SumWeight,
	//} {
	//	err = f.SetCellValue("Sheet1", "W"+strconv.Itoa(45+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//for idx, value := range []interface{}{DesignResults.HighVolSumWireWeight,
	//	DesignResults.LowVolSumWireWeight,
	//	DesignResults.StackSumWeight,
	//	DesignResults.ResinSumWeight,
	//} {
	//	err = f.SetCellValue("Sheet1", "X"+strconv.Itoa(45+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//for idx, value := range []interface{}{DesignResults.HighVolWirePerPrice,
	//	DesignResults.LowVolWirePerPrice,
	//	DesignResults.StackPerPrice,
	//	DesignResults.ResinPerPrice,
	//} {
	//	err = f.SetCellValue("Sheet1", "Y"+strconv.Itoa(45+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 模具信息
	//_ = f.SetCellStr("Sheet1", "W53", DesignResults.LowVolInMod)
	//for idx, value := range []interface{}{DesignResults.LowVolOutMod,
	//	DesignResults.HighVolInMod,
	//	DesignResults.HighVolOutMod,
	//} {
	//	err = f.SetCellValue("Sheet1", "W"+strconv.Itoa(55+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//// 短路阻抗计算
	//for idx, value := range []interface{}{DesignResults.Delta,
	//	DesignResults.Delta1,
	//	DesignResults.Delta2,
	//	DesignResults.SumAn,
	//	DesignResults.DeltaR,
	//	DesignResults.Lw,
	//	DesignResults.Dm,
	//	DesignResults.kVAlue,
	//	DesignResults.KS,
	//	DesignResults.UX,
	//	DesignResults.UR,
	//	DesignResults.UK,
	//} {
	//	err = f.SetCellValue("Sheet1", "AA"+strconv.Itoa(44+idx), value)
	//	if err != nil {
	//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
	//		return
	//	}
	//}
	//f.SaveAs("./upload/DesignResultsTemplate.xlsx")
	//c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8")
	//_ = f.Write(c.Writer)
}
