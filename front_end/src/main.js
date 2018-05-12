// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
// import store from './vuex.js'
import App from './App'
import router from './router'
import VueResource from 'vue-resource'
import jQuery from 'jquery'

Vue.use(VueResource)
Vue.config.productionTip = false
Vue.http.options.xhr = {withCredentials: true}

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})

// router.beforeEach((to, from, next) => {
//   if (to.meta.requireAuth) {
//     if (!isEmptyObject(store.state.user)) {
//       next()
//     } else {
//       next({
//         path: '/login',
//         query: {
//           redirect: {
//             redirect: to.fullPath
//           }
//         }
//       })
//     }
//   } else {
//     next()
//   }
// })

// function isEmptyObject (obj) {
//   for (var key in obj) {
//     return false
//   }
//   return true
// }
