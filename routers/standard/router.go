// Package 标准库 用户信息管理路由设置
package standard

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/standard/calculation"
	"transformer_mz/api/standard/coil_structure"
	"transformer_mz/api/standard/coil_structure_preliminary_selection"
	"transformer_mz/api/standard/core_iron_performance"
	"transformer_mz/api/standard/end_insulation"
	"transformer_mz/api/standard/end_insulation_voltage_type"
	"transformer_mz/api/standard/high_voltage_coil_structure_parameters"
	"transformer_mz/api/standard/iron_heart"
	"transformer_mz/api/standard/lib_air_duct_bar_ledger"
	"transformer_mz/api/standard/lib_boss"
	"transformer_mz/api/standard/lib_coil_air_duct_insulate"
	"transformer_mz/api/standard/lib_coil_connect"
	"transformer_mz/api/standard/lib_coil_from_iron_yoke"
	"transformer_mz/api/standard/lib_coil_shape"
	"transformer_mz/api/standard/lib_coil_wire_type"
	"transformer_mz/api/standard/lib_cool_type"
	"transformer_mz/api/standard/lib_core_k"
	"transformer_mz/api/standard/lib_current_density"
	"transformer_mz/api/standard/lib_eddy_loss_factor"
	"transformer_mz/api/standard/lib_flat_wire_size"
	"transformer_mz/api/standard/lib_foil_low_vol_mod_ledger"
	"transformer_mz/api/standard/lib_foil_oblong_mod_ledger"
	"transformer_mz/api/standard/lib_high_vol_out_mod_ledger"
	"transformer_mz/api/standard/lib_industry_freq_hold_vol"
	"transformer_mz/api/standard/lib_inner_coil_form_iron_heart"
	"transformer_mz/api/standard/lib_iron_core_oblong"
	"transformer_mz/api/standard/lib_iron_core_round"
	"transformer_mz/api/standard/lib_layer_insulate"
	"transformer_mz/api/standard/lib_lead_loss"
	"transformer_mz/api/standard/lib_loss"
	"transformer_mz/api/standard/lib_low_vol_out_mod_ledger"
	"transformer_mz/api/standard/lib_magnet_density"
	"transformer_mz/api/standard/lib_main_air_duct"
	"transformer_mz/api/standard/lib_noise_predict"
	"transformer_mz/api/standard/lib_phase_dist"
	"transformer_mz/api/standard/lib_phase_num"
	"transformer_mz/api/standard/lib_rated_freq"
	"transformer_mz/api/standard/lib_regulate_range"
	"transformer_mz/api/standard/lib_regulate_way"
	"transformer_mz/api/standard/lib_resin"
	"transformer_mz/api/standard/lib_resistivity"
	"transformer_mz/api/standard/lib_round_wire_size"
	"transformer_mz/api/standard/lib_shock_hold_vol"
	"transformer_mz/api/standard/lib_stack"
	"transformer_mz/api/standard/lib_stack_factor"
	"transformer_mz/api/standard/lib_stack_perform_data"
	"transformer_mz/api/standard/lib_tech_factor"
	"transformer_mz/api/standard/lib_temp_rise"
	"transformer_mz/api/standard/lib_terminal_insulate"
	"transformer_mz/api/standard/lib_usage"
	"transformer_mz/api/standard/lib_wind_in_mod_ledger"
	"transformer_mz/api/standard/lib_wind_insulate_media"
	"transformer_mz/api/standard/lib_wire"
	"transformer_mz/api/standard/libraries"
	"transformer_mz/api/standard/low_voltage_coil_structure_parameters"
	"transformer_mz/api/standard/product_model"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	standardRouter := rg.Group("/std")

	// 标准库文件
	standardRouter.GET("/libraries", libraries.FindLibrariess)
	standardRouter.POST("/library", libraries.CreateLibraries)
	standardRouter.GET("/library:id", libraries.GetSpecLibraries)
	standardRouter.PATCH("/library:id", libraries.UpdateLibraries)
	standardRouter.DELETE("/library:id", libraries.DeleteLibraries)

	/*
		standardRouter.GET("/product_model", product_model.FindProductModels)
		standardRouter.POST("/product_model", product_model.CreateProductModel)
		standardRouter.GET("/product_model:id", product_model.GetSpecProductModel)
		standardRouter.PATCH("/product_model:id", product_model.UpdateProductModel)
		standardRouter.DELETE("/product_model:id", product_model.DeleteProductModel)
	*/

	//---------线圈结构
	standardRouter.GET("/lcst/", coil_structure.FindCoilStructures)
	standardRouter.POST("/lcst", coil_structure.CreateCoilStructure)
	standardRouter.GET("/lcst/:id", coil_structure.GetSpecCoilStructure)
	standardRouter.PATCH("/lcst/:id", coil_structure.UpdateCoilStructure)
	standardRouter.DELETE("/lcst/:id", coil_structure.DeleteCoilStructure)

	//----线圈结构初选
	standardRouter.GET("/lcsps/", coil_structure_preliminary_selection.FindCoilStructurePreliminarySelections)
	standardRouter.POST("/lcsps", coil_structure_preliminary_selection.CreateCoilStructurePreliminarySelection)
	standardRouter.GET("/lcsps/:id", coil_structure_preliminary_selection.GetSpecCoilStructurePreliminarySelection)
	standardRouter.PATCH("/lcsps/:id", coil_structure_preliminary_selection.UpdateCoilStructurePreliminarySelection)
	standardRouter.DELETE("/lcsps/:id", coil_structure_preliminary_selection.DeleteCoilStructurePreliminarySelection)

	//------ 铁心性能----
	standardRouter.GET("/lcirp/", core_iron_performance.FindCoreIronPerformances)
	standardRouter.POST("/lcirp", core_iron_performance.CreateCoreIronPerformance)
	standardRouter.GET("/lcirp/:id", core_iron_performance.GetSpecCoreIronPerformance)
	standardRouter.PATCH("/lcirp/:id", core_iron_performance.UpdateCoreIronPerformance)
	standardRouter.DELETE("/lcirp/:id", core_iron_performance.DeleteCoreIronPerformance)

	//------ 铁心----
	standardRouter.GET("/lirh/", iron_heart.FindIronHearts)
	standardRouter.POST("/lirh", iron_heart.CreateIronHeart)
	standardRouter.GET("/lirh/:id", iron_heart.GetSpecIronHeart)
	standardRouter.PATCH("/lirh/:id", iron_heart.UpdateIronHeart)
	standardRouter.DELETE("/lirh/:id", iron_heart.DeleteIronHeart)

	//---------端绝缘
	standardRouter.GET("/lei/", end_insulation.FindEndInsulations)
	standardRouter.POST("/lei", end_insulation.CreateEndInsulation)
	standardRouter.GET("/lei/:id", end_insulation.GetSpecEndInsulation)
	standardRouter.PATCH("/lei/:id", end_insulation.UpdateEndInsulation)
	standardRouter.DELETE("/lei/:id", end_insulation.DeleteEndInsulation)

	//---------端绝缘电压类型
	standardRouter.GET("/leivt/", end_insulation_voltage_type.FindEndInsulationVTypes)
	standardRouter.POST("/leivt", end_insulation_voltage_type.CreateEndInsulationVType)
	standardRouter.GET("/leivt/:id", end_insulation_voltage_type.GetSpecEndInsulationVType)
	standardRouter.PATCH("/leivt/:id", end_insulation_voltage_type.UpdateEndInsulationVType)
	standardRouter.DELETE("/leivt/:id", end_insulation_voltage_type.DeleteEndInsulationVType)

	//---------高压线圈结构参数
	standardRouter.GET("/lhvcsp/", high_voltage_coil_structure_parameters.FindHighVcoilStructureParams)
	standardRouter.POST("/lhvcsp", high_voltage_coil_structure_parameters.CreateHighVcoilStructureParam)
	standardRouter.GET("/lhvcsp/:id", high_voltage_coil_structure_parameters.GetSpecHighVcoilStructureParam)
	standardRouter.PATCH("/lhvcsp/:id", high_voltage_coil_structure_parameters.UpdateHighVcoilStructureParam)
	standardRouter.DELETE("/lhvcsp/:id", high_voltage_coil_structure_parameters.DeleteHighVcoilStructureParam)

	//-----------风道条台账
	standardRouter.GET("/ladbl/", lib_air_duct_bar_ledger.FindAirDuctBarLedgers)
	standardRouter.POST("/ladbl", lib_air_duct_bar_ledger.CreateAirDuctBarLedger)
	standardRouter.GET("/ladbl/:id", lib_air_duct_bar_ledger.GetSpecAirDuctBarLedger)
	standardRouter.PATCH("/ladbl/:id", lib_air_duct_bar_ledger.UpdateAirDuctBarLedger)
	standardRouter.DELETE("/ladbl/:id", lib_air_duct_bar_ledger.DeleteAirDuctBarLedger)
	standardRouter.POST("/ladblexport", lib_air_duct_bar_ledger.ExcelExport)

	//-----------线圈接法
	standardRouter.GET("/lcc/", lib_coil_connect.FindCoilConnects)
	standardRouter.POST("/lcc", lib_coil_connect.CreateCoilConnect)
	standardRouter.GET("/lcc/:id", lib_coil_connect.GetSpecCoilConnect)
	standardRouter.PATCH("/lcc/:id", lib_coil_connect.UpdateCoilConnect)
	standardRouter.DELETE("/lcc/:id", lib_coil_connect.DeleteCoilConnect)
	standardRouter.POST("/lcc/export", lib_coil_connect.ExcelExport)

	//-----------线圈端部距铁轭距离
	standardRouter.GET("/lcfiy/", lib_coil_from_iron_yoke.FindCoilFromIronYokes)
	standardRouter.POST("/lcfiy", lib_coil_from_iron_yoke.CreateCoilFromIronYoke)
	standardRouter.GET("/lcfiy/:id", lib_coil_from_iron_yoke.GetSpecCoilFromIronYoke)
	standardRouter.PATCH("/lcfiy/:id", lib_coil_from_iron_yoke.UpdateCoilFromIronYoke)
	standardRouter.DELETE("/lcfiy/:id", lib_coil_from_iron_yoke.DeleteCoilFromIronYoke)
	standardRouter.POST("/lcfiy/export", lib_coil_from_iron_yoke.ExcelExport)

	//-----------线圈形状
	standardRouter.GET("/lcs/", lib_coil_shape.FindCoilShapes)
	standardRouter.POST("/lcs", lib_coil_shape.CreateCoilShape)
	standardRouter.GET("/lcs/:id", lib_coil_shape.GetSpecCoilShape)
	standardRouter.PATCH("/lcs/:id", lib_coil_shape.UpdateCoilShape)
	standardRouter.DELETE("/lcs/:id", lib_coil_shape.DeleteCoilShape)
	standardRouter.POST("/lcs/export", lib_coil_shape.ExcelExport)

	//-----------线圈导线类型
	standardRouter.GET("/lcwt/", lib_coil_wire_type.FindCoilWireTypes)
	standardRouter.POST("/lcwt", lib_coil_wire_type.CreateCoilWireType)
	standardRouter.GET("/lcwt/:id", lib_coil_wire_type.GetSpecCoilWireType)
	standardRouter.PATCH("/lcwt/:id", lib_coil_wire_type.UpdateCoilWireType)
	standardRouter.DELETE("/lcwt/:id", lib_coil_wire_type.DeleteCoilWireType)
	standardRouter.POST("/lcwt/export", lib_coil_wire_type.ExcelExport)

	//-----------冷却方式
	standardRouter.GET("/lct/", lib_cool_type.FindCoolTypes)
	standardRouter.POST("/lct", lib_cool_type.CreateCoolType)
	standardRouter.GET("/lct/:id", lib_cool_type.GetSpecCoolType)
	standardRouter.PATCH("/lct/:id", lib_cool_type.UpdateCoolType)
	standardRouter.DELETE("/lct/:id", lib_cool_type.DeleteCoolType)
	standardRouter.POST("/lct/export", lib_cool_type.ExcelExport)

	//-----------干式变压器铁心直径经验系数K
	standardRouter.GET("/lck/", lib_core_k.FindCoreKs)
	standardRouter.POST("/lck", lib_core_k.CreateCoreK)
	standardRouter.GET("/lck/:id", lib_core_k.GetSpecCoreK)
	standardRouter.PATCH("/lck/:id", lib_core_k.UpdateCoreK)
	standardRouter.DELETE("/lck/:id", lib_core_k.DeleteCoreK)
	standardRouter.POST("/lck/export", lib_core_k.ExcelExport)

	//-----------电流密度
	standardRouter.GET("/lcd/", lib_current_density.FindCurrentDensitys)
	standardRouter.POST("/lcd", lib_current_density.CreateCurrentDensity)
	standardRouter.GET("/lcd/:id", lib_current_density.GetSpecCurrentDensity)
	standardRouter.PATCH("/lcd/:id", lib_current_density.UpdateCurrentDensity)
	standardRouter.DELETE("/lcd/:id", lib_current_density.DeleteCurrentDensity)
	standardRouter.POST("/lcd/export", lib_current_density.ExcelExport)

	//-----------涡流损耗系数
	standardRouter.GET("/lelf/", lib_eddy_loss_factor.FindEddyLossFactors)
	standardRouter.POST("/lelf", lib_eddy_loss_factor.CreateEddyLossFactor)
	standardRouter.GET("/lelf/:id", lib_eddy_loss_factor.GetSpecEddyLossFactor)
	standardRouter.PATCH("/lelf/:id", lib_eddy_loss_factor.UpdateEddyLossFactor)
	standardRouter.DELETE("/lelf/:id", lib_eddy_loss_factor.DeleteEddyLossFactor)
	standardRouter.POST("/lelf/export", lib_eddy_loss_factor.ExcelExport)

	//-----------公式参数
	standardRouter.GET("/calculation/", calculation.FindCalculations)
	standardRouter.POST("/calculation", calculation.CreateCalculation)
	standardRouter.GET("/calculation/:id", calculation.GetSpecCalculation)
	standardRouter.PATCH("/calculation/:id", calculation.UpdateCalculation)
	standardRouter.DELETE("/calculation/:id", calculation.DeleteCalculation)
	standardRouter.POST("/calculation/export", calculation.ExcelExport)

	//-----------工频耐压
	standardRouter.GET("/lifhv/", lib_industry_freq_hold_vol.FindIndustryFreqHoldVols)
	standardRouter.POST("/lifhv", lib_industry_freq_hold_vol.CreateIndustryFreqHoldVol)
	standardRouter.GET("/lifhv/:id", lib_industry_freq_hold_vol.GetSpecIndustryFreqHoldVol)
	standardRouter.PATCH("/lifhv/:id", lib_industry_freq_hold_vol.UpdateIndustryFreqHoldVol)
	standardRouter.DELETE("/lifhv/:id", lib_industry_freq_hold_vol.DeleteIndustryFreqHoldVol)
	standardRouter.POST("/lifhv/export", lib_industry_freq_hold_vol.ExcelExport)

	//-----------内线圈至铁芯距离
	standardRouter.GET("/licfih/", lib_inner_coil_form_iron_heart.FindInnerCoilFormIronHearts)
	standardRouter.POST("/licfih", lib_inner_coil_form_iron_heart.CreateInnerCoilFormIronHeart)
	standardRouter.GET("/licfih/:id", lib_inner_coil_form_iron_heart.GetSpecInnerCoilFormIronHeart)
	standardRouter.PATCH("/licfih/:id", lib_inner_coil_form_iron_heart.UpdateInnerCoilFormIronHeart)
	standardRouter.DELETE("/licfih/:id", lib_inner_coil_form_iron_heart.DeleteInnerCoilFormIronHeart)
	standardRouter.POST("/licfih/export", lib_inner_coil_form_iron_heart.ExcelExport)

	//-----------铁芯（长圆形）
	standardRouter.GET("/lico/", lib_iron_core_oblong.FindIronCoreOblongs)
	standardRouter.POST("/lico", lib_iron_core_oblong.CreateIronCoreOblong)
	standardRouter.GET("/lico/:id", lib_iron_core_oblong.GetSpecIronCoreOblong)
	standardRouter.PATCH("/lico/:id", lib_iron_core_oblong.UpdateIronCoreOblong)
	standardRouter.DELETE("/lico/:id", lib_iron_core_oblong.DeleteIronCoreOblong)
	standardRouter.POST("/lico/export", lib_iron_core_oblong.ExcelExport)

	//-----------铁芯（圆形）
	standardRouter.GET("/licr/", lib_iron_core_round.FindIronCoreRounds)
	standardRouter.POST("/licr", lib_iron_core_round.CreateIronCoreRound)
	standardRouter.GET("/licr/:id", lib_iron_core_round.GetSpecIronCoreRound)
	standardRouter.PATCH("/licr/:id", lib_iron_core_round.UpdateIronCoreRound)
	standardRouter.DELETE("/licr/:id", lib_iron_core_round.DeleteIronCoreRound)
	standardRouter.POST("/licr/export", lib_iron_core_round.ExcelExport)

	//-----------层间绝缘距离
	standardRouter.GET("/lli/", lib_layer_insulate.FindLayerInsulates)
	standardRouter.POST("/lli", lib_layer_insulate.CreateLayerInsulate)
	standardRouter.GET("/lli/:id", lib_layer_insulate.GetSpecLayerInsulate)
	standardRouter.PATCH("/lli/:id", lib_layer_insulate.UpdateLayerInsulate)
	standardRouter.DELETE("/lli/:id", lib_layer_insulate.DeleteLayerInsulate)
	standardRouter.POST("/lli/export", lib_layer_insulate.ExcelExport)

	//-----------引线损耗
	standardRouter.GET("/lls/", lib_lead_loss.FindLeadLosss)
	standardRouter.POST("/lls", lib_lead_loss.CreateLeadLoss)
	standardRouter.GET("/lls/:id", lib_lead_loss.GetSpecLeadLoss)
	standardRouter.PATCH("/lls/:id", lib_lead_loss.UpdateLeadLoss)
	standardRouter.DELETE("/lls/:id", lib_lead_loss.DeleteLeadLoss)
	standardRouter.POST("/lls/export", lib_lead_loss.ExcelExport)

	//-----------损耗
	standardRouter.GET("/ll/", lib_loss.FindLosss)
	standardRouter.POST("/ll", lib_loss.CreateLoss)
	standardRouter.GET("/ll/:id", lib_loss.GetSpecLoss)
	standardRouter.PATCH("/ll/:id", lib_loss.UpdateLoss)
	standardRouter.DELETE("/ll/:id", lib_loss.DeleteLoss)
	standardRouter.POST("/ll/export", lib_loss.ExcelExport)

	//-----------低压外模台账
	standardRouter.GET("/llvoml/", lib_low_vol_out_mod_ledger.FindLowVolOutModLedgers)
	standardRouter.POST("/llvoml", lib_low_vol_out_mod_ledger.CreateLowVolOutModLedger)
	standardRouter.GET("/llvoml/:id", lib_low_vol_out_mod_ledger.GetSpecLowVolOutModLedger)
	standardRouter.PATCH("/llvoml/:id", lib_low_vol_out_mod_ledger.UpdateLowVolOutModLedger)
	standardRouter.DELETE("/llvoml/:id", lib_low_vol_out_mod_ledger.DeleteLowVolOutModLedger)
	standardRouter.POST("/llvoml/export", lib_low_vol_out_mod_ledger.ExcelExport)

	//-----------铁心磁密Bm初选值
	standardRouter.GET("/lmd", lib_magnet_density.FindMagnetDensitys)
	standardRouter.POST("/lmd", lib_magnet_density.CreateMagnetDensity)
	standardRouter.GET("/lmd/:id", lib_magnet_density.GetSpecMagnetDensity)
	standardRouter.PATCH("/lmd/:id", lib_magnet_density.UpdateMagnetDensity)
	standardRouter.DELETE("/lmd/:id", lib_magnet_density.DeleteMagnetDensity)
	standardRouter.POST("/lmd/export", lib_magnet_density.ExcelExport)

	//-----------主风道
	standardRouter.GET("/lmad/", lib_main_air_duct.FindMainAirDucts)
	standardRouter.POST("/lmad", lib_main_air_duct.CreateMainAirDuct)
	standardRouter.GET("/lmad/:id", lib_main_air_duct.GetSpecMainAirDuct)
	standardRouter.PATCH("/lmad/:id", lib_main_air_duct.UpdateMainAirDuct)
	standardRouter.DELETE("/lmad/:id", lib_main_air_duct.DeleteMainAirDuct)
	standardRouter.POST("/lmad/export", lib_main_air_duct.ExcelExport)

	//-----------噪声预测
	standardRouter.GET("/lnp", lib_noise_predict.FindNoisePredicts)
	standardRouter.POST("/lnp", lib_noise_predict.CreateNoisePredict)
	standardRouter.GET("/lnp/:id", lib_noise_predict.GetSpecNoisePredict)
	standardRouter.PATCH("/lnp/:id", lib_noise_predict.UpdateNoisePredict)
	standardRouter.DELETE("/lnp/:id", lib_noise_predict.DeleteNoisePredict)
	standardRouter.POST("/lnp/export", lib_noise_predict.ExcelExport)

	//-----------相数
	standardRouter.GET("/lpn/", lib_phase_num.FindPhaseNums)
	standardRouter.POST("/lpn", lib_phase_num.CreatePhaseNum)
	standardRouter.GET("/lpn/:id", lib_phase_num.GetSpecPhaseNum)
	standardRouter.PATCH("/lpn/:id", lib_phase_num.UpdatePhaseNum)
	standardRouter.DELETE("/lpn/:id", lib_phase_num.DeletePhaseNum)
	standardRouter.POST("/lpn/export", lib_phase_num.ExcelExport)

	//-----------额定频率
	standardRouter.GET("/lrf/", lib_rated_freq.FindRatedFreqs)
	standardRouter.POST("/lrf", lib_rated_freq.CreateRatedFreq)
	standardRouter.GET("/lrf/:id", lib_rated_freq.GetSpecRatedFreq)
	standardRouter.PATCH("/lrf/:id", lib_rated_freq.UpdateRatedFreq)
	standardRouter.DELETE("/lrf/:id", lib_rated_freq.DeletRatedFreq)
	standardRouter.POST("/lrf/export", lib_rated_freq.ExcelExport)

	//-----------调压范围
	standardRouter.GET("/lrr/", lib_regulate_range.FindRegulateRanges)
	standardRouter.POST("/lrr", lib_regulate_range.CreateRegulateRange)
	standardRouter.GET("/lrr/:id", lib_regulate_range.GetSpecRegulateRange)
	standardRouter.PATCH("/lrr/:id", lib_regulate_range.UpdateRegulateRange)
	standardRouter.DELETE("/lrr/:id", lib_regulate_range.DeleteRegulateRange)
	standardRouter.POST("/lrr/export", lib_regulate_range.ExcelExport)

	//-----------调压方式
	standardRouter.GET("/lrw/", lib_regulate_way.FindRegulateWays)
	standardRouter.POST("/lrw", lib_regulate_way.CreateRegulateWay)
	standardRouter.GET("/lrw/:id", lib_regulate_way.GetSpecRegulateWay)
	standardRouter.PATCH("/lrw/:id", lib_regulate_way.UpdateRegulateWay)
	standardRouter.DELETE("/lrw/:id", lib_regulate_way.DeleteRegulateWay)
	standardRouter.POST("/lrw/export", lib_regulate_way.ExcelExport)

	//-----------冲击电压
	standardRouter.GET("/lshv/", lib_shock_hold_vol.FindShockHoldVols)
	standardRouter.POST("/lshv", lib_shock_hold_vol.CreateShockHoldVol)
	standardRouter.GET("/lshv/:id", lib_shock_hold_vol.GetSpecShockHoldVol)
	standardRouter.PATCH("/lshv/:id", lib_shock_hold_vol.UpdateShockHoldVol)
	standardRouter.DELETE("/lshv/:id", lib_shock_hold_vol.DeleteShockHoldVol)
	standardRouter.POST("/lshv/export", lib_shock_hold_vol.ExcelExport)

	//-----------迭片系数
	standardRouter.GET("/lsf/", lib_stack_factor.FindStackFactors)
	standardRouter.POST("/lsf", lib_stack_factor.CreateStackFactor)
	standardRouter.GET("/lsf/:id", lib_stack_factor.GetSpecStackFactor)
	standardRouter.PATCH("/lsf/:id", lib_stack_factor.UpdateStackFactor)
	standardRouter.DELETE("/lsf/:id", lib_stack_factor.DeleteStackFactor)
	standardRouter.POST("/lsf/export", lib_stack_factor.ExcelExport)

	//-----------硅钢片性能数据
	standardRouter.GET("/lspd/", lib_stack_perform_data.FindStackPerformDatas)
	standardRouter.POST("/lspd", lib_stack_perform_data.CreateStackPerformData)
	standardRouter.GET("/lspd/:id", lib_stack_perform_data.GetSpecStackPerformData)
	standardRouter.PATCH("/lspd/:id", lib_stack_perform_data.UpdateStackPerformData)
	standardRouter.DELETE("/lspd/:id", lib_stack_perform_data.DeleteStackPerformData)
	standardRouter.POST("/lspd/export", lib_stack_perform_data.ExcelExport)

	//-----------绕组温升
	standardRouter.GET("/ltr/", lib_temp_rise.FindTempRises)
	standardRouter.POST("/ltr", lib_temp_rise.CreateTempRise)
	standardRouter.GET("/ltr/:id", lib_temp_rise.GetSpecTempRise)
	standardRouter.PATCH("/ltr/:id", lib_temp_rise.UpdateTempRise)
	standardRouter.DELETE("/ltr/:id", lib_temp_rise.DeleteTempRise)
	standardRouter.POST("/ltr/export", lib_temp_rise.ExcelExport)

	//-----------端绝缘距离
	standardRouter.GET("/lti/", lib_terminal_insulate.FindTerminalInsulates)
	standardRouter.POST("/lti", lib_terminal_insulate.CreateTerminalInsulate)
	standardRouter.GET("/lti/:id", lib_terminal_insulate.GetSpecTerminalInsulate)
	standardRouter.PATCH("/lti/:id", lib_terminal_insulate.UpdateTerminalInsulate)
	standardRouter.DELETE("/lti/:id", lib_terminal_insulate.DeleteTerminalInsulate)
	standardRouter.POST("/lti/export", lib_terminal_insulate.ExcelExport)

	//-----------变压器用途
	standardRouter.GET("/lus/", lib_usage.FindUsages)
	standardRouter.POST("/lus", lib_usage.CreateUsage)
	standardRouter.GET("/lus/:id", lib_usage.GetSpecUsage)
	standardRouter.PATCH("/lus/:id", lib_usage.UpdateUsage)
	standardRouter.DELETE("/lus/:id", lib_usage.DeleteUsage)
	standardRouter.POST("/lus/export", lib_usage.ExcelExport)

	//-----------绕线内模台账
	standardRouter.GET("/lwiml/", lib_wind_in_mod_ledger.FindWindInModLedgers)
	standardRouter.POST("/lwiml", lib_wind_in_mod_ledger.CreateWindInModLedger)
	standardRouter.GET("/lwiml/:id", lib_wind_in_mod_ledger.GetSpecWindInModLedger)
	standardRouter.PATCH("/lwiml/:id", lib_wind_in_mod_ledger.UpdateWindInModLedger)
	standardRouter.DELETE("/lwiml/:id", lib_wind_in_mod_ledger.DeleteWindInModLedger)
	standardRouter.POST("/lwiml/export", lib_wind_in_mod_ledger.ExcelExport)

	//-----------绕组外绝缘介质
	standardRouter.GET("/lwim/", lib_wind_insulate_media.FindWindInsulateMedias)
	standardRouter.POST("/lwim", lib_wind_insulate_media.CreateWindInsulateMedia)
	standardRouter.GET("/lwim/:id", lib_wind_insulate_media.GetSpecWindInsulateMedia)
	standardRouter.PATCH("/lwim/:id", lib_wind_insulate_media.UpdateWindInsulateMedia)
	standardRouter.DELETE("/lwim/:id", lib_wind_insulate_media.DeleteWindInsulateMedia)
	standardRouter.POST("/lwim/export", lib_wind_insulate_media.ExcelExport)

	//-----------线圈和风道绝缘
	standardRouter.GET("/lcadi/", lib_coil_air_duct_insulate.FindCoilAirDuctInsulates)
	standardRouter.POST("/lcadi", lib_coil_air_duct_insulate.CreateCoilAirDuctInsulate)
	standardRouter.GET("/lcadi/:id", lib_coil_air_duct_insulate.GetSpecCoilAirDuctInsulate)
	standardRouter.PATCH("/lcadi/:id", lib_coil_air_duct_insulate.UpdateCoilAirDuctInsulate)
	standardRouter.DELETE("/lcadi/:id", lib_coil_air_duct_insulate.DeleteCoilAirDuctInsulate)
	standardRouter.POST("/lcadi/export", lib_coil_air_duct_insulate.ExcelExport)

	//---------低压线圈结构参数
	standardRouter.GET("/llvcsp/", low_voltage_coil_structure_parameters.FindLowVCoilStructureParams)
	standardRouter.POST("/llvcsp", low_voltage_coil_structure_parameters.CreateLowVCoilStructureParam)
	standardRouter.GET("/llvcsp/:id", low_voltage_coil_structure_parameters.GetSpecLowVCoilStructureParam)
	standardRouter.PATCH("/llvcsp/:id", low_voltage_coil_structure_parameters.UpdateLowVCoilStructureParam)
	standardRouter.DELETE("/llvcsp/:id", low_voltage_coil_structure_parameters.DeleteLowVCoilStructureParam)

	//---------产品型号表
	standardRouter.GET("/product_model", product_model.FindProductModels)
	standardRouter.POST("/product_model", product_model.CreateProductModel)
	standardRouter.GET("/product_model/:id", product_model.GetSpecProductModel)
	standardRouter.PATCH("/product_model/:id", product_model.UpdateProductModel)
	standardRouter.DELETE("/product_model/:id", product_model.DeleteProductModel)

	//---------扁铜（铝）线规格
	standardRouter.GET("/lfws/", lib_flat_wire_size.FindFlatWireSize)
	standardRouter.POST("/lfws", lib_flat_wire_size.CreateFlatWireSize)
	standardRouter.GET("/lfws/:id", lib_flat_wire_size.GetSpecFlatWireSize)
	standardRouter.PATCH("/lfws/:id", lib_flat_wire_size.UpdateFlatWireSize)
	standardRouter.DELETE("/lfws/:id", lib_flat_wire_size.DeleteFlatWireSize)
	standardRouter.POST("/lfws/export", lib_flat_wire_size.ExcelExport)

	//---------箔绕低压模具台账
	standardRouter.GET("/lflvml/", lib_foil_low_vol_mod_ledger.FindFoilLowVolModLedger)
	standardRouter.POST("/lflvml", lib_foil_low_vol_mod_ledger.CreateFoilLowVolModLedger)
	standardRouter.GET("/lflvml/:id", lib_foil_low_vol_mod_ledger.GetSpecFoilLowVolModLedger)
	standardRouter.PATCH("/lflvml/:id", lib_foil_low_vol_mod_ledger.UpdateFoilLowVolModLedger)
	standardRouter.DELETE("/lflvml/:id", lib_foil_low_vol_mod_ledger.DeleteFoilLowVolModLedger)
	standardRouter.POST("/lflvml/export", lib_foil_low_vol_mod_ledger.ExcelExport)

	//---------箔绕扁形模具台账
	standardRouter.GET("/lfoml/", lib_foil_oblong_mod_ledger.FindFoilOblongModLedger)
	standardRouter.POST("/lfoml", lib_foil_oblong_mod_ledger.CreateFoilOblongModLedger)
	standardRouter.GET("/lfoml/:id", lib_foil_oblong_mod_ledger.GetSpecFoilOblongModLedger)
	standardRouter.PATCH("/lfoml/:id", lib_foil_oblong_mod_ledger.UpdateFoilOblongModLedger)
	standardRouter.DELETE("/lfoml/:id", lib_foil_oblong_mod_ledger.DeleteFoilOblongModLedger)
	standardRouter.POST("/lfoml/export", lib_foil_oblong_mod_ledger.ExcelExport)

	//---------高压外模台账
	standardRouter.GET("/lhvoml/", lib_high_vol_out_mod_ledger.FindHighVolOutModLedger)
	standardRouter.POST("/lhvoml", lib_high_vol_out_mod_ledger.CreateHighVolOutModLedger)
	standardRouter.GET("/lhvoml/:id", lib_high_vol_out_mod_ledger.GetSpecHighVolOutModLedger)
	standardRouter.PATCH("/lhvoml/:id", lib_high_vol_out_mod_ledger.UpdateHighVolOutModLedger)
	standardRouter.DELETE("/lhvoml/:id", lib_high_vol_out_mod_ledger.DeleteHighVolOutModLedger)
	standardRouter.POST("/lhvoml/export", lib_high_vol_out_mod_ledger.ExcelExport)

	//---------圆铜（铝）线规格
	standardRouter.GET("/lrws", lib_round_wire_size.FindRoundWireSize)
	standardRouter.POST("/lrws", lib_round_wire_size.CreateRoundWireSize)
	standardRouter.GET("/lrws/:id", lib_round_wire_size.GetSpecRoundWireSize)
	standardRouter.PATCH("/lrws/:id", lib_round_wire_size.UpdateRoundWireSize)
	standardRouter.DELETE("/lrws/:id", lib_round_wire_size.DeleteRoundWireSize)
	standardRouter.POST("/lrws/export", lib_round_wire_size.ExcelExport)

	//---------工艺系数
	standardRouter.GET("/ltf", lib_tech_factor.FindTechFactor)
	standardRouter.POST("/ltf", lib_tech_factor.CreateTechFactor)
	standardRouter.GET("/ltf/:id", lib_tech_factor.GetSpecTechFactor)
	standardRouter.PATCH("/ltf/:id", lib_tech_factor.UpdateTechFactor)
	standardRouter.DELETE("/ltf/:id", lib_tech_factor.DeleteTechFactor)
	standardRouter.POST("/ltf/export", lib_tech_factor.ExcelExport)

	//---------相间距离最小值
	standardRouter.GET("/lpd", lib_phase_dist.FindPhaseDist)
	standardRouter.POST("/lpd", lib_phase_dist.CreatePhaseDist)
	standardRouter.GET("/lpd/:id", lib_phase_dist.GetSpecPhaseDist)
	standardRouter.PATCH("/lpd/:id", lib_phase_dist.UpdatePhaseDist)
	standardRouter.DELETE("/lpd/:id", lib_phase_dist.DeletePhaseDist)
	standardRouter.POST("/lpd/export", lib_phase_dist.ExcelExport)

	//---------树脂规格表
	standardRouter.GET("/lr", lib_resin.FindResin)
	standardRouter.POST("/lr", lib_resin.CreateResin)
	standardRouter.GET("/lr/:id", lib_resin.GetSpecResin)
	standardRouter.PATCH("/lr/:id", lib_resin.UpdateResin)
	standardRouter.DELETE("/lr/:id", lib_resin.DeleteResin)
	standardRouter.POST("/lr/export", lib_resin.ExcelExport)

	//---------电阻率
	standardRouter.GET("/lresist", lib_resistivity.FindResistivity)
	standardRouter.POST("/lresist", lib_resistivity.CreateResistivity)
	standardRouter.GET("/lresist/:id", lib_resistivity.GetSpecResistivity)
	standardRouter.PATCH("/lresist/:id", lib_resistivity.UpdateResistivity)
	standardRouter.DELETE("/lresist/:id", lib_resistivity.DeleteResistivity)
	standardRouter.POST("/lresist/export", lib_resistivity.ExcelExport)

	//---------硅钢片
	standardRouter.GET("/lstack", lib_stack.FindStack)
	standardRouter.POST("/lstack", lib_stack.CreateStack)
	standardRouter.GET("/lstack/:id", lib_stack.GetSpecStack)
	standardRouter.PATCH("/lstack/:id", lib_stack.UpdateStack)
	standardRouter.DELETE("/lstack/:id", lib_stack.DeleteStack)
	standardRouter.POST("/lstack/export", lib_stack.ExcelExport)

	//---------导线表
	standardRouter.GET("/lwire", lib_wire.FindWire)
	standardRouter.POST("/lwire", lib_wire.CreateWire)
	standardRouter.GET("/lwire/:id", lib_wire.GetSpecWire)
	standardRouter.PATCH("/lwire/:id", lib_wire.UpdateWire)
	standardRouter.DELETE("/lwire/:id", lib_wire.DeleteWire)
	standardRouter.POST("/lwire/export", lib_wire.ExcelExport)

	//---------导线绝缘厚度表
	standardRouter.GET("/boss", lib_boss.FindBosses)
	standardRouter.POST("/boss", lib_boss.CreateBoss)
	standardRouter.GET("/boss/:id", lib_boss.GetSpecBoss)
	standardRouter.PATCH("/boss/:id", lib_boss.UpdateBoss)
	standardRouter.DELETE("/boss/:id", lib_boss.DeleteBoss)
	standardRouter.POST("/boss/export", lib_boss.ExcelExport)

	//---------凸台规格

	////---------设计结果
	//standardRouter.GET("/design_results", design.FindDesignResults)
	//standardRouter.POST("/design_results", design.CreateDesignResults)
	//standardRouter.GET("/design_results/:id", design.GetDesignResults)
	//standardRouter.PATCH("/design_results:id", design.UpdateDesignResults)
	//standardRouter.DELETE("/design_results/:id", design.DeleteDesignResults)

}
