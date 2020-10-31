import axios from 'axios'
import { Toast, Dialog } from 'vant'
import store from '@/store'
import { getToken } from '@/utils/auth'

const service = axios.create({
  headers: {
    'Content-Type': 'application/json'
  },
  baseURL: 'http://47.106.188.222:1999/',
  timeout: 12000
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['x-token'] = getToken()
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code !== undefined && res.code !== 200) {
      if (res.code === 401 || res.code === 403) {
        // to re-login
        Dialog.alert({
          title: 'Confirm logout',
          message: 'You have been logged out, you can cancel to stay on this page, or log in again',
          theme: 'round-button'
        }).then(() => {
          store.dispatch('user/resetToken').then(() => {
            location.reload()
          })
        })
      }
      return Promise.reject(new Error(res.msg || 'Error'))
    } else {
      return res
    }
  },
  error => {
    Toast.fail({
      message: error.message,
      duration: 4 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
