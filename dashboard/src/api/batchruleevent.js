import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/shenshu/event/batchrule`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function getInfo(query) {
  return request({
    url: `/shenshu/event/info/batchrule`,
    method: 'get',
    params: { ...query }
  })
}
