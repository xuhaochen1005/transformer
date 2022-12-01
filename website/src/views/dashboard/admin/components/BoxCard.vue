<!--
 * @Description: 卡片盒子，首页显示技术百分比
 * @Author: ZY
 * @Date: 2021-01-16 17:12:32
 * @LastEditors: ZY
 * @LastEditTime: 2021-01-18 09:36:19
-->

<template>
  <el-card
    class="box-card-component"
    style="margin-left:8px;height: 100%;"
  >
    <div
      class="box-card-header"
    >
      <img src="http://www.pearlelectric.com/uploads/allimg/180601/1-1P6011559420-L.jpg">
    </div>
    <div style="position:relative;">
      <Mallki
        style="margin-top: 10px;"
        class="mallki-text"
        text="产品设计进度汇总"
      />
      <el-empty
        v-if="designProjectList.length === 0"
        :image-size="100"
      />
      <div v-else>
        <div
          v-for="item in designProjectList"
          :key="item"
          style="padding-top:35px;"
          class="progress-item"
        >
          <span>{{ item.product_model }}</span>
          <el-progress
            :percentage="Math.floor((item.project_status / 10) * 100) "
            :status="item.project_status === DesignProjectStatus.DesignProjectChecked ? 'success' :
              [DesignProjectStatus.DesignProjectApproveUnaccepted,
               DesignProjectStatus.DesignProjectReviewUnaccepted,
               DesignProjectStatus.DesignProjectCheckUnAccepted
              ].includes(item.project_status) ? 'warning': ''"
          />
        </div>
      </div>
    </div>
  </el-card>
</template>

<script lang="ts">
import { useStore } from '@/store'
import { computed, defineComponent } from 'vue'
import PanThumb from '@/components/pan-thumb/index.vue'
import Mallki from '@/components/text-hover-effect/Mallki.vue'
import { roleList } from '@/utils/permission'
import { DesignProjectStatus } from '@/model/design'

export default defineComponent({
  components: {
    PanThumb,
    Mallki
  },
  props: {
    designProjectList: Array
  },
  setup() {
    const store = useStore()
    const name = computed(() => {
      return store.state.user.user_name_zh
    })

    const avatar = computed(() => {
      return ''
    })
    const roles = computed(() => {
      return [...roleList!.keys()]
    })

    return {
      DesignProjectStatus,
      name,
      avatar,
      roles
    }
  }
})
</script>

<style lang="scss">
.box-card-component {
  .el-card__header {
    padding: 0px!important;
  }
}

.panThumb {
  z-index: 100;
  height: 70px!important;
  width: 70px!important;
  position: absolute!important;
  top: -45px;
  left: 0px;
  border: 5px solid #ffffff;
  background-color: #fff;
  margin: auto;
  box-shadow: none!important;

  .pan-info {
    box-shadow: none!important;
  }
}
</style>

<style lang="scss" scoped>
.box-card-component {
  .box-card-header {
    position: relative;
    height: 220px;

    img {
      width: 100%;
      height: 100%;
      transition: all 0.2s linear;

      &:hover {
        transform: scale(1.1, 1.1);
        filter: contrast(130%);
      }
    }
  }

  .mallki-text {
    position: absolute;
    top: 0px;
    right: 0px;
    font-size: 20px;
    font-weight: bold;
  }

  .progress-item {
    margin-bottom: 10px;
    font-size: 14px;
  }

  @media only screen and (max-width: 1510px){
    .mallki-text{
      display: none;
    }
  }
}
</style>
