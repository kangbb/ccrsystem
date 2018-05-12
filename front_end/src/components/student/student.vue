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
                <td><button type='text' id='apply' v-on:click='applyClassroom(items.ClassroomNum)'>申请</button></td>
              </tr>
            </table>
          </div>

          <div id='applyWindow' v-show='showApply' class="apply-window">
            <!-- <button v-on:click='closeWindow'>x</button> -->
            <form v-on:submit='applySubmit'>
              <p><span>教室号：</span>{{this.info.ClassroomNum}}</p>
              <p><span>申请日期：</span>{{this.info.year}}<span>年</span>{{this.info.month}}<span>月</span>{{this.info.day}}<span>日</span></p>
              <p><span>时间：</span>{{this.info.begin}}<span>至</span>{{this.info.end}}</p>
              <p><span>参与人数：</span>{{this.info.Capacity}}<span>人</span></p>
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
              <td>{{items.Date.Year}}<span>年</span>{{items.Date.Month}}<span>月</span>{{items.Date.Day}}<span>日</span></td>
              <td><span>第</span>{{items.Time[0]}}<span>节至第</span>{{items.Time[1]}}<span>节</span></td>
              <td>{{items.Capacity}}<span>人</span></td>
              <td><button type="text" v-on:click="deleteApply(items)">取消申请</button></td>
            </tr>
          </table>
        </div>
      </div>
      <p v-if='showError'>没有找到符合条件的教室<br/>尝试其他的筛选条件吧~</p>
    </div>

  </div>
</template>

<script>
import {getCookie, delCookie} from '@/common/js/cookie.js'
export default {
  name: 'student',
  data () {
    return {
      userId: '',
      date: '',
      oneToTen: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
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
      reservation: []

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
    },
    toReservation () {
      this.showSearch = false
      this.showResvation = true

      this.reservation = []
       var apiStr = 'http://www.kangblog.top/api/users/student/reservations';
       this.$http.get(apiStr).then(res=>{
        console.log(res);
      })
  /*    this.$http.get('/api/student/reservation').then(res => {
        res = res.body.data
        //  console.log(res)
        for (var item in res) {
          var temp = res[item]
          if (temp.StudentId == this.userId) {
            this.reservation.push(temp)
          }
        }
        if (this.reservation.length != 0) {
          this.showResList = true
        }
      })*/
    },
    deleteApply (items) {
      // 后台交互 删除预订
      alert('确认取消该申请？')
      this.reservation.splice(this.reservation.indexOf(items), 1)
    },
    formSubmit (e) {
      e.preventDefault()
      if (this.date != '' && this.begin != '请选择' && this.end != '请选择' && this.capacity != '容量/人') {
        if (this.date < this.currentDate()) {
          alert('日期选择错误！')
        } else if (this.compareTime(this.begin) > this.compareTime(this.end)) {
          alert('时间选择错误！')
        } else {
          var beginTime = this.date + ' ' + this.begin;
          var endTime = this.date + ' ' + this.end;
          var apiStr = 'http://www.kangblog.top/api/student/reservations';
          this.$http.post(apiStr, {StartTime: beginTime, EndTime: endTime, ResReason: this.reservationInfo, ClassroomId: this.info.ClassroomId}).then(res=>{
            console.log(res);
          })
       /*   this.$http.get('/api/student/reservation').then(res => {
            res = res.body.data
            this.reservations = res
            this.$http.get('/api/student/classroom').then(res => {
              res = res.body.data
              this.classrooms = res
              console.log('res:' + this.reservations)
              console.log('class' + this.classrooms)
              this.chooseList = this.chooseClassroom(this.reservations, this.classrooms)

              if (this.chooseList.length == 0) {
                this.showError = true
              } else {
                this.showList = true
              }
            })
          })*/
        }
      } else {
        alert('请选择日期、时间、容量！')
      }
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
      if (time[2] == '0') {
        str = time.substring(1, 3)
      } else {
        str = time[1]
      }
      //  console.log(parseInt(str))
      return parseInt(str)
    },
    chooseClassroom (reservations, classrooms) {
      var tempClassrooms, tempReservations
      var chooseList = []
      for (var item in classrooms) {
        tempClassrooms = classrooms[item]
        if (this.capacity == tempClassrooms.Capacity) {
          chooseList.push(tempClassrooms)
        }
      }

      for (let item1 in reservations) {
        var flag = false
        tempReservations = reservations[item1]
        if (this.capacity == tempReservations.Capacity) {
          if (this.dateToStr(tempReservations.Date) == this.date) {
            var tempTime = tempReservations.Time
            if (!(this.compareTime(this.end) < tempTime[0] || this.compareTime(this.begin) > tempTime[1])) {
              flag = true
            }
          }
        }
        if (flag) {
        //  console.log('0')
          for (var it in classrooms) {
            tempClassrooms = classrooms[it]
            if (tempReservations.ClassroomId == tempClassrooms.ClassroomId) {
              chooseList.splice(chooseList.indexOf(tempClassrooms), 1)
              break
            }
          }
        }
      }
      // var nodeTd = document.createElement('tr')

      //  console.log(chooseList)
      return chooseList
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
    applyClassroom (classroomNum) {
      this.info.StudentId = this.userId
      for (var item in this.chooseList) {
        var temp = this.chooseList[item]
        if (temp.ClassroomNum == classroomNum) {
          this.info.ClassroomId = temp.ClassroomId
          break
        }
      }
      this.info.Capacity = this.capacity
      this.info.ClassroomNum = classroomNum
      this.info.year = this.date.substring(0, 4)
      this.info.month = this.date.substring(5, 7)
      this.info.day = this.date.substring(8, 10)
      this.info.begin = this.begin
      this.info.end = this.end
      this.showApply = true
    // console.log(this.info)
    },
    applySubmit (e) {
      e.preventDefault()
      this.organization = ''
      this.reservationInfo = ''
      // 检查
      if (document.getElementById('organization').value == ' ' || document.getElementById('reservationInfo').value == '') {
        alert('请填写使用方隶属组织和申请教室用途！')
      } else {
        // 后台交互 增加预订
        console.log('submit succeed!')
        this.showApply = false
        // this.init()
      }
      // 写入json 路径不知道要怎么搞
    },
    init () {
      // this.showList = false
    },
    closeWindow () {
      this.showApply = false
    },
    quit () {
      delCookie('StudentId')
      this.$router.push('/')
    }
  }
}
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
  }
  li.activeTag {
    background-color: #046E5F;
  }
  div.nav .exit{
    position: absolute;
    bottom: 0px;
    right: 10px;
  }
  div.nav  a{
    text-decoration:none;
    diplay:block;
    font-size:16px;
    height: 20px;
    line-height:20px;
    color: rgba(4, 110, 97, 1);
  }
  div.line{
    position: absolute;
    bottom: 0;
    length:0;
    width:900px;
    border: 1px solid rgba(187, 187, 187, 1)
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

  //弹框
  .apply-window
  {
      position:absolute;
      top:30px;
      left:300px;
      width:300px;
      padding:20px 10px;
      // margin:100px auto;
      text-align: center;
      border: 3px solid #eee;
      border-radius: 5px;
      background-color: #eee;
      box-shadow: 0 0 10px #000 inset;
  }
</style>
