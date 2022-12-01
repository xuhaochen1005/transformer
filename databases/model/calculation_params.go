// Package model
// 产品设计
package model

type CalculationParams struct {
	ID          int    `json:"id" gorm:"column:id;primary_key"` // 部门编号
	ProductMode string //产品型号

	RatedCapacity         float32   //额定容量 kVA float    如"50.0"
	RatedHighVoltage      float32   //额定高压 kV  float    如"10.0"
	RegulatingRange       []float32 //调压范围 %   float[]  如5级调压"-5.0,-2.5,0.0,2.5,5.0"
	RatedLowVoltage       float32   //额定低压 kV  float    如"0.4"
	RatedFrequency        int       //额定频率 -   int      如"50"
	PhaseNumber           int       //相数     -   int      如"3"
	InsulationLevel       string    //绝缘等级 -   String   如"H"
	AssociatedGroupType   string    //联结组别 -   String   如"Dyn11"
	ColdType              string    //冷却方式 -  String   如"AN"
	UsageCondition        string    //使用条件 -  String   如"户外式"
	IndFreqWithStdVoltage string    //工频耐压 -  String   如"35/0"
	ImpactHighVoltage     string    //冲击高压 -  String   如"75/0"
	InductionHighVoltage  float32   //感应高压 %  float    如"200.0"
	//
	//# 其他
	PullingBoard int ////有无拉板 - String "有"或"无"
	SeamType     int //接缝形式 - String 如"全斜"
	//
	//# 性能参数标准值
	CoilLossSTD            float32 //线圈损耗标准值 W  float 如"800.0"
	LeadLossSTD            float32 //引线损耗标准值 W  float 如"200.0"
	LoadLossSTD            float32 //负载损耗标准值 W  float 如"1000.0"
	NonLoadLossSTD         float32 //空载损耗标准值 W  float 如"270.0"
	TotalLossSTD           float32 //总损耗标准值   W  float 如"1270.0"
	NoneLoadCurrentSTD     float32 //空载电流标准值 %  float 如"1.8"
	ImpedanceVoltage       float32 //阻抗电压标准值 %  float 如"4.0"
	IronCoreTempRiseSTD    float32 //铁心温升标准值 K  float 如"100.0"
	LowVoltageTempRiseSTD  float32 //低压温升标准值 K  float 如"100.0"
	HighVoltageTempRiseSTD float32 //高压温升标准值 K  float 如"100.0"
	PredictedNoiseSTD      float32 //噪音预测标准值 dB float 如"48.0"
	//
	//# 性能参数允许偏差值（有默认值）
	CoilLossBiasMax         float32 //线圈损耗偏差上限 % float 如"5.0"
	CoilLossBiasMin         float32 //线圈损耗偏差下限 % float 如"-5.0"
	LeadLossBiasMax         float32 //引线损耗偏差上限 % float 如"5.0"
	LeadLossBiasMin         float32 //引线损耗偏差下限 % float 如"-5.0"
	LoadLossBiasMax         float32 //负载损耗偏差上限 % float 如"5.0"
	LoadLossBiasMin         float32 //负载损耗偏差下限 % float 如"-5.0"
	NoneLoadLossBiasMax     float32 //空载损耗偏差上限 % float 如"5.0"
	NoneLoadLossBiasMin     float32 //空载损耗偏差下限 % float 如"-5.0"
	TotalLossBiasMax        float32 //总损耗偏差上限   % float 如"5.0"
	TotalLoassBiasMin       float32 //总损耗偏差下限   % float 如"-5.0"
	NoneLoadCurrentBiasMax  float32 //空载电流偏差上限 % float 如"5.0"
	NoneLoadCurrentBiasMin  float32 //空载电流偏差下限 % float 如"-5.0"
	ImpedanceVoltageBiasMax float32 //阻抗电压偏差上限 % float 如"5.0"
	ImpedanceVoltageBiasMin float32 //阻抗电压偏差下限 % float 如"-5.0"
	CoreIronTempRiseMax     float32 //铁心温升偏差上限 % float 如"5.0"
	CoreIronTempRiseMin     float32 //铁心温升偏差下限 % float 如"-5.0"
	LowVoltageTempRiseMax   float32 //低压温升偏差上限 % float 如"5.0"
	LowVoltageTempRiseMin   float32 //低压温升偏差下限 % float 如"-5.0"
	HighVoltageTempRiseMin  float32 //高压温升偏差下限 % float 如"-5.0"
	PredictNoiseMax         float32 //噪音预测偏差上限 % float 如"5.0"
	PredictNoiseMin         float32 //噪音预测偏差下限 % float 如"-5.0"
	//
	//# 经验参数（有默认值）

	StackingFactor        float32 //迭片系数     - float 如"0.96"
	ProcessFactor         float32 //工艺系数     - float 如"1.2"
	KS                    float32 //Ks          - float 如"1.0"
	EddyCurrentLossFactor float32 //涡流损耗系数 % float 如"6.0"
	//
	//# 约束（有默认值）
	LayerVolatageMax   float32 //层间电压上限           v     float  如"650.0"
	LayerVolategeMin   float32 //层间电压上限           v     float  如"400.0"
	SegmentVolatageMax float32 //段间电压上限           v/mm  float  如"140"
	SegmentVolatageMin float32 //段间电压下限           v/mm  float  如"140"
	//轴向裕度上限           %     float  如"3.0"
	//轴向裕度下限           %     float  如"1.5"
	//辐向裕度上限           %     float  如"3.0"
	//辐向裕度下限           %     float  如"1.5"
	//层间绝缘取整位         -    float   如"0.1"
	//箔绕层间绝缘上限       mm    float  如"0.2"
	//箔绕层间绝缘下限       mm    float  如"0.2"
	//线绕层间绝缘上限       mm    float  如"0.5"
	//线绕层间绝缘下限       mm    float  如"0.4"
	//内外绝缘取整位         -     float  如"0.1"
	//线绕内绝缘上限         mm    float  如"1.6"
	//线绕内绝缘下限         mm    float  如"1.6"
	//线绕外绝缘上限         mm    float  如"3.0"
	//线绕外绝缘下限         mm    float  如"3.0"
	//有风道时线绕内绝缘上限  mm    float  如"1.0"
	//有风道时线绕内绝缘下限  mm    float  如"1.0"
	//有风道时线绕外绝缘上限  mm    float  如"1.5"
	//有风道时线绕外绝缘下限  mm    float  如"1.5"
	//箔绕内绝缘上限         mm    float  如"1.0"
	//箔绕内绝缘下限         mm    float  如"0.5"
	//箔绕外绝缘上限         mm    float  如"2.0"
	//箔绕外绝缘下限         mm    float  如"1.5"
	//距铁轭取整位           -     float  如"5"
	//距铁轭上限             mm    float  如""
	//距铁轭下限             mm    float  如"40"
	//端绝缘取整位           -     float  如"5"
	//端绝缘上限             mm    float  如""
	//端绝缘下限             mm    float  如"25"
	//主风道取整位           -     float  如"5"
	//主风道上限             mm    float  如""
	//主风道下限             mm    float  如"30"
	//相间距离取整位         -     float  如"5"
	//相间距离上限           mm    float  如""
	//相间距离下限           mm    float  如"30"
	//低压风道条间距取整位    -     float   如"1"
	//低压风道条间距上限      mm    float  如"45"
	//低压风道条间距下限      mm    float   如"45"
	//高压风道条间间距取整位  -      float   如"1"
	//高压风道条间间上限      mm    float   如"5"
	//高压风道条间间下限      mm    float   如"3"
	//Name      string    `json:"name" gorm:"column:name"`             // 部门名称
	//Status    int       `json:"status" gorm:"column:status"`         // 部门状态
	CreatedAt Timestamp `json:"created_at" gorm:"column:created_at"` // 部门信息添加时间
	UpdatedAt Timestamp `json:"updated_at" gorm:"column:updated_at"` // 部门信息更新时间
}

func (CalculationParams) TableName() string {
	return "calculation_params"
}
