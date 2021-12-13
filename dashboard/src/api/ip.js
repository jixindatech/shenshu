import request from '@/utils/request'

export function getList(id, query, current = 1, size = 10) {
  return request({
    url: `/shenshu/site/${id}/ip`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(id, data) {
  return request({
    url: `/shenshu/site/${id}/ip`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/shenshu/site/ip/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/shenshu/site/ip/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/shenshu/site/ip/${id}`,
    method: 'delete'
  })
}
