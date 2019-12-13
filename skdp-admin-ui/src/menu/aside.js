// 菜单 侧边栏
export default [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '部门管理',
    icon: 'user',
    children: [
      { path: '/department/xzq', icon: 'microchip', title: '行政区管理' },
      { path: '/department/role', icon: 'car', title: '角色管理' },
      { path: '/department/user', icon: 'user', title: '租户管理' }
    ]
  },
  {
    title: '消费报表',
    icon: 'list-ol',
    path: '/report/index'
  },
  {
    title: '应用市场',
    icon: 'shopping-cart',
    path: '/market/index'
  },
  {
    title: '系统管理',
    icon: 'cog',
    path: '/system',
    children: [
      { path: '/system/notice', icon: 'bullhorn', title: '通知公告' },
      { path: '/system/article', icon: 'file-audio-o', title: '文章管理' },
      { path: '/system/dict', icon: 'list-alt', title: '数据字典' },
      { path: '/system/watch', icon: 'eye', title: '系统监控' },
      { path: '/system/config', icon: 'cogs', title: '配置管理' }
    ]
  },
  {
    title: 'test',
    icon: 'shopping-cart',
    path: '/test'
  }
]
