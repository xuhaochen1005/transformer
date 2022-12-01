// Package department 部门信息管理
package department

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

// CreateDepartment 添加部门信息
func CreateDepartment(c *gin.Context) {
	err := permission.Check(c, resource.Department, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	newDepartment := &model.Department{}
	err = errors.New(c.BindJSON(newDepartment), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newDepartment).Error, "部门信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "部门信息添加成功"})
}

// UpdateDepartment 更新部门信息
func UpdateDepartment(c *gin.Context) {
	var department model.Department
	err := permission.Check(c, resource.Department, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&department), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.Department{}).Where("id=?", id).Updates(&department).Error,
		"部门信息更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "部门信息更新成功"})
}

// DeleteDepartment 删除部门信息
func DeleteDepartment(c *gin.Context) {
	err := permission.Check(c, resource.Department, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.Department{}).Error,
		"部门信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "部门信息删除成功"})
}

// GetSpecDepartment 查找特定部门信息
func GetSpecDepartment(c *gin.Context) {
	var department model.Department
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&department).Error,
		"部门信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if self.DepartmentID != department.ID {
		err = permission.Check(c, resource.Department, action.Read)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
			return
		}
	}
	c.Set("response", &model.Response{Spec: department})
}

// FindDepartments 查找部门信息
func FindDepartments(c *gin.Context) {
	var departments []model.Department
	err := permission.Check(c, resource.Department, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource, &departments)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: departments})
}

// ExcelExport 导出部门信息
func ExcelExport(c *gin.Context) {
	err := permission.Check(c, resource.Department, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	sheet := "Sheet1"
	file := excelize.NewFile()
	file.SetActiveSheet(file.NewSheet(sheet))
	for idx, value := range []string{"部门编号", "部门名称", "部门状态"} {
		err = file.SetCellValue(sheet, string(rune('A'+idx))+"1", value)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	var departments []model.Department
	_, err = connector.GetRecords(c, connector.DataSource, &departments)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	for index, department := range departments {
		curIndex := index + 2
		strIndex := strconv.Itoa(curIndex)
		for idx, value := range []interface{}{department.ID, department.Name, department.Status} {
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
