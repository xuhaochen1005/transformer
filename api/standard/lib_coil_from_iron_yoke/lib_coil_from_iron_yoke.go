// Package lib_coil_from_iron_yoke 线圈端部距铁轭距离API
package lib_coil_from_iron_yoke

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

// CreateLCoilFromIronYoke 添加线圈端部距铁轭距离规则
func CreateCoilFromIronYoke(c *gin.Context) {
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
	newCoilFromIronYoke := &model.CoilFromIronYoke{}
	newCoilFromIronYoke.LcfiyCreator = self.ID
	err = errors.New(c.BindJSON(newCoilFromIronYoke), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newCoilFromIronYoke).Error, "线圈端部距铁轭距离记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "线圈端部距铁轭距离记录信息添加成功"})
}

// UpdateCoilFromIronYoke 更新线圈端部距铁轭距离规则
func UpdateCoilFromIronYoke(c *gin.Context) {
	var CoilFromIronYoke *model.CoilFromIronYoke
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
	err = errors.New(c.BindJSON(&CoilFromIronYoke), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.CoilFromIronYoke{}).Select("rated_high_vol_min", "rated_low_vol_min", "rated_high_vol_max", "rated_low_vol_max", "low_vol_coil_from_iron_yoke", "high_vol_coil_from_iron_yoke").Where("id=?", id).
		Updates(&CoilFromIronYoke).Error, "线圈端部距铁轭距离记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "线圈端部距铁轭距离记录信息更新成功"})
}

// DeleteCoilFromIronYoke 删除线圈端部距铁轭距离规则
func DeleteCoilFromIronYoke(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.CoilFromIronYoke{}).Error,
		"线圈端部距铁轭距离记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "线圈端部距铁轭距离记录信息删除成功"})
}

// GetSpecCoilFromIronYoke 查找线圈端部距铁轭距离规则
func GetSpecCoilFromIronYoke(c *gin.Context) {
	var CoilFromIronYoke model.CoilFromIronYoke
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&CoilFromIronYoke).Error,
		"线圈端部距铁轭距离记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: CoilFromIronYoke})
}

// FindCoilFromIronYokes 查找线圈端部距铁轭距离规则
func FindCoilFromIronYokes(c *gin.Context) {
	var CoilFromIronYokes []model.CoilFromIronYoke
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &CoilFromIronYokes)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: CoilFromIronYokes})
}

// ExcelExport 导出线圈端部距铁轭距离规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"线圈端部距铁轭距离编号", "额定高压下界（kV）", "额定高压上界（kV）", "额定低压下界（kV）", "额定低压上界（kV）", "低压线圈端部距铁轭距离（mm）", "高压线圈端部距铁轭距离（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var coilFromIronYokes []model.CoilFromIronYoke
	_, err = connector.GetRecords(c, connector.DataSource, &coilFromIronYokes)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, coilFromIronYoke := range coilFromIronYokes {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{coilFromIronYoke.ID,
			utils.FloatToString(coilFromIronYoke.RatedHighVolMin) + "kV",
			utils.FloatToString(coilFromIronYoke.RatedHighVolMax) + "kV",
			utils.FloatToString(coilFromIronYoke.RatedLowVolMin) + "kV",
			utils.FloatToString(coilFromIronYoke.RatedLowVolMax) + "kV",
			utils.FloatToString(coilFromIronYoke.LowVolCoilFromIronYoke) + "mm",
			utils.FloatToString(coilFromIronYoke.HighVolCoilFromIronYoke) + "mm"} {
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
