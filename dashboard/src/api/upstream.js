import request from '@/utils/request'

export function add(data) {
  return request({
    url: `/nginx/upstream`,
    method: 'post',
    data
  })
}

export function put(id, data) {
  return request({
    url: `/nginx/upstream/${id}`,
    method: 'put',
    data
  })
}

export function get(id) {
  return request({
    url: `/nginx/upstream/${id}`,
    method: 'get'
  })
}

export function getList(query, current = 1, size = 20) {
  return request({
    url: `/nginx/upstream`,
    method: 'get',
    params: { ...query, current, size }
  })
}

export function deleteById(id) {
  return request({
    url: `/nginx/upstream/${id}`,
    method: 'delete'
  })
}
