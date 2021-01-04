import Vue from 'vue'
import VueRouter from 'vue-router'
// import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  // {
  //   path: '/index',
  //   name: 'index',
  //   component: () => import(/* webpackChunkName: "about" */ '../views/index.vue')
  // },
  // {
  //   path: '/',
  //   redirect: '/index/home',
  // },
  {
    path: '/',
    name: 'index',
    component: () => import(/* webpackChunkName: "about" */ '../views/index.vue'),
    redirect: '/home',
    children: [
      {
        path: 'home',
        name: '首页',
        component: resolve => require(['../views/pages/home.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/home.vue')
      },
      {
        path: 'category/xuanhuan',
        name: '玄幻',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'category/xiuzhen',
        name: '修真',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'category/dushi',
        name: '都市',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'category/chuanyue',
        name: '穿越',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'category/wangyou',
        name: '网游',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'category/kehuan',
        name: '科幻',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'category/nvpin',
        name: '女频',
        component: resolve => require(['../views/pages/category.vue'], resolve)
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/category.vue')
      },
      {
        path: 'chapter/:book_id/:page?/:page_size?/:sort?',
        component: resolve => require(['../views/pages/chapter.vue'], resolve)
        // name: '科幻',
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/chapter.vue')
      },
      {
        path: 'reader/:book_id/:id',
        // name: '科幻',
        component: resolve => require(['../views/pages/reader.vue'], resolve)
      },
      {
        path: 'search',
        // name: '科幻',
        component: resolve => require(['../views/pages/search.vue'], resolve),
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/search.vue')
      },
      {
        path: 'latest_view',
        // name: '科幻',
        component: resolve => require(['../views/pages/latest_view.vue'], resolve),
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/search.vue')
      },
      {
        path: 'rank/:type?',
        // name: '科幻',
        component: resolve => require(['../views/pages/rank.vue'], resolve),
        // component: () => import(/* webpackChunkName: "about" */ '../views/pages/search.vue')
      },
      {
        path: 'global_search',
        component: resolve => require(['../views/pages/global_search.vue'], resolve),
      },
      {
        path: 'global_search/search_chapter/:novel_url/:novel_name',
        component: resolve => require(['../views/pages/search_chapter.vue'], resolve),
      },
      {
        path: 'global_search/search_reader/:content_url/:novel_url/:novel_name',
        component: resolve => require(['../views/pages/search_reader.vue'], resolve),
      }
    ]
  },

]

const router = new VueRouter({
  mode: 'history',
  // base: "/mb",
  routes,
  scrollBehavior(to, from, savedPosition) {
    return { x: 0, y: 0 }
  }
})
// process.env.BASE_URL
export default router
