// Package model
// 干变绕线POM条台账
package model

type DryChangeWindingPomStrip struct {
	ID        int     `json:"id" gorm:"column:id;primary_key"`   //编号
	Thickness float32 `json:"thickness" gorm:"column:thickness"` //厚度
	Width     float32 `json:"width" gorm:"column:width"`         //宽度
	Length    float32 `json:"length" gorm:"column:length"`       //长度
	Number    int     `json:"number" gorm:"column:number"`       //数量
}

func (DryChangeWindingPomStrip) TableName() string {
	return "dry_change_winding_pom_strip"
}
