import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import 'normalize.css/normalize.css'

import { Lazyload, Toast } from 'vant'
Vue.use(Lazyload)
Toast.setDefaultOptions('loading', { forbidClick: true })

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
