<template>
  <div class='student'>
    <div class="nav">
      <div class="title">
        <img src="../../images/search.png">
        <h1>教室预订系统</h1>
        <h3 class="username">学生 {{userId}}</h3>
      </div>
      <ul class="navigationBar">
        <li @click="toSearch" :class="{'activeTag':showSearch}">教室状态查询</li>
        <li @click="toReservation" :class="{'activeTag':showResvation}">我的申请</li>
      </ul>
      <div class="studentPwd">
         <button @click="reviseStudentPwd">修改密码</button>
      </div>
      <div class="exit">
         <a href='#' @click='quit'>退出登录</a>
      </div>
      <div class="line"></div>
    </div>

    <div class="content">
      <div class="panel">
        <div class="search" v-show="showSearch">
          <h1>筛选条件:</h1>
          <form id='selectForm' v-on:submit='formSubmit'>
            <div class="date select">
              <img src="../../images/calendar.png">
              <input type='date' v-model='date'>
            </div>
            <div class='time select'>
              <img src="../../images/time.png">
              <select v-model='begin'>
                <option selected='selected' disabled='disabled'>请选择</option>
                <option v-for='(item,index) in oneToTen' v-bind:key="index">第{{item}}节</option>
              </select>
              <span>至</span>
              <select v-model='end'>
                <option selected='selected' disabled='disabled'>请选择</option>
                <option v-for='(items,index) in oneToTen' v-bind:key="index">第{{items}}节</option>
              </select>
            </div>
            <div class="capacity select">
              <img src="../../images/user.png">
              <select v-model='capacity'>
                <option selected='selected' disabled='disabled'>容量/人</option>
                <option value='50'>50</option>
                <option value='100'>100</option>
                <option value='200'>200</option>
              </select>
            </div>
            <div class="submit select"><input type='submit' value='筛选'></div>
          </form>

          <div class="table1">
             <table v-if='showList' class="table1">
              <tr>
                <th>序列</th>
                <th>容量</th>
                <th>教室号</th>
                <th>操作</th>
              </tr>
              <tr v-for='(items,index) in this.chooseList' v-bind:key="index">
                <td>{{index+1}}</td>
                <td>{{items.Capacity}}</td>
                <td>{{items.ClassroomNum}}</td>
                <td><button type='text' id='apply' v-on:click='applyClassroom(index,items)'>申请</button></td>
              </tr>
            </table>
          </div>

          <div id='applyWindow' v-show='showApply' class="apply-window">
            <!-- <button v-on:click='closeWindow'>x</button> -->
            <form v-on:submit='applySubmit'>
              <p><span>教室号：</span>{{this.info.ClassroomNum}}</p>
              <p><span>申请日期：</span>{{this.info.year}}<span>年</span>{{this.info.month}}<span>月</span>{{this.info.day}}<span>日</span></p>
              <p><span>时间：</span>{{begin}}<span>至</span>{{end}}</p>
              <p><span>参与人数：</span>{{capacity}}<span>人</span></p>
              <p>使用方隶属组织：<input type='text' placeholder='如，数据院15级软工二班' v-model='organization' id='organization'></p>
              <p>申请教室用途：<input type='text' placeholder='必须填写用途100字以内' v-model='reservationInfo' id='reservationInfo'></p>
            </form>
            <button @click="applySubmit">提交</button>
            <button @click="closeWindow">取消</button>
          </div>
        </div>
        <div class="myRes" v-show="showResvation">
          <h1>我的申请:</h1>
          <table v-show="showResList">
            <tr>
              <th>序列</th>
              <th>教室号</th>
              <th>日期</th>
              <th>时间</th>
              <th>容量</th>
              <th>操作</th>
            </tr>
            <tr v-for="(items, index) in this.reservation" v-bind:key="index">
              <td>{{index+1}}</td>
              <td>{{items.ClassroomNum}}</td>
              <td>{{items.year}}<span>年</span>{{items.month}}<span>月</span>{{items.day}}<span>日</span></td>
              <td>{{items.begin}}<span>至</span>{{items.end}}</td>
              <td>{{items.Capacity}}<span>人</span></td>
              <td><button @click="showDetail(items)">详情</button></td>
            </tr>
          </table>
        </div>
        <div id='applyWindow' v-show='showRes' class="apply-window">
            <!-- <button v-on:click='closeWindow'>x</button> -->
            <form>
              <p><span>教室号：</span>{{this.res.ClassroomNum}}</p>
              <p><span>申请日期：</span>{{this.res.year}}<span>年</span>{{this.res.month}}<span>月</span>{{this.res.day}}<span>日</span></p>
              <p><span>时间：</span>{{this.res.begin}}<span>至</span>{{this.res.end}}</p>
              <p><span>参与人数：</span>{{this.res.Capacity}}<span>人</span></p>
              <p><span>使用方隶属组织：{{this.res.OrganizationName}}</span></p>
              <p><span>申请教室用途：{{this.res.ResReason}}</span></p>
              <p><span>审核状态：</span>{{state}}</p>
              <p v-show="notPass"><span>不通过原因：</span>{{this.res.ApprovalNote}}</p>
            </form>
            <button @click="deleteApply">删除申请</button>
            <button @click="closeRes">取消</button>
            <button @click="reviseApply" v-show="canRevise">修改申请</button>
          </div>

          <div id='applyWindow' v-show='showRevise' class="apply-window">
            <!-- <button v-on:click='closeWindow'>x</button> -->
            <form>
              <p><span>教室号：</span>{{this.res.ClassroomNum}}</p>
              <p><span>申请日期：</span>{{this.res.year}}<span>年</span>{{this.res.month}}<span>月</span>{{this.res.day}}<span>日</span></p>
              <p><span>时间：</span>{{this.res.begin}}<span>至</span>{{this.res.end}}</p>
              <p><span>参与人数：</span>{{this.res.Capacity}}<span>人</span></p>
              <p>使用方隶属组织：<input type='text' v-model='organization' id='organization'></p>
              <p>申请教室用途：<input type='text' v-model='reservationInfo' id='reservationInfo'></p>
            </form>
            <button @click="reviseSubmit">提交修改</button>
            <button @click="closeRevise">取消修改</button>
          </div>

          <div id='applyWindow' v-show='showStudentInfo' class="apply-window">
            <!-- <button v-on:click='closeWindow'>x</button> -->
            <form>
              <p><span>姓名：</span>{{this.username}}</p>
              <p><span>学号：</span>{{this.userId}}</p>
              <p>密码：<input type='text' placeholder='输入新密码6位' v-model='password'></p>
            </form>
            <button @click="revisePwd">修改密码</button>
            <button @click="closeStudentInfo">取消</button>
          </div>
      </div>
      <p v-if='showError'>没有找到符合条件的教室<br/>尝试其他的筛选条件吧~</p>
      <p v-if='showNoRes'>没有教室申请</p>
    </div>

  </div>
