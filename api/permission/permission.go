// Package permission 权限规则管理
package permission

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

// CreateRolePermission 添加角色权限
func CreateRolePermission(c *gin.Context) {
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	form := &struct {
		RolePermissions []*model.RolePermission `json:"role_permissions"`
	}{}
	err = errors.New(c.BindJSON(form), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	var newPolicys []model.RolePermission
	err = connector.DataSource.Transaction(func(tx *gorm.DB) error {
		for _, rp := range form.RolePermissions {
			if rp.Role == "" || rp.Resource == "" || rp.Action == "" {
				continue
			}
			err := tx.Where(rp).Find(rp).Limit(1).Error
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusBadRequest,
					Err: errors.New(err, "角色权限添加失败，请检查请求数据是否正确")})
				return err
			}
			// 添加前先做存在查询
			if rp.ID == 0 {
				err = errors.New(connector.DataSource.Create(rp).Error, "角色权限添加失败，请稍后再试")
				if err != nil {
					c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
					return err
				}
				_, err = permission.Permission.AddPolicy(rp.Role, rp.Resource, rp.Action)
				if err != nil {
					for _, addedPolicy := range newPolicys {
						permission.Permission.RemovePolicy(addedPolicy.Role, addedPolicy.Resource, addedPolicy.Action)
					}
					c.Set("response", &model.Response{Code: http.StatusBadRequest,
						Err: errors.New(err, "角色权限添加失败，请检查请求数据是否正确")})
					return err
				}
				newPolicys = append(newPolicys, *rp)
			}
		}
		return nil
	})
	if err != nil {
		return
	}

	c.Set("response", &model.Response{Spec: "角色权限添加成功"})
}

// DeleteRolePermission 删除角色权限
func DeleteRolePermission(c *gin.Context) {
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	idSplitStr := c.Param("id")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	var idArr []int
	idStrArr := strings.Split(idSplitStr, ",")
	for _, idStr := range idStrArr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusBadRequest,
				Err: errors.New(err, "角色权限删除失败，请检查请求数据是否正确")})
			return
		}
		idArr = append(idArr, id)
	}
	for _, id := range idArr {
		var rolePermission model.RolePermission
		err = connector.DataSource.Transaction(func(tx *gorm.DB) error {
			err = errors.New(connector.DataSource.Where("id=?", id).First(&rolePermission).Error,
				"角色权限查找失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return err
			}
			_, err = permission.Permission.DeletePermissionForUser(rolePermission.Role, rolePermission.Resource, rolePermission.Action)
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusBadRequest,
					Err: errors.New(err, "角色权限删除失败，请检查请求数据是否正确")})
				return err
			}
			err = errors.New(connector.DataSource.Where("id=?", id).Delete(&rolePermission).Error,
				"角色权限删除失败，请稍后再试")
			if err != nil {
				c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
				return err
			}
			return nil
		})
		// transaction commit
		if err != nil {
			return
		}
	}

	c.Set("response", &model.Response{Spec: "角色权限删除成功"})
}

// GetSpecRolePermission 查找特定角色权限
func GetSpecRolePermission(c *gin.Context) {
	var rolePermission model.RolePermission
	err := permission.Check(c, resource.PermissionRule, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&rolePermission).Error,
		"角色权限查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: rolePermission})
}

// FindRolePermission 查找角色权限
func FindRolePermission(c *gin.Context) {
	var rolePermissions []model.RolePermission
	err := permission.Check(c, resource.PermissionRule, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource, &rolePermissions)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: rolePermissions})
}

// CreateUserRole 添加用户角色关联
func CreateUserRole(c *gin.Context) {
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	newUserRole := &model.UserRole{}
	err = errors.New(c.BindJSON(newUserRole), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	_, err = permission.Permission.AddGroupingPolicy(newUserRole.Username, newUserRole.RoleName)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest,
			Err: errors.New(err, "用户角色关联添加失败，请检查请求数据是否正确")})
		return
	}
	err = errors.New(connector.DataSource.Create(newUserRole).Error, "用户角色关联添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "用户角色关联添加成功"})
}

