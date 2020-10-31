<template>
  <div class="login">
    <van-form @submit="doLogin">
      <!-- identifier Field -->
      <van-field
        v-model="stuNum"
        name="stuNum"
        label="学号"
        placeholder="学号"
        :rules="[{ required: true, message: '请输入学号' }]"
      />
      <van-field
        v-model="password"
        :type="pass.type"
        name="password"
        label="密码"
        placeholder="身份证后六位"
        :rules="[{ required: true, message: '请输入密码' }]"
        :right-icon="pass.rightEyeIco"
        @click-right-icon="showOrHidePassword"
      />
      <!-- captcha Field -->
      <van-field
        v-model="captcha.captchaCode"
        type="text"
        name="captcha"
        center
        clearable
        label="验证码"
        placeholder="请输入验证码"
        :rules="[{ required: true, message: '请输入验证码' }]"
      >
        <van-image
          slot="button"
          width="120"
          height="40"
          :src="captcha.captchaImg"
          @click="refreshCaptcha"
        />
      </van-field>
      <div style="margin: 16px;">
        <van-button round block type="info" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import Vue from 'vue'
import { Form, Field, Button, Image, Toast } from 'vant'
import { getCaptcha } from '@/api/auth'

Vue.use(Form)
Vue.use(Field)
Vue.use(Button)
Vue.use(Image)

export default {
  name: 'Login',
  data() {
    return {
      captcha: {
        captchaId: '',
        captchaImg: '',
        captchaCode: ''
      },
      stuNum: '',
      password: '',
      pass: {
        credential: '',
        type: 'password',
        rightEyeIco: 'closed-eye'
      },
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  created() {
    this.refreshCaptcha()
  },
  methods: {
    // 刷新验证码
    refreshCaptcha() {
      getCaptcha().then(response => {
        const { data } = response
        this.captcha.captchaId = data.captchaId
        this.captcha.captchaImg = data.picPath
      })
    },
    doLogin(values) {
      Toast.loading({
        message: '登录中...',
        duration: 0,
        onOpened: () => {
          this.$store.dispatch('stu/login', {
            stuNum: this.stuNum,
            password: this.password,
            captchaId: this.captcha.captchaId,
            captcha: this.captcha.captchaCode
          }).then(() => {
            this.$router.push({ path: this.redirect || '/' })
            Toast.clear()
          }).catch(err => {
            console.log(err)
            this.refreshCaptcha()
            Toast.fail({
              message: '登录失败, 请重试',
              duration: 3 * 1000
            })
          })
        }
      })
    },
    showOrHidePassword() {
      if (this.pass.rightEyeIco === 'closed-eye') {
        this.pass.rightEyeIco = 'eye-o'
        this.pass.type = 'text'
      } else {
        this.pass.rightEyeIco = 'closed-eye'
        this.pass.type = 'password'
      }
    }
  }
}
</script>

<style scoped lang="scss">
.login {
  width: 100vw;
  height: 100vh;
  font-size: 14px;
  position: relative;
}
.van-form {
  position: absolute;
  width: 100vw;
  left: 50%;
  top: 50%;
  transform: translateX(-50%) translateY(-50%);
}
</style>
