package model

type TodoList struct {
	ID        int        `json:"id" gorm:"column:id;primaryKey;autoIncrement;type:integer;"`
	Title     string     `json:"title" gorm:"column:title;type:varchar(255);not null;comment:标题"`
	Category  string     `json:"category" gorm:"column:category;type:varchar(255);not null;comment:分类"`                       // 分类
	ReferType string     `json:"refer_type" gorm:"column:refer_type;type:varchar(255);not null;default:project;comment:实体类型"` //实体类型
	ReferID   int        `json:"refer_id" gorm:"column:refer_id;type:integer;not null;comment:实体id"`                          // 实体id
	CreatedBy int        `json:"created_by" gorm:"column:created_by;type:integer;not null;comment:创建者"`                       // 创建者
	CheckedBy int        `json:"checked_by" gorm:"column:checked_by;type:integer;not null;comment:创建者"`                       // 确认者/完成者
	CheckedAt *Timestamp `json:"checked_at" gorm:"column:checked_at;type:bigint;comment:确认时间"`                                // 确认时间
	CreatedAt Timestamp  `json:"created_at" gorm:"column:created_at;type:bigint"`                                             // 创建时间
	UpdatedAt Timestamp  `json:"updated_at" gorm:"column:updated_at;type:bigint"`                                             // 更新时间
	DeletedAt Timestamp  `json:"-" gorm:"column:deleted_at;type:bigint"`                                                      // 删除时间
}

func (*TodoList) TableName() string {
	return "todo_list"
}
