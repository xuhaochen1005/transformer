// Package lib_iron_core_oblong 铁芯（长圆形）API
package lib_iron_core_oblong

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

// CreateIronCoreOblong 添加铁芯（长圆形）规则
func CreateIronCoreOblong(c *gin.Context) {
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
	newIronCoreOblong := &model.IronCoreOblong{}
	newIronCoreOblong.LicoCreator = self.ID
	err = errors.New(c.BindJSON(newIronCoreOblong), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newIronCoreOblong).Error, "铁芯(长圆形)记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁芯(长圆形)记录信息添加成功"})
}

// UpdateIronCoreOblong 更新铁芯（长圆形）规则
func UpdateIronCoreOblong(c *gin.Context) {
	var IronCoreOblong *model.IronCoreOblong
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
	err = errors.New(c.BindJSON(&IronCoreOblong), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.IronCoreOblong{}).Select("core_diameter", "pull_plate", "section_area", "iron_corner_weight").Where("id=?", id).
		Updates(&IronCoreOblong).Error, "铁芯(长圆形)记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁芯(长圆形)记录信息更新成功"})
}

// DeleteIronCoreOblong 删除铁芯（长圆形）规则
func DeleteIronCoreOblong(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.IronCoreOblong{}).Error,
		"铁芯(长圆形)记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁芯(长圆形)记录信息删除成功"})
}

// GetSpecIronCoreOblong 查找铁芯（长圆形）规则
func GetSpecIronCoreOblong(c *gin.Context) {
	var IronCoreOblong model.IronCoreOblong
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&IronCoreOblong).Error,
		"铁芯(长圆形)记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: IronCoreOblong})
}

// FindIronCoreOblongs 查找铁芯（长圆形）规则
func FindIronCoreOblongs(c *gin.Context) {
	var IronCoreOblongs []model.IronCoreOblong
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &IronCoreOblongs)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: IronCoreOblongs})
}

// ExcelExport 导出铁芯（长圆形）规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"铁芯（长圆形）编号", "铁心直径（mm）", "有无拉板", "截面积（mm）",
		"角重（mm）", "片宽（0级）（mm）", "片宽（1级）（mm）", "片宽（2级）（mm）", "片宽（3级）（mm）",
		"片宽（4级）（mm）", "片宽（5级）（mm）", "片宽（6级）（mm）", "片宽（7级）（mm）", "片宽（8级）（mm）",
		"迭厚（0级）（mm）", "迭厚（1级）（mm）", "迭厚（2级）（mm）", "迭厚（3级）（mm）", "迭厚（4级）（mm）",
		"迭厚（5级）（mm）", "迭厚（6级）（mm）", "迭厚（7级）（mm）", "迭厚（8级）（mm）", "总迭厚（mm）", "主级实迭厚（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var ironCoreOblongs []model.IronCoreOblong
	_, err = connector.GetRecords(c, connector.DataSource, &ironCoreOblongs)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, ironCoreOblong := range ironCoreOblongs {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{ironCoreOblong.ID, ironCoreOblong.CoreDiameter, ironCoreOblong.PullPlate,
			ironCoreOblong.SectionArea, ironCoreOblong.IronCornerWeight, ironCoreOblong.StackWidth0, ironCoreOblong.StackWidth1,
			ironCoreOblong.StackWidth2, ironCoreOblong.StackWidth3, ironCoreOblong.StackWidth4, ironCoreOblong.StackWidth5,
			ironCoreOblong.StackWidth6, ironCoreOblong.StackWidth7, ironCoreOblong.StackWidth8, ironCoreOblong.StackThick0,
			ironCoreOblong.StackThick1, ironCoreOblong.StackThick2, ironCoreOblong.StackThick3, ironCoreOblong.StackThick4,
			ironCoreOblong.StackThick5, ironCoreOblong.StackThick6, ironCoreOblong.StackThick7, ironCoreOblong.StackThick8,
			ironCoreOblong.StackThickSum, ironCoreOblong.MainLevelRealStackThick} {
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
