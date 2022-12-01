// Package model 产品设计文档数据库模型
package model

type Product struct {
	ID        int       `json:"id" gorm:"column:id;primary_key"`     // 产品设计文档编号
	Name      string    `json:"name" gorm:"column:name"`             // 产品设计文档名称
	Status    int       `json:"status" gorm:"column:status"`         // 产品设计文档状态
	Info      string    `json:"info" gorm:"column:info"`             // 产品设计文档备注信息
	Location  string    `json:"location" gorm:"column:location"`     // 产品设计文档存储路径
	Creator   int       `json:"creator" gorm:"column:creator_id"`    // 产品设计文档创建者编号
	CreatedAt Timestamp `json:"created_at" gorm:"column:created_at"` // 产品设计文档创建时间
	UpdatedAt Timestamp `json:"updated_at" gorm:"column:updated_at"` // 产品设计文档更新时间
}

func (Product) TableName() string {
	return "product"
}
