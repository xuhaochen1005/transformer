package model

type Feedback struct {
	Id          int       `json:"id" gorm:"column:id"`
	Creator     int       `json:"creator" gorm:"column:creator"`
	CreatorUser *User     `json:"creator_user" gorm:"<-:false;foreignKey:Creator"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content" gorm:"column:content"`
	CreatedAt   Timestamp `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   Timestamp `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   Timestamp `json:"deleted_at" gorm:"column:deleted_at"`
}

func (Feedback) TableName() string {
	return "feedback"
}
