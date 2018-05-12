import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login'
import Student from '@/components/student/student'
import Approver from '@/components/approver/approver'
import Admin from '@/components/admin/admin'
import Query from '@/components/student/query'
import Classroom from '@/components/student/classroom'
import Reservation from '@/components/student/reservation'
import NotApprove from '@/components/approver/notApprove'
import HasApprove from '@/components/approver/HasApprove'

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
  }, {
    path: '/student/query',
    component: Query
  }, {
    path: '/student/classroom',
    component: Classroom
  }, {
    path: '/student/reservation',
    component: Reservation
  }, {
    path: '/approver/notApprove',
    component: NotApprove
  }, {
    path: '/approver/hasApprove',
    component: HasApprove
  }]
})
