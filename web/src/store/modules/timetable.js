import { getTimetable } from '@/api/timetable'
import { getWeekly, setWeekly, getCurrentWeekly } from '@/utils/weekly'
import { getSemesterStartTime, setSemesterStartTime } from '@/utils/semester-start-time'
import { Toast } from 'vant'

const getDefaultYearly = () => {
  const now = new Date()
  const y = now.getFullYear()
  const m = now.getMonth() + 1
  if (m > 6) {
    return `${y}-${y + 1}-1`
  }
  return `${y - 1}-${y}-2`
}

const new3DArray = (i, j) => {
  const arr = new Array(i)
  for (let k = 0; k < i; k++) {
    arr[k] = new Array(j)
    for (let m = 0; m < j; m++) {
      arr[k][m] = []
    }
  }
  return arr
}

const getDefaultState = () => {
  const start = getSemesterStartTime()
  if (!start) {
    setSemesterStartTime()
  }
  setWeekly(getCurrentWeekly(getSemesterStartTime()))
  return {
    timetable: {
      yearly: getDefaultYearly(),
      weekly: getWeekly(),
      courses: undefined
    },
    semesterStartTime: getSemesterStartTime().format('YYYY/MM/DD'),
    courses: new3DArray(5, 8)
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: state => {
    Object.assign(state, getDefaultState())
  },
  SET_TIMETABLE: (state, tt) => {
    state.timetable = tt
  },
  SET_WEEKLY: (state, weekly) => {
    state.timetable.weekly = weekly
    setWeekly(weekly)
  },
  SET_COURSE_ITEM: (state, { i, j, item }) => {
    state.courses[i][j].push(item)
  },
  SET_SEMESTER_START_TIME: (state, time) => {
    state.semesterStartTime = time.format('YYYY/MM/DD')
    setSemesterStartTime(time.format('YYYY/MM/DD'))
  },
  RESET_COURSES: state => {
    state.courses = new3DArray(5, 8)
  }
}

const actions = {
  fetchTimetable({ commit }, t) {
    return new Promise((resolve, reject) => {
      Toast.loading({
        message: 'Loading',
        duration: 0
      })

      getTimetable(t)
        .then(response => {
          const { data } = response
          data.weekly = t.weekly
          commit('SET_TIMETABLE', data)
          const c = data.courses
          if (c) {
            commit('RESET_COURSES')
            for (let i = 0, l = c.length; i < l; i++) {
              commit('SET_COURSE_ITEM', {
                i: c[i].time - 1,
                j: c[i].week,
                item: c[i]
              })
            }
          }
          Toast.clear()
          resolve()
        })
        .catch(err => {
          Toast.fail({ message: err.message, duration: 2000 })
          reject(err)
        })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
