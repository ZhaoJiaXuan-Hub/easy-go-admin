<template>
  <div class="login">
    <a-row align="stretch" class="login-box">
      <a-col :xs="24" :sm="24" :md="24">
        <div class="login-right">
          <a-form ref="formRef" :size="isMobile() ? 'large' : 'medium'" :model="form" :rules="rules"
            :style="{ width: '74%' }" :label-col-style="{ display: 'none' }" :wrapper-col-style="{ flex: 1 }">
            <h3 class="login-right__title">
              <span>用户登录</span>
            </h3>
            <div style="margin-bottom: 1.5rem">
              <a-radio-group type="button" v-model="checkedLoginType">
                <a-radio value="pass">密码登录</a-radio>
                <a-radio value="qrcode" disabled>扫码登录</a-radio>
              </a-radio-group>
            </div>
            <a-form-item field="username">
              <a-input v-model="form.username" placeholder="账号 admin/user" allow-clear>
                <template #prefix><icon-user :stroke-width="1" :style="{ fontSize: '20px' }" /></template>
              </a-input>
            </a-form-item>
            <a-form-item field="password">
              <a-input-password v-model="form.password" placeholder="密码" allow-clear>
                <template #prefix><icon-lock :stroke-width="1" :style="{ fontSize: '20px' }" /></template>
              </a-input-password>
            </a-form-item>
            <a-form-item>
              <a-row justify="space-between" align="center" class="w-full">
                <a-checkbox v-model="checked">记住密码</a-checkbox>
              </a-row>
            </a-form-item>
            <a-form-item>
              <a-space direction="vertical" fill class="w-full">
                <a-button type="primary" size="large" long :loading="loading" @click="login">登录</a-button>
              </a-space>
            </a-form-item>
          </a-form>
        </div>
      </a-col>
    </a-row>

    <GiThemeBtn class="theme-btn"></GiThemeBtn>

    <LoginBg></LoginBg>
  </div>
</template>

<script setup lang="ts">
import { type FormInstance, Message } from '@arco-design/web-vue'
import LoginBg from './components/LoginBg/index.vue'
import { useUserStore } from '@/stores'
import { useLoading } from '@/hooks'
import * as Regexp from '@/utils/regexp'
import { isMobile } from '@/utils'

defineOptions({ name: 'Login' })
const router = useRouter()
const userStore = useUserStore()

const form = reactive({
  username: 'admin',
  password: '123456'
})

const rules: FormInstance['rules'] = {
  username: [{ required: true, message: '请输入账号' }],
  password: [
    { required: true, message: '请输入密码' },
    { match: Regexp.Password, message: '输入密码格式不正确' }
  ]
}

// 登录方式
const checkedLoginType = ref("pass")
// 记住密码
const checked = ref(true)
// 登录加载
const { loading, setLoading } = useLoading()
const errorMessage = ref('')

const formRef = ref<FormInstance>()
// 点击登录
const login = async () => {
  try {
    const valid = await formRef.value?.validate()
    if (valid) return
    setLoading(true)
    await userStore.login(form)
    const { redirect, ...othersQuery } = router.currentRoute.value.query
    router.push({
      path: (redirect as string) || '/',
      query: {
        ...othersQuery
      }
    })
    Message.success('登录成功')
  } catch (error) {
    errorMessage.value = (error as Error).message
  } finally {
    setLoading(false)
  }
}
</script>

<style lang="scss" scoped>
.login {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--color-bg-5);
  background-image: url(src/assets/images/login-bg-BkIjQ0FB.png);
  background-repeat: no-repeat;
  background-size: 100% 100%;

  &-box {
    width: 86%;
    max-width: 320px;
    height: 380px;
    display: flex;
    z-index: 999;
    box-shadow: 0 0 1px 0 rgba(0, 0, 0, 0.08);
    border-radius: 10px;
  }
}

.login-right {
  width: 100%;
  height: 100%;
  background: var(--color-bg-1);
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
  border-radius: 10px;

  &__title {
    color: var(--color-text-1);
    font-size: 24px;
    line-height: 32px;
    margin-bottom: 20px;
    font-weight: bold;
    display: flex;
    justify-content: flex-start;
    align-items: center;

    .logo {
      width: 32px;
      height: 32px;
      margin-right: 6px;
    }
  }
}

.theme-btn {
  position: fixed;
  top: 20px;
  left: 30px;
  z-index: 9999;
}
</style>
