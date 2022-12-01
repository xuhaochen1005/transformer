// Package lib_magnet_density 铁心磁密Bm初选值API
package lib_magnet_density

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

// CreateMagnetDensity 添加铁心磁密Bm初选值规则
func CreateMagnetDensity(c *gin.Context) {
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
	newMagnetDensity := &model.MagnetDensity{}
	newMagnetDensity.LmdCreator = self.ID
	err = errors.New(c.BindJSON(newMagnetDensity), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newMagnetDensity).Error, "铁心磁密Bm初选值记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁心磁密Bm初选值记录信息添加成功"})
}

// UpdateMagnetDensity 更新铁心磁密Bm初选值规则
func UpdateMagnetDensity(c *gin.Context) {
	var MagnetDensity *model.MagnetDensity
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
	err = errors.New(c.BindJSON(&MagnetDensity), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.MagnetDensity{}).Select("rated_capacity_min", "rated_capacity_max", "magnet_density_min", "magnet_density_max").Where("id=?", id).
		Updates(&MagnetDensity).Error, "铁心磁密Bm初选值记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁心磁密Bm初选值记录信息更新成功"})
}

// DeleteMagnetDensity 删除铁心磁密Bm初选值规则
func DeleteMagnetDensity(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.MagnetDensity{}).Error,
		"铁心磁密Bm初选值记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "铁心磁密Bm初选值记录信息删除成功"})
}

// GetSpecMagnetDensity 查找铁心磁密Bm初选值规则
func GetSpecMagnetDensity(c *gin.Context) {
	var MagnetDensity model.MagnetDensity
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&MagnetDensity).Error,
		"铁心磁密Bm初选值记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: MagnetDensity})
}

// FindMagnetDensitys 查找铁心磁密Bm初选值规则
func FindMagnetDensitys(c *gin.Context) {
	var MagnetDensitys []model.MagnetDensity
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &MagnetDensitys)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: MagnetDensitys})
}

// ExcelExport 导出铁心磁密Bm初选值规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"铁心磁密Bm初选值编号", "额定容量下界（kVA）", "额定容量上界（kVA）", "铁心磁密Bm下界（T）", "铁心磁密Bm上界（T）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var magnetDensitys []model.MagnetDensity
	_, err = connector.GetRecords(c, connector.DataSource, &magnetDensitys)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, magnetDensity := range magnetDensitys {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{magnetDensity.ID, strconv.FormatFloat(float64(magnetDensity.RatedCapacityMin), 'f', 4, 64) + "kVA", magnetDensity.RatedCapacityMax,
			magnetDensity.MagnetDensityMin, magnetDensity.MagnetDensityMax} {
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
