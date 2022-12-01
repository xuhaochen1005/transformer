<template>
  <el-dialog
    v-model="show"
    v-loading="paramLoading"
    destroy-on-close
    width="80%"
    title="产品设计任务书"
  >
    <div class="project_form">
      <el-card
        class="box-card"
      >
        <div>
          <el-form
            ref="validateForm"
            :model="tempDetail"
            :rules="formRules"
            label-width="180px"
          >
            <el-row>
              <el-col :span="24">
                <div class="grid-content bg-purple-dark">
                  <el-row>
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="项目名称"
                          prop="project_name"
                        >
                          <el-input
                            v-model="tempDetail.project_name"
                            placeholder="请输入项目名称"
                            clearable
                            :style="{width: '100%'}"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark" />
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="产品名称"
                          prop="product_name"
                        >
                          <el-input
                            v-model="tempDetail.product_name"
                            placeholder="请输入产品名称"
                            clearable
                            :style="{width: '100%'}"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark" />
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="产品型号"
                          prop="product_model"
                        >
                          <el-input
                            v-model="tempDetail.product_model"
                            placeholder="请输入产品型号"
                            clearable
                            :style="{width: '100%'}"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="设计完成日期"
                          prop="design_finish_at"
                        >
                          <el-date-picker
                            v-model="tempDetail.design_finish_at"
                            type="date"
                            size="large"
                            style="width: 100%;"
                            value-format="X"
                            placeholder="请输入设计完成日期"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="产品代号"
                          prop="product_code"
                        >
                          <el-input
                            v-model="tempDetail.product_code"
                            placeholder="产品代号"
                            clearable
                            :style="{width: '100%'}"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="产品交货日期"
                          prop="deliver_at"
                        >
                          <el-date-picker
                            v-model="tempDetail.deliver_at"
                            type="date"
                            size="large"
                            style="width: 100%;"
                            value-format="X"
                            placeholder="请输入交货完成日期"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="24">
                      <el-form-item
                        label="用&emsp;途&emsp;"
                        prop="product_usage"
                      >
                        <el-select
                          v-model="tempDetail.product_usage"
                          filterable
                          allow-create
                          placeholder="用途"
                        >
                          <el-option
                            v-for="item in productUsageOptions"
                            :key="item"
                            :label="item.label"
                            :value="item.value"
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark">
                        <el-divider content-position="center">
                          <el-tag type="success">
                            技术规范
                          </el-tag>
                        </el-divider>
                        <el-form-item
                          label="1、容量(kVA)&emsp;&ensp;"
                          prop="product_rated_capacity"
                        >
                          <el-input-number
                            v-model="tempDetail.product_rated_capacity"
                            :precision="2"
                            placeholder="容量"
                          />
                        </el-form-item>
                        <el-form-item
                          label="2、相数&emsp;&emsp;&emsp;&emsp;"
                          prop="product_phrase"
                        >
                          <el-select
                            v-model="tempDetail.product_phrase"
                            filterable
                            allow-create
                            placeholder="相数"
                            style="width:100%"
                          >
                            <el-option
                              v-for="item in phaseOptions"
                              :key="item"
                              :label="item.label"
                              :value="item.value"
                            />
                          </el-select>
                        </el-form-item>
                        <el-form-item
                          label="3、频率(Hz)&emsp;&emsp;"
                          prop="product_frequency"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_frequency"
                                filterable
                                allow-create
                                placeholder="频率"
                              >
                                <el-option
                                  v-for="item in productFrequencyOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-row :gutter="20">
                          <el-col :span="12">
                            <el-form-item
                              label="4、额定高压(kV)"
                              prop="product_rated_v_high"
                            >
                              <el-input-number
                                v-model="tempDetail.product_rated_v_high"
                                :precision="2"
                                placeholder="额定高压"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="12">
                            <el-form-item
                              label="额定低压(kV)"
                              prop="product_rated_v_low"
                            >
                              <el-input-number
                                v-model="tempDetail.product_rated_v_low"
                                :precision="2"
                                placeholder="额定低压"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>

                        <el-form-item
                          label="5、调压范围(%)&ensp;"
                          prop="product_tap_range"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_tap_range"
                                filterable
                                allow-create
                                style="width: 100%;"
                                remote
                                placeholder="调压范围"
                                no-data-text="数据不符合条件"
                              >
                                <el-option
                                  v-for="item in specRangeOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="6、工频耐压(kV)"
                          prop="product_industry_freq_hold_vol"
                        >
                          <el-row>
                            <el-col :span="11">
                              <el-input-number
                                v-model.number="product_industry_freq_hold_vol_high"
                                :controls="false"
                                :precision="2"
                                placeholder="高压"
                              />
                            </el-col>
                            <el-col
                              :span="2"
                              style="text-align: center;"
                            >
                              /
                            </el-col>
                            <el-col :span="11">
                              <el-input-number
                                v-model.number="product_industry_freq_hold_vol_low"
                                :controls="false"
                                :precision="2"
                                placeholder="低压"
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="7、冲击电压(kV)"
                          prop="product_spec_shock_vol"
                        >
                          <el-row>
                            <el-col :span="11">
                              <el-input-number
                                v-model.number="product_spec_shock_vol_high"
                                :controls="false"
                                :precision="2"
                                placeholder="高压"
                              />
                            </el-col>
                            <el-col
                              :span="2"
                              style="text-align: center;"
                            >
                              /
                            </el-col>
                            <el-col :span="11">
                              <el-input-number
                                v-model.number="product_spec_shock_vol_low"
                                :controls="false"
                                :precision="2"
                                placeholder="低压"
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="8、感应电压(%)&ensp;"
                          prop="product_induct_high_vol"
                        >
                          <el-input-number
                            v-model="tempDetail.product_induct_high_vol"
                            :precision="2"
                            placeholder="感应电压"
                          />
                        </el-form-item>
                        <el-form-item
                          label="9、绝缘等级&emsp;&emsp;"
                          prop="product_insulation_level"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_insulation_level"
                                filterable
                                allow-create
                                placeholder="绝缘等级"
                                @change="tempDetail.product_temp_limit = productTempRiseLimitOptions.find(item => item.temp_sign === tempDetail.product_insulation_level).low_vol_temp_rise"
                              >
                                <el-option
                                  v-for="item in productTempRiseLimitOptions.filter(item => item.temp_sign !== '/')"
                                  :key="item.id"
                                  :label="item.temp_sign"
                                  :value="item.temp_sign"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="10、温升限值(k)&emsp;"
                          prop="product_temp_limit"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input-number
                                v-model="tempDetail.product_temp_limit"
                                :precision="2"
                                :controls="false"
                                placeholder="温升限值"
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="11、联结组别&emsp;&emsp;"
                          prop="product_connection_group"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_connection_group"
                                filterable
                                allow-create
                                placeholder="联结组别"
                              >
                                <el-option
                                  v-for="item in productConnectionGroupOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="12、IP等级&emsp;&emsp;&emsp;"
                          prop="product_ip_level"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input
                                v-model="tempDetail.product_ip_level"
                                placeholder="IP等级"
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="13、阻抗电压(%)"
                          prop="product_short_circuit_resistance"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_short_circuit_resistance"
                                filterable
                                allow-create
                                placeholder="阻抗电压"
                              >
                                <el-option
                                  v-for="item in productShortCircuitResistanceOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="14、海拔(m)&emsp;&emsp;"
                          prop="product_altitude"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input-number
                                v-model="tempDetail.product_altitude"
                                :precision="2"
                                placeholder="海拔"
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="15、冷却方式&emsp;&ensp;"
                          prop="product_cooling_method"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_cooling_method"
                                filterable
                                allow-create
                                placeholder="冷却方式"
                              >
                                <el-option
                                  v-for="item in productCoolingMethodOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="16、绕制类型&emsp;&ensp;"
                          prop="product_wind_type"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_wind_type"
                                filterable
                                allow-create
                                placeholder="绕制类型"
                              >
                                <el-option
                                  v-for="item in ['全箔绕','全线绕','低压箔绕高压线绕']"
                                  :key="item"
                                  :label="item"
                                  :value="item"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="17、线圈形状&emsp;&ensp;"
                          prop="product_wire_shape"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_wire_shape"
                                filterable
                                allow-create
                                placeholder="线圈形状"
                              >
                                <el-option
                                  v-for="item in shapeOptions"
                                  :key="item.id"
                                  :label="item.coil_shape"
                                  :value="item.coil_shape"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="18、损耗水平&emsp;&ensp;"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-form-item
                                prop="product_load_loss"
                                label="负载损耗(W) (120°C) ≤"
                              >
                                <el-input-number
                                  v-model="tempDetail.product_load_loss"
                                  :precision="2"
                                  placeholder="负载损耗"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                          <el-row
                            :gutter="20"
                            style="margin-top: 10px;"
                          >
                            <el-col :span="24">
                              <el-form-item
                                prop="product_noload_loss"
                                label="空载损耗(W) ≤"
                              >
                                <el-input-number
                                  v-model="tempDetail.product_noload_loss"
                                  :precision="2"
                                  placeholder="空载损耗"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                          <el-row
                            :gutter="20"
                            style="margin-top: 10px;"
                          >
                            <el-col :span="24">
                              <el-form-item
                                prop="product_total_loss"
                                label="总损耗(W) ≤"
                              >
                                <el-input-number
                                  v-model="tempDetail.product_total_loss"
                                  :precision="2"
                                  placeholder="总损耗"
                                />
                              </el-form-item>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="19、导线材质&emsp;&ensp;"
                          prop="product_wire_material"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-select
                                v-model="tempDetail.product_wire_material"
                                filterable
                                allow-create
                                placeholder="导向材质"
                              >
                                <el-option
                                  v-for="item in coilWireOptions"
                                  :key="item"
                                  :label="item.coil_wire_type"
                                  :value="item.coil_wire_type"
                                />
                              </el-select>
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="20、其他特殊要求"
                        >
                          <el-input
                            v-model="tempDetail.tech_requirments"
                            type="textarea"
                            :rows="5"
                            placeholder="请输入产品技术指标"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark">
                        <aside style="text-align: center;width: 80% ;margin-left: 10%">
                          任务进程
                        </aside>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="8">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="设计人员"
                          prop="designer"
                        >
                          <user-search-input
                            v-model="tempDetail.designer"
                            :roles="[RoleList.Designer,RoleList.Demander]"
                            :placeholder="'设计人员'"
                            :style="{width: '100%'}"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="8">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="设计日期"
                          prop="design_at"
                        >
                          <el-date-picker
                            v-model="tempDetail.design_at"
                            type="date"
                            size="large"
                            style="width: 100%;"
                            value-format="X"
                            placeholder="请输入设计日期"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="8">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="是否签名"
                        >
                          <el-radio-group v-model="tempDetail.designer_signed">
                            <el-radio
                              :label="1"
                              :value="1"
                            >
                              已完成签名确认
                            </el-radio>
                            <el-radio
                              :label="0"
                              :value="0"
                            >
                              暂未完成签名确认
                            </el-radio>
                          </el-radio-group>
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="8">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="校核人员"
                          prop="reviewer"
                        >
                          <user-search-input
                            v-model="tempDetail.reviewer"
                            :roles="[RoleList.Demander,RoleList.Reviewer]"
                            :placeholder="'校核人员'"
                            :style="{width: '100%'}"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="8">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="校核日期"
                          prop="review_at"
                        >
                          <el-date-picker
                            v-model="tempDetail.review_at"
                            type="date"
                            size="large"
                            style="width: 100%;"
                            value-format="X"
                            placeholder="请输入校核日期"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="8">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="是否签名"
                        >
                          <el-radio-group v-model="tempDetail.reviewer_signed">
                            <el-radio
                              :label="1"
                              value="1"
                            >
                              已完成签名确认
                            </el-radio>
                            <el-radio
                              :label="0"
                              type="button"
                            >
                              暂未完成签名确认
                            </el-radio>
                          </el-radio-group>
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="备注信息"
                        >
                          <el-input
                            v-model="tempDetail.design_comment"
                            type="textarea"

                            :rows="3"
                            placeholder="请输入任务书备注信息"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="编制人员"
                        >
                          <span>{{ state.user.user_name_zh }}</span>
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="编制日期"
                          prop="drafting_at"
                        >
                          <el-date-picker
                            v-model="tempDetail.drafting_at"
                            type="date"
                            size="large"
                            style="width: 100%;"
                            value-format="X"
                            placeholder="请输入编制日期"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="审核人员"
                        >
                          <!-- <user-search-input
                            v-model="tempDetail.check_by"
                            :placeholder="'审核人员'"
                            :style="{width: '100%'}"
                          /> -->
                          <span>{{ state.user.user_name_zh }}</span>
                        </el-form-item>
                      </div>
                    </el-col>
                    <el-col :span="12">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="审核日期"
                          prop="checked_at"
                        >
                          <el-date-picker
                            v-model="tempDetail.checked_at"
                            type="date"
                            size="large"
                            style="width: 100%;"
                            value-format="X"
                            placeholder="请输入审核日期"
                          />
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="24">
                      <div class="grid-content bg-purple-dark">
                        <el-form-item
                          label="总工审核"
                        >
                          <el-radio-group v-model="tempDetail.need_master_approve">
                            <el-radio
                              :label="1"
                              :value="1"
                            >
                              需要总工审核
                            </el-radio>
                            <el-radio
                              :label="0"
                              :value="0"
                              type="button"
                            >
                              无需总工审核
                            </el-radio>
                          </el-radio-group>
                        </el-form-item>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
            </el-row>
          </el-form>
        </div>
      </el-card>
    </div>

    <template #footer>
      <span class="dialog-footer">
        <el-button
          type="primary"
          @click="confirmCreate"
        >确 定</el-button>
        <el-button @click="show = false">取 消</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, PropType, reactive, ref, Ref, watch, watchEffect } from 'vue'