</template>

<script>
import {getCookie, delCookie} from '@/common/js/cookie.js'
export default {
  name: 'student',
  data () {
    return {
      username: '',
      password: '',
      userId: '',
      date: '',
      oneToTen: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15],
      classrooms: {},
      reservations: {},
      begin: '请选择',
      end: '请选择',
      capacity: '容量/人',
      chooseList: [],
      classroomNum: '',
      showError: false,
      showList: false,
      showApply: false,
      info: {},
      organization: '',
      reservationInfo: '',
      showSearch: true,
      showResvation: false,
      showResList: false,
      reservation: [],
      showRes: false,
      showNoRes: false,
      notPass: false,
      res: [],
      state: '',
      canRevise: false,
      showRevise: false,
      showStudentInfo: false
    }
  },
  created () {
    let uname = getCookie('StudentId')
    console.log('uname', uname)
    this.userId = uname
    if (uname === '') {
      this.$router.push('/')
    }
  },
  methods: {
    toSearch () {
      this.showSearch = true
      this.showResvation = false
      this.showNoRes = false
    },
    toReservation () {
      this.showSearch = false
      this.showResvation = true
      this.showError = false

      this.reservation = []
       var apiStr = '/api/api/users/student/reservations';
       this.$http.get(apiStr).then(res=>{
        res = res.body
        console.log(res)
        var count = 0
        for (var item in res) {
          var temp = res[item]
          var obj = {}
          obj.ClassroomNum = temp.ClassroomNum
          obj.year = this.parseDateTime(temp.StartTime, 'year')
          obj.month = this.parseDateTime(temp.StartTime, 'month')
          obj.day = this.parseDateTime(temp.StartTime, 'day')
          obj.begin = this.parseDateTime(temp.StartTime, 'time')
          obj.end = this.parseDateTime(temp.EndTime, 'time')
          obj.OrganizationName = temp.OrganizationName
          obj.Capacity = temp.Capacity
          obj.ResReason = temp.ResReason
          obj.ApprovalNote = temp.ApprovalNote
          obj.ResState = temp.ResState
          obj.indexer = count
          obj.ResId = temp.ResId
          count = count + 1
          this.reservation.push(obj)
        //  console.log(obj)
        }
        this.showResList = true
      }).catch(err=>{
        this.showNoRes = true
      })
    },
    parseDateTime(dateTime, choice) {
      var str = ''
      if (choice == 'year') {
        str = dateTime.substring(0, 4)
      } else if (choice == 'month') {
        str = dateTime.substring(5, 7)
      } else if (choice == 'day') {
        str = dateTime.substring(8, 10)
      } else {
        str = dateTime.substring(11, dateTime.length)
      }
      return str
    },
    showDetail(items) {
      this.res = items
      if (this.res.ResState == 0) {
        this.state = '未审核'
        this.notPass = false
        this.canRevise = true
      } else if (this.res.ResState == 1) {
        this.state = '审核中'
        this.notPass = false
        this.canRevise = true
      } else if (this.res.ResState == 2) {
        this.state = '审核通过'
        this.notPass = false
        this.canRevise = false
      } else {
        this.state = '审核不通过'
        this.notPass = true
        this.canRevise = false
      }
      this.showRes = true
    },
    closeRes() {
      this.showRes = false
    },
    deleteApply() {
      if (confirm('确认删除该申请？')) {
        var id = parseInt(this.res.ResId)
        var apiStr = '/api/api/reservations/' + id
        this.$http.delete(apiStr, {body:{id: id}}).then(res=>{
          console.log(res)
          this.reservation.splice(parseInt(this.res.indexer), 1)
          this.closeRes()
          alert('删除成功')
          if (this.res.length == 0) {
            this.showNoRes = true
          }
        }).catch(err=>{
          this.closeRes()
          alert('删除失败')
        })
      }
    },
    reviseApply() {
      this.showRevise = true
      this.organization = this.res.OrganizationName
      this.reservationInfo = this.res.ResReason
      this.closeRes()
    },
    reviseSubmit() {
      this.res.OrganizationName = this.organization
      this.res.ResReason = this.reservationInfo
      var id = parseInt(this.res.ResId)
      var apiStr = '/api/api/reservations/'+ id
      this.$http.put(apiStr, {ResReason: this.res.ResReason, OrganizationName: this.res.OrganizationName, Id: id}).then(res=>{
          console.log(res)
          alert('修改成功')
      }).catch(err=>{
        alert('修改失败')
      })
      this.closeRevise()
    },
    closeRevise() {
      this.showRevise = false
    },
    formSubmit (e) {
      this.showError = false
      e.preventDefault()
      if (this.date != '' && this.begin != '请选择' && this.end != '请选择' && this.capacity != '容量/人') {
        if (this.date < this.currentDate()) {
          alert('日期选择错误！')
        } else if (this.compareTime(this.begin) > this.compareTime(this.end)) {
          alert('时间选择错误！')
        } else {
          var beginTime = this.getDateTime(this.date, this.begin)
          var endTime = this.getDateTime(this.date, this.end)
          var apiStr = '/api/api/classrooms/state'
          var cap = parseInt(this.capacity)

          this.$http.get(apiStr, {params:{ClassroomCampus: '东校园', ClassroomBuilding: '公教楼', StartTime: beginTime, EndTime: endTime, Capacity: cap}}).then(res=>{
            res = res.body
            console.log("res: ",res)
            this.chooseList = res
            if (this.chooseList.length != 0) {
              this.showList = true
            } else {
              this.showList = false
              this.showError = true
            }
          }).catch(err=>{
             this.showError = true
          })
        }
      } else {
        alert('请选择日期、时间、容量！')
      }
    },
    getDateTime(date, time) {
      var arr = ['零', '一', '二', '三', '四', '五', '六', '七', '八', '九', '十', '十一', '十二', '十三', '十四', '十五']
      var str = date + ' ' + '第' + arr[this.compareTime(time)] + '节课'
      return str
    },
    currentDate () {
      var newdate = new Date()
      var year = newdate.getFullYear()
      var month = newdate.getMonth() + 1
      // 必须+1
      var day = newdate.getDate()
      month = this.addZero(month)
      day = this.addZero(day)
      var currentdate = year + '-' + month + '-' + day
      // console.log(currentdate)
      return currentdate
    },
    compareTime (time) {
      var str = ''
      if (time[2] != '') {
        str = time.substring(1, 3)
      } else {
        str = time[1]
      }
      return parseInt(str)
    },
    dateToStr (date) {
      var year = date.Year.toString()
      var month = date.Month
      month = this.addZero(month)
      var day = date.Day
      day = this.addZero(day)
      var str = year + '-' + month + '-' + day
      return str
    },
    addZero (number) {
      if (number >= 0 && number <= 9) {
        number = '0' + number
      }
      return number
    },
    applyClassroom (index,items) {
      this.info.indexer = index
      this.info.ClassroomId = items.ClassroomId
      this.info.ClassroomNum = items.ClassroomNum
      this.info.year = this.parseDateTime(this.date, 'year')
      this.info.month = this.parseDateTime(this.date, 'month')
      this.info.day = this.parseDateTime(this.date, 'day')
      this.showApply = true
    // console.log(this.info)
    },
    applySubmit (e) {
      // 检查
      if (document.getElementById('organization').value == ' ' || document.getElementById('reservationInfo').value == '') {
        alert('请填写使用方隶属组织和申请教室用途！')
      } else {
        // 后台交互 增加预订
        var apiStr = '/api/api/users/student/reservations'
        var obj = {};
        this.info.ResReason = this.reservationInfo
        this.info.OrganizationName = this.organization
        obj.StartTime = this.getDateTime(this.date, this.begin)
        obj.EndTime = this.getDateTime(this.date, this.end)
        obj.ResReason = this.reservationInfo
        obj.ClassroomId = parseInt(this.info.ClassroomId)
        obj.OrganizationName = this.organization

        this.$http.post(apiStr, obj).then(res=>{
          alert('submission succeeds!')
          this.chooseList.splice(this.info.indexer, 1)
          this.closeWindow()
          console.log(res)
          if (this.chooseList.length == 0) {
            this.showList = false
            this.showError = true
          }
        }).catch(err=>{
          alert('submission fails!')
          this.closeWindow()
        })
      }
      e.preventDefault()
    },
    closeWindow () {
      this.showApply = false
    },
    quit () {
      delCookie('StudentId')
      var apiStr = '/api/signout'
      this.$http.post(apiStr).then(res=>{
        console.log(res)
        alert('注销成功')
        this.$router.push('/')
      }).catch(err=>{
        alert('注销失败')
      })
    },
    reviseStudentPwd() {
      this.showStudentInfo = true
      var apiStr = '/api/api/users/student'
      this.$http.get(apiStr).then(res=>{
        console.log(res)
        res = res.body
        this.username = res.StudentName
      })
    },
    revisePwd() {
      var apiStr = '/api/api/users/student'
      this.$http.put(apiStr, {StudentPwd: this.password}).then(res=>{
        console.log(res)
        alert('修改密码成功！')
      }).catch(err=>{
        alert('修改密码失败！')
      })
      this.closeStudentInfo()
    },
    closeStudentInfo() {
      this.showStudentInfo = false
    }
  }
};
</script>

