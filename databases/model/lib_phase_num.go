// Package model
//  相数
package model

type PhaseNum struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`             // 相数编号
	Phase        int       `json:"phase" gorm:"column:phase"`                   //相数
	PhaseNum     string    `json:"phase_num" gorm:"column:phase_num"`           // 相数,三相/单相
	PhaseNumSign string    `json:"phase_num_sign" gorm:"column:phase_num_sign"` // 相数字母代号
	LpnCreator   int       `json:"lpn_creator" gorm:"column:lpn_creator"`       // 创建者
	CreatorUser  *User     `json:"creator_user,omitempty" gorm:"ForeignKey:LpnCreator;AssociationForeignKey:ID"`
	CreatedAt    Timestamp `json:"created_at" gorm:"column:created_at"` // 相数添加时间
	UpdatedAt    Timestamp `json:"updated_at" gorm:"column:updated_at"` // 相数更新时间
	DeletedAt    Timestamp `json:"deleted_at" gorm:"column:deleted_at"` // 相数删除时间
}

func (PhaseNum) TableName() string {
	return "lib_phase_num"
}
