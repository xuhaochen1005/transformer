// Package model
// 标准库表
package model

type Libraries struct {
	ID                   int       `json:"id" gorm:"column:id;primary_key"`
	LibrariesName        string    `json:"libraries_name" gorm:"column:libraries_name"`
	LibrariesShort       string    `json:"libraries_short" gorm:"column:libraries_short"`
	LibrariesDescription string    `json:"libraries_description" gorm:"column:libraries_description"`
	LibrariesSource      string    `json:"libraries_source" gorm:"column:libraries_source"`
	LibrariesCreator     int       `json:"libraries_creator" gorm:"column:libraries_creator"`
	LibrariesStatus      int       `json:"libraries_status" gorm:"column:libraries_status"`
	CreatedAt            Timestamp `json:"created_at" gorm:"column:created_at"`
	UpdatedAt            Timestamp `json:"updated_at" gorm:"column:updated_at"`
	Creator              User      `json:"creator" gorm:"foreignKey:LibrariesCreator"`
}

func (Libraries) TableName() string {
	return "libraries"
}
