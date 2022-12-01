// Package model
// 干变绕线外膜台账
package model

type DryChangeWindingOuterMold struct {
	ID                      int     `json:"id" gorm:"column:id;primary_key"`           //外模具直径
	MoldDiameter            string  `json:"mold_diameter" gorm:"column:mold_diameter"` //类型
	MoldType                string  `json:"mold_type" gorm:"column:mold_type"`
	MoldHeight              float32 `json:"mold_height" gorm:"column:mold_height"`                               //模高
	MoldWidth               float32 `json:"mold_width" gorm:"column:mold_width"`                                 //模宽
	Number                  int     `json:"number" gorm:"column:number"`                                         //数量(个)
	MoldDrawings            string  `json:"mold_drawings" gorm:"column:mold_drawings"`                           //模具图号
	MoldNumber              int     `json:"mold_number" gorm:"column:mold_number"`                               //模具编号
	Info                    string  `json:"info" gorm:"column:info"`                                             //备注
	TabWidthCoverPlatSize   string  `json:"tab_width_cover_plate_size" gorm:"column:tab_width_cover_plate_size"` //凸台高度(封板尺寸)
	MoleSize                string  `json:"mold_size" gorm:"column:mold_size"`                                   // 部门信息更新时间
	Slotting                string  `json:"slotting" gorm:"column:slotting"`                                     //是否开槽
	Nuts                    string  `json:"nuts" gorm:"column:nuts"`                                             //螺母
	AheadSize               string  `json:"ahead_size" gorm:"column:ahead_size"`                                 //A头尺寸
	TapHoleDistance         string  `json:"tap_hole_distance" gorm:"column:tap_hole_distance"`                   //抽头孔距
	CoverPlateDrawingNumber string  `json:"cover_plate_drawing_number" gorm:"column:cover_plate_drawing_number"` //封板图号

}

func (DryChangeWindingOuterMold) TableName() string {
	return "dry_change_winding_outer_mold"
}
