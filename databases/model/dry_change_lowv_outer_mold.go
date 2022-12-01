// Package model
// 干变低压外膜台账
package model

type DryChangeLowVOuterMold struct {
	ID                int       `json:"id" gorm:"column:id;primary_key"`                       // 编号
	MoldOuterDiameter string    `json:"mold_outer_diameter" gorm:"column:mold_outer_diameter"` //外模外径
	MoldType          int       `json:"mold_type" gorm:"column:mold_type"`                     //模具类型
	MoldHeight        Timestamp `json:"mold_height" gorm:"column:mold_height"`                 //模具高度
	MoldDrawings      Timestamp `json:"mold_drawings" gorm:"column:mold_drawings"`             //模具图号
	Number            Timestamp `json:"number" gorm:"column:number"`                           //模具数量
	Info              Timestamp `json:"info" gorm:"column:info"`                               //备注
	MoldStatus        Timestamp `json:"mold_status" gorm:"column:mold_status"`                 //是否报废

}

func (DryChangeLowVOuterMold) TableName() string {
	return "dry_change_low_voltage_outer_mold"
}
