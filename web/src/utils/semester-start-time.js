import dayjs from 'dayjs'

const key = 'ccsu_semester_start_time'

export function setSemesterStartTime(t) {
  if (!t) {
    const now = new Date()
    const y = now.getFullYear()
    const m = now.getMonth() + 1
    if (m > 6) {
      setTime(`${y}/09/09`)
    } else {
      setTime(`${y}/02/08`)
    }
  } else {
    setTime(t)
  }
}

function setTime(t) {
  localStorage.setItem(key, t)
}

export function getSemesterStartTime() {
  const d = localStorage.getItem(key)
  if (d) {
    return dayjs(d)
  }
  return null
}
