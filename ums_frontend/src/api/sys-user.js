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

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}

export function changeUserStatus() {}

export function createUser(data) {
  return request({
    url: '/user/AddSysUser',
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

export function resetUserPassword(data) {
  return request({
    url: '/user/ResetSysPassword',
    method: 'post',
    data
  })
}

