import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data
  })
}

export function getUserInfo(username) {
  return request({
    url: '/user/getUserInfo',
    method: 'get',
    params: { username }
  })
}

export function getUserList(data) {
  return request({
    url: '/user/getUserList',
    method: 'post',
    data
  })
}

export function changeUserStatus(data) {
  return request({
    url: '/user/changeUserStatus',
    method: 'post',
    data
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}

export function createUser(data) {
  return request({
    url: '/user/addUser',
    method: 'post',
    data
  })
}

export function getExistUser(data) {
  return request({
    url: '/user/getExistUser',
    method: 'post',
    data
  })
}

export function resetUser(data) {
  return request({
    url: '/user/resetPassword',
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
    url: '/user/getUserTotal',
    method: 'get'
  })
}
