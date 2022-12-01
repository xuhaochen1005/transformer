// Package end_insulation_voltage_type 端绝缘电压类型
package end_insulation_voltage_type

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"

	"github.com/gin-gonic/gin"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// CreateEndInsulationVType 添加端绝缘电压类型
func CreateEndInsulationVType(c *gin.Context) {
	err := permission.Check(c, resource.Libraries, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	newEndInsulationVType := &model.EndInsulationVType{}
	err = errors.New(c.BindJSON(newEndInsulationVType), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newEndInsulationVType).Error, "端绝缘电压类型信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "端绝缘电压类型信息添加成功"})
}

// UpdateEndInsulationVType 更新端绝缘电压类型
func UpdateEndInsulationVType(c *gin.Context) {
	var endInsulationVType *model.EndInsulationVType
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
	err = errors.New(c.BindJSON(endInsulationVType), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.EndInsulationVType{}).Where("id=?", id).
		Updates(&endInsulationVType).Error, "端绝缘电压类型更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "端绝缘电压类型信息更新成功"})
}

// DeleteEndInsulationVType 删除端绝缘电压类型
func DeleteEndInsulationVType(c *gin.Context) {
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
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.EndInsulationVType{}).Error,
		"端绝缘电压类型信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "端绝缘电压类型信息删除成功"})
}

// GetSpecEndInsulationVType 查找特定端绝缘电压类型
func GetSpecEndInsulationVType(c *gin.Context) {
	var endInsulationVType model.EndInsulationVType
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
	err = errors.New(connector.DataSource.Where("id=?", id).First(&endInsulationVType).Error,
		"端绝缘电压类型信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: endInsulationVType})
}

// FindEndInsulationVTypes 查找端绝缘电压类型
func FindEndInsulationVTypes(c *gin.Context) {
	var endInsulationVTypes []model.EndInsulationVType
	err := permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource, &endInsulationVTypes)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: endInsulationVTypes})
}

/**
//todo 文件excel后端下载
使用后端excel 库为
https://github.com/qax-os/excelize
回写到web 流媒体 形成下载
*/
func ExcelExport(c *gin.Context) {
	//todo step1 解析前端参数，获取文件请求下载对象
	// todo step2 构造sql语句，查询出目标数据
	//todo step3 构造excel 文件，将数据写入目标文件
	//todo step4 文件流写入
	// 参考核心代码如下
	/*
		f := excelize.NewFile()
		// Create a new sheet.
		index := f.NewSheet("Sheet2")
		// Set value of a cell.
		f.SetCellValue("Sheet2", "A2", "Hello world.")
		f.SetCellValue("Sheet1", "B2", 100)
		// Set active sheet of the workbook.
		f.SetActiveSheet(index)
		// Save spreadsheet by the given path.
		if err := f.SaveAs("Book1.xlsx"); err != nil {
			fmt.Println(err)
		}
		f.WriteTo(c.Writer)



	*/

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	f.WriteTo(c.Writer)

}