import { DesignProject } from '@/model/design'
import UserSearchInput from '@/components/user-search-input/index.vue'
import { CreateDesignProject } from '@/api/projects'
import { ElMessage } from 'element-plus'
import { GetStdLrfLibraries, Lrf } from '@/api/standard_libraries/lrf'

import { GetStdLusLibraries, Lus } from '@/api/standard_libraries/lus'
import { GetStdLpnLibraries, Lpn } from '@/api/standard_libraries/lpn'
import { GetStdlctLibraries, Lct } from '@/api/standard_libraries/lct'
import { GetStdLtrLibraries, Ltr } from '@/api/standard_libraries/ltr'
import { RoleList } from '@/model/permission'
import { useStore } from 'vuex'
import { GetStdLrrLibraries, ListLrrQuery, Lrr } from '@/api/standard_libraries/lrr'
import { getValidateFormDefaultRules } from '@/utils/validate'
import { GetStdLshvLibraries, ListLshvQuery, Lshv } from '@/api/standard_libraries/lshv'
import { GetStdlifhvLibraries, Lifhv, ListLifhvQuery } from '@/api/standard_libraries/lifhv'
import { GetStdlcsLibraries, Lcs, ListLcsQuery } from '@/api/standard_libraries/lcs'
import lcs from '@/views/standard/libraries/lcs/index.vue'
import { GetStdLlLibraries, ListLlQuery } from '@/api/standard_libraries/ll'
import { GetStdLnpLibraries, ListLnpQuery } from '@/api/standard_libraries/lnp'
import { GetStdlcwtLibraries, Lcwt, ListLcwtQuery } from '@/api/standard_libraries/lcwt'

