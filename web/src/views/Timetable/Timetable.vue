<template>
  <div class="timetable">
    <van-nav-bar
      :title="timetable.weekly === 0 ? '总课表' : `第${timetable.weekly}周`"
      @click-right="onNavBarClickRight"
    >
      <template #right>
        <van-icon name="calendar-o" color="#000000" />
      </template>
    </van-nav-bar>

    <van-popup v-model="showSelectWeek" round position="bottom">
      <van-picker
        title="请选择"
        show-toolbar
        :default-index="timetable.weekly"
        :columns="columns"
        @confirm="onConfirm"
      />
    </van-popup>
    <van-grid :column-num="8">
      <van-grid-item v-for="(v,i) in weekdays" :key="'w_' + i" :text="v" />
    </van-grid>
    <van-grid v-for="(row, i) in courses" id="courses-body" :key="'r_' + i" center clickable :column-num="8">
      <van-grid-item
        v-for="(col, j) in row"
        :key="`c_${i}${j}`"
        :class="['course-item', col.length !== 0 ? 'has-course' : '']"
        @click="onClickCourseItem(i,j)"
      >
        <template v-if="j === 0" #default>
          <div v-html="filterDayTime(daytime[i])" />
        </template>
        <template v-else-if="col.length !== 0" #default>
          <div class="course-item-body" :style="`background: ${randomBgColor()}`">
            <div>{{ filterCourseName(col[0].name) }}</div>
            <div>{{ col[0].location }}</div>
          </div>
        </template>
      </van-grid-item>
    </van-grid>
    <p class="remarks">备注:{{ timetable.remarks }}</p>
    <p class="padding-bottom" />

    <van-popup v-model="showCourseItem" class="course-detail" style="height: 88.93333vw;width: 80vw" round>
      <van-tabs v-model="active" swipeable>
        <van-tab v-for="(item, i) in tmpCourses" :key="'details_' + i" :title="item.name">
          <van-cell-group>
            <van-cell title="课程" :value="item.name" />
            <van-cell title="教室" :value="item.location" />
            <van-cell title="时间" :value="`第${item.time}大节(${daytime[item.time-1]})`" />
            <van-cell title="班级" :value="item.class" />
            <van-cell title="教师" :value="item.teacher" />
            <van-cell title="周次" :value="item.weeks" />
          </van-cell-group>
        </van-tab>
      </van-tabs>
    </van-popup>
  </div>
</template>

<script>
import Vue from 'vue'
import { NavBar, Icon, Picker, Popup, Grid, GridItem, CellGroup, Cell, Tab, Tabs } from 'vant'

Vue.use(NavBar)
Vue.use(Icon)
Vue.use(Popup)
Vue.use(Picker)
Vue.use(Grid)
Vue.use(GridItem)
Vue.use(CellGroup)
Vue.use(Cell)
Vue.use(Tabs)
Vue.use(Tab)
export default {
  name: 'Timetable',
  data() {
    return {
      showSelectWeek: false,
      columns: [],
      times: ['一', '二', '三', '四', '五'],
      weekdays: ['', '周一', '周二', '周三', '周四', '周五', '周六', '周日'],
      daytime: ['8:00-9:40', '10:00-11:40', '14:00-15:40', '16:00-17:40', '19:00-20:40'],
      bgColors: ['#ee3f4d', '#806d9e', '#10aec2', '#1ba784',
        '#fbda41', '#fa5d19', '#2bae85', '#5bae23',
        '#ad6598', '#d2357d', '#ed2f6a', '#f07c82',
        '#2177b8', '#93b5cf', '#2f90b9', '#5cb3cc'],
      showCourseItem: false,
      tmpCourses: [{
        name: '',
        class: '',
        time: '',
        teacher: '',
        location: '',
        weeks: ''
      }],
      active: 0
    }
  },
  computed: {
    timetable() {
      return this.$store.getters.timetable
    },
    courses() {
      return this.$store.getters.courses
    }
  },
  created() {
    for (let i = 0, l = 25; i < l; i++) {
      if (i === 0) {
        this.columns.push('总课表')
        continue
      }
      this.columns.push(`第${i}周`)
    }
    if (!this.timetable.courses) {
      this.fetchTimetable()
    }
  },
  methods: {
    onNavBarClickRight() {
      this.showSelectWeek = !this.showSelectWeek
    },
    onConfirm(val, index) {
      this.showSelectWeek = !this.showSelectWeek
      if (this.timetable.weekly !== index) {
        this.$store.commit('tt/SET_WEEKLY', index)
        this.fetchTimetable()
      }
    },
    fetchTimetable() {
      this.$store.dispatch('tt/fetchTimetable', {
        yearly: this.timetable.yearly,
        weekly: this.timetable.weekly,
        stuId: this.$store.getters.stuId
      })
    },
    onClickCourseItem(i, j) {
      const c = this.courses[i][j]
      if (c.length !== 0) {
        this.showCourseItem = !this.showCourseItem
        this.tmpCourses = c
      }
    },
    filterCourseName(name) {
      if (name.length > 12) {
        return name.substring(0, 10).trim() + '...'
      }
      return name
    },
    randomBgColor() {
      return this.bgColors[this.randomNum(0, this.bgColors.length - 1)]
    },
    randomNum(minNum, maxNum) {
      switch (arguments.length) {
        case 1:
          return parseInt(Math.random() * minNum + 1, 10)
        case 2:
          return parseInt(Math.random() * (maxNum - minNum + 1) + minNum, 10)
        default:
          return 0
      }
    },
    filterDayTime(time) {
      const t = time.split('-')
      return `<div>${t[0]}</div><div>|</div><div>${t[1]}</div>`
    }
  }
}
</script>

<style scoped lang="scss">
  .timetable {
    width: 100vw;
    height: 100vh;
    overflow-y: auto;
  }

  #courses-body {
    font-size: 12px;

    & .course-item {
      width: 46.88px;
      height: 97.5px;
      text-align: center;
      overflow: hidden;
      white-space: normal;
      text-overflow: ellipsis;
    }

    & div.van-grid-item__content {
      padding: 0;
    }
  }

  /deep/ .has-course .van-grid-item__content {
    font-size: 12px;
    padding: 0;
    color: #fff;
    & .course-item-body {
      width: 100%;
      height: 100%;
      padding: 5px 0;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }
  }

  .course-detail /deep/ .van-cell__value {
    -webkit-box-flex: 3;
    -webkit-flex: 3;
    flex: 3;
  }

  .remarks {
    font-size: 12px;
    text-align: center;
  }

  .padding-bottom {
    height: 50px;
  }

  .van-hairline--top::after {
    border-top-width: 0;
  }

  /*
  .course-detail .van-cell-group {
    width: 100%;
    left: 50%;
    top: 50%;
    transform: translateX(-50%) translateY(-50%);
  }*/
</style>
