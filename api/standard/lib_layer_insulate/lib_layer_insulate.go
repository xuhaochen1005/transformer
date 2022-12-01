// Package lib_layer_insulate 层间绝缘距离API
package lib_layer_insulate

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

// CreateLayerInsulate 添加间绝缘距离规则
func CreateLayerInsulate(c *gin.Context) {
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
	newLayerInsulate := &model.LayerInsulate{}
	newLayerInsulate.LliCreator = self.ID
	err = errors.New(c.BindJSON(newLayerInsulate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newLayerInsulate).Error, "层间绝缘距离记录信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "层间绝缘距离记录信息添加成功"})
}

// UpdateLayerInsulate 更新层间绝缘距离规则
func UpdateLayerInsulate(c *gin.Context) {
	var LayerInsulate *model.LayerInsulate
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
	err = errors.New(c.BindJSON(&LayerInsulate), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.LayerInsulate{}).Select("winding_type", "rated_low_vol_min", "rated_low_vol_max", "layer_vol_min", "layer_vol_max", "layer_insulate").Where("id=?", id).
		Updates(&LayerInsulate).Error, "层间绝缘距离记录更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "层间绝缘距离记录信息更新成功"})
}

// DeleteLayerInsulate 删除层间绝缘距离规则
func DeleteLayerInsulate(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.LayerInsulate{}).Error,
		"层间绝缘距离记录信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "层间绝缘距离记录信息删除成功"})
}

// GetSpecLayerInsulate 查找层间绝缘距离规则
func GetSpecLayerInsulate(c *gin.Context) {
	var LayerInsulate model.LayerInsulate
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&LayerInsulate).Error,
		"层间绝缘距离记录信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: LayerInsulate})
}

// FindLayerInsulates 查找层间绝缘距离规则
func FindLayerInsulates(c *gin.Context) {
	var LayerInsulates []model.LayerInsulate
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser"), &LayerInsulates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: LayerInsulates})
}

// ExcelExport 导出层间绝缘距离规则
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"层间绝缘距离编号", "绕制类型，箔绕/线绕", "额定低压下界（kV）",
		"额定低压上界（kV）", "层间电压下界（V）", "层间电压上界（V）", "层间绝缘距离（mm）"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var layerInsulates []model.LayerInsulate
	_, err = connector.GetRecords(c, connector.DataSource, &layerInsulates)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, layerInsulate := range layerInsulates {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{layerInsulate.ID, layerInsulate.WindingType, layerInsulate.RatedLowVolMin,
			layerInsulate.RatedLowVolMax, layerInsulate.LayerVolMin, layerInsulate.LayerVolMax,
			layerInsulate.LayerInsulate} {
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
