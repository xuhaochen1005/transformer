// Package model
//产品型号
package model

type ProductModel struct {
	ID                              int    `json:"id" gorm:"column:id;primary_key"`                                                        // 编号
	Name                            string `json:"name" gorm:"column:name"`                                                                // 产品型号名称
	Info                            string `json:"info" gorm:"column:info"`                                                                // 备注信息
	DryChangeLowVOuterMoldID        int    `json:"dry_change_lowv_wind_outer_mold_id" gorm:"column:dry_change_lowv_wind_outer_mold_id;"`   //干变低压外模ID
	DryChangeWindingInternalrMoldID int    `json:"dry_change_winding_internal_mold_id" gorm:"column:dry_change_winding_internal_mold_id;"` //干变高压内模ID
	DryChangeWindingOuterMoldID     int    `json:"dry_change_winding_outer_mold_id" gorm:"column:dry_change_winding_outer_mold_id;"`       //干变高压外模ID
	DryChangeLowvWindMoldID         int    `json:"dry_change_lowv_wind_mold_id" gorm:"column:dry_change_lowv_wind_mold_id;"`               // 干变低压箔绕模具ID
}

func (ProductModel) TableName() string {
	return "product_model"
}
