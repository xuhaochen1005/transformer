// Package lib_iron_core_round 铁芯（圆形）API
package lib_iron_core_round

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

// CreateIronCoreRound 添加铁芯（圆形）规则
func CreateIronCoreRound(c *gin.Context) {
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
	newIronCoreRound := &model.IronCoreRound{}
	newIronCoreRound.LicrCreator = self.ID
	err = errors.New(c.BindJSON(newIronCoreRound), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newIronCoreRound).Error, "铁芯(圆形)记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁芯(圆形)记录信息添加成功"})
}

// UpdateIronCoreRound 更新铁芯（圆形）规则
func UpdateIronCoreRound(c *gin.Context) {
	var IronCoreRound *model.IronCoreRound
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
	err = errors.New(c.BindJSON(&IronCoreRound), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.IronCoreRound{}).Select("core_diameter", "pull_plate", "section_area", "iron_corner_weight").Where("id=?", id).
		Updates(&IronCoreRound).Error, "铁芯(圆形)记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁芯(圆形)记录信息更新成功"})
}

// DeleteIronCoreRound 删除铁芯（圆形）规则
func DeleteIronCoreRound(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.IronCoreRound{}).Error,
		"铁芯(圆形)记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁芯(圆形)记录信息删除成功"})
}

// GetSpecIronCoreRound 查找铁芯（圆形）规则
func GetSpecIronCoreRound(c *gin.Context) {
	var IronCoreRound model.IronCoreRound
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&IronCoreRound).Error,
		"铁芯(圆形)记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: IronCoreRound})
}

// FindIronCoreRounds 查找铁芯（圆形）规则
func FindIronCoreRounds(c *gin.Context) {
	var IronCoreRounds []model.IronCoreRound
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &IronCoreRounds)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: IronCoreRounds})
}

// ExcelExport 导出铁芯（圆形）规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"铁芯（圆形）编号", "铁心直径（mm）", "有无拉板", "截面积（mm）",
		"角重（mm）", "片宽（1级）（mm）", "片宽（2级）（mm）", "片宽（3级）（mm）",
		"片宽（4级）（mm）", "片宽（5级）（mm）", "片宽（6级）（mm）", "片宽（7级）（mm）", "片宽（8级）（mm）",
		"迭厚（1级）（mm）", "迭厚（2级）（mm）", "迭厚（3级）（mm）", "迭厚（4级）（mm）",
		"迭厚（5级）（mm）", "迭厚（6级）（mm）", "迭厚（7级）（mm）", "迭厚（8级）（mm）", "总迭厚（mm）", "主级实迭厚（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var ironCoreRounds []model.IronCoreRound
	_, err = connector.GetRecords(c, connector.DataSource, &ironCoreRounds)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, ironCoreRound := range ironCoreRounds {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{ironCoreRound.ID, ironCoreRound.CoreDiameter, ironCoreRound.PullPlate,
			ironCoreRound.SectionArea, ironCoreRound.IronCornerWeight, ironCoreRound.StackWidth1,
			ironCoreRound.StackWidth2, ironCoreRound.StackWidth3, ironCoreRound.StackWidth4, ironCoreRound.StackWidth5,
			ironCoreRound.StackWidth6, ironCoreRound.StackWidth7, ironCoreRound.StackWidth8,
			ironCoreRound.StackThick1, ironCoreRound.StackThick2, ironCoreRound.StackThick3, ironCoreRound.StackThick4,
			ironCoreRound.StackThick5, ironCoreRound.StackThick6, ironCoreRound.StackThick7, ironCoreRound.StackThick8,
			ironCoreRound.StackThickSum, ironCoreRound.MainLevelRealStackThick} {
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
