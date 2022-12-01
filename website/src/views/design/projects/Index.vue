<template>
  <div class="app-container">
    <aside class="page-header">
      任务书管理
    </aside>
    <div class="filter-container">
      <el-input
        v-model="query.field_lk_project_name"
        placeholder="项目名称"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <el-input
        v-model="query.field_lk_product_model"
        placeholder="产品型号"
        clearable
        class="filter-item"
        style="width: 200px"
      />
      <el-input
        v-model="query.field_lk_serial_code"
        placeholder="任务书编号"
        clearable
        class="filter-item"
        style="width: 200px"
      />
      <el-select
        v-model="query.field_eq_project_status"
        placeholder="任务状态"
        clearable
        :value-key="'value'"
        style="width: 200px"
        class="filter-item"
      >
        <el-option
          v-for="item in selectStatusOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索
      </el-button>
      <el-button
        v-role="[RoleList.Root, RoleList.ChiefEngineer, RoleList.Demander]"
        class="filter-item"
        style="margin-left: 10px"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate()"
      >
        添加
      </el-button>
      <!--      <el-button-->
      <!--        v-waves-->
      <!--        :loading="downloadLoading"-->
      <!--        class="filter-item"-->
      <!--        type="primary"-->
      <!--        icon="el-icon-download"-->
      <!--        @click="handleDownload"-->
      <!--      >-->
      <!--        导出-->
      <!--      </el-button>-->
    </div>
    <div class="main-table">
      <el-table
        key="design-project"
        v-loading="listLoading"
        :data="listData"
        border
        fit
        highlight-current-row
        style="width: 100%"
        @sort-change="sortChange"
      >
        <el-table-column type="expand">
          <template #default="props">
            <el-form
              ref="validateForm"
              label-position="left"
              inline
              class="project-table-expand"
              label-width="160px"
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
                              v-model="props.row.project_name"

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
                              v-model="props.row.product_name"

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
                              v-model="props.row.product_model"

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
                              v-model="props.row.design_finish_at"
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
                              v-model="props.row.product_code"

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
                              v-model="props.row.deliver_at"
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
                          <el-input v-model="props.row.product_usage" />
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
                              v-model="props.row.product_rated_capacity"
                              :precision="2"
                            />
                          </el-form-item>
                          <el-form-item
                            label="2、相数&emsp;&emsp;&emsp;&emsp;"
                            prop="product_phrase"
                          >
                            <el-input v-model.number="props.row.product_phrase" />
                          </el-form-item>
                          <el-form-item
                            label="3、频率(Hz)&emsp;&emsp;"
                            prop="product_frequency"
                          >
                            <el-row :gutter="20">
                              <el-col :span="24">
                                <el-input
                                  v-model="props.row.product_frequency"
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
                                  v-model="props.row.product_rated_v_high"
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
                                  v-model="props.row.product_rated_v_low"
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
                                  v-model="props.row.product_tap_range"
                                />
                              </el-col>
                            </el-row>
                          </el-form-item>
                          <el-form-item
                            label="6、工频耐压(kV)"
                            prop="product_industry_freq_hold_vol"
                          >
                            <el-input v-model="props.row.product_industry_freq_hold_vol" />
                          </el-form-item>
                          <el-form-item
                            label="7、冲击电压(kV)"
                            prop="product_spec_shock_vol"
                          >
                            <el-input v-model="props.row.product_spec_shock_vol" />
                          </el-form-item>
                          <el-form-item
                            label="8、感应电压(%)&ensp;"
                            prop="product_induct_high_vol"
                          >
                            <el-input-number
                              v-model="props.row.product_induct_high_vol"
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
                                  v-model="props.row.product_insulation_level"
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
                                  v-model="props.row.product_temp_limit"
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
                                  v-model="props.row.product_connection_group"
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
                                  v-model="props.row.product_ip_level"
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
                                  v-model="props.row.product_short_circuit_resistance"
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
                                  v-model="props.row.product_altitude"
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
                                  v-model="props.row.product_cooling_method"
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
                                  v-model="props.row.product_wind_type"
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
                                  v-model="props.row.product_wire_shape"
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
                                    v-model="props.row.product_load_loss"
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
                                    v-model="props.row.product_noload_loss"
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
                                    v-model="props.row.product_total_loss"
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
                                  v-model="props.row.product_wire_material"
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
                              v-model="props.row.tech_requirments"
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
                            <span>{{ props.row.designer_user.user_name_zh }}</span>
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
                              v-model="props.row.design_at"
                              type="date"
                              size="large"
                              style="width: 100%;"
                              value-format="X"
                            />
                          </el-form-item>
                        </div>
                      </el-col>
                      <el-col :span="8">
                        <div class="grid-content bg-purple-dark">
                          <el-form-item
                            label="是否签名"
                          >
                            <el-radio-group v-model="props.row.designer_signed">
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
                            {{ props.row.reviewer_user.user_name_zh }}
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
                              v-model="props.row.review_at"
                              type="date"
                              size="large"
                              style="width: 100%;"
                              value-format="X"
                            />
                          </el-form-item>
                        </div>
                      </el-col>
                      <el-col :span="8">
                        <div class="grid-content bg-purple-dark">
                          <el-form-item
                            label="是否签名"
                          >
                            <el-radio-group v-model="props.row.reviewer_signed">
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
                              v-model="props.row.design_comment"
                              type="textarea"

                              :rows="3"
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
                            <span>{{ props.row.drafting_user.user_name_zh }}</span>
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
                              v-model="props.row.drafting_at"
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
                            label="审核人员"
                          >
                            <!-- <user-search-input
                              v-model="props.row.check_by"
                              :
                              :style="{width: '100%'}"
                            /> -->
                            {{ props.row.reviewer_user.user_name_zh }}
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
                              v-model="props.row.checked_at"
                              type="date"
                              size="large"
                              style="width: 100%;"
                              value-format="X"
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
                            <el-radio-group v-model="props.row.need_master_approve">
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
          </template>
        </el-table-column>
        <el-table-column
          label="ID"
          sortable="custom"
          align="center"
          min-width="100px"
          prop="id"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="项目名称"
          sortable="custom"
          align="center"
          prop="project_name"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.project_name }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="产品型号"
          min-width="200px"
          sortable="custom"
          align="center"
          prop="product_model"
        >
          <template #default="{row}">
            <span>{{ row.product_model }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="产品代号"
          align="center"
          sortable="custom"
          prop="product_code"
          width="130px"
        >
          <template #default="{row}">
            <span>{{ row.product_code }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="任务书编号"
          align="center"
          sortable="custom"
          prop="serial_code"
          width="130px"
        >
          <template #default="{row}">
            <span>{{ row.serial_code }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="产品交货日期"
          prop="deliver_at"
          sortable="custom"
          align="center"
          width="170px"
        >
          <template #default="{row}">
            <span>{{ row.deliver_at?UnixTime2HumanWithUnix(row.design_finish_at):'' }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="设计完成日期"
          prop="design_finish_at"
          sortable="custom"
          align="center"
          width="170px"
        >
          <template #default="{row}">
            <span>{{ row.design_finish_at?UnixTime2HumanWithUnix(row.design_finish_at):'' }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="设计人员"
          prop="designer"
          align="center"
          width="80px"
        >
          <template #default="{row}">
            <span>{{ row.designer_user?.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="校核人员"
          prop="checked_by"
          align="center"
          width="80px"
        >
          <template #default="{row}">
            <span>{{ row.reviewer_user?.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="编制人员"
          prop="drafting_by"
          align="center"
          width="80px"
        >
          <template #default="{row}">
            <span>{{ row.drafting_user?.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="审核人员"
          prop="reviewer"
          align="center"
          width="80px"
        >
          <template #default="{row}">
            <span>{{ row.check_user?.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="项目状态"
          prop="project_status"
          sortable="custom"
          align="center"
          width="120px"
        >
          <template #default="{row}">
            <span>
              <el-tag
                type="primary"
              >
                {{ statusOptions.find(i => i.value === row.project_status)?.label }}
              </el-tag>
            </span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          min-width="400px"
          fixed="right"
          class-name="fixed-width"
        >
          <template #default="{row}">
            <el-row style="display: flex;flex-wrap: nowrap;gap: 5px;justify-content: space-evenly;">
              <el-button
                type="primary"
                size="mini"
                @click="handleDetail(row)"
              >
                查看详情
              </el-button>
              <span
                v-permission="[ResourceList.DesignProject,ActionList.Write]"
              >
                <el-button
                  size="mini"
                  type="success"
                  @click="handleUpdate(row)"
                >
                  修改
                </el-button>
              </span>
              <span
                v-permission="[ResourceList.DesignProject,ActionList.Write]"
              >
                <el-button
                  size="mini"
                  type="warning"
                  @click="handleCreate(row)"
                >
                  复制
                </el-button>
              </span>
              <el-button
                size="mini"
                type="primary"
                :loading="downloading"
                @click="handleExport(row,row.id)"
              >
                导出
              </el-button>
              <span
                v-permission="[ResourceList.DesignProject,ActionList.Write]"
              >
                <el-button
                  size="mini"
                  type="danger"
                  @click="handleDelete(row.id)"
                >
                  删除
                </el-button>
              </span>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div
      class="pagination-footer"
      style="margin-top: 1%"
    >
      <el-pagination
        v-show="total > 0"
        v-model:current-page="query.page"
        v-model:page-size="query.limit"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleFilter"
        @current-change="handleFilter"
      />
    </div>
    <div class="dialogs-main">
      <ProjectCreate
        v-model="showCreate"
        :handle-value="tempCreate"
        @refresh="handleFilter"
      />
      <ProjectView
        v-model="showDetail"
        :handle-value="currentRow"
      />
      <ProjectUpdate
        v-model="showUpdate"
        :handle-value="currentRow"
        @refresh="handleFilter"
      />
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, ref, Ref } from 'vue'
import { DesignProject, DesignProjectQuery, DesignProjectStatus, DesignProjectUsage, PhaseType } from '@/model/design'
import { DeleteDesignProject, ExportDesignProject, GetDesignProjects } from '@/api/projects'
import { UnixTime2HumanWithUnix } from '@/utils/timeutils'
import ProjectCreate from './create.vue'
import ProjectView from './view.vue'
import ProjectUpdate from './update.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { stringify } from '@/utils/jsonutils'
import { useRoute } from 'vue-router'
import { GetStdlctLibraries, Lct } from '@/api/standard_libraries/lct'
import { GetStdLpnLibraries, Lpn } from '@/api/standard_libraries/lpn'
import { GetStdLrfLibraries, Lrf } from '@/api/standard_libraries/lrf'
import { GetStdLtrLibraries, Ltr } from '@/api/standard_libraries/ltr'
import { GetStdLusLibraries, Lus } from '@/api/standard_libraries/lus'
import { ActionList, ResourceList, RoleList } from '@/model/permission'
import { Response } from '@/model'

export default defineComponent({
  components: {
    ProjectCreate,
    ProjectView,
    ProjectUpdate
  },
  setup() {
    const query: DesignProjectQuery = reactive({
      field_lk_project_name: '',
      field_lk_product_model: '',
      limit: 10,
      order: 'desc',
      order_by: 'created_at',
      page: 1
    })
    const route = useRoute()
    const total = ref(0)
    const listLoading = ref(false)
    const downloading = ref(false)
    const listData : Ref<DesignProject[]> = ref([])
    async function getList() {
      listLoading.value = true
      const res = await GetDesignProjects(query)
      if (res.data.code === 200) {
        total.value = res.data.total
        listData.value = res.data.spec
      }
      listLoading.value = false
    }
    async function handleExport(designProject: DesignProject, id:number) {
      downloading.value = true
      const res = await ExportDesignProject(id)
      if ((res.data as Blob).type === 'application/json') {
        const data : Response<any> = await new Promise((resolve, reject) => {
          const reader = new FileReader()
          reader.readAsText(res.data, 'utf-8')
          reader.onload = function() {
            try {
              resolve(JSON.parse(reader.result as string))
            } catch (error) {
              resolve({ code: 500, error: '文件解析失败', spec: undefined, total: 0 })
            }
          }
        })
        if (data.code !== 200) {
          downloading.value = false
          ElMessage.error(data.error)
        }
        return
      }
      // console.log(res)
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data)
      } else {
        const objectUrl = URL.createObjectURL(new Blob([res.data]))
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = designProject.project_name + '.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      downloading.value = false
    }
    function handleFilter() {
      getList()
    }
    const statusOptions = [
      { label: '导入成功', value: DesignProjectStatus.DesignProjectImported },
      { label: '已取消', value: DesignProjectStatus.DesignProjectCreated },
      { label: '开始执行', value: DesignProjectStatus.DesignProjectStarted },
      { label: '执行完毕', value: DesignProjectStatus.DesignProjectFinished },
      { label: '复核不通过', value: DesignProjectStatus.DesignProjectReviewUnaccepted },
      { label: '复核通过', value: DesignProjectStatus.DesignProjectReviewed },
      { label: '审核不通过', value: DesignProjectStatus.DesignProjectApproveUnaccepted },
      { label: '审核通过', value: DesignProjectStatus.DesignProjectApproved },
      { label: '审核不通过', value: DesignProjectStatus.DesignProjectCheckUnAccepted },
      { label: '审核通过', value: DesignProjectStatus.DesignProjectChecked }
    ]
    const selectStatusOptions = [
      { label: '导入成功', value: DesignProjectStatus.DesignProjectImported },
      { label: '已取消', value: DesignProjectStatus.DesignProjectCreated },
      { label: '开始执行', value: DesignProjectStatus.DesignProjectStarted },
      { label: '执行完毕', value: DesignProjectStatus.DesignProjectFinished },
      { label: '复核不通过', value: DesignProjectStatus.DesignProjectReviewUnaccepted },
      { label: '复核通过', value: DesignProjectStatus.DesignProjectReviewed },
      { label: '审核不通过', value: [DesignProjectStatus.DesignProjectApproveUnaccepted, DesignProjectStatus.DesignProjectCheckUnAccepted] },
      { label: '审核通过', value: [DesignProjectStatus.DesignProjectApproved, DesignProjectStatus.DesignProjectChecked] }
    ]
    const emptyDesignProject : DesignProject = {
      product_spec_shock_vol: '',
      product_wind_type: '',
      product_wire_shape: '',
      product_industry_freq_hold_vol: '',
      product_induct_high_vol: 0,
      need_master_approve: 0,
      product_frequency: 50,
      product_frequency_special: '',
      product_altitude: 0,
      product_connection_group: 'Dyn11',
      product_cooling_method: 'AN',
      product_insulation_high_ac: 10,
      product_insulation_high_li: 75,
      product_insulation_level: 'F',
      product_insulation_level_special: '',
      product_insulation_low_ac: 10,
      product_insulation_low_li: 0,
      product_load_loss: 1000,
      product_noload_loss: 500,
      product_phrase: PhaseType.Three,
      product_rated_capacity: 800,
      product_rated_v_high: 10,
      product_rated_v_low: 0.4,
      product_short_circuit_resistance: 6,
      product_tap_range: '',
      product_tap_range_special: '',
      product_temp_limit: 100,
      product_total_loss: 1400,
      product_usage: DesignProjectUsage.DesignProjectUsageDefault,
      product_wire_material: '全铜',

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
      reviewer_signed: 0
    }

    const showCreate = ref(false)
    const tempCreate : Ref<DesignProject> = ref(JSON.parse(stringify(emptyDesignProject)) as DesignProject)
    function handleCreate(row ?: DesignProject) {
      if (row) {
        const copy = JSON.parse(stringify(row)) as DesignProject
        console.log(row, copy)
        copy.id = 0
        tempCreate.value = copy
      }
      showCreate.value = true
    }
    const showDetail = ref(false)
    const currentRow : Ref<DesignProject> = ref(JSON.parse(JSON.stringify(emptyDesignProject)))
    function handleDetail(row: DesignProject) {
      currentRow.value = row
      showDetail.value = true
    }
    const showUpdate = ref(false)
    function handleUpdate(row: DesignProject) {
      currentRow.value = row
      showUpdate.value = true
    }
    function sortChange(column: any) {
      query.order_by = column.prop
      query.order = column.order === 'descending' ? 'desc' : 'asc'
      getList()
    }
    async function handleDelete(id: number) {
      await ElMessageBox.confirm('是否删除此设计任务书', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      const res = await DeleteDesignProject(id)
      if (res.data.code === 200) {
        ElMessage.success('删除任务设计书成功')
        await getList()
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
    onMounted(async () => {
      if (route.query['code']) {
        query.field_lk_serial_code = route.query['code'] as string
      }
      Promise.all([
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
        getList()
      ])
    })
    // if (checkRole([RoleList.Designer])) {
    //   CurrentRole.value = 'editor-dashboard'
    // }
    // if (checkRole([RoleList.Reviewer])) {
    //   CurrentRole.value = 'reviewer-dashboard'
    // }
    return {
      downloading,
      handleExport,
      RoleList,
      ResourceList,
      ActionList,
      UnixTime2HumanWithUnix,
      query,
      total,
      listLoading,
      sortChange,
      listData,
      handleFilter,
      selectStatusOptions,
      statusOptions,
      showCreate,
      handleCreate,
      tempCreate,
      currentRow,
      showDetail,
      handleDetail,
      showUpdate,
      handleUpdate,
      handleDelete,
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

<style lang="scss" scoped>
.page-header {
line-height: 24px;
font-size: 24px;
color: #409eff;
background: none;
padding: 0;
}

.filter-container {
  .container-row:first-child {
    margin-bottom: 10px;
  }
  .filter-item:not(button) {
    margin-right: 10px;
  }
}
:deep {
  .project-table-expand {
    &.el-form {
      width: 1100px;
    }
    .el-form-item {
      width: 100%;
    }
    input,textarea,.el-input,.el-radio {
          cursor: text;
          pointer-events: none;
          border: 0 !important;
    }
    .el-select {
      pointer-events: none;
    }
    .el-input__suffix {
      display: none;
    }
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
      .el-form-item__label {
        width: 140px !important;
      }

  }

}
</style>
