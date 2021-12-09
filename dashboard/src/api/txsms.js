import request from '@/utils/request'

export function get() {
  return request({
    url: `/system/txsms`,
    method: 'get'
  })
}

export function add(data) {
  return request({
    url: `/system/txsms`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/system/txsms/${id}`,
    method: 'put',
    data
  })
}
