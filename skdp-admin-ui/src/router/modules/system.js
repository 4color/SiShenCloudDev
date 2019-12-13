import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true, cache: true }

export default {
  path: '/system',
  name: 'router-system',
  meta,
  redirect: { name: 'notice' },
  component: layoutHeaderAside,
  children: (pre => [
    {
      path: 'notice',
      name: `${pre}index`,
      component: _import('system/notice/index'),
      meta: {
        ...meta,
        title: '通知公告'
      }
    },
    {
      path: 'noticedetail/:id',
      name: `${pre}noticedetail`,
      component: _import('system/notice/detail'),
      meta: {
        ...meta,
        title: '公告详情'
      }
    },
    {
      path: 'article',
      name: `${pre}article`,
      component: _import('system/notice/index'),
      meta: {
        ...meta,
        title: '文章'
      }
    },
    {
      path: 'dict',
      name: `system-dict`,
      component: _import('system/dict/index'),
      meta: {
        ...meta,
        title: '数据字典'
      }
    },
    {
      path: 'watch',
      name: `${pre}watch`,
      component: _import('system/watch/index'),
      meta: {
        ...meta,
        title: '系统监控'
      }
    },
    {
      path: 'config',
      name: `${pre}config`,
      component: _import('system/config/index'),
      meta: {
        ...meta,
        title: '配置'
      }
    }
  ])('router-system-')
}