<style>

  .student{
    width: 100%;
  }
  .nav{
    position:relative;
    height:150px;
    width:900px;
    margin:0px auto;
  }
  div.nav .title{
    position: absolute;
    bottom: 0px;
    left: 0px;
    width: 400px;
    height: 120px;
    text-align: center;
    border: 1px solid rgba(187, 187, 187, 1);
  }
  div.nav img{
    display:block;
    float:left;
    margin-top:10px;
    margin-left: 20px;
    height:100px;
    width:100px;
  }
  div.nav .title h1{
    position: relative;
    left:50px;
    top:20px;
    width: 300px;
    height: 50px;
    line-height: 52px;
    color: rgba(4, 110, 97, 1);
    font-size: 30px;
    margin:0;
  }
  div.nav .title h3{
    position: relative;
    left:50px;
    top:30px;
    width: 300px;
    height: 30px;
    line-height: 26px;
    color: rgba(4, 110, 97, 1);
    font-size: 20px;
    margin:0;
  }

  div.nav .navigationBar {
    position: absolute;
    bottom: 0px;
    left: 480px;
    margin: 0px;
    height:50px;
  }
  .navigationBar li{
    float: left;
    display: block;
    line-height: 50px;
    height:50px;
    color:#FFF;
    padding: 0 15px;
    font-size:24px;
    text-align:center;
    text-decoration:none;
    background-color:#33B7A4;
    cursor: pointer;
  }
  li.activeTag {
    background-color: #046E5F;
  }
  div.nav .studentPwd {
    position: absolute;
    right: 5px;
    bottom: 30px;
  }
  div.nav .exit{
    position: absolute;
    bottom: 0px;
    right: 10px;
  }
  div.nav  a{
    text-decoration:none;
    display:block;
    font-size:16px;
    height: 20px;
    line-height:20px;
    color: rgba(4, 110, 97, 1);
  }
  div.line{
    position: absolute;
    bottom: 0;
    width:900px;
    border: 1px solid rgba(187, 187, 187, 1);
  }
  .content{
    position: relative;
    width:900px;
    margin:0px auto;
  }
  .content h1{
    width: 141px;
    height: 41px;
    margin:10px;
    line-height: 41px;
    color: rgba(16, 16, 16, 1);
    font-size: 20px;
    text-align: center;

  }
  .content #selectForm {
    width:100%;
    height:50px;
    margin-top: 10px;
  }

  .select {
    float:left;
    height:30px;
    margin:10px 50px;
  }

  .select img{
    width:20px;
    height:20px;
    vertical-align: middle;
  }
  .submit input{
    width:50px;
    height:30px;
    font-size: 12px;
    background-color:#1E9205;
    color:#fff;
    border-radius: 5px;
  }
  table {
    font-size:18px;
    color:#333;
    border-width:1px;
    border-color:#eee;
    border-collapse:collapse;
    margin:20px auto;
  }
  table th {
    border-width:1px;
    padding:5px 10px;
    border-style:solid;
    border-color:#eee;
    background-color:#dedede;
  }
  table td{
    border-width: 1px;
    padding:5px 10px ;
    border-style: solid;
    border-color:#eee;
    background-color:#fff;
  }

  .apply-window
  {
      position:absolute;
      top:30px;
      left:300px;
      width:300px;
      padding:20px 10px;
      text-align: center;
      border: 3px solid #eee;
      border-radius: 5px;
      background-color: #eee;
      box-shadow: 0 0 10px #000 inset;
  }
</style>
