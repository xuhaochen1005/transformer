// Package model 角色权限模型
package model

type RolePermission struct {
	ID       int    `json:"id" gorm:"column:id;primary_key"`                                  // 角色权限编号
	Role     string `json:"role" validate:"required,min=4,max=32" gorm:"column:role"`         // 角色名
	Resource string `json:"resource" validate:"required,min=4,max=32" gorm:"column:resource"` // 资源名
	Action   string `json:"action" validate:"required,min=4,max=32" gorm:"column:action"`     // 操作名
}

func (RolePermission) TableName() string {
	return "role_permission"
}
