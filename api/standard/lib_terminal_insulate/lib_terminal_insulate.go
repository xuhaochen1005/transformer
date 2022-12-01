// Package lib_terminal_insulate 端绝缘距离API
package lib_terminal_insulate

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

// CreateTerminalInsulate 添加端绝缘距离规则
func CreateTerminalInsulate(c *gin.Context) {
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
	newTerminalInsulate := &model.TerminalInsulate{}
	newTerminalInsulate.LtiCreator = self.ID
	err = errors.New(c.BindJSON(newTerminalInsulate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newTerminalInsulate).Error, "端绝缘距离记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "绕端绝缘距离记录信息添加成功"})
}

// UpdateTerminalInsulate 更新端绝缘距离规则
func UpdateTerminalInsulate(c *gin.Context) {
	var TerminalInsulate *model.TerminalInsulate
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
	err = errors.New(c.BindJSON(&TerminalInsulate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.TerminalInsulate{}).Select("rated_high_vol_min", "rated_high_vol_max", "rated_low_vol_min", "rated_low_vol_max", "terminal_insulate").Where("id=?", id).
		Updates(&TerminalInsulate).Error, "端绝缘距离记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "端绝缘距离记录信息更新成功"})
}

// DeleteTerminalInsulate 删除端绝缘距离规则
func DeleteTerminalInsulate(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.TerminalInsulate{}).Error,
		"端绝缘距离记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "端绝缘距离记录信息删除成功"})
}

// GetSpecTerminalInsulate 查找端绝缘距离规则
func GetSpecTerminalInsulate(c *gin.Context) {
	var TerminalInsulate model.TerminalInsulate
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&TerminalInsulate).Error,
		"端绝缘距离记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: TerminalInsulate})
}

// FindTerminalInsulates 查找端绝缘距离规则
func FindTerminalInsulates(c *gin.Context) {
	var TerminalInsulates []model.TerminalInsulate
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &TerminalInsulates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: TerminalInsulates})
}

/// ExcelExport 导出端绝缘距离规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"端绝缘距离编号", "额定高压下界（kV）", "额定高压上界（kV）", "额定低压下界（kV）",
		"额定低压上界（kV）", "端绝缘距离（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var terminalInsulates []model.TerminalInsulate
	_, err = connector.GetRecords(c, connector.DataSource, &terminalInsulates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, terminalInsulate := range terminalInsulates {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{terminalInsulate.ID, terminalInsulate.RatedHighVolMin, terminalInsulate.RatedHighVolMax,
			terminalInsulate.RatedLowVolMin, terminalInsulate.RatedLowVolMax, terminalInsulate.TerminalInsulate} {
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
