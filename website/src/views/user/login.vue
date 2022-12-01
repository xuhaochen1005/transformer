<template>
  <div
    class="login-container"
    @keyup.enter="handleLogin"
  >
    <el-form
      ref="loginFormRef"
      :model="loginForm"
      :rules="loginRules"
      class="login-form"
      autocomplete="on"
      label-position="left"
    >
      <el-row class="login-header">
        <img
          src="../../assets/images/user/logo.png"
          class="login-logo"
        >
        <p class="divider" />
        <p class="login-title">
          明珠变压器设计平台
        </p>
      </el-row>
      <el-form-item
        prop="username"
        label="用户名"
        class="el-icon-user"
      >
        <el-input
          ref="userNameRef"
          v-model="loginForm.username"
          placeholder="请输入用户名"
          name="username"
          type="text"
          tabindex="1"
        />
      </el-form-item>
      <el-tooltip
        v-model="capsTooltip"
        content="Caps lock is On"
        placement="right"
        manual
      >
        <el-form-item
          prop="password"
          label="密码"
          class="el-icon-lock"
        >
          <el-input
            :key="passwordType"
            ref="passwordRef"
            v-model="loginForm.password"
            :type="passwordType"
            placeholder="请输入密码"
            name="password"
            tabindex="1"
            autocomplete="on"
            @keyup="checkCapslock"
            @blur="capsTooltip = false"
          />
          <span
            class="show-pwd"
            @click="showPwd"
          >
            <svg-icon
              :name="passwordType === 'password' ? 'eye-off' : 'eye-on'"
              :size="20"
              color="#409EFF"
            />
          </span>
        </el-form-item>
      </el-tooltip>
      <el-form-item
        prop="verifyCode"
        label="验证码"
        class="el-icon-lock erifyCodeItemCss"
        style="justify-content: space-between;"
      >
        <el-input
          ref="verifyCodeRef"
          v-model="loginForm.verifyCode"
          placeholder="请输入验证码"
          name="verifyCode"
          type="text"
          tabindex="2"
        />
        <p
          class="v_container"
          style="width:90px;height:35px;font-size:30px;color:white;background-color:gray;text-align:center;"
          @click="switchCode"
        >
          {{ verifyCode }}
        </p>
      </el-form-item>
      <el-button
        :loading="loading"
        type="primary"
        @click.prevent="handleLogin"
      >
        登录
      </el-button>
    </el-form>
  </div>
</template>

<script lang="ts">
import { defineComponent, nextTick, onMounted, reactive, ref, toRefs, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { isValidUsername } from '@/utils/validate'
import { LocationQuery, useRoute, useRouter } from 'vue-router'
import { useStore } from '@/store'
import { UserAction } from '@/store/user/action'
import { verifyCode } from '@/utils/verifyCode'

export default defineComponent({
  setup() {
    const userNameRef = ref(null)
    const passwordRef = ref(null)
    const loginFormRef = ref(null)
    const router = useRouter()
    const route = useRoute()
    const store = useStore()
    const state = reactive({
      loginForm: {
        username: '',
        password: '',
        verifyCode: ''
      },
      loginRules: {
        username: [{ validator: userNameRef, trigger: 'blur' }],
        password: [{ validator: passwordRef, trigger: 'blur' }]
      },
      passwordType: 'password',
      loading: false,
      showDialog: false,
      capsTooltip: false,
      redirect: '',
      otherQuery: {},
      verifyCode: verifyCode()
    })
    const methods = reactive({
      validateUsername: (rule: any, value: string, callback: Function) => {
        if (!isValidUsername(value)) {
          callback(new Error('Please enter the correct user name'))
        } else {
          callback()
        }
      },

      validatePassword: (rule: any, value: string, callback: Function) => {
        if (value.length < 6) {
          callback(new Error('The password can not be less than 6 digits'))
        } else {
          callback()
        }
      },

      checkCapslock: (e: KeyboardEvent) => {
        const { key } = e
        state.capsTooltip =
          key !== null && key.length === 1 && key >= 'A' && key <= 'Z'
      },

      showPwd: () => {
        if (state.passwordType === 'password') {
          state.passwordType = ''
        } else {
          state.passwordType = 'password'
        }
        nextTick(() => {
          (passwordRef.value as any).focus()
        })
      },
      handleLogin: () => {
        (loginFormRef.value as any).validate(async (valid: boolean) => {
          if (valid) {
            if (state.loginForm.verifyCode.toLocaleLowerCase() !== state.verifyCode.toLowerCase()) {
              methods.switchCode()
              ElMessage.error('验证码输入错误，请重试')
              return
            }
            state.loading = true
            await store.dispatch(UserAction.LOGIN, state.loginForm)
            router
              .push({
                path: state.redirect || '/',
                query: state.otherQuery
              })
              .catch(err => {
                ElMessage.error(err)
              })
            // Just to simulate the time of the request
            setTimeout(() => {
              state.loading = false
            }, 0.5 * 1000)
          } else {
            return false
          }
        })
      },
      switchCode() {
        state.verifyCode = verifyCode()
      }
    })

    function getOtherQuery(query: LocationQuery) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {} as LocationQuery)
    }

    watch(
      () => route.query,
      query => {
        if (query) {
          state.redirect = query.redirect?.toString() ?? ''
          state.otherQuery = getOtherQuery(query)
        }
      }
    )

    onMounted(() => {
      if (state.loginForm.username === '') {
        (userNameRef.value as any).focus()
      } else if (state.loginForm.password === '') {
        (passwordRef.value as any).focus()
      }
    })

    return {
      userNameRef,
      passwordRef,
      loginFormRef,
      ...toRefs(state),
      ...toRefs(methods)
    }
  }
})
</script>

