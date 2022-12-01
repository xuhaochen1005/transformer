// Package lib_eddy_loss_factor 涡流损耗系数API
package lib_eddy_loss_factor

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

// CreateEddyLossFactor 添加涡流损耗系数规则
func CreateEddyLossFactor(c *gin.Context) {
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
	newEddyLossFactor := &model.EddyLossFactor{}
	newEddyLossFactor.LelfCreator = self.ID
	err = errors.New(c.BindJSON(newEddyLossFactor), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newEddyLossFactor).Error, "涡流损耗系数记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "涡流损耗系数记录信息添加成功"})
}

// UpdateEddyLossFactor 更新涡流损耗系数规则
func UpdateEddyLossFactor(c *gin.Context) {
	var EddyLossFactor *model.EddyLossFactor
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
	err = errors.New(c.BindJSON(&EddyLossFactor), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.EddyLossFactor{}).Select("rated_capacity", "eddy_loss_factor").Where("id=?", id).
		Updates(&EddyLossFactor).Error, "涡流损耗系数记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "涡流损耗系数记录信息更新成功"})
}

// DeleteEddyLossFactor 删除涡流损耗系数规则
func DeleteEddyLossFactor(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.EddyLossFactor{}).Error,
		"涡流损耗系数记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "涡流损耗系数记录信息删除成功"})
}

// GetSpecEddyLossFactor 查找涡流损耗系数规则
func GetSpecEddyLossFactor(c *gin.Context) {
	var EddyLossFactor model.EddyLossFactor
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&EddyLossFactor).Error,
		"涡流损耗系数记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: EddyLossFactor})
}

// FindEddyLossFactors 查找涡流损耗系数规则
func FindEddyLossFactors(c *gin.Context) {
	var EddyLossFactors []model.EddyLossFactor
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &EddyLossFactors)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: EddyLossFactors})
}

// ExcelExport 导出涡流损耗系数规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"涡流损耗系数编号", "额定容量（kVA）", "涡流损耗系数（%）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var eddyLossFactors []model.EddyLossFactor
	_, err = connector.GetRecords(c, connector.DataSource, &eddyLossFactors)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, eddyLossFactor := range eddyLossFactors {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{eddyLossFactor.ID, eddyLossFactor.RatedCapacity, eddyLossFactor.EddyLossFactor} {
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
