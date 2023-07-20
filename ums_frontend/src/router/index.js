import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  {
    'path': '/',
    'component': Layout,
    'redirect': '/dashboard',
    'children': [{
      'path': 'dashboard',
      'name': 'Dashboard',
      'component': () => import('@/views/dashboard/index'),
      'meta': { 'title': '数据面板', 'icon': 'dashboard' }
    }]
  },
  {
    path: '/user-manager',
    component: Layout,
    redirect: 'noRedirect',
    alwaysShow: true,
    name: '用户管理',
    meta: { title: '用户管理', icon: 'el-icon-set-up' },
    children: [
      {
        path: 'ldap-user',
        name: 'LDAP管理',
        component: () => import('@/views/ldap-manage/user'),
        meta: { title: 'LDAP管理', icon: 'el-icon-user-solid' }
      }
    ]
  },
  {
    path: '/oms-manage',
    redirect: 'noRedirect',
    alwaysShow: true,
    component: Layout,
    name: '平台管理',
    meta: { title: '平台管理', icon: 'el-icon-s-tools' },
    children: [
      {
        path: 'user-manager',
        name: 'user-manager',
        component: () => import('@/views/ums-manage/user'),
        meta: { title: '用户管理', icon: 'el-icon-user' }
      }
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

let createRouter = ''

if (process.env.NODE_ENV === 'development') {
  createRouter = () => new Router({
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes
  })
} else {
  createRouter = () => new Router({
    mode: 'history',
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes
  })
}

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
