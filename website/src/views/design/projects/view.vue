<template>
  <el-dialog
    v-model="show"
    v-loading="paramLoading"
    destroy-on-close
    width="80%"
    title="产品设计任务书"
  >
    <div class="project_form">
      <el-card class="box-card">
        <div>
          <el-form
            ref="validateForm"
            disabled
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
                        <el-input v-model="tempDetail.product_usage" />
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
                          />
                        </el-form-item>
                        <el-form-item
                          label="2、相数&emsp;&emsp;&emsp;&emsp;"
                          prop="product_phrase"
                        >
                          <el-input v-model.number="tempDetail.product_phrase" />
                        </el-form-item>
                        <el-form-item
                          label="3、频率(Hz)&emsp;&emsp;"
                          prop="product_frequency"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input
                                v-model="tempDetail.product_frequency"
                              />
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
                              <el-input
                                v-model="tempDetail.product_tap_range"
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="6、工频耐压(kV)"
                          prop="product_industry_freq_hold_vol"
                        >
                          <el-input v-model="tempDetail.product_industry_freq_hold_vol" />
                        </el-form-item>
                        <el-form-item
                          label="7、冲击电压(kV)"
                          prop="product_spec_shock_vol"
                        >
                          <el-input v-model="tempDetail.product_spec_shock_vol" />
                        </el-form-item>
                        <el-form-item
                          label="8、感应电压(%)&ensp;"
                          prop="product_induct_high_vol"
                        >
                          <el-input-number
                            v-model="tempDetail.product_induct_high_vol"
                            :precision="2"
                          />
                        </el-form-item>
                        <el-form-item
                          label="9、绝缘等级&emsp;&emsp;"
                          prop="product_insulation_level"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input
                                v-model="tempDetail.product_insulation_level"
                                filterable
                                allow-create
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="10、温升限值(k)&emsp;"
                          prop="product_temp_limit"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input
                                v-model="tempDetail.product_temp_limit"
                                filterable
                                allow-create
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
                              <el-input
                                v-model="tempDetail.product_connection_group"
                                filterable
                                allow-create
                              />
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
                              <el-input
                                v-model="tempDetail.product_short_circuit_resistance"
                                filterable
                                allow-create
                              />
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
                              <el-input
                                v-model="tempDetail.product_cooling_method"
                                filterable
                                allow-create
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="16、绕制类型&emsp;&ensp;"
                          prop="product_wind_type"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input
                                v-model="tempDetail.product_wind_type"
                                filterable
                                allow-create
                              />
                            </el-col>
                          </el-row>
                        </el-form-item>
                        <el-form-item
                          label="17、线圈形状&emsp;&ensp;"
                          prop="product_wire_shape"
                        >
                          <el-row :gutter="20">
                            <el-col :span="24">
                              <el-input
                                v-model="tempDetail.product_wire_shape"
                                filterable
                                allow-create
                              />
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
                              <el-input
                                v-model="tempDetail.product_wire_material"
                                filterable
                                allow-create
                              />
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
                          <user-search-input
                            v-model="tempDetail.drafting_by"
                            :placeholder="'编制人员'"
                            :style="{width: '100%'}"
                          />
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
                          <user-search-input
                            v-model="tempDetail.check_by"
                            :placeholder="'审核人员'"
                            :style="{width: '100%'}"
                          />
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
        <el-button @click="show = false">关 闭</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref, Ref, watch } from 'vue'
import { DesignProject } from '@/model/design'
import UserSearchInput from '@/components/user-search-input/index.vue'
import { GetStdlctLibraries, Lct } from '@/api/standard_libraries/lct'
import { GetStdLrfLibraries, Lrf } from '@/api/standard_libraries/lrf'
import { GetStdLtrLibraries, Ltr } from '@/api/standard_libraries/ltr'
import { GetStdLusLibraries, Lus } from '@/api/standard_libraries/lus'
import { GetStdLpnLibraries, Lpn } from '@/api/standard_libraries/lpn'
import { GetStdLrrLibraries, ListLrrQuery, Lrr } from '@/api/standard_libraries/lrr'
import { RoleList } from '@/model/permission'

