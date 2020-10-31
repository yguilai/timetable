import Vue from 'vue'
import Vuex from 'vuex'
import getters from '@/store/getters'
import stu from './modules/stu'
import tt from './modules/timetable'
import version from './modules/version'

Vue.use(Vuex)

export default new Vuex.Store({
  getters,
  modules: {
    stu,
    tt,
    version
  }
})
