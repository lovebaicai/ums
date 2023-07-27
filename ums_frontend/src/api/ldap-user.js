import request from '@/utils/request'

export function getUserList(data) {
  return request({
    url: '/ldap/getUserList',
    method: 'post',
    data
  })
}

export function changeUserStatus(data) {
  return request({
    url: '/ldap/changeUserStatus',
    method: 'post',
    data
  })
}

export function createUser(data) {
  return request({
    url: '/ldap/addLdapUser',
    method: 'post',
    data
  })
}

export function getExistUser(data) {
  return request({
    url: '/ldap/getExistUser',
    method: 'post',
    data
  })
}

export function resetUser(data) {
  return request({
    url: '/ldap/resetPassword',
    method: 'post',
    data
  })
}

export function syncUser() {
  return request({
    url: '/ldap/sync',
    method: 'get'
  })
}

export function getUserTotal() {
  return request({
    url: '/ldap/getUserTotal',
    method: 'get'
  })
}
