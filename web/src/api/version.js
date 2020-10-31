import request from '@/utils/request'

export function getVersion() {
  return request({
    url: `http://47.106.188.222:11088/config.json?${new Date().getTime()}`,
    method: 'get'
  })
}
