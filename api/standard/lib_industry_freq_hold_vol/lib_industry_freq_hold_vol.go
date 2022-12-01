// Package lib_industry_freq_hold_vol 工频耐压API
package lib_industry_freq_hold_vol

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

// CreateIndustryFreqHoldVol 添加工频耐压规则
func CreateIndustryFreqHoldVol(c *gin.Context) {
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
	newIndustryFreqHoldVol := &model.IndustryFreqHoldVol{}
	newIndustryFreqHoldVol.LifhvCreator = self.ID
	err = errors.New(c.BindJSON(newIndustryFreqHoldVol), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newIndustryFreqHoldVol).Error, "工频耐压记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "工频耐压记录信息添加成功"})
}

// UpdateIndustryFreqHoldVol 更新工频耐压规则
func UpdateIndustryFreqHoldVol(c *gin.Context) {
	var IndustryFreqHoldVol *model.IndustryFreqHoldVol
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
	err = errors.New(c.BindJSON(&IndustryFreqHoldVol), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.IndustryFreqHoldVol{}).Select("rated_high_vol_min", "rated_high_vol_max", "industry_freq_hold_vol").Where("id=?", id).
		Updates(&IndustryFreqHoldVol).Error, "工频耐压记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "工频耐压记录信息更新成功"})
}

// DeleteIndustryFreqHoldVol 删除工频耐压规则
func DeleteIndustryFreqHoldVol(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.IndustryFreqHoldVol{}).Error,
		"工频耐压记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "工频耐压记录信息删除成功"})
}

// GetSpecIndustryFreqHoldVol 查找工频耐压规则
func GetSpecIndustryFreqHoldVol(c *gin.Context) {
	var IndustryFreqHoldVol model.IndustryFreqHoldVol
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&IndustryFreqHoldVol).Error,
		"工频耐压记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: IndustryFreqHoldVol})
}

// FindIndustryFreqHoldVols 查找工频耐压规则
func FindIndustryFreqHoldVols(c *gin.Context) {
	var IndustryFreqHoldVols []model.IndustryFreqHoldVol
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &IndustryFreqHoldVols)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: IndustryFreqHoldVols})
}

// ExcelExport 导出工频耐压规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"工频耐压编号", "额定高压下界（kV）", "额定高压上界（kV）", "工频耐压（kV）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var industryFreqHoldVols []model.IndustryFreqHoldVol
	_, err = connector.GetRecords(c, connector.DataSource, &industryFreqHoldVols)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, industryFreqHoldVol := range industryFreqHoldVols {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{industryFreqHoldVol.ID, industryFreqHoldVol.RatedHighVolMin, industryFreqHoldVol.RatedHighVolMax, industryFreqHoldVol.IndustryFreqHoldVol} {
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
