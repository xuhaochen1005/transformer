// Package lib_loss 损耗API
package lib_loss

import (
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// CreateLoss 添加损耗规则
func CreateLoss(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	newLoss := &model.Loss{}
	newLoss.LlCreator = self.ID
	err = errors.New(c.BindJSON(newLoss), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newLoss).Error, "损耗记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "损耗记录信息添加成功"})
}

// UpdateLoss 更新损耗规则
func UpdateLoss(c *gin.Context) {
	var Loss *model.Loss
	err := permission.Check(c, resource.Libraries, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&Loss), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.Loss{}).Select("regulate_way", "rated_capacity_min", "rated_capacity_max", "rated_high_vol_min", "rated_high_vol_max", "rated_low_vol_min",
		"rated_low_vol_max", "regulate_range_min", "regulate_range_max", "regulate_range_step", "temperature", "load_loss", "no_load_loss", "no_load_current", "short_circuit_imped").Where("id=?", id).
		Updates(&Loss).Error, "损耗记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "损耗记录信息更新成功"})
}

// DeleteLoss 删除损耗规则
func DeleteLoss(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.Loss{}).Error,
		"损耗记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "损耗记录信息删除成功"})
}

// GetSpecLoss 查找损耗规则
func GetSpecLoss(c *gin.Context) {
	var Loss model.Loss
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&Loss).Error,
		"损耗记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: Loss})
}

// FindLosss 查找损耗规则
func FindLosss(c *gin.Context) {
	var Losss []model.Loss
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &Losss)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: Losss})
}

// ExcelExport 导出损耗规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"损耗编号", "调压方式（无励磁调压/有载调压）", "额定容量下界（kVA）", "额定容量上界（kVA）",
		"额定高压下界（kV）", "额定高压上界（kV）", "额定低压下界（kV）", "额定低压上界（kV）",
		"调压范围下界（%）", "调压范围上界（%）", "调压范围步长（%）", "绝缘温度（℃）", "负载损耗（kW）",
		"空载损耗（kW）", "空载电流（%）", "短路阻抗（%）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var losss []model.Loss
	_, err = connector.GetRecords(c, connector.DataSource, &losss)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, loss := range losss {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{loss.ID, loss.RegulateWay, loss.RatedCapacityMin,
			loss.RatedCapacityMax, loss.RatedHighVolMin, loss.RatedHighVolMax,
			loss.RatedLowVolMin, loss.RatedLowVolMax, loss.RegulateRangeMin, loss.RegulateRangeMax,
			loss.RegulateRangeStep, loss.Temperature, loss.LoadLoss,
			loss.NoLoadLoss, loss.NoLoadCurrent, loss.ShortCircuitImped} {
			err = file.SetCellValue(sheet, string(rune('A'+idx))+strIndex, value)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return
			}
		}
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8")
	_ = file.Write(c.Writer)
}
