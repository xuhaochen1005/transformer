// Package model 文档数据库模型
package model

type DocumentCategory string

const (
	DocumentCategoryDesign   DocumentCategory = "设计文档"
	DocumentCategorySystem   DocumentCategory = "系统文档"
	DocumentCategoryTraining DocumentCategory = "培训文档"
)

type Documents struct {
	ID              int              `json:"id" gorm:"column:id;primary_key"`                          // 文档编号
	Name            string           `json:"name" validate:"required,min=4,max=24" gorm:"column:name"` // 文档名称
	UUID            string           `json:"uuid" validate:"required,gte=1,lte=2" gorm:"column:uuid"`  // 文档uuid
	Location        string           `json:"location" gorm:"column:location"`                          // 文档路径
	DocsCreator     int              `json:"docs_creator" gorm:"column:docs_creator"`                  // 创建者
	CreatorUser     *User            `json:"creator_user,omitempty" gorm:"ForeignKey:DocsCreator;AssociationForeignKey:ID"`
	DocumentType    string           `json:"document_type" gorm:"column:document_type"`         //文档类型
	FileSize        float64          `json:"file_size" gorm:"column:file_size"`                 //文档大小
	Category        DocumentCategory `json:"category" gorm:"column:category"`                   //文档归属
	DesignProjectID int              `json:"design_project_id" gorm:"column:design_project_id"` //文件关联设计项目
	DesignProject   *DesignProject   `json:"design_project" gorm:"<-:false;ForeignKey:DesignProjectID"`
	CreatedAt       Timestamp        `json:"created_at" gorm:"column:created_at"` // 用户创建时间
	UpdatedAt       Timestamp        `json:"updated_at" gorm:"column:updated_at"` //  文档更新时间
	DeletedAt       Timestamp        `json:"deleted_at" gorm:"column:deleted_at"` // 文档删除时间

}

func (Documents) TableName() string {
	return "documents"
}
