// Package lib_stack_perform_data 硅钢片性能数据API
package lib_stack_perform_data

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

// CreateStackPerformData 添加硅钢片性能数据规则
func CreateStackPerformData(c *gin.Context) {
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
	newStackPerformData := &model.StackPerformData{}
	newStackPerformData.LspdCreator = self.ID
	err = errors.New(c.BindJSON(newStackPerformData), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newStackPerformData).Error, "硅钢片性能数据记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "硅钢片性能数据记录信息添加成功"})
}

// UpdateStackPerformData 更新硅钢片性能数据规则
func UpdateStackPerformData(c *gin.Context) {
	var StackPerformData *model.StackPerformData
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
	err = errors.New(c.BindJSON(&StackPerformData), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.StackPerformData{}).Select("stack_type", "core_flux_density", "unit_iron_loss", "unit_mass_magnet_capacity", "unit_area_seam_va").Where("id=?", id).
		Updates(&StackPerformData).Error, "硅钢片性能数据记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "硅钢片性能数据记录信息更新成功"})
}

// DeleteStackPerformData 删除硅钢片性能数据规则
func DeleteStackPerformData(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.StackPerformData{}).Error,
		"硅钢片性能数据记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "硅钢片性能数据记录信息删除成功"})
}

// GetSpecStackPerformData 查找硅钢片性能数据规则
func GetSpecStackPerformData(c *gin.Context) {
	var StackPerformData model.StackPerformData
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&StackPerformData).Error,
		"硅钢片性能数据记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: StackPerformData})
}

// FindStackPerformDatas 查找硅钢片性能数据规则
func FindStackPerformDatas(c *gin.Context) {
	var StackPerformDatas []model.StackPerformData
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &StackPerformDatas)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: StackPerformDatas})
}

// ExcelExport 导出硅钢片性能数据规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"硅钢片性能数据编号", "硅钢片型号", "铁心磁密（T）", "单位铁损（W/kg）", "单位质量磁化容量（VA/kg）", "单位面积接缝伏安值（VA/cm2）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var stackPerformDatas []model.StackPerformData
	_, err = connector.GetRecords(c, connector.DataSource, &stackPerformDatas)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, stackPerformData := range stackPerformDatas {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{stackPerformData.ID, stackPerformData.StackType, stackPerformData.CoreFluxDensity,
			stackPerformData.UnitIronLoss, stackPerformData.UnitMassMagnetCapacity, stackPerformData.UnitAreaSeamVa} {
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
