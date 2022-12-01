<template>
  <div class="app-container">
    <aside class="page-header">
      产品设计结果
    </aside>
    <div class="design__result-container">
      <el-row>
        <el-col :span="24">
          <el-form
            label-position="right"
            label-width="160px"
          >
            <el-row>
              <el-col :span="24">
                <div class="grid-content bg-purple-dark" />
                <el-row>
                  <el-col>
                    <el-collapse
                      v-model="activeCollapseNames"
                    >
                      <el-collapse-item
                        title="项目信息"
                        name="1"
                        style="alignment: center"
                      >
                        <el-row gutter="20">
                          <el-col :span="24">
                            <el-form-item
                              label="项目名称:"
                            >
                              <span>{{ designResults.design_project.project_name }}</span>
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="12">
                            <el-form-item
                              label="产品型号:"
                            >
                              <span>{{ designResults.design_project.product_model }}</span>
                            </el-form-item>
                          </el-col>
                          <el-col :span="12">
                            <el-form-item
                              label="产品代号:"
                            >
                              <span>{{ designResults.design_project.product_code }}</span>
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-divider content-position="center">
                          <el-tag type="success">
                            技术规范
                          </el-tag>
                        </el-divider>
                        <div>
                          <el-row>
                            <el-col :span="6">
                              <el-form-item label="额定容量(kVA):">
                                <!--                              <span>3250</span>-->
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.rated_capacity"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="额定高压(kV):">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.rated_high_vol"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="额定低压(kv):">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.rated_low_vol"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item
                                label="调压范围:"
                                class="multiple-select"
                              >
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.connect_type"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                          <el-row>
                            <el-col :span="6">
                              <el-form-item label="额定频率(HZ):">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.rated_frequency"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="相数:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.phase_number === 3 ? '三相' : '单相'"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="绝缘等级:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.insulation_class"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="联结组别:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_task.input.connect_type"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                          <el-row>
                            <el-col :span="6">
                              <el-form-item label="工频耐压(kV):">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.product_industry_freq_hold_vol"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="冲击电压(kV):">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.product_spec_shock_vol"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="感应电压(%):">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.product_induct_high_vol"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="冷却方式:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.product_cooling_method"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                          <el-row>
                            <el-col :span="12">
                              <el-form-item label="IP等级:">
                                <el-input
                                  :model-value="designResults.design_project.product_ip_level"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="12">
                              <el-form-item label="使用条件:">
                                <el-input
                                  :model-value="designResults.design_project.product_usage"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                        </div>
                        <el-divider content-position="center">
                          <el-tag type="success">
                            参与人员
                          </el-tag>
                        </el-divider>
                        <div>
                          <el-row>
                            <el-col :span="12">
                              <el-form-item label="设计人员:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.designer_user?.user_name_zh"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="12">
                              <el-form-item label="校核人员:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.reviewer_user?.user_name_zh"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                          <el-row>
                            <el-col :span="12">
                              <el-form-item label="审核人员:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.check_user?.user_name_zh"
                                />
                              </el-form-item>
                            </el-col>
                            <el-col :span="12">
                              <el-form-item label="设计主管:">
                                <el-input
                                  readonly
                                  :model-value="designResults.design_project.creator_user?.user_name_zh"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                        </div>
                      </el-collapse-item>

                      <!--                  性能参数标准 start-->
                      <el-collapse-item

                        name="2"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">性能参数</span>
                        </template>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="线圈损耗(W):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['线圈损耗']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="引线损耗(W):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['引线损耗']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="负载损耗(W):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['负载损耗']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="空载损耗(W):">
                              <span />
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['空载损耗']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="&emsp;总损耗(W):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['总损耗']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="铁心温升(K):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['铁心温升']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="低压温升(K):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['低压温升']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="高压温升(K):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['高压温升']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="空载电流(A):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['空载电流']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="阻抗电压(V):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['阻抗电压']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="噪音预测(dB):">
                              <el-input-number
                                :controls="false"
                                precision="2"
                                readonly
                                :model-value="designResults.performance_result['噪音预测']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>

                      <el-collapse-item

                        name="3"
                        style="alignment: center"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">温升</span>
                        </template>
                        <el-row>
                          <el-col :span="24">
                            <el-form-item label="铁心温升:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.temp_result['铁心温升']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="12">
                            <el-form-item label="低压温升:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.temp_result['低压温升']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="12">
                            <el-form-item label="高压温升:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.temp_result['高压温升']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="12">
                            <el-form-item label="绝缘筒温升1:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.temp_result['绝缘筒温升1']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="12">
                            <el-form-item label="绝缘筒温升2:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.temp_result['绝缘筒温升2']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>

                      <el-collapse-item

                        name="4"
                        style="alignment: center"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">低压线圈</span>
                        </template>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="额定容量:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['额定容量']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="相数:">
                              <el-input
                                readonly
                                :model-value="designResults.lv_result['相数']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="相电压:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['相电压']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="相电流:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['相电流']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="导线材质:">
                              <el-input
                                readonly
                                :model-value="designResults.lv_result['导线材质']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="导线重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['导线重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="树脂重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['树脂重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="低压接法:">
                              <el-input
                                readonly
                                :model-value="designResults.lv_result['低压接法']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="低压匝数:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['低压匝数']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="匝电压:">
                              <el-input-number
                                :precision="2"
                                :controls="false"
                                readonly
                                :model-value="designResults.lv_result['匝电压']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="线圈高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['线圈高']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="距铁轭:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['距铁轭']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="端绝缘:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['端绝缘']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="窗高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['窗高']"
                              />
                            </el-form-item>
                          </el-col>

                          <el-col :span="6">
                            <el-form-item label="内模半径:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['内模半径']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="外模半径:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['外模半径']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="低压到铁心距离:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['低压到铁心距离']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="纵向漏磁:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['纵向漏磁']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="铁心直径:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['铁心直径']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="绕制高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.lv_result['绕制高']"
                              />
                            </el-form-item>
                          </el-col>

                          <div style="padding-left: 20px;width: 100%;">
                            <el-collapse>
                              <el-collapse-item
                                v-for="(pack,i) in designResults.lv_result['包封']"
                                :key="i"
                                :title="'包封' + (i + 1)"
                                :name="i + 1"
                              >
                                <div style="padding-left: 20px;">
                                  <el-collapse>
                                    <el-collapse-item
                                      title="线规"
                                      name="1"
                                    >
                                      <div
                                        style="padding-left: 20px;"
                                      >
                                        <el-collapse
                                          v-for="(stacked,index) in pack['线规']['迭联']"
                                          :key="index"
                                          :model-value="1"
                                          accordion
                                        >
                                          <el-collapse-item
                                            :title="'迭联' + (index + 1)"
                                            :name="index + 1"
                                          >
                                            <div style="padding-left: 20px;">
                                              <div class="table-title">
                                                并联
                                              </div>
                                              <div>
                                                <el-table
                                                  style="width: 100%"
                                                  :data="stacked['并联']"
                                                >
                                                  <el-table-column
                                                    prop="线型"
                                                    label="线型"
                                                  />
                                                  <el-table-column
                                                    prop="线厚"
                                                    label="线厚"
                                                  >
                                                    <template #default="scope">
                                                      {{ scope.row['线厚'].toFixed(2) }}
                                                    </template>
                                                  </el-table-column>
                                                  <el-table-column
                                                    prop="线宽"
                                                    label="线宽"
                                                  >
                                                    <template #default="scope">
                                                      {{ scope.row['线宽'].toFixed(2) }}
                                                    </template>
                                                  </el-table-column>
                                                </el-table>
                                              </div>
                                            </div>
                                          </el-collapse-item>
                                        </el-collapse>
                                      </div>
                                    </el-collapse-item>
                                    <el-collapse-item
                                      title="排列"
                                      name="2"
                                    >
                                      <div style="padding-left: 20px">
                                        <el-collapse :model-value="1">
                                          <el-collapse-item
                                            title="段"
                                            :name="1"
                                          >
                                            <div style="padding-left: 20px;">
                                              <el-table
                                                :data="pack['段']"
                                                style="width: 100%"
                                              >
                                                <el-table-column
                                                  prop="层数"
                                                  label="层数"
                                                />
                                                <el-table-column
                                                  prop="匝数"
                                                  label="匝数"
                                                />
                                                <el-table-column
                                                  prop="各层匝数"
                                                  label="各层匝数"
                                                />
                                              </el-table>
                                            </div>
                                          </el-collapse-item>
                                        </el-collapse>
                                      </div>
                                    </el-collapse-item>
                                    <el-collapse-item
                                      title="包封数据"
                                      name="3"
                                    >
                                      <div style="padding-left: 20px;">
                                        <el-table :data="[pack]">
                                          <el-table-column
                                            prop="内侧半径"
                                            label="内侧半径"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['内侧半径'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="内绝缘"
                                            label="内绝缘"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['内绝缘'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝厚"
                                            label="匝厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝宽"
                                            label="匝宽"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝宽'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝截面积"
                                            label="匝截面积"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝截面积'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝数"
                                            label="匝数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="外侧半径"
                                            label="外侧半径"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['外侧半径'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="外绝缘"
                                            label="外绝缘"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['外绝缘'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="导线重"
                                            label="导线重"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['导线重'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="导线长度"
                                            label="导线长度"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['导线长度'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="层间绝缘"
                                            label="层间绝缘"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['层间绝缘'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="平均匝长"
                                            label="平均匝长"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['平均匝长'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="平均半径"
                                            label="平均半径"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['平均半径'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="并联线数"
                                            label="并联线数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['并联线数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="总线厚"
                                            label="总线厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['总线厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="最大层数"
                                            label="最大层数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['最大层数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="树脂体积"
                                            label="树脂体积"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['树脂体积'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="涡流损耗"
                                            label="涡流损耗"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['涡流损耗'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="电流密度"
                                            label="电流密度"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['电流密度'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="电阻损耗"
                                            label="电阻损耗"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['电阻损耗'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="直流电阻"
                                            label="直流电阻"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['直流电阻'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="绕制净厚"
                                            label="绕制净厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['绕制净厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="绕制厚"
                                            label="绕制厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['绕制厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="轴向净高"
                                            label="轴向净高"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['轴向净高'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="辐向裕度"
                                            label="辐向裕度"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['辐向裕度'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="迭联线数"
                                            label="迭联线数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['迭联线数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                        </el-table>
                                      </div>
                                    </el-collapse-item>
                                  </el-collapse>
                                </div>
                              </el-collapse-item>
                            </el-collapse>
                          </div>
                        </el-row>
                      </el-collapse-item>
                      <el-collapse-item

                        name="5"
                        style="alignment: center"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">高压线圈</span>
                        </template>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="额定容量:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['额定容量']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="相电压:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['相电压']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="相电流:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['相电流']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="导线材质:">
                              <el-input
                                readonly
                                :model-value="designResults.hv_result['导线材质']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="导线重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['导线重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="树脂重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['树脂重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="高压接法:">
                              <el-input
                                readonly
                                :model-value="designResults.hv_result['高压接法']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="匝电压:">
                              <el-input-number
                                :precision="2"
                                :controls="false"
                                readonly
                                :model-value="designResults.hv_result['匝电压']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="线圈高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['线圈高']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="距铁轭:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['距铁轭']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="端绝缘:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['端绝缘']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="窗高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['窗高']"
                              />
                            </el-form-item>
                          </el-col>

                          <el-col :span="6">
                            <el-form-item label="内模半径:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['内模半径']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="外模半径:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['外模半径']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="绕制高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.hv_result['绕制高']"
                              />
                            </el-form-item>
                          </el-col>

                          <div style="padding-left: 20px;width: 100%;">
                            <el-collapse>
                              <el-collapse-item
                                v-for="(pack,i) in designResults.hv_result['包封']"
                                :key="i"
                                :title="'包封' + (i + 1)"
                                :name="i + 1"
                              >
                                <div style="padding-left: 20px;">
                                  <el-collapse>
                                    <el-collapse-item
                                      title="线规"
                                      name="1"
                                    >
                                      <div
                                        style="padding-left: 20px;"
                                      >
                                        <el-collapse
                                          v-for="(stacked,index) in pack['线规']['迭联']"
                                          :key="index"
                                          :model-value="1"
                                          accordion
                                        >
                                          <el-collapse-item
                                            :title="'迭联' + (index + 1)"
                                            :name="index + 1"
                                          >
                                            <div style="padding-left: 20px;">
                                              <div class="table-title">
                                                并联
                                              </div>
                                              <div>
                                                <el-table
                                                  style="width: 100%"
                                                  :data="stacked['并联']"
                                                >
                                                  <el-table-column
                                                    prop="线型"
                                                    label="线型"
                                                  />
                                                  <el-table-column
                                                    prop="线厚"
                                                    label="线厚"
                                                  >
                                                    <template #default="scope">
                                                      {{ scope.row['线厚'].toFixed(2) }}
                                                    </template>
                                                  </el-table-column>
                                                  <el-table-column
                                                    prop="线宽"
                                                    label="线宽"
                                                  >
                                                    <template #default="scope">
                                                      {{ scope.row['线宽'].toFixed(2) }}
                                                    </template>
                                                  </el-table-column>
                                                </el-table>
                                              </div>
                                            </div>
                                          </el-collapse-item>
                                        </el-collapse>
                                      </div>
                                    </el-collapse-item>
                                    <el-collapse-item
                                      title="排列"
                                      name="2"
                                    >
                                      <div style="padding-left: 20px">
                                        <el-collapse :model-value="1">
                                          <el-collapse-item
                                            title="段"
                                            :name="1"
                                          >
                                            <div style="padding-left: 20px;">
                                              <el-table
                                                :data="pack['段']"
                                                style="width: 100%"
                                              >
                                                <el-table-column
                                                  prop="层数"
                                                  label="层数"
                                                />
                                                <el-table-column
                                                  prop="匝数"
                                                  label="匝数"
                                                />
                                                <el-table-column
                                                  prop="首层匝数"
                                                  label="首层匝数"
                                                />
                                                <el-table-column
                                                  prop="净高"
                                                  label="净高"
                                                >
                                                  <template #default="scope">
                                                    {{ scope.row['净高'].toFixed(2) }}
                                                  </template>
                                                </el-table-column>
                                                <el-table-column
                                                  prop="绕制高"
                                                  label="绕制高"
                                                >
                                                  <template #default="scope">
                                                    {{ scope.row['绕制高'].toFixed(2) }}
                                                  </template>
                                                </el-table-column>
                                                <el-table-column
                                                  prop="轴向裕度"
                                                  label="轴向裕度"
                                                >
                                                  <template #default="scope">
                                                    {{ scope.row['轴向裕度'].toFixed(2) }}
                                                  </template>
                                                </el-table-column>
                                              </el-table>
                                            </div>
                                          </el-collapse-item>
                                        </el-collapse>
                                      </div>
                                    </el-collapse-item>
                                    <el-collapse-item
                                      title="包封数据"
                                      name="3"
                                    >
                                      <div style="padding-left: 20px;">
                                        <el-table :data="[pack]">
                                          <el-table-column
                                            prop="内侧半径"
                                            label="内侧半径"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['内侧半径'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="内绝缘"
                                            label="内绝缘"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['内绝缘'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝厚"
                                            label="匝厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝宽"
                                            label="匝宽"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝宽'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝截面积"
                                            label="匝截面积"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝截面积'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="匝数"
                                            label="匝数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['匝数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="外侧半径"
                                            label="外侧半径"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['外侧半径'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="外绝缘"
                                            label="外绝缘"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['外绝缘'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="导线重"
                                            label="导线重"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['导线重'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="导线长度"
                                            label="导线长度"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['导线长度'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="层间绝缘"
                                            label="层间绝缘"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['层间绝缘'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="平均匝长"
                                            label="平均匝长"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['平均匝长'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="平均半径"
                                            label="平均半径"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['平均半径'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="并联线数"
                                            label="并联线数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['并联线数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="总线厚"
                                            label="总线厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['总线厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="最大层数"
                                            label="最大层数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['最大层数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="树脂体积"
                                            label="树脂体积"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['树脂体积'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="涡流损耗"
                                            label="涡流损耗"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['涡流损耗'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="电流密度"
                                            label="电流密度"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['电流密度'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="电阻损耗"
                                            label="电阻损耗"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['电阻损耗'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="直流电阻"
                                            label="直流电阻"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['直流电阻'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="绕制净厚"
                                            label="绕制净厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['绕制净厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="绕制厚"
                                            label="绕制厚"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['绕制厚'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="辐向裕度"
                                            label="辐向裕度"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['辐向裕度'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="迭联线数"
                                            label="迭联线数"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['迭联线数'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                          <el-table-column
                                            prop="额定线长"
                                            label="额定线长"
                                          >
                                            <template #default="scope">
                                              {{ scope.row['额定线长'].toFixed(2) }}
                                            </template>
                                          </el-table-column>
                                        </el-table>
                                      </div>
                                    </el-collapse-item>
                                  </el-collapse>
                                </div>
                              </el-collapse-item>
                            </el-collapse>
                          </div>
                        </el-row>
                      </el-collapse-item>

                      <el-collapse-item
                        name="7"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">铁心</span>
                        </template>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="中心距:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['中心距']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="单位质量磁化容量:">
                              <span />
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['单位质量磁化容量']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="单位铁损:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['单位铁损']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="单位面积接缝伏安值:">
                              <span />
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['单位面积接缝伏安值']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="工艺系数:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['工艺系数']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="总重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['总重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="接缝形式:">
                              <el-input
                                readonly
                                :model-value="designResults.core_result['接缝形式']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="有无拉板:">
                              <el-input
                                readonly
                                :model-value="designResults.core_result['有无拉板']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="柱重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['柱重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="磁通密度:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['磁通密度']"
                              />
                            </el-form-item>
                          </el-col><el-col :span="6">
                            <el-form-item label="空载损耗:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['空载损耗']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="窗高:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['窗高']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="角重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['角重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="轭重:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['轭重']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="迭片系数:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['迭片系数']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="钢片厚度:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['钢片厚度']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="钢片片数:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['钢片片数']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="钢片牌号:">
                              <el-input
                                readonly
                                :model-value="designResults.core_result['钢片牌号']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="铁心截面:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['铁心截面']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="铁心直径:">
                              <el-input-number
                                :controls="false"
                                :precision="2"
                                readonly
                                :model-value="designResults.core_result['铁心直径']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>
                      <!--                      <el-collapse-item-->
                      <!--                        name="8"-->
                      <!--                      >-->
                      <!--                        <template-->
                      <!--                          #title-->
                      <!--                        >-->
                      <!--                          <span class="collapse-title">模具</span>-->
                      <!--                        </template>-->
                      <!--                        <el-row>-->
                      <!--                          <el-col :span="6">-->
                      <!--                            <el-form-item label="低压内模半径(mm):">-->
                      <!--                              <el-input-->
                      <!--                                readonly-->
                      <!--                                :model-value="designResults.low_vol_in_diameter"-->
                      <!--                              />-->
                      <!--                            </el-form-item>-->
                      <!--                          </el-col>-->
                      <!--                          <el-col :span="6">-->
                      <!--                            <el-form-item label="低压外模半径(mm):">-->
                      <!--                              <span />-->
                      <!--                              <el-input-->

                      <!--                                readonly-->
                      <!--                                :model-value="designResults.low_vol_out_diameter"-->
                      <!--                              />-->
                      <!--                            </el-form-item>-->
                      <!--                          </el-col>-->
                      <!--                          <el-col :span="6">-->
                      <!--                            <el-form-item label="高压内模半径(mm):">-->
                      <!--                              <el-input-->

                      <!--                                readonly-->
                      <!--                                :model-value="designResults.high_vol_in_diameter"-->
                      <!--                              />-->
                      <!--                            </el-form-item>-->
                      <!--                          </el-col>-->
                      <!--                          <el-col :span="6">-->
                      <!--                            <el-form-item label="高压外模半径(mm):">-->
                      <!--                              <span />-->
                      <!--                              <el-input-->

                      <!--                                readonly-->
                      <!--                                :model-value="designResults.high_vol_out_diameter"-->
                      <!--                              />-->
                      <!--                            </el-form-item>-->
                      <!--                          </el-col>-->
                      <!--                        </el-row>-->
                      <!--                        <el-row>-->
                      <!--                          <el-col :span="6">-->
                      <!--                            <el-form-item label="铁心半径(mm):">-->
                      <!--                              <el-input-->
                      <!--                                readonly-->
                      <!--                                :model-value="designResults.iron_core_radius"-->
                      <!--                              />-->
                      <!--                            </el-form-item>-->
                      <!--                          </el-col>-->
                      <!--                          <el-col :span="18">-->
                      <!--                            <el-form-item label="低压到铁心距离(mm):">-->
                      <!--                              <el-input-->

                      <!--                                readonly-->
                      <!--                                :model-value="designResults.iron_core_from_coil"-->
                      <!--                              />-->
                      <!--                            </el-form-item>-->
                      <!--                          </el-col>-->
                      <!--                        </el-row>-->
                      <!--                      </el-collapse-item>-->

                      <!--                      <el-collapse-item name="9">-->
                      <!--                        <template-->
                      <!--                          #title-->
                      <!--                        >-->
                      <!--                          <span class="collapse-title">变比校核</span>-->
                      <!--                        </template>-->
                      <!--                        <el-row v-if="designResults.transformation_ratio_check">-->
                      <!--                          <el-col :span="24">-->
                      <!--                            <div style="padding-left: 20px;">-->
                      <!--                              <div class="table-title">-->
                      <!--                                分接电压-->
                      <!--                              </div>-->
                      <!--                              <el-table-->
                      <!--                                style="width: 100%;"-->
                      <!--                                :data="designResults.transformation_ratio_check.tapping_voltage"-->
                      <!--                              >-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="tapping_voltage"-->
                      <!--                                  label="分接电压"-->
                      <!--                                />-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="turn_number"-->
                      <!--                                  label="匝数"-->
                      <!--                                />-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="fact_voltage"-->
                      <!--                                  label="实际电压"-->
                      <!--                                />-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="deviation"-->
                      <!--                                  label="偏差"-->
                      <!--                                />-->
                      <!--                              </el-table>-->
                      <!--                            </div>-->
                      <!--                          </el-col>-->
                      <!--                        </el-row>-->
                      <!--                        <el-row v-if="designResults.transformation_ratio_check">-->
                      <!--                          <el-col :span="24">-->
                      <!--                            <div style="padding-left: 20px;">-->
                      <!--                              <div class="table-title">-->
                      <!--                                分接间电压-->
                      <!--                              </div>-->
                      <!--                              <el-table-->
                      <!--                                style="width: 100%;"-->
                      <!--                                :data="designResults.transformation_ratio_check.tapping_room_voltage"-->
                      <!--                              >-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="tapping_voltage"-->
                      <!--                                  label="分接电压"-->
                      <!--                                />-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="turn_number"-->
                      <!--                                  label="匝数"-->
                      <!--                                />-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="fact_voltage"-->
                      <!--                                  label="实际电压"-->
                      <!--                                />-->
                      <!--                                <el-table-column-->
                      <!--                                  prop="deviation"-->
                      <!--                                  label="偏差"-->
                      <!--                                />-->
                      <!--                              </el-table>-->
                      <!--                            </div>-->
                      <!--                          </el-col>-->
                      <!--                        </el-row>-->
                      <!--                      </el-collapse-item>-->
                      <el-collapse-item name="10">
                        <template
                          #title
                        >
                          <span class="collapse-title">阻抗电压</span>
                        </template>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item label="delta:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                :model-value="designResults.impedance_result['delta']"
                                readonly
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="delta_pi1:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['delta_pi1']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="delta_pi2:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['delta_pi2']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="an3:">
                              <span />
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['an3']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="delta_pi:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['delta_pi']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="delta_r:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['delta_r']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="lw:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['lw']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="dm:">
                              <span />
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['dm']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="k:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['k']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="ks:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['ks']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="ux:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['ux']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="ur:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['ur']"
                              />
                            </el-form-item>
                          </el-col>

                          <el-col :span="6">
                            <el-form-item label="uk:">
                              <el-input-number
                                :precision="4"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['uk']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="引线损耗:">
                              <el-input-number
                                :precision="2"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['引线损耗']"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item label="线圈损耗:">
                              <el-input-number
                                :precision="2"
                                :controls="false"
                                readonly
                                :model-value="designResults.impedance_result['线圈损耗']"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>
                      <el-collapse-item name="11">
                        <template
                          #title
                        >
                          <span class="collapse-title">成本</span>
                        </template>
                        <el-row
                          v-if="designResults.cost_result['成本']"
                          style="padding-left: 20px;"
                        >
                          <el-col :span="24">
                            <div class="table-title">
                              明细
                            </div>
                            <div
                              style="padding-left: 20px;"
                            >
                              <el-table :data="designResults.cost_result['成本'].map(item => item[0].split(' ').filter(item => item !== ''))">
                                <el-table-column
                                  :label="designResults.cost_result['标题'].split(' ').filter(item => item !== '')[0]"
                                >
                                  <template #default="scope">
                                    {{ scope.row[0] }}
                                  </template>
                                </el-table-column>
                                <el-table-column
                                  :label="designResults.cost_result['标题'].split(' ').filter(item => item !== '')[1]"
                                >
                                  <template #default="scope">
                                    {{ scope.row[1] }}
                                  </template>
                                </el-table-column>
                                <el-table-column
                                  :label="designResults.cost_result['标题'].split(' ').filter(item => item !== '')[2]"
                                >
                                  <template #default="scope">
                                    {{ scope.row[2] }}
                                  </template>
                                </el-table-column>
                                <el-table-column
                                  :label="designResults.cost_result['标题'].split(' ').filter(item => item !== '')[3]"
                                >
                                  <template #default="scope">
                                    {{ scope.row[3] }}
                                  </template>
                                </el-table-column>
                              </el-table>
                            </div>
                          </el-col>
                        </el-row>
                        <el-row
                          v-if="designResults.cost_result['总成本']"
                          style="margin-top: 20px;"
                        >
                          <el-col :span="24">
                            <el-form-item label="总成本:">
                              <el-input
                                readonly
                                :model-value="designResults.cost_result['总成本'].toFixed(2)"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>
                    </el-collapse>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
          </el-form>
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, ref, PropType, onMounted, Ref, watch, computed } from 'vue'
import { DecodeDesignTask, DesignProject, DesignResults, DesignTask } from '@/model/design'
import { useRoute } from 'vue-router'
import { GetDesignResults } from '@/api/design'
import { decodeBytesData, decodeObjectBytes } from '@/utils/jsonutils'

export default defineComponent({
  props: {
    handleValue: {
      type: Object as PropType<DesignResults> | undefined,
      default: undefined
    }
  },
  emits: [],
  setup(props, context) {
    const route = useRoute()
    const activeCollapseNames = ref(['1'])
    const designResults : Ref<DesignResults> = ref({
      continue_vars: {} as Object,
      core_result: {} as Object,
      cost_result: {} as Object,
      design_project: {
        check_by: 0,
        check_user: {} as Object,
        checked_at: 0,
        comment: '',
        created_at: '',
        creator: 0,
        deleted_at: '',
        deliver_at: 0,
        design_at: 0,
        design_comment: '',
        design_finish_at: 0,
        designer: 0,
        designer_signed: 0,
        designer_user: {} as Object,
        drafting_at: 0,
        drafting_by: 0,
        drafting_user: {} as Object,
        drawing_at: 0,
        id: 0,
        need_master_approve: 0,
        order_at: 0,
        product_altitude: 0,
        product_code: '',
        product_connection_group: '',
        product_cooling_method: '',
        product_frequency: 0,
        product_frequency_special: '',
        product_induct_high_vol: 0,
        product_industry_freq_hold_vol: '',
        product_insulation_high_ac: 0,
        product_insulation_high_li: 0,
        product_insulation_level: '',
        product_insulation_level_special: '',
        product_insulation_low_ac: 0,
        product_insulation_low_li: 0,
        product_ip_level: '',
        product_load_loss: 0,
        product_model: '',
        product_name: '',
        product_noload_loss: 0,
        product_phrase: '',
        product_rated_capacity: 0,
        product_rated_v_high: 0,
        product_rated_v_low: 0,
        product_short_circuit_resistance: 0,
        product_spec_shock_vol: '',
        product_tap_range: '',
        product_tap_range_special: '',
        product_temp_limit: 0,
        product_total_loss: 0,
        product_usage: '',
        product_wind_type: '',
        product_wire_material: '',
        product_wire_shape: '',
        project_name: '',
        project_status: 0,
        review_at: 0,
        reviewer: 0,
        reviewer_signed: 0,
        serial_code: '',
        tech_requirments: '',
        updated_at: ''
      },
      design_project_id: 0,
      design_task: {
        result_progress: 0,
        approve: 0,
        approve_by_creator: false,
        approve_node_instance: undefined,
        comment: '',
        created_at: '',
        creator: 0,
        deleted_at: '',
        design_project: undefined,
        design_project_id: 0,
        design_results: '',
        designer: 0,
        designer_user: undefined,
        final_design_results: [],
        id: 0,
        name: '',
        reviewer: 0,
        task_status: 0,
        updated_at: '',
        input: {
          boss_spec: '',
          connect_type: '',
          deviation_hv_temp_rise: 0,
          deviation_impedance_vol: 0,
          deviation_load_loss: 0,
          deviation_lv_temp_rise: 0,
          deviation_no_load_current: 0,
          deviation_no_load_loss: 0,
          deviation_noise_prediction: 0,
          deviation_total_loss: 0,
          eddy_current_loss_rate: 0,
          has_pull_plate: '',
          hv_wire_type: { flat: ['', 0], round: ['', 0] },
          id: 0,
          insulation_class: '',
          insulation_temp: 0,
          insulation_tube1: '',
          insulation_tube2: '',
          ks: 0,
          lead_loss: 0,
          limit_hv_temp_rise: 0,
          limit_impedance_vol: 0,
          limit_load_loss: 0,
          limit_lv_temp_rise: 0,
          limit_no_load_current: 0,
          limit_no_load_loss: 0,
          limit_noise_prediction: 0,
          limit_total_loss: 0,
          lv_wire_type: { flat: ['', 0], foil: ['', 0], round: ['', 0] },
          phase_number: 0,
          radiating_surface_coeff: {
            hv: [[0, 0], [0, 0], [0, 0], [0, 0]],
            insula_tube1: [0, 0],
            insula_tube2: [0, 0],
            lv: [[0, 0], [0, 0], [0, 0], [0, 0]]
          },
          rated_capacity: 0,
          rated_frequency: 0,
          rated_high_vol: 0,
          rated_low_vol: 0,
          resin_type: ['', 0],
          steel_type: [],
          temp_rise_coeff: { hv: [0, 0, 0, 0], insula_tube1: 0, insula_tube2: 0, lv: [0, 0, 0, 0] },
          vol_regulation_max: 0,
          vol_regulation_min: 0,
          vol_regulation_step: 0,
          wire_material: ''
        }
      } as DecodeDesignTask,
      design_task_id: 0,
      hv_result: {} as Object,
      id: 0,
      impedance_result: {} as Object,
      lv_result: {} as Object,
      performance_result: {} as Object,
      product_name: '',
      project_code: '',
      project_name: '',
      temp_result: {} as Object,
      total_cost: 0
    } as DesignResults)
    watch(() => props.handleValue, function() {
      designResults.value = JSON.parse(JSON.stringify(props.handleValue))
    })
    onMounted(async () => {
      if (route.query.id && !props.handleValue) {
        const res = await GetDesignResults({
          'field_eq_design_results.id': parseInt(String(route.query.id)),
          limit: 0,
          order: '',
          order_by: '',
          page: 0
        })
        if (res.data.code === 200) {
          if (res.data.spec[0]) {
            decodeObjectBytes(res.data.spec[0])
            console.log(res.data.spec[0])
            designResults.value = res.data.spec[0] as DesignResults
          }
        }
      } else if (props.handleValue) {
        console.log(props.handleValue)
        designResults.value = props.handleValue
      }
    })
    return {
      activeCollapseNames,
      designResults
    }
  }
})
</script>

<style lang="scss" scoped>
:deep {
  .design__result-container {
      input {
        text-align: center;
        color: $--color-primary;
        width: 100% !important;
      }
      .el-select,.el-input,.el-input-number {
        width: 100% !important;
      }
      .el-input-number__decrease {
        display: none;
      }
      .el-input-number__increase  {
        display: none;
      }
    }
}
</style>
