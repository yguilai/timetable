import request from '@/utils/request'

export function getTimetable(data) {
  return request({
    url: `/stu/${data.stuId}/timetable`,
    method: 'get',
    params: {
      weekly: data.weekly,
      yearly: data.yearly
    }
  })
}
