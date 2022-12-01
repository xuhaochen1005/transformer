// Package model 用户角色关联模型
package model

type UserRole struct {
	ID        int    `json:"id" gorm:"column:id;primary_key"`                                    // 用户角色关联编号
	Username  string `json:"username" validate:"required,min=4,max=24" gorm:"column:username"`   // 用户名
	RoleName  string `json:"role_name" validate:"required,min=4,max=32" gorm:"column:role_name"` // 角色名
	Temporary int    `json:"temporary" gorm:"column:temporary"`                                  //是否是指派权限
}

func (UserRole) TableName() string {
	return "user_role"
}
