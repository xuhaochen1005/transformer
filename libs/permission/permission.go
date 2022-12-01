// Package permission 基于casbin的权限校验
package permission

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gin-gonic/gin"
	"transformer_mz/databases/connector"
	model2 "transformer_mz/databases/model"
	"transformer_mz/libs/errors"
)

var Permission *casbin.Enforcer

// Check 检查是否有操作权限
func Check(subject interface{}, resource, action string) error {
	username := ""
	switch subject.(type) {
	case *gin.Context:
		userIf, ok := subject.(*gin.Context).Get("user")
		if !ok {
			return errors.New("内部服务出错，请联系开发人员处理")
		}
		username = userIf.(*model2.User).Name
	case string:
		username = subject.(string)
	default:
		return errors.New("不支持的权限检查操作")
	}
	hasPermission, err := Permission.Enforce(username, resource, action)
	if err != nil {
		return errors.New(err, "内部服务出错，请联系开发人员处理")
	}
	if hasPermission {
		return nil
	} else {
		return errors.New("您没有此操作的权限")
	}
}

// Init 初始化权限模型，除了用户与角色关系，角色与权限关系来自数据库，其他的直接配置好
func Init() error {
	var (
		rolePermissions []*model2.RolePermission
		userRoles       []*model2.UserRole
	)
	// 从数据库加载权限规则集合
	err := connector.DataSource.Find(&rolePermissions).Error
	if err != nil {
		return errors.New(err)
	}
	err = connector.DataSource.Find(&userRoles).Error
	if err != nil {
		return errors.New(err)
	}
	// 设置RBAC模型
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", `g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || g(r.sub, "root")`)
	Permission, err = casbin.NewEnforcer(m)
	if err != nil {
		return errors.New(err)
	}
	// 将权限规则添加进模型中
	for _, rolePermission := range rolePermissions {
		// 添加角色对指定对象的操作权限
		_, err = Permission.AddPolicy(rolePermission.Role, rolePermission.Resource, rolePermission.Action)
		if err != nil {
			return errors.New(err)
		}
	}
	for _, userRole := range userRoles {
		// 添加用户与角色的关系
		_, err = Permission.AddGroupingPolicy(userRole.Username, userRole.RoleName)
		if err != nil {
			return errors.New(err)
		}
	}
	return nil
}