export default defineComponent({
  components: {
    UserSearchInput
  },
  props: {
    modelValue: Boolean,
    handleValue: {
      type: Object as PropType<DesignProject>
    }
  },
  emits: ['update:modelValue', 'refresh'],
  setup(props, context) {
    const validateForm = ref(null)
    const formProps = ['product_wire_shape', 'product_wind_type', 'project_name', 'product_name', 'product_model', 'design_finish_at', 'product_code', 'deliver_at', 'product_usage', 'product_rated_capacity', 'product_phrase', 'product_frequency.number', 'product_rated_v_high', 'product_rated_v_low', 'product_tap_range.number.array', 'product_industry_freq_hold_vol', 'product_induct_high_vol', 'product_insulation_level', 'product_temp_limit.number', 'product_connection_group', 'product_short_circuit_resistance.number', 'product_altitude', 'product_cooling_method', 'product_load_loss', 'product_noload_loss', 'product_total_loss', 'product_wire_material', 'designer.user', 'design_at', 'reviewer.user', 'review_at', 'drafting_at', 'checked_at']
    const formRules = getValidateFormDefaultRules(formProps)

    const state = useStore().state
    const emptyDesignProject : DesignProject = {
      product_spec_shock_vol: '',
      product_wind_type: '',
      product_wire_shape: '',
      product_ip_level: '',
      product_frequency: 50,
      product_frequency_special: '',
      product_altitude: 0,
      product_connection_group: '',
      product_cooling_method: '',
      product_insulation_high_ac: 0,
      product_insulation_high_li: 0,
      product_insulation_level: '',
      product_insulation_level_special: '',
      product_insulation_low_ac: 0,
      product_insulation_low_li: 0,
      product_load_loss: 0,
      product_noload_loss: 0,
      product_phrase: '三相',
      product_rated_capacity: 50,
      product_rated_v_high: 0,
      product_rated_v_low: 0,
      product_short_circuit_resistance: 4,
      product_tap_range_special: '',
      product_temp_limit: 0,
      product_total_loss: 0,
      product_usage: '',
      product_wire_material: '',

      // approve: 0,
      // approve_at: 0,
      review_at: 0,
      check_by: 0,
      checked_at: 0,
      design_comment: '',
      comment: '',
      creator: 0,
      deliver_at: 0,
      design_at: 0,
      design_finish_at: 0,
      designer: 0,
      drafting_at: 0,
      drafting_by: 0,
      drawing_at: 0,
      id: 0,
      project_name: '',
      order_at: 0,
      product_code: '',
      product_model: '',
      product_name: '',
      reviewer: 0,
      serial_code: '',
      tech_requirments: '',
      designer_signed: 0,
      reviewer_signed: 0,
      product_industry_freq_hold_vol: '',
      product_induct_high_vol: 0,
      need_master_approve: 0
    }
    const tempDetail : Ref<DesignProject> = ref(emptyDesignProject)
    if (props.handleValue) {
      tempDetail.value = JSON.parse(JSON.stringify(props.handleValue)) as DesignProject
    }

    const show = computed({
      get: () => props.modelValue,
      set: (v) => {
        context.emit('update:modelValue', v)
        if (!v) {
          tempDetail.value = JSON.parse(JSON.stringify(emptyDesignProject)) as DesignProject
        }
      }
    })

    watch(() => tempDetail.value.product_model, (value) => {
      const nameLike = value.trim().match(/(ZL|ZB)?([DS])(C)(Z?)(B?)(10|11|12|13|15|16)-([0-9.]*)\/([0-9.]*)\/([0-9.]*)/)
      console.log(nameLike)
      if (nameLike) {
        // 用途
        if (nameLike[1]) {
          for (const item of productUsageOptions.value) {
            if (item.label === nameLike[1]) {
              tempDetail.value.product_usage = item.value
              break
            }
          }
        } else {
          tempDetail.value.product_usage = '/'
        }
        // 相数
        if (nameLike[2]) {
          for (const item of phaseOptions.value) {
            console.log(item)
            if (item.value === nameLike[2]) {
              tempDetail.value.product_phrase = item.label
              break
            }
          }
        }
        // 额定容量
        if (nameLike[7]) {
          tempDetail.value.product_rated_capacity = parseFloat(nameLike[7])
        }
        // 额定高压侧电压
        if (nameLike[8]) {
          tempDetail.value.product_rated_v_high = parseFloat(nameLike[8])
        }
        // 额定低压侧电压
        if (nameLike[9]) {
          tempDetail.value.product_rated_v_low = parseFloat(nameLike[9])
        }
      }
    })
    // 工频耐压高低
    const product_industry_freq_hold_vol_low: Ref<number> = ref(0)
    const product_industry_freq_hold_vol_high: Ref<number> = ref(0)
    watchEffect(() => {
      tempDetail.value.product_industry_freq_hold_vol = `${product_industry_freq_hold_vol_high.value}/${product_industry_freq_hold_vol_low.value}`
    })
    const product_industry_freq_hold_vol_options: Ref<Lifhv[]> = ref([])
    const getStdLifhvLibraries = async () => {
      const res = await GetStdlifhvLibraries({
        limit: 99999,
        order: '',
        order_by: '',
        page: 1
      } as ListLifhvQuery)
      if (res.data.code === 200) {
        product_industry_freq_hold_vol_options.value = res.data.spec
      }
    }
    watchEffect(() => {
      product_industry_freq_hold_vol_high.value = product_industry_freq_hold_vol_options.value.find(item => item.rated_high_vol_min! <= tempDetail.value.product_rated_v_high && item.rated_high_vol_max! > tempDetail.value.product_rated_v_high)?.industry_freq_hold_vol || 0
    })
    watchEffect(() => {
      product_industry_freq_hold_vol_low.value = product_industry_freq_hold_vol_options.value.find(item => item.rated_high_vol_min! <= tempDetail.value.product_rated_v_low && item.rated_high_vol_max! > tempDetail.value.product_rated_v_low)?.industry_freq_hold_vol || 0
    })
    // 冲击电压高压
    const product_spec_shock_vol_low : Ref<number> = ref(0)
    const product_spec_shock_vol_high : Ref<number> = ref(0)
    watchEffect(() => {
      tempDetail.value.product_spec_shock_vol = `${product_spec_shock_vol_high.value}/${product_spec_shock_vol_low.value}`
    })
    const product_spec_shock_vol_options : Ref<Lshv[]> = ref([])
    const getStdLshvLibraries = async () => {
      const res = await GetStdLshvLibraries({
        limit: 9999,
        order: '',
        order_by: '',
        page: 1
      } as ListLshvQuery)
      if (res.data.code === 200) {
        product_spec_shock_vol_options.value = res.data.spec
      }
    }
    watchEffect(() => {
      product_spec_shock_vol_low.value = product_spec_shock_vol_options.value.find(item => item.rated_high_vol_min! <= tempDetail.value.product_rated_v_low && item.rated_high_vol_max! > tempDetail.value.product_rated_v_low)?.shock_hold_vol || 0
    })
    watchEffect(() => {
      product_spec_shock_vol_high.value = product_spec_shock_vol_options.value.find(item => item.rated_high_vol_min! <= tempDetail.value.product_rated_v_high && item.rated_high_vol_max! > tempDetail.value.product_rated_v_high)?.shock_hold_vol || 0
    })
    // 线圈形状
    const shapeOptions: Ref<Lcs[]> = ref([])
    const getStdShapeLibraries = async () => {
      const res = await GetStdlcsLibraries({
        limit: 9999,
        order: '',
        order_by: '',
        page: 1
      } as ListLcsQuery)
      if (res.data.code === 200) {
        shapeOptions.value = res.data.spec
      }
    }
    // 损耗
    watchEffect(async () => {
      if (tempDetail.value.product_tap_range && tempDetail.value.product_rated_capacity &&
        tempDetail.value.product_rated_v_high && tempDetail.value.product_rated_v_low &&
        tempDetail.value.product_insulation_level &&
      tempDetail.value.product_model) {
        let regulate_way = '无励磁调压'
        // 10/11/12/13型
        let lossLevel = ''
        const nameLike = tempDetail.value.product_model.trim().match(/(ZL|ZB)?([DS])(C)(Z?)(B?)(10|11|12|13|15|16)-([0-9.]*)\/([0-9.]*)\/([0-9.]*)/)
        if (nameLike) {
          if (nameLike[4] === 'Z') {
            regulate_way = '有载调压'
          }
          if (!nameLike[6]) {
            return
          }
          lossLevel = nameLike[6]
        }
        const regulationArr = []
        try {
          regulationArr.push(...JSON.parse(tempDetail.value.product_tap_range))
        } catch (e) {
          console.log(e)
          return
        }
        const llStdQuery = reactive({
          limit: 1,
          order: '',
          order_by: '',
          page: 1,
          field_eq_regulate_way: regulate_way,
          field_le_rated_capacity_min: tempDetail.value.product_rated_capacity,
          field_gt_rated_capacity_max: tempDetail.value.product_rated_capacity,
          field_le_rated_high_vol_min: tempDetail.value.product_rated_v_high,
          field_gt_rated_high_vol_max: tempDetail.value.product_rated_v_high,
          field_le_rated_low_vol_min: tempDetail.value.product_rated_v_low,
          field_gt_rated_low_vol_max: tempDetail.value.product_rated_v_low,
          field_eq_temperature: productTempRiseLimitOptions.value.find(item => item.temp_sign === tempDetail.value.product_insulation_level)?.temp,
          field_eq_regulate_range_min: Math.min(...regulationArr),
          field_eq_regulate_range_max: Math.max(...regulationArr),
          field_eq_regulate_range_step: Math.abs(regulationArr[1] - regulationArr[0]),
          filed_eq_loss_level: lossLevel
        } as ListLlQuery)
        const res = await GetStdLlLibraries(llStdQuery)
        if (res.data.code === 200) {
          if (res.data.spec.length === 1) {
            tempDetail.value.product_load_loss = res.data.spec[0].load_loss!
            tempDetail.value.product_noload_loss = res.data.spec[0].no_load_loss!
            tempDetail.value.product_total_loss = res.data.spec[0].load_loss! + res.data.spec[0].no_load_loss!
          }
        }
      }
    })

    watch(() => props.handleValue, function() {
      tempDetail.value = JSON.parse(JSON.stringify(props.handleValue)) as DesignProject
      if (tempDetail.value.product_industry_freq_hold_vol) {
        product_industry_freq_hold_vol_low.value = parseFloat(tempDetail.value.product_industry_freq_hold_vol.split('/')[0])
        product_industry_freq_hold_vol_high.value = parseFloat(tempDetail.value.product_industry_freq_hold_vol.split('/')[1])
      }
      if (tempDetail.value.product_spec_shock_vol) {
        product_spec_shock_vol_low.value = parseFloat(tempDetail.value.product_spec_shock_vol.split('/')[0])
        product_spec_shock_vol_high.value = parseFloat(tempDetail.value.product_spec_shock_vol.split('/')[1])
      }
    })

    async function confirmCreate() {
      (validateForm.value as any).validate(async(valid: Boolean) => {
        console.log(valid)
        if (!valid) {
          return
        }
        tempDetail.value.design_at *= 1
        tempDetail.value.review_at *= 1
        tempDetail.value.checked_at *= 1
        tempDetail.value.deliver_at *= 1
        tempDetail.value.drafting_at *= 1
        tempDetail.value.design_finish_at *= 1
        // tempDetail.value. = dayjs(emptyDesignProject.design_finish_at).unix()
        const res = await CreateDesignProject(tempDetail.value)
        if (res.data.code === 200) {
          ElMessage.success('创建设计任务成功')
          context.emit('refresh', true)
          show.value = false
        }
      })
    }
    const specRangeOptions: Ref<{label: string, value: string}[]> = ref([])
    const getSpecRange = async () => {
      const listLrrQuery : ListLrrQuery = {
        field_eq_regulate_range_min: 0,
        field_eq_regulate_range_max: 0,
        field_eq_regulate_range_step: 0,
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      }
      const res = await GetStdLrrLibraries(listLrrQuery)
      if (res.data.code === 200) {
        specRangeOptions.value = res.data.spec.map((item:Lrr) => {
          const result: number[] = []
          const rangeLength = item.regulate_range_max / item.regulate_range_step
          for (let i = 0; i < rangeLength; i++) {
            result.push(item.regulate_range_max - item.regulate_range_step * i)
          }
          const label = `[${result.join(', ')}, 0, ${result.map(s => '-' + s).join(', ')}]`
          return {
            label: label,
            value: label
          }
        })
      }
    }
    // 相数
    const phaseOptions : Ref<{label: string, value: string}[]> = ref([])
    const getStdLpnLibraries = async () => {
      const res = await GetStdLpnLibraries({
        field_eq_phase_num: null,
        field_eq_phase_num_sign: null,
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        phaseOptions.value = res.data.spec.map((i: Lpn) => {
          return {
            label: i.phase_num as string,
            value: i.phase_num as string
          }
        })
      }
    }

    // 用途
    const productUsageOptions : Ref<{label: string, value: string}[]> = ref([])
    const getStdLusLibraries = async () => {
      const res = await GetStdLusLibraries({
        field_eq_usage: '',
        field_eq_usage_sign: '',
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        productUsageOptions.value = res.data.spec.map((i: Lus) => {
          return {
            label: `${i.usage} ${i.usage_sign}`,
            value: i.usage_sign
          }
        })
      }
    }
    // 频率
    const productFrequencyOptions : Ref<{label: string, value: number}[]> = ref([])
    const getStdLrfLibraries = async () => {
      const res = await GetStdLrfLibraries({
        field_eq_rated_freq: null,
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        productFrequencyOptions.value = res.data.spec.map((i: Lrf) => {
          return {
            label: i.rated_freq + 'Hz',
            value: i.rated_freq
          }
        })
      }
    }
    const productTempRiseLimitOptions : Ref<Ltr[]> = ref([])
    const getStdLtrLibraries = async () => {
      const res = await GetStdLtrLibraries({
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        productTempRiseLimitOptions.value = res.data.spec
      }
    }
    const productConnectionGroupOptions = [
      { label: 'Dyn11', value: 'Dyn11' },
      { label: 'Yyn0', value: 'Yyn0' },
      { label: 'Dd0', value: 'Dd0' },
      { label: 'Yd11', value: 'Yd11' }
    ]
    const productShortCircuitResistanceOptions = [
      { label: '4%', value: 4 },
      { label: '6%', value: 6 },
      { label: '8%', value: 8 },
      { label: '10%', value: 10 }
    ]
    const productCoolingMethodOptions : Ref<{label: string, value: string}[]> = ref([])
    const getStdlctLibraries = async () => {
      const res = await GetStdlctLibraries({
        field_eq_cool_type: '',
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        productCoolingMethodOptions.value = res.data.spec.map((i: Lct) => {
          return {
            label: i.cool_type,
            value: i.cool_type
          }
        })
      }
    }
    // 线圈材质
    const coilWireOptions : Ref<Lcwt[]> = ref([])
    const getStdlcwtLibraries = async () => {
      const res = await GetStdlcwtLibraries({
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      } as ListLcwtQuery)
      if (res.data.code === 200) {
        coilWireOptions.value = res.data.spec
      }
    }
    const paramLoading = ref(true)
    onMounted(async () => {
      await Promise.all([
        // 分接范围
        getSpecRange(),
        // 相数
        getStdLpnLibraries(),
        // 额定频率
        getStdLrfLibraries(),
        // 用途
        getStdLusLibraries(),
        // 冷却方式
        getStdlctLibraries(),
        // 温升限值
        getStdLtrLibraries(),
        // 工频耐压， 冲击电压
        getStdLifhvLibraries(),
        getStdLshvLibraries(),
        // 线圈形状
        getStdShapeLibraries(),
        // 线圈材质
        getStdlcwtLibraries()
      ])
      paramLoading.value = false
    })

    return {
      validateForm,
      formRules,
      paramLoading,
      state,
      RoleList,
      product_industry_freq_hold_vol_low,
      product_industry_freq_hold_vol_high,
      product_spec_shock_vol_low,
      product_spec_shock_vol_high,
      show,
      tempDetail,
      confirmCreate,
      specRangeOptions,
      phaseOptions,
      productUsageOptions,
      productFrequencyOptions,
      productTempRiseLimitOptions,
      productConnectionGroupOptions,
      productShortCircuitResistanceOptions,
      productCoolingMethodOptions,
      shapeOptions,
      coilWireOptions
    }
  }
})
</script>

<style scoped lang="scss">
:deep {
  .el-select,.el-input,.el-input-number {
    width: 100% !important;
    input {
      text-align: left !important;
      padding-left: 30px !important;
    }
  }
  .el-input-number__decrease {
    display: none;
  }
  .el-input-number__increase  {
    display: none;
  }
}
</style>
