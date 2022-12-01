// Package lib_foil_oblong_mod_ledger 箔绕扁形模具台账API
package lib_foil_oblong_mod_ledger

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

// CreateFoilOblongModLedger 添加箔绕扁形模具台账规则
func CreateFoilOblongModLedger(c *gin.Context) {
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
	newFoilOblongModLedger := &model.FoilOblongModLedger{}
	newFoilOblongModLedger.LfomlCreator = self.ID
	err = errors.New(c.BindJSON(newFoilOblongModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newFoilOblongModLedger).Error, "箔绕扁形模具台账记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "箔绕扁形模具台账记录信息添加成功"})
}

// UpdateFoilOblongModLedger 更新箔绕扁形模具台账规则
func UpdateFoilOblongModLedger(c *gin.Context) {
	var FoilOblongModLedger *model.FoilOblongModLedger
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
	err = errors.New(c.BindJSON(&FoilOblongModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.FoilOblongModLedger{}).Select("mod_size", "mod_height").Where("id=?", id).
		Updates(&FoilOblongModLedger).Error, "箔绕扁形模具台账记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "箔绕扁形模具台账记录信息更新成功"})
}

// DeleteFoilOblongModLedger 删除箔绕扁形模具台账规则
func DeleteFoilOblongModLedger(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.FoilOblongModLedger{}).Error,
		"箔绕扁形模具台账记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "箔绕扁形模具台账记录信息删除成功"})
}

// GetSpecFoilOblongModLedger 查找箔绕低压模具台账规则
func GetSpecFoilOblongModLedger(c *gin.Context) {
	var FoilOblongModLedger model.FoilOblongModLedger
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&FoilOblongModLedger).Error,
		"箔绕扁形模具台账记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: FoilOblongModLedger})
}

// FindFoilOblongModLedger 查找箔绕扁形模具台账规则
func FindFoilOblongModLedger(c *gin.Context) {
	var FoilOblongModLedger []model.FoilOblongModLedger
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &FoilOblongModLedger)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: FoilOblongModLedger})
}

// ExcelExport 导出箔绕扁形模具台账规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"箔绕扁形模具台账编号", "模具尺寸", "模具高度(mm)", "模具图号", "模具编号", "模具数量(个)", "适用产品规格", "备注"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var foilOblongModLedgers []model.FoilOblongModLedger
	_, err = connector.GetRecords(c, connector.DataSource, &foilOblongModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, foilOblongModLedger := range foilOblongModLedgers {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{foilOblongModLedger.ID, foilOblongModLedger.ModSize, foilOblongModLedger.ModPic,
			foilOblongModLedger.ModNum, foilOblongModLedger.ModAmount, foilOblongModLedger.ModSuit, foilOblongModLedger.Remark} {
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
