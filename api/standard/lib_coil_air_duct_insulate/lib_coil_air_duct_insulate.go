// Package lib_coil_air_duct_insulate 线圈与风道绝缘API
package lib_coil_air_duct_insulate

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

// CreateCoilAirDuctInsulate 添加线圈与风道绝缘规则
func CreateCoilAirDuctInsulate(c *gin.Context) {
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
	newCoilAirDuctInsulate := &model.CoilAirDuctInsulate{}
	newCoilAirDuctInsulate.LcadiCreator = self.ID
	err = errors.New(c.BindJSON(newCoilAirDuctInsulate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newCoilAirDuctInsulate).Error, "线圈与风道绝缘记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "线圈与风道绝缘记录信息添加成功"})
}

// UpdateCoilAirDuctInsulate 更新线圈与风道绝缘规则
func UpdateCoilAirDuctInsulate(c *gin.Context) {
	var CoilAirDuctInsulate *model.CoilAirDuctInsulate
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
	err = errors.New(c.BindJSON(&CoilAirDuctInsulate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.CoilAirDuctInsulate{}).Select("wind_type", "coil_inner_insulate", "coil_outer_insulate", "air_duct_thick", "air_duct_in_out_insulate").Where("id=?", id).
		Updates(&CoilAirDuctInsulate).Error, "线圈与风道绝缘记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "线圈与风道绝缘记录信息更新成功"})
}

// DeleteCoilAirDuctInsulate 删除线圈与风道绝缘规则
func DeleteCoilAirDuctInsulate(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.CoilAirDuctInsulate{}).Error,
		"线圈与风道绝缘记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "线圈与风道绝缘记录信息删除成功"})
}

// GetSpecCoilAirDuctInsulate 查找线圈与风道绝缘规则
func GetSpecCoilAirDuctInsulate(c *gin.Context) {
	var CoilAirDuctInsulate model.CoilAirDuctInsulate
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&CoilAirDuctInsulate).Error,
		"线圈与风道绝缘记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: CoilAirDuctInsulate})
}

// FindCoilAirDuctInsulates 查找线圈与风道绝缘规则
func FindCoilAirDuctInsulates(c *gin.Context) {
	var CoilAirDuctInsulates []model.CoilAirDuctInsulate
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &CoilAirDuctInsulates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: CoilAirDuctInsulates})
}

// ExcelExport 导出线圈与风道绝缘规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"线圈与风道绝缘编号", "绕制类型", "线圈内层绝缘（mm）", "线圈外层绝缘（mm）", "风道厚度可选值（mm）", "风道内外层绝缘（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var coilAirDuctInsulates []model.CoilAirDuctInsulate
	_, err = connector.GetRecords(c, connector.DataSource, &coilAirDuctInsulates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, coilAirDuctInsulate := range coilAirDuctInsulates {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{coilAirDuctInsulate.ID, coilAirDuctInsulate.WindType,
			strconv.FormatFloat(float64(coilAirDuctInsulate.CoilInnerInsulate), 'f', 4, 64) + "mm",
			strconv.FormatFloat(float64(coilAirDuctInsulate.CoilOuterInsulate), 'f', 4, 64) + "mm",
			coilAirDuctInsulate.AirDuctThick + "mm",
			strconv.FormatFloat(float64(coilAirDuctInsulate.AirDuctInOutInsulate), 'f', 4, 64) + "mm"} {
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
