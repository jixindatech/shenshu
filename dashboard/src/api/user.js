import request from '@/utils/request'

export function login(data) {
  return request({
    url: `/login`,
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: `/system/user/info`,
    method: 'get'
  })
}

export function logout() {
  return request({
    url: `/system/user/logout`,
    method: 'post'
  })
}

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/system/user`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/system/user`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/system/user/${id}`,
    method: 'put',
    data
  })
}

export function updateInfo(data) {
  return request({
    url: `/system/user`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/system/user/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/system/user/${id}`,
    method: 'delete'
  })
}

export function resetUserPasswordById(id) {
  return request({
    url: `/system/user/password/${id}`,
    method: 'put'
  })
}
