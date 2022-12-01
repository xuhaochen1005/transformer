// Package model
//  公式参数
package model

type FormulaParam struct {
	ID                int       `json:"id" gorm:"column:id;primary_key"`                       // 公式参数编号
	FormulaNameZh     string    `json:"formula_name_zh" gorm:"column:formula_name_zh"`         //公式中文名
	FormulaExpress    string    `json:"formula_express" gorm:"column:formula_express"`         //公式中文名
	FormulaName       string    `json:"formula_name" gorm:"column:formula_name"`               // 公式名
	FomulaDescription string    `json:"formula_description" gorm:"column:formula_description"` // 公式
	FormulaParam      string    `json:"formula_param" gorm:"column:formula_param"`             // 参数
	Remark            string    `json:"remark" gorm:"column:remark"`                           // 备注
	CreatedAt         Timestamp `json:"created_at" gorm:"column:created_at"`                   // 系数K添加时间
	UpdatedAt         Timestamp `json:"updated_at" gorm:"column:updated_at"`                   // 系数K更新时间

}

func (FormulaParam) TableName() string {
	return "lib_formula_param"
}
