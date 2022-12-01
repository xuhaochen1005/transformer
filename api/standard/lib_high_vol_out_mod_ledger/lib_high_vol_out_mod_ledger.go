// Package lib_high_vol_out_mod_ledger 高压外模台账API
package lib_high_vol_out_mod_ledger

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

// CreateHighVolOutModLedger 添加高压外模台账规则
func CreateHighVolOutModLedger(c *gin.Context) {
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
	newHighVolOutModLedger := &model.HighVolOutModLedger{}
	newHighVolOutModLedger.LhvomlCreator = self.ID
	err = errors.New(c.BindJSON(newHighVolOutModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newHighVolOutModLedger).Error, "高压外模台账记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "高压外模台账记录信息添加成功"})
}

// UpdateHighVolOutModLedger 更新高压外模台账规则
func UpdateHighVolOutModLedger(c *gin.Context) {
	var HighVolOutModLedger *model.HighVolOutModLedger
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
	err = errors.New(c.BindJSON(&HighVolOutModLedger), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.HighVolOutModLedger{}).
		Select("mod_size", "mod_type", "mod_height", "mod_amount", "mod_pic", "mod_num", "boss_width", "groove", "nut_size", "a_size", "tap_hole_distance").Where("id=?", id).
		Updates(&HighVolOutModLedger).Error, "高压外模台账记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "高压外模台账记录信息更新成功"})
}

// DeleteHighVolOutModLedger 删除高压外模台账规则
func DeleteHighVolOutModLedger(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.HighVolOutModLedger{}).Error,
		"高压外模台账记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "高压外模台账记录信息删除成功"})
}

// GetSpecHighVolOutModLedger 查找高压外模台账规则
func GetSpecHighVolOutModLedger(c *gin.Context) {
	var HighVolOutModLedger model.HighVolOutModLedger
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&HighVolOutModLedger).Error,
		"高压外模台账记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: HighVolOutModLedger})
}

// FindHighVolOutModLedger 查找高压外模台账规则
func FindHighVolOutModLedger(c *gin.Context) {
	var HighVolOutModLedger []model.HighVolOutModLedger
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &HighVolOutModLedger)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: HighVolOutModLedger})
}

// ExcelExport 导出高压外模台账规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"高压外模台账编号", "模具尺寸", "模具类型", "模具高度(mm)", "模具数量(个)", "模具图号", "模具编号", "凸台高度（封板尺寸）", "是否开槽",
		"螺母尺寸(mm)", "A头尺寸(mm)", "抽头孔距(mm)", "封板图号", "适用产品", "备注"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var highVolOutModLedgers []model.HighVolOutModLedger
	_, err = connector.GetRecords(c, connector.DataSource, &highVolOutModLedgers)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, highVolOutModLedger := range highVolOutModLedgers {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{highVolOutModLedger.ID, highVolOutModLedger.ModSize, highVolOutModLedger.ModType, highVolOutModLedger.ModHeight,
			highVolOutModLedger.ModAmount, highVolOutModLedger.ModPic, highVolOutModLedger.ModNum, highVolOutModLedger.BossWidth, highVolOutModLedger.Groove, highVolOutModLedger.NutSize, highVolOutModLedger.ASize, highVolOutModLedger.TapHoleDistance, highVolOutModLedger.ClosurePic, highVolOutModLedger.ModSuit, highVolOutModLedger.Remark} {
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
