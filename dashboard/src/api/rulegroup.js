import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/shenshu/rulegroup`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/shenshu/rulegroup`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/shenshu/rulegroup/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/shenshu/rulegroup/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/shenshu/rulegroup/${id}`,
    method: 'delete'
  })
}

export function GetSiteRuleGroup(id) {
  return request({
    url: `/shenshu/site/${id}/rulegroup`,
    method: 'get'
  })
}

export function UpdateSiteRuleGroup(id, data) {
  return request({
    url: `/shenshu/site/${id}/rulegroup`,
    method: 'put',
    data
  })
}

export function EnableSiteRuleGroup(id) {
  return request({
    url: `/shenshu/site/${id}/rulegroup`,
    method: 'post'
  })
}
