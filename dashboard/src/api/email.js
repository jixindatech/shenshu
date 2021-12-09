import request from '@/utils/request'

export function get() {
  return request({
    url: `/system/email`,
    method: 'get'
  })
}

export function add(data) {
  return request({
    url: `/system/email`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/system/email/${id}`,
    method: 'put',
    data
  })
}
