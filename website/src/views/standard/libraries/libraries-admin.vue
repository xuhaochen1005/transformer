<template>
  <div class="main-container">
    <el-container style="height: 100%;  border: 1px solid #eee">
      <el-main>
        <aside class="page-header">
          明珠电气公司变压器设计管理系统标准库管理页面
        </aside>
        <el-tabs
          v-model="activeName"
          @tab-click="handleClick"
        >
          <el-tab-pane
            label="圆铜(铝)线规格"
            name="four_1"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="线圈结构初选"
            name="four_2"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="线圈结构"
            name="four_3"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="铁心性能表"
            name="four_4"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="铁心表"
            name="four_5"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="高压线圈参数"
            name="four_6"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="干变绕线外模台账"
            name="four_7"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="干变绕线内模台账"
            name="four_8"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="干变绕线POM"
            name="four_9"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="干变低压箔绕模具"
            name="four_10"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="端绝缘电压"
            name="four_11"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="低压线圈参数"
            name="four_12"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="产品型号表"
            name="four_13"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="干变低压外模"
            name="four_14"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
          <el-tab-pane
            label="扁铜(铝)线标尺寸"
            name="four"
          >
            <div class="main-manage">
              <div class="search-head">
                <el-form
                  :inline="true"
                  :model="formInline"
                  class="demo-form-inline"
                >
                  <el-form-item style="width:100px">
                    <el-input
                      v-model="formInline.user"
                      placeholder="类型"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V1"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V2"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V3"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V4"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-input
                      v-model="formInline.user"
                      placeholder="V5"
                    />
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-waves
                      class="filter-item"
                      type="primary"
                      icon="el-icon-search"
                      @click="onSubmit"
                    >
                      {{ $t('table.search') }}
                    </el-button>
                    <el-button
                      class="filter-item"
                      style="margin-left: 10px;"
                      type="primary"
                      icon="el-icon-edit"
                      @click="onCreate"
                    >
                      {{ $t('table.add') }}
                    </el-button>
                    <el-button
                      v-waves
                      :loading="downloadLoading"
                      class="filter-item"
                      type="primary"
                      icon="el-icon-download"
                      @click="onExport"
                    >
                      {{ $t('table.export') }}
                    </el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div class="main-table">
                <el-table
                  :key="tableKey"
                  v-loading="listLoading"
                  :data="list"
                  border
                  fit
                  highlight-current-row
                  style="width: 100%;"
                  @sort-change="sortChange"
                >
                  <el-table-column
                    :label="$t('table.id')"
                    prop="id"
                    sortable="custom"
                    align="center"
                    width="80"
                    :class-name="getSortClass('id')"
                  >
                    <template #default="{row}">
                      <span>{{ row.id }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.ptype')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v0')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('casbin_table.v1')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v2')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v3')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v4')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column
                    :label="$t('casbin_table.v5')"
                    width="180px"
                    align="center"
                  >
                    <template #default="{row}">
                      <span>{{ row.author }}</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('user_table.user_info')"
                    align="center"
                    width="95"
                  >
                    <template #default="{row}">
                      <span
                        v-if="row.pageviews"
                        class="link-type"
                        @click="handleGetPageviews(row.pageviews)"
                      >{{ row.pageviews }}</span>
                      <span v-else>0</span>
                    </template>
                  </el-table-column>

                  <el-table-column
                    :label="$t('table.actions')"
                    align="center"
                    width="230"
                    class-name="fixed-width"
                  >
                    <template #default="{row, $index}">
                      <el-button
                        type="primary"
                        size="mini"
                        @click="handleUpdate(row)"
                      >
                        {{ $t('table.edit') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='published'"
                        size="mini"
                        type="success"
                        @click="handleModifyStatus(row,'published')"
                      >
                        {{ $t('table.publish') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='draft'"
                        size="mini"
                        @click="handleModifyStatus(row,'draft')"
                      >
                        {{ $t('table.draft') }}
                      </el-button>
                      <el-button
                        v-if="row.status!=='deleted'"
                        size="mini"
                        type="danger"
                        @click="handleDelete(row, $index)"
                      >
                        {{ $t('table.delete') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="footer-pagination">
                <pagination
                  v-show="total>0"
                  v-model:page="listQuery.page"
                  v-model:limit="listQuery.limit"
                  :total="total"
                  @pagination="getList"
                />
              </div>
            </div>
          </el-tab-pane>
        </el-tabs>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import LineChart from '@/components/Charts/LineChart.vue'
import { createArticle, defaultArticleData, getArticles, getPageviews, updateArticle } from '@/api/articles'
import Pagination from '@/components/Pagination/index.vue'
import { IArticleData } from '@/api/types'
import { cloneDeep } from 'lodash'
import { Form } from 'element-ui'
import { formatJson } from '@/utils'
import { exportJson2Excel } from '@/utils/excel'
@Component({
  name: 'LineChartDemo',
  components: {
    LineChart,
    Pagination
  },
  data() {
    return {
      activeName: 'third',
      formInline: {
        user: '',
        region: ''
      },
      tableKey: 0,
      IArticleData: [],
      listLoading: true,
      downloadLoading: false,
      // listLoading:false,
      total: 30,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      }
    }
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event)
    },
    onSubmit() {
      console.log('submit!')
    },
    onRefresh() {
      console.log('refresh')
    },
    onCreate() {
      console.log('create')
    },
    onExport() {
      console.log('export')
    },
    getList() {
      /*   this.listLoading = true
      const { data } = await getArticles(this.listQuery)
      this.list = data.items
      this.total = data.total
      // Just to simulate the time of the request
      setTimeout(() => {
        this.listLoading = false
      }, 0.5 * 1000) */
    }
  }
})
export default class extends Vue {
  private tableKey = 0
  private list: IArticleData[] = []
  private total = 0
  private listLoading = true
  private listQuery = {
    page: 1,
    limit: 20,
    importance: undefined,
    title: undefined,
    type: undefined,
    sort: '+id'
  }

  private importanceOptions = [1, 2, 3]
  // private calendarTypeOptions = calendarTypeOptions
  private sortOptions = [
    { label: 'ID Ascending', key: '+id' },
    { label: 'ID Descending', key: '-id' }
  ]

  private statusOptions = ['published', 'draft', 'deleted']
  private showReviewer = false
  private dialogFormVisible = false
  private dialogStatus = ''
  private textMap = {
    update: 'Edit',
    create: 'Create'
  }

  private dialogPageviewsVisible = false
  private pageviewsData = []
  private rules = {
    type: [{ required: true, message: 'type is required', trigger: 'change' }],
    timestamp: [{ required: true, message: 'timestamp is required', trigger: 'change' }],
    title: [{ required: true, message: 'title is required', trigger: 'blur' }]
  }

  private downloadLoading = false
  private tempArticleData = defaultArticleData

  created() {
    this.getList()
  }

  private async getList() {
    this.listLoading = true
    const { data } = await getArticles(this.listQuery)
    this.list = data.items
    this.total = data.total
    // Just to simulate the time of the request
    setTimeout(() => {
      this.listLoading = false
    }, 0.5 * 1000)
  }

  private handleFilter() {
    this.listQuery.page = 1
    this.getList()
  }

  private handleModifyStatus(row: any, status: string) {
    this.$message({
      message: '操作成功',
      type: 'success'
    })
    row.status = status
  }

  private sortChange(data: any) {
    const { prop, order } = data
    if (prop === 'id') {
      this.sortByID(order)
    }
  }

  private sortByID(order: string) {
    if (order === 'ascending') {
      this.listQuery.sort = '+id'
    } else {
      this.listQuery.sort = '-id'
    }
    this.handleFilter()
  }

  private getSortClass(key: string) {
    const sort = this.listQuery.sort
    return sort === `+${key}` ? 'ascending' : 'descending'
  }

  private resetTempArticleData() {
    this.tempArticleData = cloneDeep(defaultArticleData)
  }

  private handleCreate() {
    this.resetTempArticleData()
    this.dialogStatus = 'create'
    this.dialogFormVisible = true
    this.$nextTick(() => {
      (this.$refs.dataForm as Form).clearValidate()
    })
  }

  private createData() {
    (this.$refs.dataForm as Form).validate(async(valid) => {
      if (valid) {
        const articleData = this.tempArticleData
        articleData.id = Math.round(Math.random() * 100) + 1024 // mock a id
        articleData.author = 'vue-typescript-admin'
        const { data } = await createArticle({ article: articleData })
        data.article.timestamp = Date.parse(data.article.timestamp)
        this.list.unshift(data.article)
        this.dialogFormVisible = false
        this.$notify({
          title: '成功',
          message: '创建成功',
          type: 'success',
          duration: 2000
        })
      }
    })
  }

  private handleUpdate(row: any) {
    this.tempArticleData = Object.assign({}, row)
    this.tempArticleData.timestamp = +new Date(this.tempArticleData.timestamp)
    this.dialogStatus = 'update'
    this.dialogFormVisible = true
    this.$nextTick(() => {
      (this.$refs.dataForm as Form).clearValidate()
    })
  }

  private updateData() {
    (this.$refs.dataForm as Form).validate(async(valid) => {
      if (valid) {
        const tempData = Object.assign({}, this.tempArticleData)
        tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
        const { data } = await updateArticle(tempData.id, { article: tempData })
        const index = this.list.findIndex(v => v.id === data.article.id)
        this.list.splice(index, 1, data.article)
        this.dialogFormVisible = false
        this.$notify({
          title: '成功',
          message: '更新成功',
          type: 'success',
          duration: 2000
        })
      }
    })
  }

  private handleDelete(row: any, index: number) {
    this.$notify({
      title: 'Success',
      message: 'Delete Successfully',
      type: 'success',
      duration: 2000
    })
    this.list.splice(index, 1)
  }

  private async handleGetPageviews(pageviews: string) {
    const { data } = await getPageviews({ pageviews })
    this.pageviewsData = data.pageviews
    this.dialogPageviewsVisible = true
  }

  private handleDownload() {
    this.downloadLoading = true
    const tHeader = ['timestamp', 'title', 'type', 'importance', 'status']
    const filterVal = ['timestamp', 'title', 'type', 'importance', 'status']
    const data = formatJson(filterVal, this.list)
    exportJson2Excel(tHeader, data, 'table-list')
    this.downloadLoading = false
  }
}
</script>

<style lang="scss" scoped>
.main-container {
  margin-left: 1%;
  margin-right: 1%;
  /** width: 95%; */
  height: 100%;
  margin-top: 30px;

}
</style>
