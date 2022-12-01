// Package lib_wind_in_mod_ledger 绕线内模台账API
package lib_wind_in_mod_ledger

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

// CreateWindInModLedger 添加绕线内模台账规则
func CreateWindInModLedger(c *gin.Context) {
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
	newWindInModLedger := &model.WindInModLedger{}
	newWindInModLedger.LwimlCreator = self.ID
	err = errors.New(c.BindJSON(newWindInModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newWindInModLedger).Error, "绕线内模台账记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "绕线内模台账记录信息添加成功"})
}

// UpdateWindInModLedger 更新绕线内模台账规则
func UpdateWindInModLedger(c *gin.Context) {
	var WindInModLedger *model.WindInModLedger
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
	err = errors.New(c.BindJSON(&WindInModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.WindInModLedger{}).Select("mod_diameter", "mod_size", "mod_amount", "mod_pic", "mod_num").Where("id=?", id).
		Updates(&WindInModLedger).Error, "绕线内模台账记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "绕线内模台账记录信息更新成功"})
}

// DeleteWindInModLedger 删除绕线内模台账规则
func DeleteWindInModLedger(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.WindInModLedger{}).Error,
		"绕线内模台账记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "绕线内模台账记录信息删除成功"})
}

// GetSpecWindInModLedger 查找绕线内模台账规则
func GetSpecWindInModLedger(c *gin.Context) {
	var WindInModLedger model.WindInModLedger
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&WindInModLedger).Error,
		"绕线内模台账记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: WindInModLedger})
}

// FindWindInModLedgers 查找绕线内模台账规则
func FindWindInModLedgers(c *gin.Context) {
	var WindInModLedgers []model.WindInModLedger
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &WindInModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: WindInModLedgers})
}

// ExcelExport 导出绕线内模台账规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"绕线内模台账编号", "模具直径（mm）", "模具规格（直径*高度）",
		"模具数量", "模具图号", "模具编号", "适用产品型号", "备注"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var windInModLedgers []model.WindInModLedger
	_, err = connector.GetRecords(c, connector.DataSource, &windInModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, windInModLedger := range windInModLedgers {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{windInModLedger.ID, windInModLedger.ModDiameter, windInModLedger.ModSize,
			windInModLedger.ModAmount, windInModLedger.ModPic, windInModLedger.ModNum,
			windInModLedger.ModSuit, windInModLedger.Remark} {
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
