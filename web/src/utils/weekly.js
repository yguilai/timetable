import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import weekOfYear from 'dayjs/plugin/weekOfYear'
dayjs.locale('zh-cn')
dayjs.extend(weekOfYear)

const key = 'ccsu_current_weekly'

export function getWeekly() {
  const w = localStorage.getItem(key)
  if (w) {
    return parseInt(w)
  }
  return 0
}

export function setWeekly(weekly) {
  return localStorage.setItem(key, weekly.toString())
}

export function getCurrentWeekly(semesterStartTime) {
  const r = dayjs().week() - semesterStartTime.week() + 1
  return r <= 0 ? 0 : r
}

export function removeWeekly() {
  return localStorage.setItem(key, '0')
}
