// Package lib_low_vol_out_mod_ledger 低压外模台账API
package lib_low_vol_out_mod_ledger

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

// CreateLowVolOutModLedger 添加低压外模台账规则
func CreateLowVolOutModLedger(c *gin.Context) {
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
	newLowVolOutModLedger := &model.LowVolOutModLedger{}
	newLowVolOutModLedger.LlvomlCreator = self.ID
	err = errors.New(c.BindJSON(newLowVolOutModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newLowVolOutModLedger).Error, "低压外模台账记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "低压外模台账记录信息添加成功"})
}

// UpdateLowVolOutModLedger 更新低压外模台账规则
func UpdateLowVolOutModLedger(c *gin.Context) {
	var LowVolOutModLedger *model.LowVolOutModLedger
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
	err = errors.New(c.BindJSON(&LowVolOutModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.LowVolOutModLedger{}).Select("mod_diameter", "mod_type", "mod_height").Where("id=?", id).
		Updates(&LowVolOutModLedger).Error, "低压外模台账记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "低压外模台账记录信息更新成功"})
}

// DeleteLowVolOutModLedger 删除低压外模台账规则
func DeleteLowVolOutModLedger(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.LowVolOutModLedger{}).Error,
		"低压外模台账记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "低压外模台账记录信息删除成功"})
}

// GetSpecLowVolOutModLedger 查找低压外模台账规则
func GetSpecLowVolOutModLedger(c *gin.Context) {
	var LowVolOutModLedger model.LowVolOutModLedger
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&LowVolOutModLedger).Error,
		"低压外模台账记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: LowVolOutModLedger})
}

// FindLowVolOutModLedgers 查找低压外模台账规则
func FindLowVolOutModLedgers(c *gin.Context) {
	var LowVolOutModLedgers []model.LowVolOutModLedger
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &LowVolOutModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: LowVolOutModLedgers})
}

// ExcelExport 导出低压外模台账规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"低压外模台账编号", "低压外模外径（mm）", "模具类型", "模具高度", "模具数量",
		"模具编号", "适用产品", "备注"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var lowVolOutModLedgers []model.LowVolOutModLedger
	_, err = connector.GetRecords(c, connector.DataSource, &lowVolOutModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, lowVolOutModLedger := range lowVolOutModLedgers {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{lowVolOutModLedger.ID, lowVolOutModLedger.ModDiameter, lowVolOutModLedger.ModType,
			lowVolOutModLedger.ModHeight, lowVolOutModLedger.ModAmount, lowVolOutModLedger.ModNum,
			lowVolOutModLedger.ModSuit, lowVolOutModLedger.Remark} {
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