<style lang="scss">
.el-form--label-left .el-form-item__label {
  padding-left: 21px;
  font-weight: normal;
  font-size: 17px;
  height: 30px;
  color: black;
}

.el-form-item__content {
  border-bottom: 1px solid $loginbordercolor !important;
}

.el-icon-lock:before {
  position: absolute;
  top: 12px;
  color: $loginfontcolor;
}
.el-icon-user:before {
  position: absolute;
  top: 12px;
  color: $loginfontcolor;
  width: 36px;
}

.login-container {
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      height: 47px;
      border: 0;
      border-radius: 0;
      padding: 6px 5px 6px 0px;
      -webkit-appearance: none;
      font-size: 18px;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $loginBg inset !important;
        -webkit-text-fill-color: #fff !important;
      }
    }
  }

  .el-form-item {
    background-color: #fff !important;
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 5px;
    color: #454545;
    flex-direction: column;
    position: relative;
    width: 100%;
    padding: 0 11%;
  }
  .erifyCodeItemCss .el-form-item__content {
    display: flex;
  }
}
</style>

<style lang="scss" scoped>
.login-container {
  height: 100%;
  min-height: 600px;
  width: 100%;
  overflow: hidden;
  background-image: url(../../assets/images/user/cover.png);
  background-repeat: no-repeat;
  background-size: auto;
  display: flex;
  align-items: center;
  justify-content: flex-end;

  .login-form {
    min-width: 600px;
    height: 1400px;
    overflow: hidden;
    border-top-left-radius: 1200px 1400px;
    border-bottom-left-radius: 1200px 1400px;
    background-color: white;
    display: flex;
    align-items: center;
    flex-direction: column;
    justify-content: center;
    padding: 0 25px 0 75px;

    .login-header {
      width: 100%;
      height: 50px;
      margin-bottom: 75px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      .divider {
        height: 100%;
        width: 1px;
        background-color: black;
      }
      .login-logo {
        height: 100%;
      }
      .login-title {
        line-height: 50px;
        color: $--color-primary;
        font-size: 26px;
        letter-spacing: 5px;
        text-align: right;
        font-weight: 800;
      }
    }

    .el-button {
      width: 79%;
      border-radius: 25px;
      height: 49px;
      font-size: 1.4rem;
      letter-spacing: 7px;
      font-weight: 500;
    }
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $darkGray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $darkGray;
    cursor: pointer;
    user-select: none;
  }

  .v_container {
    width: 60px;
    height: 25px;
    display: block;
    background-color: #c4c4c4;
    color: lightgoldenrodyellow;
    line-height: 25px;
    text-align: center;
    text-decoration: overline;
    font-style: italic;
    font-size: medium;
    cursor: pointer;
    margin: 17px 0 0 0;
  }

  @media only screen and (max-width: 470px) {
    .thirdparty-button {
      display: none;
    }
  }
}
</style>