export default defineComponent({
  components: {
    UserSearchInput
  },
  props: {
    modelValue: Boolean,
    handleValue: Object
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const show = computed({
      get: () => props.modelValue,
      set: (v) => {
        context.emit('update:modelValue', v)
      }
    })
    const emptyDesignProject : DesignProject = {
      product_spec_shock_vol: '',
      product_wind_type: '',
      product_wire_shape: '',
      product_industry_freq_hold_vol: '',
      product_induct_high_vol: 0,
      need_master_approve: 0,
      product_frequency: 0,
      product_frequency_special: '',
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
      product_phrase: '',
      product_rated_capacity: 0,
      product_rated_v_high: 0,
      product_rated_v_low: 0,
      product_short_circuit_resistance: 0,
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
      tech_requirments: ''
    }
    const tempDetail : Ref<DesignProject> = ref(emptyDesignProject)
    watch(() => props.handleValue, () => {
      tempDetail.value = JSON.parse(JSON.stringify(props.handleValue as DesignProject))
    })

    onMounted(() => {
      if (props.handleValue) {
        tempDetail.value = JSON.parse(JSON.stringify(props.handleValue as DesignProject))
      }
    })
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
            label: i.phase_num,
            value: String(i.phase_num)
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
            label: i.usage,
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
    // 绝缘水平雷电LI
    const productInsulationLIOptions = [
      { label: '0kV', value: 0 },
      { label: '40kV', value: 40 },
      { label: '60kV', value: 60 },
      { label: '75kV', value: 75 },
      { label: '95kV', value: 95 },
      { label: '125kV', value: 125 },
      { label: '170V', value: 170 }
    ]
    // 绝缘水平工频耐压AC
    const productInsulationACOptions = [
      { label: '3kV', value: 3 },
      { label: '5kV', value: 5 },
      { label: '10kV', value: 10 },
      { label: '20kV', value: 20 },
      { label: '35kV', value: 35 },
      { label: '38kV', value: 38 },
      { label: '50kV', value: 50 },
      { label: '70kV', value: 70 }
    ]
    const productInsulationLevelOptions = [
      { label: 'F级', value: 'F' },
      { label: 'H级', value: 'H' }
    ]
    const productTempRiseLimitOptions : Ref<{label: string, value: number}[]> = ref([])
    const getStdLtrLibraries = async () => {
      const res = await GetStdLtrLibraries({
        field_eq_temp: null,
        field_eq_temp_sign: null,
        field_eq_low_vol_temp_rise: null,
        field_eq_high_vol_temp_rise: null,
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        productTempRiseLimitOptions.value = res.data.spec.map((i: Ltr) => {
          return {
            label: `${i.high_vol_temp_rise}k`,
            value: i.high_vol_temp_rise
          }
        })
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
    const productWireMaterialOptions = [
      { label: '全铜', value: '全铜' },
      { label: '全铝', value: '全铝' },
      { label: '高铜低铝', value: '高铜低铝' },
      { label: '高铝低铜', value: '高铝低铜' }
    ]
    // TODO: 产品型号根据参数自动调整
    // watch(() => [
    //     tempDetail.value.product_usage,
    //     tempDetail.value.product_phrase,
    //     tempDetail.value.
    // ])
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
        getStdLtrLibraries()
      ])
      paramLoading.value = false
    })
    return {
      RoleList,
      paramLoading,
      show,
      tempDetail,
      specRangeOptions,
      phaseOptions,
      productUsageOptions,
      productFrequencyOptions,
      productInsulationLIOptions,
      productInsulationACOptions,
      productInsulationLevelOptions,
      productTempRiseLimitOptions,
      productConnectionGroupOptions,
      productShortCircuitResistanceOptions,
      productCoolingMethodOptions,
      productWireMaterialOptions
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
  .is-disabled {
    .el-input__inner,.el-textarea__inner,.el-radio__label {
      background: #fff;
      color: #000 !important;
    }
    .is-checked {
      .el-radio__inner {
        background: $__color-primary !important;
        color: #fff !important;
        border-color: $__color-primary !important;
        &::after {
          background: #fff !important;
        }
      }
    }
  }
}
</style>
