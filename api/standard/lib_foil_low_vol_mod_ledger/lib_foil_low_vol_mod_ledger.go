// Package lib_foil_low_vol_mod_ledger 箔绕低压模具台账API
package lib_foil_low_vol_mod_ledger

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

// CreateFoilLowVolModLedger 添加箔绕低压模具台账规则
func CreateFoilLowVolModLedger(c *gin.Context) {
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
	newFoilLowVolModLedger := &model.FoilLowVolModLedger{}
	newFoilLowVolModLedger.LflvmlCreator = self.ID
	err = errors.New(c.BindJSON(newFoilLowVolModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newFoilLowVolModLedger).Error, "箔绕低压模具台账记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "箔绕低压模具台账记录信息添加成功"})
}

// UpdateFoilLowVolModLedger 更新箔绕低压模具台账规则
func UpdateFoilLowVolModLedger(c *gin.Context) {
	var FoilLowVolModLedger *model.FoilLowVolModLedger
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
	err = errors.New(c.BindJSON(&FoilLowVolModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.FoilLowVolModLedger{}).Select("mod_size", "platform_height", "mod_height", "mod_pic", "mod_num", "mod_amount").Where("id=?", id).
		Updates(&FoilLowVolModLedger).Error, "箔绕低压模具台账记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "箔绕低压模具台账记录信息更新成功"})
}

// DeleteFoilLowVolModLedger 删除箔绕低压模具台账规则
func DeleteFoilLowVolModLedger(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.FoilLowVolModLedger{}).Error,
		"箔绕低压模具台账记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "箔绕低压模具台账记录信息删除成功"})
}

// GetSpecFoilLowVolModLedger 查找箔绕低压模具台账规则
func GetSpecFoilLowVolModLedger(c *gin.Context) {
	var FoilLowVolModLedger model.FoilLowVolModLedger
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&FoilLowVolModLedger).Error,
		"箔绕低压模具台账记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: FoilLowVolModLedger})
}

// FindFoilLowVolModLedger 查找箔绕低压模具台账规则
func FindFoilLowVolModLedger(c *gin.Context) {
	var FoilLowVolModLedger []model.FoilLowVolModLedger
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &FoilLowVolModLedger)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: FoilLowVolModLedger})
}

// ExcelExport 导出箔绕低压模具台账规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"箔绕低压模具台账编号", "模具尺寸", "平台高度", "模具高度(mm)", "模具图号", "模具编号", "模具数量(个)", "适用产品型号", "备注"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var foilLowVolModLedgers []model.FoilLowVolModLedger
	_, err = connector.GetRecords(c, connector.DataSource, &foilLowVolModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, foilLowVolModLedger := range foilLowVolModLedgers {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{foilLowVolModLedger.ID, foilLowVolModLedger.ModSize, foilLowVolModLedger.PlatformHeight, foilLowVolModLedger.ModPic,
			foilLowVolModLedger.ModNum, foilLowVolModLedger.ModAmount, foilLowVolModLedger.ModSuit, foilLowVolModLedger.Remark} {
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
