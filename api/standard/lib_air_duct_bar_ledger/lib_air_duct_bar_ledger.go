// Package lib_air_duct_bar_ledger 风道条台账API
package lib_air_duct_bar_ledger

import (
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// CreateAirDuctBarLedger 添加风道条台账规则
func CreateAirDuctBarLedger(c *gin.Context) {
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
	newAirDuctBarLedger := &model.AirDuctBarLedger{}
	err = errors.New(c.BindJSON(newAirDuctBarLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	newAirDuctBarLedger.LadblCreator = self.ID
	err = errors.New(connector.DataSource.Create(newAirDuctBarLedger).Error, "风道条台账记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "风道条台账记录信息添加成功"})
}

// UpdateAirDuctBarLedger 更新风道条台账规则
func UpdateAirDuctBarLedger(c *gin.Context) {
	var AirDuctBarLedger *model.AirDuctBarLedger
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
	err = errors.New(c.BindJSON(&AirDuctBarLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.AirDuctBarLedger{}).Select("air_duct_bar_width", "air_duct_bar_thickness", "air_duct_bar_length", "air_duct_bar_num").Where("id=?", id).
		Updates(&AirDuctBarLedger).Error, "风道条台账记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "风道条台账记录信息更新成功"})
}

// DeleteAirDuctBarLedger 删除风道条台账规则
func DeleteAirDuctBarLedger(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.AirDuctBarLedger{}).Error,
		"风道条台账记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "风道条台账记录信息删除成功"})
}

// GetSpecAirDuctBarLedger 查找风道条台账规则
func GetSpecAirDuctBarLedger(c *gin.Context) {
	var AirDuctBarLedger model.AirDuctBarLedger
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&AirDuctBarLedger).Error,
		"风道条台账记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: AirDuctBarLedger})
}

// FindAirDuctBarLedgers 查找风道条台账规则
func FindAirDuctBarLedgers(c *gin.Context) {
	var AirDuctBarLedgers []model.AirDuctBarLedger
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &AirDuctBarLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: AirDuctBarLedgers})
}

// ExcelExport 导出风道条台账规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"风道条台账编号", "风道条厚度(mm)", "风道条宽度(mm)", "风道条长度(mm)", "风道条数量(个)"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}

	var airDuctBarLedgers []model.AirDuctBarLedger
	_, err = connector.GetRecords(c, connector.DataSource, &airDuctBarLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, airDuctBarLedger := range airDuctBarLedgers {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{airDuctBarLedger.ID, strconv.Itoa(airDuctBarLedger.AirDuctBarThickness) + "mm",
			strconv.Itoa(airDuctBarLedger.AirDuctBarWidth) + "mm", strconv.Itoa(airDuctBarLedger.AirDuctBarLength) + "mm", airDuctBarLedger.AirDuctBarNum} {
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
