<template>
  <div class="app-container">
    <aside class="page-header">
      {{ headerText }}
    </aside>
    <div
      v-loading="paramLoading"
      :class="['design__create-container', dialogAction]"
    >
      <el-row>
        <el-col :span="24">
          <el-form
            ref="validateForm"
            :model="t"
            :rules="formRules"
            :disabled="dialogAction === 'detail'"
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
                      @change="handleCollapseChange"
                    >
                      <!--                  技术规范 start-->
                      <el-collapse-item
                        title="新产品设计"
                        name="1"
                        style="alignment: center"
                      >
                        <el-row :gutter="20">
                          <el-col :span="12">
                            <el-form-item
                              label="项目名称:"
                              prop="design_project.project_name"
                            >
                              <el-select
                                :disabled="dialogAction !== 'create'"
                                :model-value="t.design_project.project_name"
                                placeholder="项目/用户名称"
                                remote
                                filterable
                                style="width: 100%;"
                                :loading="searchProjectLoading"
                                :remote-method="searchProject"
                                @change="handleSelectProject"
                              >
                                <el-option
                                  v-for="item in designProjectOptions"
                                  :key="item"
                                  :label="item.project_name"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col :span="12">
                            <el-form-item
                              label="产品型号:"
                              class="product-model"
                              prop="name"
                            >
                              <el-select
                                v-model="t.name"
                                :disabled="dialogAction !== 'create'"
                                placeholder="产品型号"
                                remote
                                filterable
                                style="width: 100%;"
                                :loading="searchProjectLoading"
                                :remote-method="searchProject"
                              >
                                <el-option
                                  v-for="item in designProjectOptions"
                                  :key="item.value"
                                  :label="item.label"
                                  :value="item.label"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-divider content-position="center">
                          <el-tag type="success">
                            技术规范要求
                          </el-tag>
                        </el-divider>
                        <el-row>
                          <el-col :span="6">
                            <el-form-item
                              label="额定容量(kVA):"
                              prop="input.rated_capacity"
                            >
                              <el-input-number
                                v-model="t.input.rated_capacity"
                                :precision="2"
                                :controls="false"
                                placeholder="额定容量"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="额定高压(kV):"
                              prop="input.rated_high_vol"
                            >
                              <el-input-number
                                v-model="t.input.rated_high_vol"
                                :precision="2"
                                :controls="false"
                                placeholder="额定高压"
                                style="width: 100%"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="额定低压(kV):"
                              prop="input.rated_low_vol"
                            >
                              <el-input-number
                                v-model="t.input.rated_low_vol"
                                :precision="2"
                                :controls="false"
                                placeholder="额定低压"
                                style="width: 100%"

                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="调压范围(%):"
                              class="multiple-select"
                              prop="input.vol_regulation_min"
                            >
                              <el-select
                                v-model="currentRegulation"
                                filterable
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
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="相数:"
                              prop="input.phase_number"
                            >
                              <el-select
                                v-model="currentPhase"
                                filterable
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
                          </el-col>
                          <el-col
                            :span="6"
                            allow-create
                          >
                            <el-form-item
                              label="联结组别:"
                              prop="input.connect_type"
                            >
                              <el-select
                                v-model="t.input.connect_type"
                                filterable
                                allow-create
                                placeholder="联结组别"
                                style="width: 100%"
                              >
                                <el-option
                                  v-for="item in productConnectionGroupOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col
                            :span="6"
                            allow-create
                          >
                            <el-form-item
                              label="额定频率(Hz):"
                              prop="input.rated_frequency"
                            >
                              <el-select
                                v-model="t.input.rated_frequency"
                                filterable
                                allow-create
                                placeholder="额定频率"
                                style="width: 100%"
                                type="number"
                              >
                                <el-option
                                  v-for="item in productFrequencyOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="绝缘等级:"
                              prop="input.insulation_class"
                            >
                              <el-select
                                v-model="t.input.insulation_class"
                                filterable
                                placeholder="绝缘等级"
                                style="width: 100%"
                              >
                                <el-option
                                  v-for="item in libTempRiseOptions"
                                  :key="item"
                                  :label="item.temp_sign"
                                  :value="item.temp_sign"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                        </el-row>
                        <el-row style="margin-top: 20px;padding-top: 20px;border-top: 1px solid #ddd">
                          <el-col :span="6">
                            <el-form-item
                              label="有无拉板"
                              prop="input.has_pull_plate"
                            >
                              <el-select
                                v-model="t.input.has_pull_plate"
                                filterable
                                style="width:100%"
                              >
                                <el-option
                                  v-for="item in ['有','无']"
                                  :key="item"
                                  :label="item"
                                  :value="item"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col
                            :span="6"
                            autofill
                          >
                            <el-form-item
                              label="绝缘温度(℃):"
                              prop="input.insulation_temp"
                            >
                              <el-input-number
                                v-model="t.input.insulation_temp"
                                :precision="2"
                                :controls="false"
                                placeholder="绝缘温度"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="导线材质:"
                              prop="input.wire_material"
                            >
                              <el-select
                                v-model="t.input.wire_material"
                                style="width: 100%"
                              >
                                <el-option
                                  v-for="(item,index) in coilWireOptions"
                                  :key="index"
                                  :label="item.coil_wire_type"
                                  :value="item.coil_wire_type"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="绝缘筒1:"
                              prop="input.insulation_tube1"
                            >
                              <el-select
                                v-model="t.input.insulation_tube1"
                                style="width: 100%"
                              >
                                <el-option
                                  v-for="(item,index) in [{label: '是',value: true},{label: '否',value: false}]"
                                  :key="index"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="绝缘筒2:"
                              prop="input.insulation_tube2"
                            >
                              <el-select
                                v-model="t.input.insulation_tube2"
                                style="width: 100%"
                              >
                                <el-option
                                  v-for="(item,index) in [{label: '是',value: true},{label: '否',value: false}]"
                                  :key="index"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col
                            :span="6"
                            allow-create
                          >
                            <el-form-item
                              label="凸台规格:"
                              prop="input.boss_spec"
                            >
                              <el-select
                                v-model="t.input.boss_spec"
                                filterable
                                allow-create
                                placeholder="凸台规格"
                                style="width: 100%"
                              >
                                <el-option
                                  v-for="item in bossSpecOptions"
                                  :key="item"
                                  :label="item.label"
                                  :value="item.value"
                                />
                              </el-select>
                            </el-form-item>
                          </el-col>
                          <el-col :span="12">
                            <el-row>
                              <el-col :span="12">
                                <el-form-item
                                  style="width: 100%"
                                  label="树脂类型:"
                                  prop="input.resin_type[0]"
                                >
                                  <el-select
                                    v-model="t.input.resin_type[0]"
                                    filterable
                                    placeholder="树脂类型"
                                    @change="
                                      t.input.resin_type[1] = resinTypeOptions.find(item => item.type === t.input.resin_type[0]).price;
                                      t.input.resin_type = [...t.input.resin_type]"
                                  >
                                    <el-option
                                      v-for="item in resinTypeOptions"
                                      :key="item"
                                      :label="item.type"
                                      :value="item.type"
                                    />
                                  </el-select>
                                </el-form-item>
                              </el-col>
                              <el-col
                                :span="12"
                                autofill
                              >
                                <el-form-item
                                  required
                                  style="width: 100%"
                                  label="价格"
                                  prop="input.resin_type[1]"
                                >
                                  <el-input-number
                                    v-model="t.input.resin_type[1]"
                                    :disabled="t.input.resin_type[0] === ''"
                                    placeholder="价格"
                                    :controls="false"
                                    :precision="2"
                                    @change="t.input.resin_type = [...t.input.resin_type]"
                                  />
                                </el-form-item>
                              </el-col>
                            </el-row>
                          </el-col>
                          <el-col :span="24">
                            <el-form-item
                              label="低压导线:"
                            >
                              <el-row style="display: flex;gap: 5px;">
                                <el-col :span="24">
                                  <el-row>
                                    <el-col :span="12">
                                      <el-form-item
                                        label="圆线"
                                        prop="input.lv_wire_type.round[0]"
                                      >
                                        <el-select
                                          v-model="t.input.lv_wire_type.round[0]"
                                          placeholder="圆线"
                                          @change="t.input.lv_wire_type.round[1] = wireTypeOptions.find(item => item.wire_type === $event).wire_price;
                                                   t.input.lv_wire_type.round = [...t.input.lv_wire_type.round]"
                                        >
                                          <el-option
                                            v-for="(item,index) in wireTypeOptions.filter((item) => item.wire_shape === '圆'&& item.wire_material.includes('线'))"
                                            :key="index"
                                            :label="item.wire_type"
                                            :value="item.wire_type"
                                          />
                                        </el-select>
                                      </el-form-item>
                                    </el-col>
                                    <el-col
                                      autofill
                                      :span="12"
                                    >
                                      <el-form-item
                                        label="价格"
                                        prop="input.lv_wire_type.round[1]"
                                      >
                                        <el-input-number
                                          v-model="t.input.lv_wire_type.round[1]"
                                          :disabled="t.input.lv_wire_type.round[0] === ''"
                                          placeholder="价格"
                                          :controls="false"
                                          :precision="2"
                                          @change="t.input.lv_wire_type.round = [...t.input.lv_wire_type.round]"
                                        />
                                      </el-form-item>
                                    </el-col>
                                  </el-row>
                                </el-col>
                                <el-col :span="24">
                                  <el-row>
                                    <el-col :span="12">
                                      <el-form-item
                                        label="扁线"
                                        prop="input.lv_wire_type.flat[0]"
                                      >
                                        <el-select
                                          v-model="t.input.lv_wire_type.flat[0]"
                                          placeholder="扁线"
                                          @change="
                                            t.input.lv_wire_type.flat[1] = wireTypeOptions.find(item => item.wire_type === t.input.lv_wire_type.flat[0]).wire_price;
                                            t.input.lv_wire_type.flat = [...t.input.lv_wire_type.flat]"
                                        >
                                          <el-option
                                            v-for="(item,index) in wireTypeOptions.filter((item) => item.wire_shape === '扁'&& item.wire_material.includes('线'))"
                                            :key="index"
                                            :label="item.wire_type"
                                            :value="item.wire_type"
                                          />
                                        </el-select>
                                      </el-form-item>
                                    </el-col>
                                    <el-col
                                      :span="12"
                                      autofill
                                    >
                                      <el-form-item
                                        label="价格"
                                        prop="input.lv_wire_type.flat[1]"
                                      >
                                        <el-input-number
                                          v-model="t.input.lv_wire_type.flat[1]"
                                          :disabled="t.input.lv_wire_type.flat[0] === ''"
                                          placeholder="价格"
                                          :controls="false"
                                          :precision="2"
                                          @change="t.input.lv_wire_type.flat = [...t.input.lv_wire_type.flat]"
                                        />
                                      </el-form-item>
                                    </el-col>
                                  </el-row>
                                </el-col>
                                <el-col :span="24">
                                  <el-row>
                                    <el-col :span="12">
                                      <el-form-item
                                        label="箔"
                                        prop="input.lv_wire_type.foil[0]"
                                      >
                                        <el-select
                                          v-model="t.input.lv_wire_type.foil[0]"
                                          placeholder="箔"
                                          @change="t.input.lv_wire_type.foil[1] = wireTypeOptions.find(item => item.wire_type === t.input.lv_wire_type.foil[0]).wire_price;
                                                   t.input.lv_wire_type.foil = [...t.input.lv_wire_type.foil]"
                                        >
                                          <el-option
                                            v-for="(item,index) in wireTypeOptions.filter((item) => item.wire_material.includes('箔'))"
                                            :key="index"
                                            :label="item.wire_type"
                                            :value="item.wire_type"
                                          />
                                        </el-select>
                                      </el-form-item>
                                    </el-col>
                                    <el-col
                                      :span="12"
                                      autofill
                                    >
                                      <el-form-item
                                        label="价格"
                                        prop="input.lv_wire_type.foil[1]"
                                      >
                                        <el-input-number
                                          v-model="t.input.lv_wire_type.foil[1]"
                                          :disabled="t.input.lv_wire_type.foil[0] === ''"
                                          placeholder="价格"
                                          :controls="false"
                                          :precision="2"
                                          @change="t.input.lv_wire_type.foil = [...t.input.lv_wire_type.foil]; t.input.lv_wire_type.foil[1] = wireTypeOptions.find(item => item.wire_type === t.input.lv_wire_type.foil[0]).wire_price"
                                        />
                                      </el-form-item>
                                    </el-col>
                                  </el-row>
                                </el-col>
                              </el-row>
                            </el-form-item>
                          </el-col>
                          <el-col :span="24">
                            <el-form-item
                              label="高压导线:"
                            >
                              <el-row style="display: flex;flex-wrap: wrap;gap: 5px;">
                                <el-col :span="24">
                                  <el-row>
                                    <el-col :span="12">
                                      <el-form-item
                                        label="圆线"
                                        prop="input.hv_wire_type.round[0]"
                                      >
                                        <el-select
                                          v-model="t.input.hv_wire_type.round[0]"
                                          placeholder="圆线"
                                          @change="
                                            t.input.hv_wire_type.round[1] = wireTypeOptions.find(item => item.wire_type === t.input.hv_wire_type.round[0]).wire_price;
                                            t.input.hv_wire_type.round = [...t.input.hv_wire_type.round]"
                                        >
                                          <el-option
                                            v-for="(item,index) in wireTypeOptions.filter((item) => item.wire_shape === '圆'&& item.wire_material.includes('线'))"
                                            :key="index"
                                            :label="item.wire_type"
                                            :value="item.wire_type"
                                          />
                                        </el-select>
                                      </el-form-item>
                                    </el-col>
                                    <el-col
                                      :span="12"
                                      autofill
                                    >
                                      <el-form-item
                                        label="价格"
                                        prop="input.hv_wire_type.round[1]"
                                      >
                                        <el-input-number
                                          v-model="t.input.hv_wire_type.round[1]"
                                          :disabled="t.input.hv_wire_type.round[0] === ''"
                                          placeholder="价格"
                                          :controls="false"
                                          :precision="2"
                                          @change="t.input.hv_wire_type.round = [...t.input.hv_wire_type.round]"
                                        />
                                      </el-form-item>
                                    </el-col>
                                  </el-row>
                                </el-col>
                                <el-col :span="24">
                                  <el-row>
                                    <el-col :span="12">
                                      <el-form-item
                                        label="扁线"
                                        prop="input.hv_wire_type.flat[0]"
                                      >
                                        <el-select
                                          v-model="t.input.hv_wire_type.flat[0]"
                                          placeholder="扁线"
                                          @change="
                                            t.input.hv_wire_type.flat[1] = wireTypeOptions.find(item => item.wire_type === t.input.hv_wire_type.flat[0]).wire_price;
                                            t.input.hv_wire_type.flat = [...t.input.hv_wire_type.flat]"
                                        >
                                          <el-option
                                            v-for="(item,index) in wireTypeOptions.filter((item) => item.wire_shape === '扁' && item.wire_material.includes('线'))"
                                            :key="index"
                                            :label="item.wire_type"
                                            :value="item.wire_type"
                                          />
                                        </el-select>
                                      </el-form-item>
                                    </el-col>
                                    <el-col
                                      :span="12"
                                      autofill
                                    >
                                      <el-form-item
                                        label="价格"
                                        prop="input.hv_wire_type.flat[1]"
                                      >
                                        <el-input-number
                                          v-model="t.input.hv_wire_type.flat[1]"
                                          :disabled="t.input.hv_wire_type.flat[0] === ''"
                                          placeholder="价格"
                                          :controls="false"
                                          :precision="2"
                                          @change="t.input.hv_wire_type.flat = [...t.input.hv_wire_type.flat]; t.input.hv_wire_type.flat[1] = wireTypeOptions.find(item => item.wire_type === t.input.hv_wire_type.flat[0]).wire_price"
                                        />
                                      </el-form-item>
                                    </el-col>
                                  </el-row>
                                </el-col>
                              </el-row>
                            </el-form-item>
                          </el-col>
                          <el-col :span="24">
                            <el-form-item
                              label="钢片牌号:"
                            >
                              <el-row style="height: 40px;" />
                              <el-row
                                v-for="(item,index) in t.input.steel_type"
                                :key="index"
                                style="margin-bottom: 20px;"
                              >
                                <el-col :span="7">
                                  <el-form-item
                                    label="钢片牌号"
                                    label-width="100px"
                                  >
                                    <el-select
                                      v-model="t.input.steel_type[index][0]"
                                      placeholder=""
                                      @change="t.input.steel_type = [...t.input.steel_type];t.input.steel_type[index][1] = stackOptions.find(item => item.type === $event)?.price"
                                    >
                                      <el-option
                                        v-for="(item,index) in stackOptions"
                                        :key="index"
                                        :label="item.type"
                                        :value="item.type"
                                      />
                                    </el-select>
                                  </el-form-item>
                                </el-col>
                                <el-col
                                  autofill
                                  :span="7"
                                >
                                  <el-form-item
                                    label="价格"
                                    label-width="70px"
                                  >
                                    <el-input-number
                                      v-model="t.input.steel_type[index][1]"
                                      :disabled="t.input.steel_type[index][0] === ''"
                                      placeholder="价格"
                                      :controls="false"
                                      :precision="2"
                                      @change="t.input.steel_type = [...t.input.steel_type]"
                                    />
                                  </el-form-item>
                                </el-col>
                                <el-col
                                  :span="7"
                                  autofill
                                >
                                  <el-form-item
                                    label="迭片系数"
                                    label-width="100px"
                                  >
                                    <el-input-number
                                      v-model="t.input.steel_type[index][2]"
                                      :disabled="t.input.steel_type[index][0] === ''"
                                      placeholder="迭片系数"
                                      :controls="false"
                                      :precision="2"
                                      @change="t.input.steel_type = [...t.input.steel_type]"
                                    />
                                  </el-form-item>
                                </el-col>
                                <el-col
                                  :span="2"
                                  :offset="1"
                                >
                                  <el-button
                                    v-if="index !==0 "
                                    type="text"
                                    @click="t.input.steel_type.splice(index,1)"
                                  >
                                    <el-icon :size="20">
                                      <circle-close-filled />
                                    </el-icon>
                                  </el-button>
                                </el-col>
                              </el-row>
                              <el-row>
                                <el-button
                                  type="primary"
                                  size="small"
                                  @click="t.input.steel_type.push(['',0,0])"
                                >
                                  <el-icon class="el-icon--right">
                                    <CirclePlusFilled />
                                  </el-icon>
                                  添加钢片牌号
                                </el-button>
                              </el-row>
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>
                      <!--                  性能参数标准 start-->
                      <el-collapse-item

                        name="3"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">性能参数标准值(如损耗、温升、噪音预测等)<i class="header-icon el-icon-edit" /></span>
                        </template>
                        <el-row autofill>
                          <el-col :span="6">
                            <el-form-item
                              label="总损耗(W):"
                              prop="input.limit_total_loss"
                            >
                              <el-input-number
                                v-model="t.input.limit_total_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="总损耗"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="负载损耗(W):"
                              prop="input.limit_load_loss"
                            >
                              <el-input-number
                                v-model="t.input.limit_load_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="负载损耗"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="空载损耗(W):"
                              prop="input.limit_no_load_loss"
                            >
                              <el-input-number
                                v-model="t.input.limit_no_load_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="空载损耗"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="空载电流(%):"
                              prop="input.limit_no_load_current"
                            >
                              <el-input-number
                                v-model="t.input.limit_no_load_current"
                                :precision="2"
                                :controls="false"
                                placeholder="空载电流"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="阻抗电压(%):"
                              prop="input.limit_impedance_vol"
                            >
                              <el-input-number
                                v-model="t.input.limit_impedance_vol"
                                :precision="2"
                                :controls="false"
                                placeholder="阻抗电压"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="低压温升(K):"
                              prop="input.limit_lv_temp_rise"
                            >
                              <el-input-number
                                v-model="t.input.limit_lv_temp_rise"
                                :controls="false"
                                :precision="2"
                                placeholder="低压温升"
                                style="width: 100%"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="高压温升(K):"
                              prop="input.limit_hv_temp_rise"
                            >
                              <el-input-number
                                v-model="t.input.limit_hv_temp_rise"
                                :controls="false"
                                :precision="2"
                                placeholder="高压温升"
                                style="width: 100%"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="噪音预测(dB):"
                              prop="input.limit_noise_prediction"
                            >
                              <el-input-number
                                v-model="t.input.limit_noise_prediction"
                                :controls="false"
                                :precision="2"
                                placeholder="噪音预测"
                                style="width: 100%"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                        </el-row>
                      </el-collapse-item>
                      <el-collapse-item

                        name="4"
                      >
                        <template
                          #title
                        >
                          <span class="collapse-title">允许的性能参数偏差范围<i class="header-icon el-icon-edit" /></span>
                        </template>
                        <el-row>
                          <el-form-item
                            label="参数偏差范围标准："
                            style="width: 100%;"
                          >
                            <el-select
                              v-model="deviationStd"
                            >
                              <el-option
                                v-for="(item,index) in ['国标','自定义']"
                                :key="index"
                                :label="item"
                                :value="item"
                              />
                            </el-select>
                          </el-form-item>
                        </el-row>
                        <el-row autofill>
                          <el-col :span="6">
                            <el-form-item
                              label="总损耗(%):"
                              prop="input.deviation_total_loss"
                            >
                              <el-input-number
                                v-model="t.input.deviation_total_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="总损耗"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="负载损耗(%):"
                              prop="input.deviation_load_loss"
                            >
                              <el-input-number
                                v-model="t.input.deviation_load_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="负载损耗"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="空载损耗(%):"
                              prop="input.deviation_no_load_loss"
                            >
                              <el-input-number
                                v-model="t.input.deviation_no_load_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="空载损耗"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="空载电流(%):"
                              prop="input.deviation_no_load_current"
                            >
                              <el-input-number
                                v-model="t.input.deviation_no_load_current"
                                :precision="2"
                                :controls="false"
                                placeholder="空载电流"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="阻抗电压(%):"
                              prop="input.deviation_impedance_vol"
                            >
                              <el-input-number
                                v-model="t.input.deviation_impedance_vol"
                                :precision="2"
                                :controls="false"
                                placeholder="阻抗电压"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="低压温升(%):"
                              prop="input.deviation_lv_temp_rise"
                            >
                              <el-input-number
                                v-model="t.input.deviation_lv_temp_rise"
                                :precision="2"
                                :controls="false"
                                placeholder="低压温升"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="高压温升(%):"
                              prop="input.deviation_hv_temp_rise"
                            >
                              <el-input-number
                                v-model="t.input.deviation_hv_temp_rise"
                                :precision="2"
                                :controls="false"
                                placeholder="高压温升"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="6">
                            <el-form-item
                              label="噪音预测(%):"
                              prop="input.deviation_noise_prediction"
                            >
                              <el-input-number
                                v-model="t.input.deviation_noise_prediction"
                                :precision="2"
                                :controls="false"
                                placeholder="噪音预测"
                                type="number"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col
                            :span="6"
                            autofill
                          >
                            <el-form-item
                              label="涡流损耗系数(%):"
                              prop="input.eddy_current_loss_rate"
                            >
                              <el-input-number
                                v-model="t.input.eddy_current_loss_rate"
                                :precision="2"
                                :controls="false"
                                placeholder="涡流损耗系数"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col
                            :span="6"
                            autofill
                          >
                            <el-form-item
                              label="引线损耗(W):"
                              prop="input.lead_loss"
                            >
                              <el-input-number
                                v-model="t.input.lead_loss"
                                :precision="2"
                                :controls="false"
                                placeholder="引线损耗"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col
                            :span="6"
                            autofill
                          >
                            <el-form-item
                              label="ks:"
                              prop="input.ks"
                            >
                              <el-input-number
                                v-model="t.input.ks"
                                :precision="2"
                                :controls="false"
                                placeholder="ks"
                                type="number"
                                style="width: 100%"
                              />
                            </el-form-item>
                          </el-col>
                          <el-col :span="24">
                            <el-form-item
                              label="温升修正系数:"
                            >
                              <!--                              低压-->
                              <el-row style="height: 40px;" />
                              <el-row style="margin-bottom: 10px;">
                                <el-col :span="12">
                                  <el-form-item label="绝缘筒1">
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.insula_tube1"
                                      :precision="2"
                                      :controls="false"
                                    />
                                  </el-form-item>
                                </el-col>
                                <el-col :span="12">
                                  <el-form-item label="绝缘筒2">
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.insula_tube2"
                                      :precision="2"
                                      :controls="false"
                                    />
                                  </el-form-item>
                                </el-col>
                              </el-row>
                              <el-row style="margin-bottom: 10px;">
                                <el-form-item
                                  label="低压"
                                  style="width: 100%;"
                                >
                                  <el-row style="display: flex;flex-wrap: nowrap;gap: 5px;">
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[0]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[1]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[2]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[3]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                  </el-row>
                                </el-form-item>
                              </el-row>
                              <el-row>
                                <el-form-item
                                  label="高压"
                                  style="width: 100%;"
                                >
                                  <el-row style="display: flex;flex-wrap: nowrap;gap: 5px;">
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[0]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[1]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[2]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                    <el-input-number
                                      v-model="t.input.temp_rise_coeff.lv[3]"
                                      :controls="false"
                                      :precision="2"
                                      style="width: 100%"
                                      type="number"
                                      @change="t.input.temp_rise_coeff.lv = [...t.input.temp_rise_coeff.lv]"
                                    />
                                  </el-row>
                                </el-form-item>
                              </el-row>
                            </el-form-item>
                          </el-col>
                          <el-col :span="24">
                            <el-form-item
                              label="散热面修正系数:"
                            >
                              <el-row style="height: 40px;" />
                              <el-row style="margin-bottom: 10px;">
                                <el-col :span="12">
                                  <el-form-item label="绝缘筒1">
                                    <el-row
                                      style="display: flex;flex-wrap: nowrap;"
                                      :gutter="5"
                                    >
                                      <el-col :span="12">
                                        <el-input-number
                                          v-model="t.input.radiating_surface_coeff.insula_tube1[0]"
                                          :precision="2"
                                          :controls="false"
                                          @change="t.input.radiating_surface_coeff.insula_tube1 = [...t.input.radiating_surface_coeff.insula_tube1]"
                                        />
                                      </el-col>
                                      <el-col :span="12">
                                        <el-input-number
                                          v-model="t.input.radiating_surface_coeff.insula_tube1[1]"
                                          :precision="2"
                                          :controls="false"
                                          @change="t.input.radiating_surface_coeff.insula_tube1 = [...t.input.radiating_surface_coeff.insula_tube1]"
                                        />
                                      </el-col>
                                    </el-row>
                                  </el-form-item>
                                </el-col>
                                <el-col :span="12">
                                  <el-form-item label="绝缘筒2">
                                    <el-row
                                      style="display: flex;flex-wrap: nowrap;"
                                      :gutter="5"
                                    >
                                      <el-col :span="12">
                                        <el-input-number
                                          v-model="t.input.radiating_surface_coeff.insula_tube2[0]"
                                          :precision="2"
                                          :controls="false"
                                          @change="t.input.radiating_surface_coeff.insula_tube2 = [...t.input.radiating_surface_coeff.insula_tube2]"
                                        />
                                      </el-col>
                                      <el-col :span="12">
                                        <el-input-number
                                          v-model="t.input.radiating_surface_coeff.insula_tube2[1]"
                                          :precision="2"
                                          :controls="false"
                                          @change="t.input.radiating_surface_coeff.insula_tube2 = [...t.input.radiating_surface_coeff.insula_tube2]"
                                        />
                                      </el-col>
                                    </el-row>
                                  </el-form-item>
                                </el-col>
                              </el-row>
                              <el-row style="margin-bottom: 10px;">
                                <el-form-item
                                  label="低压"
                                  style="width: 100%;"
                                >
                                  <el-row
                                    style="display: flex;flex-wrap: nowrap;"
                                    :gutter="5"
                                  >
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[0][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[0][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[1][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[1][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[2][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[2][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[3][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.lv[3][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.lv = [...t.input.radiating_surface_coeff.lv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                  </el-row>
                                </el-form-item>
                              </el-row>
                              <el-row>
                                <el-form-item
                                  label="高压"
                                  style="width: 100%;"
                                >
                                  <el-row
                                    style="display: flex;flex-wrap: nowrap;"
                                    :gutter="5"
                                  >
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[0][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[0][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[1][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[1][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[2][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[2][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                    <el-col :span="6">
                                      <el-row>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[3][0]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                        <el-col :span="12">
                                          <el-input-number
                                            v-model="t.input.radiating_surface_coeff.hv[3][1]"
                                            :controls="false"
                                            :precision="2"
                                            style="width: 100%"
                                            type="number"
                                            @change="t.input.radiating_surface_coeff.hv = [...t.input.radiating_surface_coeff.hv]"
                                          />
                                        </el-col>
                                      </el-row>
                                    </el-col>
                                  </el-row>
                                </el-form-item>
                              </el-row>
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
      <el-row
        v-if="dialogAction !== 'detail'"
        style="margin-top: 1%"
      >
        <el-col>
          <el-button
            type="primary"
            @click="handleConfirmCreate"
          >
            提交设计任务
          </el-button>
          <el-button
            type="success"
            @click="handleShowRecommend()"
          >
            查看相似设计
          </el-button>
        </el-col>
      </el-row>
      <DesignRecommend
        v-model="showRecommendTask"
        v-model:handle-value="t"
      />
    </div>
  </div>
</template>
<script lang="ts">
import { ElMessage } from 'element-plus'
import { CircleCloseFilled, CirclePlusFilled } from '@element-plus/icons-vue'
import { createDesignTask, getDesignTaskList, updateDesignTask } from '@/api/design'
import { GetDesignProjects } from '@/api/projects'
import { GetStdLpnLibraries, Lpn } from '@/api/standard_libraries/lpn'
import { GetStdLrrLibraries, ListLrrQuery, Lrr } from '@/api/standard_libraries/lrr'
import {
  DecodeDesignTask,
  DesignProject,
  DesignProjectQuery,
  DesignTask, DesignTaskQuery,
  DesignTaskStatus,
  PullType,
  SeamType,
  TransformerTaskForm
} from '@/model/design'

import { GetStdLrfLibraries, Lrf } from '@/api/standard_libraries/lrf'
import { computed, defineComponent, onMounted, PropType, reactive, ref, Ref, watchEffect } from 'vue'
import { GetStdLsfLibraries, Lsf } from '@/api/standard_libraries/lsf'
import { GetStdLtfLibraries, Ltf } from '@/api/standard_libraries/ltf'
import { GetStdlelfLibraries, ListLelfQuery } from '@/api/standard_libraries/lelf'
import { useRoute, useRouter } from 'vue-router'
import { getValidateFormDefaultRules } from '@/utils/validate'
import DesignRecommend from './components/design-recommand.vue'
import { Boss, GetStdBossLibraries } from '@/api/standard_libraries/boss'
import { GetStdlrLibraries, ListLrQuery, Lr } from '@/api/standard_libraries/lr'
import { GetStdLtrLibraries, Ltr } from '@/api/standard_libraries/ltr'
import { GetStdLlLibraries, ListLlQuery } from '@/api/standard_libraries/ll'
import { decodeObjectBytes, encodeBytesData } from '@/utils/jsonutils'
import { GetStdlwireLibraries, Lwire } from '@/api/standard_libraries/lwire'
import { GetStdlstackLibraries, Lstack } from '@/api/standard_libraries/lstack'
import { GetStdLnpLibraries, ListLnpQuery } from '@/api/standard_libraries/lnp'
import { GetStdLlsLibraries, ListLlsQuery } from '@/api/standard_libraries/lls'
import lcwt from '@/views/standard/libraries/lcwt/index.vue'
import { GetStdlcwtLibraries, Lcwt, ListLcwtQuery } from '@/api/standard_libraries/lcwt'
export default defineComponent({
  components: {
    CircleCloseFilled,
    CirclePlusFilled,
    DesignRecommend
  },
  props: {
    dialogActionProps: {
      type: String as PropType<'create' | 'update' | 'detail'>,
      default: () => 'create'
    },
    taskId: {
      type: Number,
      default: () => 0
    }
  },
  setup(props) {
    const validateForm = ref(null)
    const formProps = [
      'design_project.project_name',
      'name',
      'input.rated_capacity',
      'input.rated_high_vol',
      'input.rated_low_vol',
      'input.vol_regulation_min',
      'input.phase_number',
      'input.connect_type',
      'input.rated_frequency',
      'input.insulation_class',
      'input.has_pull_plate',
      'input.insulation_temp',
      'input.wire_material',
      'input.insulation_tube1',
      'input.insulation_tube2',
      'input.boss_spec',
      'input.resin_type[0]',
      'input.resin_type[1].number.nzero',
      'input.lv_wire_type.round[0]',
      'input.lv_wire_type.round[1]',
      'input.lv_wire_type.flat[0]',
      'input.lv_wire_type.flat[1]',
      'input.lv_wire_type.foil[0]',
      'input.lv_wire_type.foil[1]',
      'input.hv_wire_type.round[0]',
      'input.hv_wire_type.round[1]',
      'input.hv_wire_type.flat[0]',
      'input.hv_wire_type.flat[1]',
      'input.limit_total_loss',
      'input.limit_load_loss',
      'input.limit_no_load_loss',
      'input.limit_no_load_current',
      'input.limit_impedance_vol',
      'input.limit_lv_temp_rise',
      'input.limit_hv_temp_rise',
      'input.limit_noise_prediction',
      'input.eddy_current_loss_rate.number.nzero',
      'input.lead_loss.number.nzero',
      'input.ks.number.nzero'

    ]
    const formRules = getValidateFormDefaultRules(formProps)

    const route = useRoute()
    const router = useRouter()
    const dialogAction = ref(props.dialogActionProps)
    const headerText = computed(() => {
      if (dialogAction.value === 'create') {
        return '新建设计'
      }
      if (dialogAction.value === 'update') {
        return '修改设计'
      }
      if (dialogAction.value === 'detail') {
        return '设计详情'
      }
      return ''
    })
    const taskStatusOptions = [
      { label: '撤销', value: DesignTaskStatus.DesignTaskCanceled },
      { label: '创建中', value: DesignTaskStatus.DesignTaskCreated },
      { label: '计算中', value: DesignTaskStatus.DesignTaskStarted },
      { label: '已完成', value: DesignTaskStatus.DesignTaskFinished },
      { label: '审核不通过', value: DesignTaskStatus.DesignTaskReviewUnaccepted },
      { label: '审核通过', value: DesignTaskStatus.DesignTaskReviewed },
      { label: '不批准', value: DesignTaskStatus.DesignTaskApproveUnaccepted },
      { label: '批准', value: DesignTaskStatus.DesignTaskApproved },
      { label: '不批准', value: DesignTaskStatus.DesignTaskCheckUnaccepted },
      { label: '批准', value: DesignTaskStatus.DesignTaskChecked }
    ]
    const defaultTableRow : DecodeDesignTask = {
      approve: 0,
      approve_by_creator: false,
      approve_node_instance: undefined,
      comment: '',
      creator: 0,
      result_progress: 0,
      design_project: {
        check_by: 0,
        check_user: undefined,
        checked_at: 0,
        comment: '',
        created_at: '',
        creator: 0,
        creator_user: undefined,
        deleted_at: '',
        deliver_at: 0,
        design_at: 0,
        design_comment: '',
        design_finish_at: 0,
        designer: 0,
        designer_signed: 0,
        designer_user: undefined,
        drafting_at: 0,
        drafting_by: 0,
        drafting_user: undefined,
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
        reviewer_user: undefined,
        serial_code: '',
        tech_requirments: '',
        updated_at: ''

      },
      design_project_id: 0,
      design_results: '',
      designer: 0,
      final_design_results: [],
      id: 0,
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
        hv_wire_type: { flat: ['', 0], round: ['', 0] },
        id: 0,
        insulation_class: '',
        insulation_temp: 0,
        has_pull_plate: '有',
        insulation_tube1: '否',
        insulation_tube2: '否',
        ks: 1,
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
        radiating_surface_coeff: { hv: [[1, 1], [1, 1], [1, 1], [1, 1]], insula_tube1: [1, 1], insula_tube2: [1, 1], lv: [[1, 1], [1, 1], [1, 1], [1, 1]] },
        rated_capacity: 0,
        rated_frequency: 0,
        rated_high_vol: 0,
        rated_low_vol: 0,
        resin_type: ['', 0],
        steel_type: [['', 0, 0]],
        temp_rise_coeff: { hv: [1, 1, 1, 1], insula_tube1: 1, insula_tube2: 1, lv: [1, 1, 1, 1] },
        vol_regulation_max: 0,
        vol_regulation_min: 0,
        vol_regulation_step: 0,
        wire_material: ''
      } as TransformerTaskForm,
      name: '',
      reviewer: 0,
      task_status: 0
    }
    const pullTypeOptions = [
      { label: '有拉板', value: PullType.WithPull },
      { label: '无拉板', value: PullType.WithoutPull }
    ]
    const seamTypeOptions = [
      { label: '全斜', value: SeamType.FullOblique },
      { label: '半斜', value: SeamType.HalfOblique }
    ]
    // 调压范围
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
    const currentRegulation = ref('')
    watchEffect(() => {
      if (!currentRegulation.value) {
        return
      }
      try {
        const regulationRange = JSON.parse(currentRegulation.value) as number[]
        t.value.input.vol_regulation_max = Math.max(...regulationRange)
        t.value.input.vol_regulation_min = Math.min(...regulationRange)
        t.value.input.vol_regulation_step = Math.abs(regulationRange[1] - regulationRange[0])
      } catch (e) {
        console.log(e)
      }
    })
    const currentPhase = ref('')
    watchEffect(() => {
      if (!currentPhase.value) {
        return
      }
      t.value.input.phase_number = currentPhase.value === '三相' ? 3 : 1
    })

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
            value: String(i.phase)
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
    // // 温升限值
    const libTempRiseOptions : Ref<Ltr[]> = ref([])
    const getStdLtrLibraries = async () => {
      const res = await GetStdLtrLibraries({
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        libTempRiseOptions.value = res.data.spec.filter(item => item.temp_sign !== '/')
      }
    }
    // 联结组别
    const productConnectionGroupOptions = [
      { label: 'Dyn11', value: 'Dyn11' },
      { label: 'Yyn0', value: 'Yyn0' },
      { label: 'Dd0', value: 'Dd0' },
      { label: 'Yd11', value: 'Yd11' }
    ]
    // 导线材质
    const wireTypeOptions : Ref<Lwire[]> = ref([])
    const getStdWireLibraries = async () => {
      const res = await GetStdlwireLibraries({
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        wireTypeOptions.value = res.data.spec
      }
    }
    // 硅钢片
    const stackOptions : Ref<Lstack[]> = ref([])
    const getStdStackLibraries = async () => {
      const res = await GetStdlstackLibraries({
        page: 1,
        limit: 9999,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        stackOptions.value = res.data.spec
      }
    }
    const expLayParamOptions : Ref<Lsf[]> = ref([])
    const getStdLsfLibraries = async () => {
      const res = await GetStdLsfLibraries({
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        expLayParamOptions.value = res.data.spec
      }
    }
    const expProParamOptions : Ref<{label: string, value: number}[]> = ref([])
    const getStdLtfLibraries = async () => {
      const res = await GetStdLtfLibraries({
        field_eq_stack_amount: null,
        field_eq_tech_factor: null,
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        expProParamOptions.value = []
        res.data.spec.map((i: Ltf) => {
          if (expProParamOptions.value.filter(item => item.value === i.tech_factor).length === 0) {
            expProParamOptions.value.push({ label: i.tech_factor!.toString(), value: i.tech_factor! })
          }
        })
      }
    }
    // 凸台规格
    const bossSpecOptions : Ref<{label: string, value: string}[]> = ref([])
    const getStdBossLibraries = async () => {
      const res = await GetStdBossLibraries({
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        bossSpecOptions.value = []
        res.data.spec.map((i: Boss) => {
          bossSpecOptions.value.push({ label: i.boss_spec, value: i.boss_spec })
        })
      }
    }
    // 树脂类型
    const resinTypeOptions : Ref<Lr[]> = ref([])
    const getStdResinLibraries = async () => {
      const res = await GetStdlrLibraries({
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      } as ListLrQuery)
      if (res.data.code === 200) {
        resinTypeOptions.value = res.data.spec
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
    // const wireTypeOptions : Ref<wire[]> = ref([])

    const t : Ref<DecodeDesignTask> = ref(JSON.parse(JSON.stringify(defaultTableRow)) as DecodeDesignTask)
    // 损耗

    watchEffect(async () => {
      if (t.value.input.vol_regulation_step && t.value.design_project_id && t.value.input.rated_capacity &&
        t.value.input.rated_high_vol && t.value.input.rated_low_vol && t.value.input.insulation_class && t.value.input.insulation_temp) {
        const project = t.value.design_project || designProjectOptions?.value.find(p => p.value === t.value.design_project_id)?.project
        if (!project || !project.product_model) {
          return
        }
        console.log(project)
        let regulate_way = '无励磁调压'
        // 10/11/12/13型
        let lossLevel = ''
        const nameLike = project?.product_model.trim().match(/(ZL|ZB)?([DS])(C)(Z?)(B?)(10|11|12|13|15|16)-([0-9.]*)\/([0-9.]*)\/([0-9.]*)/)
        if (nameLike) {
          if (nameLike[4] === 'Z') {
            regulate_way = '有载调压'
          }
          if (!nameLike[6]) {
            return
          }
          lossLevel = nameLike[6]
        }
        const llStdQuery = reactive({
          limit: 1,
          order: '',
          order_by: '',
          page: 1,
          field_eq_regulate_way: regulate_way,
          field_le_rated_capacity_min: t.value.input.rated_capacity,
          field_gt_rated_capacity_max: t.value.input.rated_capacity,
          field_le_rated_high_vol_min: t.value.input.rated_high_vol,
          field_gt_rated_high_vol_max: t.value.input.rated_high_vol,
          field_le_rated_low_vol_min: t.value.input.rated_low_vol,
          field_gt_rated_low_vol_max: t.value.input.rated_low_vol,
          field_eq_temperature: t.value.input.insulation_temp,
          field_eq_regulate_range_min: t.value.input.vol_regulation_min,
          field_eq_regulate_range_max: t.value.input.vol_regulation_max,
          field_eq_regulate_range_step: t.value.input.vol_regulation_step,
          filed_eq_loss_level: lossLevel
        } as ListLlQuery)
        const res = await GetStdLlLibraries(llStdQuery)
        if (res.data.code === 200) {
          if (res.data.spec.length === 1) {
            t.value.input.limit_load_loss = res.data.spec[0].load_loss!
            t.value.input.limit_no_load_current = res.data.spec[0].no_load_current!
            t.value.input.limit_no_load_loss = res.data.spec[0].no_load_loss!
            t.value.input.limit_total_loss = res.data.spec[0].load_loss! + res.data.spec[0].no_load_loss!
            t.value.input.limit_impedance_vol = res.data.spec[0].short_circuit_imped!
          }
        }
        const lnpRes = await GetStdLnpLibraries({
          limit: 1,
          order: '',
          order_by: '',
          page: 1,
          field_le_rated_capacity_min: t.value.input.rated_capacity,
          field_gt_rated_capacity_max: t.value.input.rated_capacity,
          field_le_rated_high_vol_min: t.value.input.rated_high_vol,
          field_gt_rated_high_vol_max: t.value.input.rated_high_vol,
          field_eq_cool_type: project.product_cooling_method
        } as ListLnpQuery)
        if (lnpRes.data.code === 200) {
          if (lnpRes.data.spec.length === 1) {
            t.value.input.limit_noise_prediction = lnpRes.data.spec[0].noise_predict!
          }
        }
      }
    })
    // 绝缘等级
    watchEffect(() => {
      const tempRise = libTempRiseOptions.value.find(item => item.temp_sign === t.value.input.insulation_class)
      t.value.input.insulation_temp = tempRise?.temp || 0
      t.value.input.limit_lv_temp_rise = tempRise?.low_vol_temp_rise || 0
      t.value.input.limit_hv_temp_rise = tempRise?.high_vol_temp_rise || 0
    })
    // 迭片系数
    watchEffect(() => {
      t.value.input.steel_type.forEach((item, index) => {
        if (item[0]) {
          item[2] = expLayParamOptions.value.find((iitem) => iitem.stack_thick === parseInt(item[0].substr(0, 2)) / 100)?.stack_factor as number
        }
      })
    })
    // 参数偏差范围
    const deviationStd = ref('国标')
    watchEffect(() => {
      if (deviationStd.value === '国标') {
        t.value.input.deviation_total_loss = 10
        t.value.input.deviation_load_loss = 15
        t.value.input.deviation_no_load_loss = 10
      } else {
        t.value.input.deviation_total_loss = 0
        t.value.input.deviation_load_loss = 0
        t.value.input.deviation_no_load_loss = 0
        t.value.input.deviation_impedance_vol = 0
        t.value.input.deviation_lv_temp_rise = 0
        t.value.input.deviation_hv_temp_rise = 0
        t.value.input.deviation_noise_prediction = 0
      }
    })
    // 涡流损耗系数，引线损耗
    watchEffect(async () => {
      const lelfList = GetStdlelfLibraries({
        field_eq_rated_capacity: t.value.input.rated_capacity,
        limit: 1,
        page: 1,
        order: '',
        order_by: ''
      } as ListLelfQuery)
      const llsList = GetStdLlsLibraries({
        limit: 1, order: '', order_by: '', page: 1
      } as ListLlsQuery)
      const [res1, res2] = await Promise.all([lelfList, llsList])
      if (res1.data.code === 200) {
        if (res1.data.spec.length > 0) {
          t.value.input.eddy_current_loss_rate = res1.data.spec[0].eddy_loss_factor! / 100
        }
      }
      if (res2.data.code === 200) {
        if (res2.data.spec.length > 0) {
          t.value.input.lead_loss = res2.data.spec[0].lead_loss
        }
      }
    })
    const activeCollapseNames = ref(['1'])
    activeCollapseNames.value = ['1']
    function handleCollapseChange(value: any) {
      console.log('collapse change', activeCollapseNames, value)
    }
    const designProjectList : Ref<DesignProject[]> = ref([])
    const designProjectOptions = computed(() => designProjectList.value.map((p) => { return { value: p.id, label: p.product_model, project_name: p.project_name, project: p } }))
    const projectQuery : DesignProjectQuery = reactive({
      field_lk_project_name: '',
      field_lk_product_model: '',
      limit: 10,
      page: 1,
      order: '',
      order_by: ''
    })
    const searchProjectLoading = ref(false)
    async function searchProject(queryString: string) {
      projectQuery.field_lk_product_model = queryString
      searchProjectLoading.value = true
      const res = await GetDesignProjects(projectQuery)
      if (res.data.code === 200) {
        designProjectList.value = res.data.spec
      }
      searchProjectLoading.value = false
    }
    function handleSelectProject(design_project_id: number) {
      t.value.design_project_id = design_project_id
      const project = designProjectOptions.value.find(p => p.value === design_project_id)
      if (project) {
        t.value.design_project!.project_name = designProjectOptions.value.find(p => p.value === design_project_id)?.project_name || ''
        t.value.name = designProjectOptions.value.find(p => p.value === design_project_id)?.label || ''
        // 设置参数  //按顺序
        const regulationRange = JSON.parse(project.project.product_tap_range!) as number[]

        t.value.input = {
          ...t.value.input,
          wire_material: project.project.product_wire_material,
          rated_capacity: project.project.product_rated_capacity,
          rated_high_vol: project.project.product_rated_v_high,
          rated_low_vol: project.project.product_rated_v_low,
          limit_impedance_vol: project.project.product_short_circuit_resistance,
          phase_number: project.project.product_phrase === '三相' ? 3 : 1,
          connect_type: project.project.product_connection_group,
          rated_frequency: project.project.product_frequency,
          insulation_class: project.project.product_insulation_level,
          limit_load_loss: project.project.product_load_loss,
          limit_no_load_loss: project.project.product_noload_loss,
          limit_total_loss: project.project.product_total_loss
        }
        currentRegulation.value = String(project.project.product_tap_range)
        currentPhase.value = project.project.product_phrase
      }
    }
    async function handleConfirmCreate() {
      (validateForm.value as any).validate(async (valid : Boolean, fields: any) => {
        if (!valid) {
          for (const key in fields) {
            ElMessage.error(fields[key][0].message)
            return
          }
        }
        paramLoading.value = true

        const designTask = JSON.parse(JSON.stringify(t.value, (k, v) => {
          if (!k) { return v }
          if (typeof v === 'object') { return null }
          return v
        })) as DesignTask
        designTask.input = encodeBytesData(t.value.input)
        if (dialogAction.value === 'create') {
          const response = await createDesignTask(designTask)
          if (response.data.code === 200) {
            ElMessage.success('设计任务创建成功')
            router.push(`/design/design_check?code=${t.value.design_project!.serial_code}`)
          }
        }
        if (dialogAction.value === 'update') {
          const response = await updateDesignTask(designTask)
          if (response.data.code === 200) {
            ElMessage.success('设计任务更新成功')
            router.push(`/design/design_check?code=${t.value.design_project!.serial_code}`)
          }
        }
        paramLoading.value = false
      })
    }
    // 查看相似设计
    const showRecommendTask = ref(false)
    function handleShowRecommend() {
      showRecommendTask.value = true
      if (!activeCollapseNames.value.includes('8')) {
        activeCollapseNames.value = [...activeCollapseNames.value, '8']
      }
    }

    // 参数加载
    const paramLoading = ref(true)
    onMounted(async () => {
      Promise.all([
      // 调压范围
        getSpecRange(),
        // 相数
        getStdLpnLibraries(),
        // 额定频率
        getStdLrfLibraries(),
        // 温升限值
        getStdLtrLibraries(),
        // 导线类型
        getStdWireLibraries(),
        // 硅钢片
        getStdStackLibraries(),
        // 迭片系数
        getStdLsfLibraries(),
        // 工艺系数
        getStdLtfLibraries(),
        // 凸台规格
        getStdBossLibraries(),
        // 树脂类型
        getStdResinLibraries(),
        // 线圈材质
        getStdlcwtLibraries()
      ]).then(() => {
        paramLoading.value = false
      })

      // 初始化
      if (route.query.id || props.taskId) {
        const id = parseInt(String(route.query.id || props.taskId))
        const res = await getDesignTaskList({
          'field_eq_design_task.id': id,
          limit: 1,
          order: '',
          order_by: '',
          page: 1
        } as DesignTaskQuery)
        if (res.data.code === 200 && res.data.spec[0]) {
          decodeObjectBytes(res.data.spec[0])
          const designTask = res.data.spec[0] as DecodeDesignTask
          designTask.name = designTask.design_project!.product_model
          const project = designTask.design_project!
          const stepRegulation :number[] = []
          const stepRegulationNum = (designTask.input.vol_regulation_max - designTask.input.vol_regulation_min) / designTask.input.vol_regulation_step - 1
          for (let i = 0; i < stepRegulationNum; i++) {
            stepRegulation.push(designTask.input.vol_regulation_min + designTask.input.vol_regulation_step * (i + 1))
          }
          currentRegulation.value = `[${[
            designTask.input.vol_regulation_min,
            ...stepRegulation,
            designTask.input.vol_regulation_max
          ].toString()}]`
          currentPhase.value = project.product_phrase
          t.value = designTask as DecodeDesignTask
        }
        if (route.query.modify) {
          dialogAction.value = 'update'
          return
        }
        dialogAction.value = 'detail'
      } else {
        // 搜索任务书
        searchProject('')
      }
      if (dialogAction.value === 'create') {
        watchEffect(() => {
          if (t.value.name && t.value.design_project) {
            t.value.design_project_id = designProjectOptions.value.find(p => p.label === t.value.name)?.value || 0
            t.value.design_project!.project_name = designProjectOptions.value.find(p => p.label === t.value.name)?.project_name || ''
          }
        })
      }
    })

    return {
      CircleCloseFilled,
      CirclePlusFilled,
      validateForm,
      formRules,
      dialogAction,
      headerText,
      paramLoading,
      taskStatusOptions,
      pullTypeOptions,
      seamTypeOptions,
      specRangeOptions,
      phaseOptions,
      productFrequencyOptions,
      libTempRiseOptions,
      productConnectionGroupOptions,
      wireTypeOptions,
      stackOptions,
      expLayParamOptions,
      expProParamOptions,
      t,
      currentRegulation,
      currentPhase,
      deviationStd,
      activeCollapseNames,
      handleCollapseChange,
      designProjectOptions,
      searchProjectLoading,
      searchProject,
      handleSelectProject,
      showRecommendTask,
      handleShowRecommend,
      projectQuery,
      handleConfirmCreate,
      bossSpecOptions,
      resinTypeOptions,
      coilWireOptions
    }
  }

})
</script>

<style lang="scss" scoped>

:deep() {
  .autofill,[autofill] {
    input {
      --el-input-background-color: #ecf5ff !important;
    }
  }
  [allow-create] {
    input {
      --el-input-background-color: #f7ffd7 !important;
    }
  }
  .product-model {
    label {
       border: 1px  ;
      line-height: 38px;
      width: 157px;
      text-align: center;
    }
    .el-input,.el-input-number {
      width: 100% !important;
      input {
        text-align: left !important;
        width: 100% !important;
      }
    }
  }
  label {
    width: 90px;
    text-align: center;
    font-size: 14px;
    margin: 0;
  }
  .is-in-pagination {
    .el-input {
      width: 50px !important;
      padding: 0 3px !important;
    }
  }
/*.el-input-number.is-without-controls .el-input__inner {
  width: 550px;
  margin-outside: 1px;
  !*line-height: 30px;
  height: 28px;*!
}*/
  .el-select {
    width: 100% !important;
  }
  .el-input-number {
    width: 100% !important;
    .el-input {
      width: 100% !important;
    }
  }
  .el-input{
    width: 100%;
    input {
      padding: 0 27px;
      text-align: center;
      color: #409eff;
    }
  }

  .data-unit {
    width: 35px;
    text-align: center;
    display: inline-block;
    font-size: 12px;
    font-weight: 700;
  }
  .value-to {
    width: 15px;
    line-height: 35px;
    text-align: center;
    font-size: 24px;
    font-weight: 700;
    margin-left: 22px;
  }
  .collapse-title{
    justify-content: center; /* 水平居中 */
  text-align: right;
    alignment: center;
 flex: 0 2 50%;
  }
  .multiple-select {
    .el-input {
      width: 100% !important;
    }
  }
  .design__create-container.update, .design__create-container.detail {
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
}

</style>
