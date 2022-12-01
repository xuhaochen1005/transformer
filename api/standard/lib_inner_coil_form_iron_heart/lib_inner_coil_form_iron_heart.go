// Package lib_inner_coil_form_iron_heart
// 内线圈至铁芯距离API
package lib_inner_coil_form_iron_heart

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

// CreateInnerCoilFormIronHeart 添加内线圈至铁芯距离规则
func CreateInnerCoilFormIronHeart(c *gin.Context) {
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
	newInnerCoilFormIronHeart := &model.InnerCoilFormIronHeart{}
	newInnerCoilFormIronHeart.LicfihCreator = self.ID
	err = errors.New(c.BindJSON(newInnerCoilFormIronHeart), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newInnerCoilFormIronHeart).Error, "内线圈至铁芯距离记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "内线圈至铁芯距离记录信息添加成功"})
}

// UpdateInnerCoilFormIronHeart 更新内线圈至铁芯距离规则
func UpdateInnerCoilFormIronHeart(c *gin.Context) {
	var InnerCoilFormIronHeart *model.InnerCoilFormIronHeart
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
	err = errors.New(c.BindJSON(&InnerCoilFormIronHeart), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.InnerCoilFormIronHeart{}).Select("rated_high_vol_min", "rated_high_vol_max", "inner_coil_form_iron_heart_min").Where("id=?", id).
		Updates(&InnerCoilFormIronHeart).Error, "内线圈至铁芯距离记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "内线圈至铁芯距离记录信息更新成功"})
}

// DeleteInnerCoilFormIronHeart 删除内线圈至铁芯距离规则
func DeleteInnerCoilFormIronHeart(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.InnerCoilFormIronHeart{}).Error,
		"内线圈至铁芯距离记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "内线圈至铁芯距离记录信息删除成功"})
}

// GetSpecInnerCoilFormIronHeart 查找内线圈至铁芯距离规则
func GetSpecInnerCoilFormIronHeart(c *gin.Context) {
	var InnerCoilFormIronHeart model.InnerCoilFormIronHeart
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&InnerCoilFormIronHeart).Error,
		"内线圈至铁芯距离记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: InnerCoilFormIronHeart})
}

// FindInnerCoilFormIronHearts 查找内线圈至铁芯距离规则
func FindInnerCoilFormIronHearts(c *gin.Context) {
	var InnerCoilFormIronHearts []model.InnerCoilFormIronHeart
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &InnerCoilFormIronHearts)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: InnerCoilFormIronHearts})
}

// ExcelExport 导出内线圈至铁芯距离规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"涡内线圈至铁芯距离编号", "额定高压下界（kV）", "额定高压上界（kV）", "内线圈至铁芯距离最小值（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var innerCoilFormIronHearts []model.InnerCoilFormIronHeart
	_, err = connector.GetRecords(c, connector.DataSource, &innerCoilFormIronHearts)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, innerCoilFormIronHeart := range innerCoilFormIronHearts {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{innerCoilFormIronHeart.ID, innerCoilFormIronHeart.RatedHighVolMin, innerCoilFormIronHeart.RatedHighVolMax, innerCoilFormIronHeart.InnerCoilFormIronHeartMin} {
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
