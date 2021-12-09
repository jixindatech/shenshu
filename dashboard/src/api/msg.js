import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/system/msg`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/system/msg`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/system/msg/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/system/msg/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/system/msg/${id}`,
    method: 'delete'
  })
}

export function sendUserMsg(id, data) {
  return request({
    url: `/system/msg/${id}/user`,
    method: 'post',
    data
  })
}
