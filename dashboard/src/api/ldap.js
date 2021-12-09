import request from '@/utils/request'

export function get() {
  return request({
    url: `/system/ldap`,
    method: 'get'
  })
}

export function add(data) {
  return request({
    url: `/system/ldap`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/system/ldap/${id}`,
    method: 'put',
    data
  })
}
