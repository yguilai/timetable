<template>
  <div class="index">
    <transition name="fade">
      <router-view />
    </transition>
    <van-tabbar v-show="true" v-model="active" active-color="#0981F9">
      <van-tabbar-item name="Timetable" icon="calendar-o" to="/timetable">
        课表
      </van-tabbar-item>
      <van-tabbar-item name="Profile" icon="user-o" to="/profile">
        我的
      </van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
import Vue from 'vue'
import { Tabbar, TabbarItem, Notify } from 'vant'
// import { Tabbar, TabbarItem } from 'vant'

Vue.use(Tabbar)
Vue.use(TabbarItem)

export default {
  name: 'Index',
  data() {
    return {
      active: 'Timetable'
    }
  },
  created() {
    Notify({
      type: 'danger',
      message: '11月4号, 服务器过期, 服务停止(以后可能会恢复)'
    })
  },
  mounted() {
    if (this.$router.currentRoute.name !== this.active) {
      this.active = this.$router.currentRoute.name
    }
  },
  methods: {}
}
</script>

<style scoped lang="scss">
  .index {
    width: 100vw;
    height: 100vh;
  }

  .fade-enter-active, .fade-leave-active {
    transition: opacity .3s
  }

  .fade-leave-to {
    display: none
  }

  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */
  {
    opacity: 0
  }
</style>

<style lang="scss">
  .van-icon {
    font-size: 20px;
  }
</style>
