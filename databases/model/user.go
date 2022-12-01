// Package model 用户信息数据库模型
package model

type UserStatus int

const (
	UserStatusNormal UserStatus = iota + 1
	UserStatusBanned
)

type User struct {
	ID             int         `json:"id" gorm:"column:id;primary_key"`                                          // 用户编号
	UserNameZh     string      `json:"user_name_zh" validate:"required,min=4,max=24" gorm:"column:user_name_zh"` //
	Name           string      `json:"name" validate:"required,min=4,max=24" gorm:"column:name"`                 // 用户名
	Password       string      `json:"-" validate:"required,min=8,max=24" gorm:"column:password"`                // 用户密码
	Telephone      string      `json:"telephone" validate:"required,len=11" gorm:"column:telephone"`             // 用户手机号码
	Email          string      `json:"email" validate:"required,email" gorm:"column:email"`                      // 用户邮箱地址
	DepartmentID   int         `json:"department_id" validate:"required" gorm:"column:department_id"`            // 用户所属部门ID
	DepartmentName string      `json:"department_name" gorm:"column:department_name;<-:false"`                   // 用户所属部门
	Department     *Department `json:"department,omitempty" gorm:"<-:false;ForeignKey:DepartmentID;AssociationForeignKey:ID"`
	Status         UserStatus  `json:"status" gorm:"column:status"`                 // 用户状态
	Info           string      `json:"info" validate:"required" gorm:"column:info"` // 用户备注信息
	CreatedAt      Timestamp   `json:"created_at" gorm:"column:created_at"`         // 用户创建时间
	UpdatedAt      Timestamp   `json:"updated_at" gorm:"column:updated_at"`         // 用户信息更新时间
}

func (User) TableName() string {
	return "user"
}
