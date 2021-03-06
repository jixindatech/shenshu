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
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: 'Dashboard', icon: 'dashboard' }
    }]
  },
  {
    path: '/nginx',
    component: Layout,
    meta: { title: '站点管理', icon: 'el-icon-s-help' },
    redirect: 'site',
    children: [{
      path: 'site',
      name: 'Site',
      component: () => import('@/views/nginx/site'),
      meta: { title: '站点配置', icon: 'dashboard' }
    },
    {
      path: 'ssl',
      name: 'SSL',
      component: () => import('@/views/nginx/ssl'),
      meta: { title: '证书配置', icon: 'dashboard' }
    },
    {
      path: 'upstream',
      name: 'Upstream',
      component: () => import('@/views/nginx/upstream'),
      meta: { title: 'Upstream配置', icon: 'dashboard' }
    }]
  },
  {
    path: '/shenshu',
    component: Layout,
    meta: { title: '规则管理', icon: 'el-icon-s-help' },
    redirect: 'ip',
    children: [{
      path: 'ip',
      name: 'IP',
      component: () => import('@/views/shenshu/ip'),
      meta: { title: 'IP管理', icon: 'dashboard' }
    },
    {
      path: 'cc',
      name: 'CC',
      component: () => import('@/views/shenshu/cc'),
      meta: { title: 'CC配置', icon: 'dashboard' }
    },
    /*
    {
      path: 'bot',
      name: 'Bot',
      component: () => import('@/views/shenshu/bot'),
      meta: { title: 'Bot管理', icon: 'dashboard' }
    },
    */
    {
      path: 'rule',
      name: 'Rule',
      component: () => import('@/views/shenshu/rule'),
      meta: { title: '规则管理', icon: 'dashboard' },
      children: [{
        path: 'batchgroup',
        name: 'BatchGroup',
        component: () => import('@/views/shenshu/rule/batchgroup'),
        meta: { title: 'Batch组配置', icon: 'dashboard' }
      },
      {
        path: 'batchrule',
        name: 'BatchRule',
        hidden: true,
        component: () => import('@/views/shenshu/rule/batchrule'),
        meta: { title: 'Batch规则', icon: 'dashboard' }
      },
      {
        path: 'specificgroup',
        name: 'SpecificGroup',
        component: () => import('@/views/shenshu/rule/specificgroup'),
        meta: { title: 'Specific组配置', icon: 'dashboard' }
      },
      {
        path: 'specificrule',
        name: 'SpecificRule',
        hidden: true,
        component: () => import('@/views/shenshu/rule/specificrule'),
        meta: { title: 'Batch规则', icon: 'dashboard' }
      }
      ]
    }]
  },
  {
    path: '/event',
    component: Layout,
    meta: { title: '日志列表', icon: 'el-icon-s-help' },
    redirect: 'ccevent',
    children: [{
      path: 'ccevent',
      name: 'CCEvent',
      component: () => import('@/views/event/cc'),
      meta: { title: 'CC日志', icon: 'dashboard' }
    },
    {
      path: 'batchruleevent',
      name: 'BatchRuleEvent',
      component: () => import('@/views/event/batchrule'),
      meta: { title: 'Batch日志', icon: 'dashboard' }
    },
    {
      path: 'specificruleevent',
      name: 'SpecificRuleEvent',
      component: () => import('@/views/event/specificrule'),
      meta: { title: 'Specific日志', icon: 'dashboard' }
    }]
  },
  {
    path: '/system',
    component: Layout,
    meta: { title: '系统管理', icon: 'el-icon-s-help' },
    redirect: 'user',
    children: [{
      path: 'user',
      name: 'User',
      component: () => import('@/views/system/user'),
      meta: { title: '用户管理', icon: 'dashboard' }
    },
    /*
    {
      path: 'msg',
      name: 'Msg',
      component: () => import('@/views/system/msg'),
      meta: { title: '短信管理', icon: 'dashboard' }
    },
    */
    {
      path: 'config',
      name: 'Config',
      component: () => import('@/views/system/config'),
      meta: { title: '系统配置', icon: 'dashboard' }
    }]
  },
  /*
  {
    path: '/example',
    component: Layout,
    redirect: '/example/table',
    name: 'Example',
    meta: { title: 'Example', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'table',
        name: 'Table',
        component: () => import('@/views/table/index'),
        meta: { title: 'Table', icon: 'table' }
      },
      {
        path: 'tree',
        name: 'Tree',
        component: () => import('@/views/tree/index'),
        meta: { title: 'Tree', icon: 'tree' }
      }
    ]
  },

  {
    path: '/form',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'Form',
        component: () => import('@/views/form/index'),
        meta: { title: 'Form', icon: 'form' }
      }
    ]
  },

  {
    path: '/nested',
    component: Layout,
    redirect: '/nested/menu1',
    name: 'Nested',
    meta: {
      title: 'Nested',
      icon: 'nested'
    },
    children: [
      {
        path: 'menu1',
        component: () => import('@/views/nested/menu1/index'), // Parent router-view
        name: 'Menu1',
        meta: { title: 'Menu1' },
        children: [
          {
            path: 'menu1-1',
            component: () => import('@/views/nested/menu1/menu1-1'),
            name: 'Menu1-1',
            meta: { title: 'Menu1-1' }
          },
          {
            path: 'menu1-2',
            component: () => import('@/views/nested/menu1/menu1-2'),
            name: 'Menu1-2',
            meta: { title: 'Menu1-2' },
            children: [
              {
                path: 'menu1-2-1',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-1'),
                name: 'Menu1-2-1',
                meta: { title: 'Menu1-2-1' }
              },
              {
                path: 'menu1-2-2',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-2'),
                name: 'Menu1-2-2',
                meta: { title: 'Menu1-2-2' }
              }
            ]
          },
          {
            path: 'menu1-3',
            component: () => import('@/views/nested/menu1/menu1-3'),
            name: 'Menu1-3',
            meta: { title: 'Menu1-3' }
          }
        ]
      },
      {
        path: 'menu2',
        component: () => import('@/views/nested/menu2/index'),
        name: 'Menu2',
        meta: { title: 'menu2' }
      }
    ]
  },
  */
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'https://github.com/jixindatech/shenshu',
        meta: { title: 'External Link', icon: 'link' }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
