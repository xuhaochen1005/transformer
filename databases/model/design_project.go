package model

type DesignProjectStatus int

const (
	DesignProjectImported DesignProjectStatus = iota + 1
	DesignProjectCreated
	DesignProjectStarted
	DesignProjectFinished //设计任务完成
	DesignProjectReviewUnaccepted
	DesignProjectReviewed
	DesignProjectApproveUnaccepted
	DesignProjectApproved
	DesignProjectCheckUnaccepted
	DesignProjectChecked //最终通过
)

type DesignProjectUsage string

const (
	DesignProjectUsageDefault DesignProjectUsage = "/"  //默认电力变压器
	DesignProjectUsageZL      DesignProjectUsage = "ZL" //励磁用变压器	ZL
	DesignProjectUsageZB      DesignProjectUsage = "ZB" //一般工业用整流变压器ZB
)

type DesignProject struct {
	ID             int    `json:"id" gorm:"column:id;primary_key"`                                      // 设计任务编号
	ProjectName    string `json:"project_name"  gorm:"column:project_name"`                             // 设计任务的名称
	Creator        int    `json:"creator" validate:"required" gorm:"column:creator"`                    // 设计任务发起者
	Designer       int    `json:"designer" validate:"required" gorm:"column:designer"`                  // 设计任务执行者
	DesignerSigned int    `json:"designer_signed" validate:"gte=0,lte=1" gorm:"column:designer_signed"` // 设计人员是否签名
	Reviewer       int    `json:"reviewer" validate:"required" gorm:"column:reviewer"`                  // 设计任务复核者
	ReviewerSigned int    `json:"reviewer_signed" validate:"gte=0,lte=1" gorm:"column:reviewer_signed"` // 复核者是否签名
	ReviewAt       int    `json:"review_at" validate:"required" gorm:"column:review_at"`                //复核时间
	//Approve        int    `json:"approve" validate:"required,nefield=Designer" gorm:"column:approve"` // 设计任务审核者
	ProjectStatus DesignProjectStatus `json:"project_status" validate:"gte=1,lte=10" gorm:"column:project_status"` // 设计任务执行状态
	//Comment         string    `json:"comment" gorm:"column:design_comment"`                                                        // 设计任务取消，复核不通过或者审批不通过的原因
	DesignComment   string    `json:"design_comment" gorm:"column:design_comment"`                         // 设计任务备注
	ProductName     string    `json:"product_name" validate:"required" gorm:"column:product_name"`         // 设计任务的名称
	ProductModel    string    `json:"product_model" validate:"required" gorm:"column:product_model"`       // 设计任务发起者
	DeliverAt       int       `json:"deliver_at" validate:"required" gorm:"column:deliver_at"`             // 设计任务执行者
	DesignFinishAt  int       `json:"design_finish_at" validate:"required" gorm:"column:design_finish_at"` // 设计任务复核者
	ProductCode     string    `json:"product_code" validate:"required" gorm:"column:product_code"`         // 设计任务审核者
	TechRequirments string    `json:"tech_requirments" validate:"required" gorm:"column:tech_requirments"` // 特殊要求
	OrderAt         int       `json:"order_at" validate:"required" gorm:"column:order_at"`                 // 设计任务发起者
	DrawingAt       int       `json:"drawing_at" validate:"required" gorm:"column:drawing_at"`             // 设计任务执行者
	DesignAt        int       `json:"design_at" validate:"required" gorm:"column:design_at"`               // 设计任务复核者
	CheckBy         int       `json:"check_by" validate:"required" gorm:"column:check_by"`                 // 设计任务审核者
	DraftingAt      int       `json:"drafting_at" validate:"required" gorm:"column:drafting_at"`           // 设计任务编制时间
	DraftingBy      int       `json:"drafting_by" validate:"required" gorm:"column:drafting_by"`           // 设计任务编制人员
	CheckedAt       int       `json:"checked_at" validate:"required" gorm:"column:checked_at"`             // 设计任务复核者
	SerialCode      string    `json:"serial_code" validate:"required" gorm:"column:serial_code"`           // 任务书编号
	CreatedAt       Timestamp `json:"created_at" gorm:"column:created_at"`                                 //添加时间
	UpdatedAt       Timestamp `json:"updated_at" gorm:"column:updated_at"`                                 // 更新时间
	DeletedAt       Timestamp `json:"deleted_at" gorm:"column:deleted_at"`                                 // 删除时间
	DesignerUser    *User     `json:"designer_user" gorm:"<-:false;foreignKey:Designer"`
	ReviewerUser    *User     `json:"reviewer_user" gorm:"<-:false;foreignKey:Reviewer"`
	CreatorUser     *User     `json:"creator_user" gorm:"<-:false;foreignKey:Creator"`
	CheckUser       *User     `json:"check_user" gorm:"<-:false;foreignKey:CheckBy"`
	DraftingUser    *User     `json:"drafting_user" gorm:"<-:false;foreignKey:DraftingBy"`
	// 设计任务审核者
	ProductUsage                  DesignProjectUsage `json:"product_usage" gorm:"column:product_usage"`                                       //用途
	ProductPhrase                 string             `json:"product_phrase" gorm:"column:product_phrase"`                                     //相数
	ProductRatedCapacity          float32            `json:"product_rated_capacity" gorm:"column:product_rated_capacity"`                     //额定容量
	ProductFrequency              float32            `json:"product_frequency" gorm:"column:product_frequency"`                               //频率
	ProductFrequencySpecial       string             `json:"product_frequency_special" gorm:"column:product_frequency_special"`               //频率特殊
	ProductRatedVHigh             float32            `json:"product_rated_v_high" gorm:"column:product_rated_v_high"`                         //高压侧额定电压
	ProductRatedVLow              float32            `json:"product_rated_v_low" gorm:"column:product_rated_v_low"`                           //低压侧额定电压
	ProductTapRange               string             `json:"product_tap_range" gorm:"column:product_tap_range"`                               //分接范围
	ProductTapRangeSpecial        string             `json:"product_tap_range_special" gorm:"column:product_tap_range_special"`               //分接范围特殊
	ProductInsulationHighLI       float32            `json:"product_insulation_high_li" gorm:"column:product_insulation_high_li"`             //绝缘水平高压雷电
	ProductInsulationLowLI        float32            `json:"product_insulation_low_li" gorm:"column:product_insulation_low_li"`               //绝缘水平低压雷电
	ProductInsulationHighAC       float32            `json:"product_insulation_high_ac" gorm:"column:product_insulation_high_ac"`             //绝缘水平高压工频耐压
	ProductInsulationLowAC        float32            `json:"product_insulation_low_ac" gorm:"column:product_insulation_low_ac"`               //绝缘水平高压工频耐压
	ProductInsulationLevel        string             `json:"product_insulation_level" gorm:"column:product_insulation_level"`                 // 绝缘等级
	ProductInsulationLevelSpecial string             `json:"product_insulation_level_special" gorm:"column:product_insulation_level_special"` // 绝缘等级特殊
	ProductTempLimit              float32            `json:"product_temp_limit" gorm:"column:product_temp_limit"`                             //温升限值
	ProductConnectionGroup        string             `json:"product_connection_group" gorm:"column:product_connection_group"`                 //联结组别
	ProductShortCircuitResistance float32            `json:"product_short_circuit_resistance" gorm:"column:product_short_circuit_resistance"` //短路阻抗
	ProductAltitude               float32            `json:"product_altitude" gorm:"column:product_altitude"`                                 //海拔
	ProductCoolingMethod          string             `json:"product_cooling_method" gorm:"column:product_cooling_method"`                     //冷却方式
	ProductLoadLoss               float32            `json:"product_load_loss" gorm:"column:product_load_loss"`                               //负载损耗
	ProductNoloadLoss             float32            `json:"product_noload_loss"  gorm:"column:product_noload_loss"`                          //空载损耗
	ProductTotalLoss              float32            `json:"product_total_loss" gorm:"column:product_total_loss"`                             //总损耗
	ProductWireMaterial           string             `json:"product_wire_material" gorm:"column:product_wire_material"`                       //导线材质
	ProductIndustryFreqHoldVol    string             `json:"product_industry_freq_hold_vol" gorm:"column:product_industry_freq_hold_vol"`     //工频耐压
	ProductSpecShockVol           string             `json:"product_spec_shock_vol"  gorm:"column:product_spec_shock_vol"`                    //冲击电压
	ProductInductHighVol          float32            `json:"product_induct_high_vol" gorm:"column:product_induct_high_vol"`                   //感应电压
	ProductWindingType            string             `json:"product_wind_type" gorm:"column:product_wind_type"`                               //绕制类型
	ProductWireShape              string             `json:"product_wire_shape" gorm:"column:product_wire_shape"`                             //线圈形状
	NeedMasterApprove             int                `json:"need_master_approve" gorm:"column:need_master_approve"`                           //是否需要总工审核
	ProductIPLevel                string             `json:"product_ip_level" gorm:"column:product_ip_level"`                                 //ip等级
}

func (DesignProject) TableName() string {
	return "design_project"
}