// DeleteUserRole 删除用户角色关联
func DeleteUserRole(c *gin.Context) {
	var userRole model.UserRole
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&userRole).Error,
		"用户角色关联查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	_, err = permission.Permission.DeleteRoleForUser(userRole.Username, userRole.RoleName)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest,
			Err: errors.New(err, "用户角色关联删除失败，请检查请求数据是否正确")})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&userRole).Error,
		"用户角色关联删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "用户角色关联删除成功"})
}

// GetSpecUserRole 查找特定用户角色关联
func GetSpecUserRole(c *gin.Context) {
	var userRole model.UserRole
	err := permission.Check(c, resource.PermissionRule, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&userRole).Error,
		"用户角色关联查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: userRole})
}

func UpdateUserRole(c *gin.Context) {
	var userRole model.UserRole
	var userRoleForm model.UserRole
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&userRoleForm), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&userRole).Updates(&userRoleForm).Error, "用户角色关联查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "用户角色关联更新成功"})
}

//// FindUserRole 查找用户角色关联
//func FindUserRole(c *gin.Context) {
//	var userRoles []model.UserRole
//	err := permission.Check(c, resource.PermissionRule, action.Read)
//	if err != nil {
//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
//		return
//	}
//	total, err := connector.GetRecords(c, connector.DataSource, &userRoles)
//	if err != nil {
//		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
//		return
//	}
//	c.Set("response", &model.Response{Total: total, Spec: userRoles})
//}

// FindUserRole 按用户分组查找角色
func FindUserRole(c *gin.Context) {
	err := permission.Check(c, resource.PermissionRule, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	type userRoleList struct {
		Username   string            `json:"username"`
		UserNameZh string            `json:"user_name_zh"`
		UserRole   []*model.UserRole `json:"user_role" gorm:"-"`
	}
	var users []model.User
	var userRoleLists []*userRoleList
	db := connector.DataSource.Joins("left join user_role on user.name = user_role.username").
		Select([]string{"user.name", "user.user_name_zh"}).
		Group("user.name").
		Order("count(user.name) DESC")
	total, err := connector.GetRecords(c, db, &users)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	isTemporary, err := strconv.Atoi(c.DefaultQuery("field_eq_temporary", "0"))
	if err != nil {
		err = errors.New(err, "获取数据失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	for _, user := range users {
		userRoleList := &userRoleList{user.Name, user.UserNameZh, []*model.UserRole{}}
		userRoleLists = append(userRoleLists, userRoleList)
		err := errors.New(connector.DataSource.Model(&model.UserRole{}).
			Where("username = ? and temporary = ?", user.Name, isTemporary).
			Find(&userRoleList.UserRole).Error, "查找角色关联失败，请稍后再试")
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			break
		}
	}
	if err != nil {
		return
	}
	c.Set("response", &model.Response{Spec: userRoleLists, Total: total})
}

// FindRoles 查找所有Role
func FindRoles(c *gin.Context) {
	var roles []model.Role
	err := permission.Check(c, resource.PermissionRule, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	db := connector.DataSource
	username := c.DefaultQuery("username", "")
	if username != "" {
		db = db.Where("name not in (?)",
			connector.DataSource.Model(&model.UserRole{}).Select("role_name").Where("username = ?", username))
	}
	total, err := connector.GetRecords(c, db, &roles)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: &roles})
}

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	newRole := &model.Role{}
	err = errors.New(c.BindJSON(newRole), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Create(newRole).Error, "用户角色关联添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "角色添加成功"})
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	var role model.Role
	var roleForm model.Role
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&roleForm), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&role).Updates(&roleForm).Error, "角色查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "角色更新成功"})
}

// DeleteRole 删除角色
func DeleteRole(c *gin.Context) {
	var role model.Role
	err := permission.Check(c, resource.PermissionRule, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&role).Error,
		"角色查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	_, err = permission.Permission.DeleteRole(role.Name)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest,
			Err: errors.New(err, "角色删除失败，请检查请求数据是否正确")})
		return
	}
	var rolePermission model.RolePermission
	err = errors.New(connector.DataSource.Where("role=?", &role.Name).Delete(&rolePermission).Error,
		"角色删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	var userRole model.UserRole
	err = errors.New(connector.DataSource.Where("role_name=?", &role.Name).Delete(&userRole).Error,
		"角色删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&role).Error,
		"角色删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "角色删除成功"})
}
