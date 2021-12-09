import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/shenshu/bot`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/shenshu/bot`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/shenshu/bot/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/shenshu/bot/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/shenshu/bot/${id}`,
    method: 'delete'
  })
}
