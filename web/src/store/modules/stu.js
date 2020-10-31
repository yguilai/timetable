import { getToken, setToken, removeToken } from '@/utils/auth'
import { login, getStuInfo } from '@/api/auth'
import { Toast } from 'vant'

const getDefaultState = () => {
  return {
    token: getToken(),
    stuNum: '',
    stuId: 0
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: state => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_STUNUM: (state, stuNum) => {
    state.stuNum = stuNum
  },
  SET_STUID: (state, stuId) => {
    state.stuId = stuId
  }
}

const actions = {
  login({ commit }, stuInfo) {
    const { stuNum, password, captchaId, captcha } = stuInfo
    return new Promise((resolve, reject) => {
      login({
        stuNum: stuNum.trim().toUpperCase(),
        password,
        captchaId,
        captcha
      })
        .then(response => {
          const { data } = response
          commit('SET_TOKEN', data.token)
          commit('SET_STUNUM', stuNum)
          commit('SET_STUID', data.stu.id)
          setToken(data.token)
          resolve()
        })
        .catch(err => {
          Toast.fail(err.message)
          reject(err)
        })
    })
  },
  getStuInfo({ commit }) {
    return new Promise((resolve, reject) => {
      getStuInfo()
        .then(response => {
          const { data } = response
          commit('SET_STUID', data.id)
          commit('SET_STUNUM', data.stuNum)
          resolve()
        })
        .catch(err => {
          Toast.fail(err.message)
          reject(err)
        })
    })
  },
  // 用户登出
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      commit('tt/RESET_STATE', null, { root: true })
      resolve()
    })
  },
  // 移除TOKEN
  resetToken({ commit }) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
