// Package user 用户信息管理
package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/libs/ssha"
	"transformer_mz/libs/token"
	"transformer_mz/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xuri/excelize/v2"
)

// CreateUser 添加用户信息
func CreateUser(c *gin.Context) {
	err := permission.Check(c, resource.User, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	newUser := &model.User{}
	err = errors.New(c.ShouldBindBodyWith(newUser, binding.JSON), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	form := &struct {
		Password string `json:"password" validate:"required,min=8,max=24"` // 用户密码
	}{
		Password: "",
	}
	err = errors.New(c.ShouldBindBodyWith(form, binding.JSON), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	newUser.Password = form.Password
	newUser.Password, err = ssha.GenerateSSHA(newUser.Password)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Omit("department_name").Create(newUser).Error, "用户信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: newUser})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
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
	user := &model.User{}
	err = errors.New(c.ShouldBindBodyWith(user, binding.JSON), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	form := &struct {
		Password string `json:"password" validate:"required,min=8,max=24"` // 用户密码
	}{
		Password: "",
	}
	err = errors.New(c.ShouldBindBodyWith(form, binding.JSON), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	user.Password = form.Password
	if id != self.ID {
		err = permission.Check(self.Name, resource.User, action.Write)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	update := []string{"name", "email", "telephone", "department_id", "status", "info"}
	if user.Password != "" {
		update = append(update, "password")
		user.Password, err = ssha.GenerateSSHA(user.Password)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	err = errors.New(connector.DataSource.Model(user).Where("id=?", id).
		Select(update).UpdateColumns(user).Error, "用户信息更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "用户信息更新成功"})
}

// DeleteUser 删除用户信息
func DeleteUser(c *gin.Context) {
	err := permission.Check(c, resource.User, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.User{}).Error,
		"用户信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "用户信息删除成功"})
}

// FindUsers 查找用户信息
func FindUsers(c *gin.Context) {
	var users []model.User
	err := permission.Check(c, resource.User, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	db := connector.DataSource.Joins("LEFT JOIN department ON department.id=user.department_id").
		Select([]string{"user.id as id", "user.name as name", "user.user_name_zh as user_name_zh", "department.name as department_name", "department_id",
			"email", "telephone", "user.status as status", "info", "user.created_at as created_at",
			"user.updated_at as updated_at"})
	role_name := c.DefaultQuery("role_name", "")
	if role_name != "" {
		db = db.Where("user.name not in (?)",
			connector.DataSource.Model(&model.UserRole{}).Select("username").Where("role_name = ?", role_name))
	}
	roles := c.QueryArray("roles")
	if len(roles) > 0 {
		db = db.Where("user.name in (?)",
			connector.DataSource.Model(&model.UserRole{}).Select("username").Where("role_name in (?)", roles))
	}
	total, err := connector.GetRecords(c, db, &users)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New(err, "用户信息查找失败，请稍候再试")})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: users})
}

// GetSpecUser 查找单个用户的信息
func GetSpecUser(c *gin.Context) {
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
	if id == self.ID {
		c.Set("response", &model.Response{Spec: self})
	} else {
		err = permission.Check(c, resource.User, action.Read)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		user := &model.User{}
		err = errors.New(connector.DataSource.Where("id=?", id).First(user).Error, "用户信息查询失败，请检查")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
		c.Set("response", &model.Response{Spec: user})
	}
}

// GetSpecUser 查找单个用户的信息
func GetBasicUser(c *gin.Context) {
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	err = permission.Check(c, resource.User, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	user := &model.User{}
	db := connector.DataSource.Joins("JOIN department ON department.id=user.department_id").
		Select([]string{"user.id as id", "user.name as name", "user.user_name_zh as user_name_zh", "department.name as department_name", "department_id",
			"email", "telephone", "user.status as status", "info", "user.created_at as created_at",
			"user.updated_at as updated_at"}).Where("user.id=?", self.ID)
	err = errors.New(db.First(user).Error, "用户信息查询失败，请检查")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}

	c.Set("response", &model.Response{Spec: user, Total: 1})
}

//func UserResetPassword(c *gin.Context) {
//
//	self, err := utils.GetUserFromRequest(c)
//	if err != nil {
//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
//		return
//	}
//	type resetUserPwdStruct struct {
//		Password string `json:"password"`
//		NewPassword string `json:"new_password"`
//		RepeatPassword string `json:"repeat_password"`
//	}
//	resetUserPwd := &resetUserPwdStruct{}
//	err = errors.New(c.BindJSON(&resetUserPwd), "请求数据有误，请联系开发人员进行处理")
//	if err != nil {
//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
//		return
//	}
//
//
//}

type User struct {
	Username string `json:"username" validate:"required,min=4,max=24"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

// CreateUserToken 为用户创建登录所需的token
func CreateUserToken(c *gin.Context) {
	userInfo := &User{}
	err := errors.New(c.BindJSON(userInfo), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	user := &model.User{}
	err = errors.New(connector.DataSource.Where("name=?", userInfo.Username).First(user).Error,
		"用户信息查询失败，请检查")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	valid, err := ssha.ValidateSSHA(userInfo.Password, user.Password)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if !valid {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: errors.New("用户密码不正确，请检查")})
		return
	}
	tokenStr, err := token.Create(user.ID)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = connector.DataSource.First(user).Error
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New("用户未找到")})
		return
	}
	if user.Status != model.UserStatusNormal {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New("账号被禁用，无法登录")})
		return
	}
	data, _ := json.Marshal(&user)
	result := make(map[string]interface{})
	json.Unmarshal(data, &result)
	result["token"] = tokenStr
	c.Set("response", &model.Response{Spec: result})
}

// ResetUserPassword 重置用户密码
func ResetUserPassword(c *gin.Context) {
	err := permission.Check(c, resource.User, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	userInfo := &User{}
	err = errors.New(c.BindJSON(userInfo), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	userInfo.Password, err = ssha.GenerateSSHA(userInfo.Password)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.User{}).Where("id=?", id).
		Update("password", userInfo.Password).Error, "重置用户密码失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "用户密码重置成功"})
}

// GetUserRoles 获取用户自己关联的角色列表
func GetUserRoles(c *gin.Context) {
	user, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	roles, err := permission.Permission.GetRolesForUser(user.Name)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New(err, "用户角色列表获取失败，请联系开发人员进行处理")})
		return
	}
	c.Set("response", &model.Response{Total: int64(len(roles)), Spec: roles})
}

// GetUserPermissions 获取用户自己关联的权限列表
func GetUserPermissions(c *gin.Context) {
	user, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	permissions, err := permission.Permission.GetImplicitPermissionsForUser(user.Name)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError,
			Err: errors.New(err, "用户权限列表获取失败，请联系开发人员进行处理")})
		return
	}
	c.Set("response", &model.Response{Total: int64(len(permissions)), Spec: permissions})
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
