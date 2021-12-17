import request from '@/utils/request'

export function add(data) {
  return request({
    url: `/nginx/site`,
    method: 'post',
    data
  })
}

export function put(id, data) {
  return request({
    url: `/nginx/site/${id}`,
    method: 'put',
    data
  })
}

export function get(id) {
  return request({
    url: `/nginx/site/${id}`,
    method: 'get'
  })
}

export function getList(query, current = 1, size = 20) {
  return request({
    url: `/nginx/site`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function deleteById(id) {
  return request({
    url: `/nginx/site/${id}`,
    method: 'delete'
  })
}

export function enable(id) {
  return request({
    url: `/shenshu/site/${id}/enable`,
    method: 'post'
  })
}

export function GetSiteRuleGroup(id, query) {
  return request({
    url: `/shenshu/site/${id}/rulegroup`,
    method: 'get',
    params: query
  })
}

export function UpdateSiteRuleGroup(id, data) {
  return request({
    url: `/shenshu/site/${id}/rulegroup`,
    method: 'put',
    data
  })
}
