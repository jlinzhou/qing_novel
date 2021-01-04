import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import fitler from './utils/filter';
import VueCookies from 'vue-cookies'
import VueLazyload from 'vue-lazyload'

import VueAnalytics from 'vue-analytics'




Vue.use(VueCookies)

Vue.config.productionTip = false
Vue.use(ElementUI);
Vue.use(VueLazyload, {
  loading: require('@/assets/images/loading.png'),//加载中图片，一定要有，不然会一直重复加载占位图
  // error: require('img/error.png')  //加载失败图片
});
router.beforeEach((to, from, next) => {
  window.scrollTo(0, 0)
  next()
});


router.afterEach((to, from, next) => {
  window.scrollTo(0, 0);
  document.body.scrollTop = 0;
  document.documentElement.scrollTop = 0;
  window.pageYOffset = 0
});

Object.keys(fitler).forEach((key) => {
  Vue.filter(key, (fitler)[key]);
});
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
