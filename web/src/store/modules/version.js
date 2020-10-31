import { getVersion } from '@/api/version'

const getDefaultVersionState = () => {
  return {
    ver: {
      author: '燕归来',
      email: 'yguilai@88.com',
      version: '1.1.0',
      versionInfo: []
    }
  }
}

const state = getDefaultVersionState()

const mutations = {
  SET_VERSION: (state, ver) => {
    state.ver = ver
  }
}

const actions = {
  getVersion({ commit }) {
    return new Promise((resolve, reject) => {
      getVersion()
        .then(response => {
          console.log(response)
          commit('SET_VERSION', response)
        })
        .catch(err => reject(err))
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
