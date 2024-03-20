/**
 * createRouter 这个为创建路由的方法
 * createWebHashHistory 这个就是vue2中路由的模式，
 *                      这里的是hash模式，这个还可以是createWebHistory等
 * RouteRecordRaw 这个为要添加的路由记录，也可以说是routes的ts类型
 */
import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
// 路由记录，这个跟vue2中用法一致，就不做过多解释了
const routes:Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: () => import("@/components/HelloWorld.vue"),
    alias: '/home',
    meta: {
      title: '首页'
    }
  },
  {
    path: '/joblist',
    name: 'joblist',
    component: () => import("@/components/JobList.vue"),
    alias: '/joblist',
    meta: {
      title: 'Job配置列表'
    }
  },
  {
    path: '/jobconfig',
    name: 'jobconfig',
    component: () => import("@/components/JobConfig.vue"),
    alias: '/jobconfig',
    meta: {
      title: 'Job配置文件'
    }
  },
  {
    path: '/setting',
    name: 'setting',
    component: () => import("@/components/Setting.vue"),
    alias: '/setting',
    meta: {
      title: '系统设置'
    }
  },
  {
    path: '/datasource',
    name: 'datasource',
    component: () => import("@/components/DataSource.vue"),
    alias: '/datasource',
    meta: {
      title: '数据源设置'
    }
  },
  {
    path: '/about',
    name: 'about',
    component: () => import("@/components/About.vue"),
    alias: '/about',
    meta: {
      title: '关于'
    }
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});
// 使用 beforeEach 钩子
// router.beforeEach((to, from, next) => {
//   // 获取刷新前的路由信息
//   const previousRoute = to

//   // 进行相应的处理
//   console.log(previousRoute) // Home

//   // 继续执行路由跳转
//   next()
// })
export default router;