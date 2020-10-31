<template>
  <div class="about">
    <van-nav-bar title="关于" left-arrow :border="false" @click-left="$router.back()">
      <template #left>
        <van-icon name="arrow-left" color="#000000" size="20" />
      </template>
    </van-nav-bar>

    <van-cell-group>
      <van-cell title="作者" :value="ver.author" />
      <van-cell title="反馈邮箱" :value="ver.email" />
      <van-cell title="版本" :value="ver.version" is-link @click="showVersionDetails = !showVersionDetails" />
    </van-cell-group>

    <van-popup v-model="showVersionDetails" round style="height: 70vw;width: 70vw">
      <p class="version-title">更新内容:</p>
      <p v-for="(item, i) in ver.versionInfo" :key="'info_' + i" v-text="item" />
    </van-popup>
  </div>
</template>

<script>
import Vue from 'vue'
import { NavBar, Icon, Cell, CellGroup, Popup } from 'vant'
Vue.use(NavBar)
Vue.use(Icon)
Vue.use(CellGroup)
Vue.use(Cell)
Vue.use(Popup)

export default {
  name: 'About',
  data() {
    return {
      showVersionDetails: false
    }
  },
  computed: {
    ver() {
      return this.$store.getters.ver
    }
  },
  created() {
    if (!this.ver.name) {
      this.$store.dispatch('version/getVersion')
    }
  },
  methods: {}
}
</script>

<style scoped lang="scss">
.about {
  width: 100vw;
  height: 100vh;
}

.van-popup {
  padding: 10px;
  font-size: 14px;
}

.version-title {
  text-align: center;
}
</style>
