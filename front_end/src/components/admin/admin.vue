<template>
  <div class="student">
    <h1>Admin Main Page</h1>
    <h1>welcome {{userId}}</h1>
    <a href="#" @click="quit">注销登录</a>
    <!-- <div v-show='user'>
      <p>welcome{{user.username}}</p>
    </div> -->
  </div>
</template>

<script>
import {getCookie, delCookie} from '@/common/js/cookie.js'
export default {
  name: 'student',
  data () {
    return {
      userId: '123'
    }
  },
  created () {
    // let uname = localStorage.getItem('userId')
    let uname = getCookie('AdminId')
    console.log('uname', uname)
    this.userId = uname
    if (uname === '') {
      this.$router.push('/')
    }
  },
  methods: {
    quit () {
      delCookie('userId')
      this.$router.push('/')
    }
  }
}
</script>

<style>
</style>
