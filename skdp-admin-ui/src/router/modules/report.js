import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true }

export default {
  path: '/report',
  name: 'router-report',
  meta,
  redirect: { name: '/report/index' },
  component: layoutHeaderAside,
  children: (pre => [
    {
      path: 'index',
      name: `${pre}index`,
      component: _import('report/index'),
      meta: {
        ...meta,
        title: '消费报表'
      }
    },
    {
      path: 'detail/:id',
      name: `${pre}detail`,
      component: _import('market/detail'),
      meta: {
        ...meta,
        title: '应用详情'
      }
    },
    {
      path: 'addapp',
      name: `${pre}addapp`,
      component: _import('market/addApp'),
      meta: {
        ...meta,
        title: '添加应用'
      }
    },
    {
      path: 'config',
      name: `${pre}config`,
      component: _import('demo/page3'),
      meta: {
        ...meta,
        title: '配置'
      }
    }
  ])('router-report-')
}
