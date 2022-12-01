// Package model 消息模板数据库模型
package model

type MessageTemplate struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`   // 消息模板编号
	Name         string    `json:"name" gorm:"column:name"`           // 消息模板名称
	Title        string    `json:"title" gorm:"column:title"`         // 消息模板标题
	Content      string    `json:"content" gorm:"column:content"`     // 消息模板内容
	CreateBy     int       `json:"create_by" gorm:"column:create_by"` // 消息模板创建者
	CreateByUser *User     `json:"create_by_user,omitempty" gorm:"ForeignKey:CreateBy;AssociationForeignKey:ID"`
	CreatedAt    Timestamp `json:"created_at" gorm:"column:created_at"` // 消息模板创建时间
	DeletedAt    Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 消息模板删除时间
	UpdatedAt    Timestamp `json:"updated_at" gorm:"column:updated_at"` // 消息模板更新时间
}

func (MessageTemplate) TableName() string {
	return "message_template"
}
