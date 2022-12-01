// Package model 消息通知数据库模型
package model

const (
	UnRead  = 1
	HasRead = 2
)

type Message struct {
	ID         int       `json:"id" gorm:"column:id;primary_key"` // 消息编号
	Title      string    `json:"title" gorm:"column:title"`       // 消息标题
	Content    string    `json:"content" gorm:"column:content"`   // 消息内容
	SendBy     int       `json:"send_by" gorm:"column:send_by"`   // 消息发送人
	SendByUser *User     `json:"send_by_user,omitempty" gorm:"<-:false;ForeignKey:SendBy;AssociationForeignKey:ID"`
	SendTo     int       `json:"send_to" gorm:"column:send_to"` // 消息接收者
	SendToUser *User     `json:"send_to_user,omitempty" gorm:"<-:false;ForeignKey:SendTo;AssociationForeignKey:ID"`
	Status     int       `json:"status" gorm:"column:status"`         // 消息状态
	CreatedAt  Timestamp `json:"created_at" gorm:"column:created_at"` // 消息发送时间
	DeletedAt  Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 消息删除时间
	UpdatedAt  Timestamp `json:"updated_at" gorm:"column:updated_at"` // 消息更新时间

}

func (Message) TableName() string {
	return "message"
}
