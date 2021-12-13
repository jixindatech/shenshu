import request from '@/utils/request'

export function getList(id, query, current = 1, size = 10) {
  return request({
    url: `/shenshu/site/${id}/cc`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(id, data) {
  return request({
    url: `/shenshu/site/${id}/cc`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/shenshu/site/cc/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/shenshu/site/cc/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/shenshu/site/cc/${id}`,
    method: 'delete'
  })
}
