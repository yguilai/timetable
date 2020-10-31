<template>
  <div class="profile">
    <van-nav-bar title="我的" :border="false" @click-right="$router.push('/about')">
      <template #right>
        <van-icon name="comment-o" color="#000000" />
      </template>
    </van-nav-bar>

    <van-cell-group>
      <van-cell title="当前账号" :value="stuNum" />
      <van-cell title="当前学期" clickable :value="yearly" />
      <van-cell title="开学日期" clickable :value="semesterStartTime" label="默认: 下: 09/09(我跟我媳妇儿的恋爱纪念日), 上: 02/08" @click="showCalendar = !showCalendar" />
      <van-cell title="当前周次" clickable :value="weekly" />
      <van-cell title="版本Q群" value="324889934" />
      <van-cell
        title="⚠注意"
        is-link
        @click="showCareful = !showCareful"
      />
    </van-cell-group>

    <van-calendar v-model="showCalendar" title="选择开学日期" :default-date="new Date(semesterStartTime)" :min-date="minDate" :max-date="maxDate" @confirm="onCalendarConfirm" />

    <div style="margin: 16px;">
      <van-button round block type="danger" native-type="button" @click="logout">
        退出登录
      </van-button>
    </div>

    <van-popup v-model="showCareful" class="careful" round style="height: 85vw;width: 70vw">
      <p class="careful-title">⚠注意</p>
      <p class="careful-content">1. 密码为身份证后六位, 仅用于登入应用和登录教务系统获取课表;</p>
      <p class="careful-content">
        2. 由于校内暂未开放外网, 课表数据需要通过反向代理隧道获取.
        因此在学校断网期间, <strong>首次登录</strong>的用户无法获取到课表, 已在白天登录并获取过课表数据的用户不受影响;
      </p>
      <p>3. 作者即将去实习, 希望有同学能够接手这个应用并维护下去. 有兴趣的同学可以联系我, 邮箱: yguilai@88.com</p>
    </van-popup>
  </div>
</template>

<script>
import Vue from 'vue'
import { Cell, CellGroup, NavBar, Icon, Dialog, Button, Popup, Calendar } from 'vant'
import dayjs from 'dayjs'
import { getCurrentWeekly } from '@/utils/weekly'

Vue.use(CellGroup)
Vue.use(Cell)
Vue.use(NavBar)
Vue.use(Icon)
Vue.use(Button)
Vue.use(Popup)
Vue.use(Calendar)

export default {
  name: 'Profile',
  data() {
    const now = new Date()
    return {
      showCareful: false,
      showCalendar: false,
      minDate: new Date(now.getFullYear(), 0, 1),
      maxDate: new Date(now.getFullYear(), 11, 30)
    }
  },
  computed: {
    stuNum() {
      return this.$store.getters.stuNum
    },
    yearly() {
      return this.$store.getters.timetable.yearly
    },
    weekly() {
      return this.$store.getters.timetable.weekly
    },
    semesterStartTime() {
      return this.$store.getters.semesterStartTime
    }
  },
  methods: {
    logout() {
      Dialog.alert({
        message: '确认退出登录',
        theme: 'round-button',
        showConfirmButton: true,
        showCancelButton: true,
        confirmButtonColor: '#ccc',
        cancelButtonColor: ''
      }).then(() => {
        this.$store.dispatch('stu/logout')
          .then(() => {
            this.$router.push('/login')
          })
      }).catch(() => {})
    },
    onCalendarConfirm(d) {
      this.showCalendar = !this.showCalendar
      this.$store.commit('tt/SET_SEMESTER_START_TIME', dayjs(d.toLocaleString()))
      this.$store.commit('tt/SET_WEEKLY', getCurrentWeekly(dayjs(d.toLocaleString())))
      this.$store.dispatch('tt/fetchTimetable', {
        stuId: this.$store.getters.stuId,
        yearly: this.yearly,
        weekly: this.weekly
      })
    }
  }
}
</script>

<style scoped lang="scss">
.profile {
  width: 100vw;
  height: 100vh;
}

.careful {
  padding: 10px;

  & .careful-title {
    text-align: center;
    color: #ee0a24;
  }
}
</style>
