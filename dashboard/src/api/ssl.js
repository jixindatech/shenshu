import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/shenshu/ssl`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/shenshu/ssl`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/shenshu/ssl/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/shenshu/ssl/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/shenshu/ssl/${id}`,
    method: 'delete'
  })
}
