<template>
  <div class="login-wrapper" v-show="showLogin">
      <h1>教室预订系统</h1>
      <div class="middle">
        <div class="id">
          <label>账号</label>
          <input type="text" placeholder="学号/工号" v-model="id" @click="clear" ><br />
        </div>
        <div class="pass">
          <label>密码</label>
          <input type="password" placeholder="请输入密码" v-model="password" @click="clear"><br />
        </div>
      </div>
      <div class="error-msg">{{error}}</div>
      <button @click="login('student')" >学生登录</button>
      <button @click="login('approver')">审核员登录</button>
      <button @click="login('admin')">管理员登录</button>
  </div>
</template>

<script>
import {setCookie} from '../common/js/cookie.js'
export default {
  data () {
    return {
      showLogin: true,
      id: '',
      password: '',
      error: '',
      showError: false
    }
  },
  methods: {
    login (param) {
      if (this.id === '' || this.password === '') {
        this.error = '账号或密码不能为空！'
        this.showError = true
      } else {
        var intId = parseInt(this.id)
        var obj = {}

        let jid = param.substring(0, 1).toUpperCase() + param.substring(1) + 'Id'
        let jpwd = param.substring(0, 1).toUpperCase() + param.substring(1) + 'Pwd'

        obj[jid] = intId;
        obj[jpwd] = this.password

        var apiStr = '/' + param + '/signin'
        this.$http.post(apiStr, obj).then(res=>{
          console.log(res)
          console.log("status: ",res.status)
          setCookie(jid, this.id, 1000 * 60)
          this.$router.push('/' + param)
        }).catch(err=>{
          this.id = ''
          this.password = ''
          if (err.status == 500) {
            this.error = '用户名与密码不匹配!'
          } else if (err.status == 404) {
            this.error = '此用户不存在！'
          }
        })
      }
    },
    clear () {
      this.error = ''
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login-wrapper {
  border: 1px solid grey;
  width: 400px;
  height: 260px;
  margin: 200px auto;
}
h1 {
  color: #046e61;
  font-size: 32px;
  margin-top: 40px;
}
label {
  color: black;
}
button {
  background-color: #046e61;
  border: none;
  font-size: 12px;
  height: 25px;
  width: 80px;
  color: white;
  margin:10px;
}
.middle {
  margin: 10px auto;
  font-size: 12px;
}

.id , .pass{
  margin:5px auto;
}

.error-msg {
  height:20px;
  line-height: 20px;
  color: red;
}
</style>
