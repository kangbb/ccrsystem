import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login'
import Student from '@/components/student/student'
import Approver from '@/components/approver/approver'
import Admin from '@/components/admin/admin'

Vue.use(Router)

export default new Router({
  routes: [{
    path: '/',
    component: Login
  }, {
    path: '/student',
    component: Student
  }, {
    path: '/approver',
    component: Approver
  }, {
    path: '/admin',
    component: Admin
  }]
})
