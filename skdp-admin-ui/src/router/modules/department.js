import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true, cache: true }

export default {
  path: '/department',
  name: 'router-department',
  meta,
  redirect: { name: 'department' },
  component: layoutHeaderAside,
  children: (pre => [
    {
      path: 'xzq',
      name: `departmentxzq`,
      component: _import('department/xzq'),
      meta: {
        ...meta,
        title: '行政区管理'
      }
    },
    {
      path: 'role',
      name: `${pre}role`,
      component: _import('department/role'),
      meta: {
        ...meta,
        title: '角色管理'
      }
    },
    {
      path: 'user',
      name: `${pre}user`,
      component: _import('department/user'),
      meta: {
        ...meta,
        title: '租户管理'
      }
    },
    {
      path: 'userdetail/:id',
      name: `${pre}userdetail`,
      component: _import('department/user/detail.vue'),
      meta: {
        ...meta,
        title: '用户详情'
      }
    }
  ])('router-department-')
}
