<template>
  <div class="approver">

    <div class="nav">
      <div class="title">
        <img src="../../images/search.png">
        <h1>教室预订系统</h1>
        <h3 class="username">审批人 {{userId}}</h3>
      </div>
      <ul class="navigationBar">
        <li @click="toWaitApprove" :class="{'activeTag':showWaitApprove}">待审批</li>
        <li @click="toAlreadyApprove" :class="{'activeTag':showAlreadyApprove}">已审批</li>
      </ul>

      <div class="exit">
         <a href='#' @click='quit'>退出登录</a>
      </div>
      <div class="line"></div>
    </div>

    <div class="content">
      <div class="panel">
        <div class="wait-approve" v-show="showWaitApprove">
          <h1 v-show="showMsg">没有已审批信息</h1>
          <table v-if="showList1">
            <tr>
              <th>序列</th>
              <th>教室号</th>
              <th>容量</th>
              <th>操作</th>
            </tr>
            <tr v-for="(items,index) in this.waitApproveList" v-bind:key="index">
              <td>{{index+1}}</td>
              <td>{{items.ClassroomNum}}</td>
              <td>{{items.Capacity}}</td>
              <td><button type="text" id="apply" v-on:click="detail(items.ClassroomNum)">详情</button></td>
            </tr>
          </table>

          <div id="detailWindow" v-show="showDetail" class="apply-window">
            <!-- <button v-on:click="closeWindow">x</button> -->
            <form>
              <p><span>申请人：</span>{{this.info.StudentId}}</p>
              <p><span>教室号：</span>{{this.info.ClassroomNum}}</p>
              <p><span>申请日期：</span>{{this.info.year}}<span>年</span>{{this.info.month}}<span>月</span>{{this.info.day}}<span>日</span></p>
              <p><span>时间：</span>{{this.info.begin}}<span>至</span>{{this.info.end}}</p>
              <p><span>参与人数：</span>{{this.info.Capacity}}<span>人</span></p>
              <p><span>使用方隶属组织：</span>{{this.info.Organization}}</p>
              <p><span>申请教室用途：</span>{{this.info.ReservationInfo}}</p>
              <!-- <input type="button" value="通过审批" v-on:click="passApprove()">
              <input type="button" value="拒绝审批" v-on:click="failApprove()"> -->
            </form>
            <button v-on:click="passApprove">通过审批</button>
            <button v-on:click="failApprove">拒绝审批</button>
            <button v-on:click="closeWindow">关闭</button>
          </div>
        </div>

        <div class="alrealy-approve" v-show="showAlreadyApprove">
          <table v-if="showList2">
            <tr>
              <th>序列</th>
              <th>申请人</th>
              <th>教室号</th>
              <th>日期</th>
              <th>时间</th>
              <th>容量</th>
              <th>使用方隶属组织</th>
              <th>申请教室用途</th>
              <th>操作</th>
            </tr>
            <tr v-for="(items,index) in this.alreadyApproveList" v-bind:key="index">
              <td>{{index+1}}</td>
              <td>{{items.StudentId}}</td>
              <td>{{items.ClassroomNum}}</td>
              <td>{{items.Date.Year}}<span>年</span>{{items.Date.Month}}<span>月</span>{{items.Date.Day}}<span>日</span></td>
              <td><span>第</span>{{items.Time[0]}}<span>节至第</span>{{items.Time[1]}}<span>节</span></td>
              <td>{{items.Capacity}}<span>人</span></td>
              <td>{{items.Organization}}</td>
              <td>{{items.ReservationInfo}}</td>
              <td v-if="items.ReservationState">审核通过</td>
              <td v-else>审核不通过</td>
            </tr>
          </table>
        </div>
      </div>
    </div>

  </div>
</template>

<script type="text/ecmascript-6">
import {getCookie, delCookie} from '@/common/js/cookie.js'
export default {
  name: 'approver',
  data () {
    return {
      userId: '',
      waitApproveList: [],
      alreadyApproveList: [],
      showWaitApprove: true,
      showAlreadyApprove: false,
      showMsg: false,
      showList1: false,
      showList2: false,
      info: {},
      showDetail: false
    }
  },
  created () {
    let uname = getCookie('ApproverId')
    console.log('uname', uname)
    this.userId = uname
    if (uname === '') {
      this.$router.push('/')
    }
    this.toWaitApprove()
  },
  methods: {
    toWaitApprove () {
      this.showWaitApprove = true
      this.showAlreadyApprove = false
      this.showList1 = true

      this.waitApproveList = []
      var apiStr = 'http://www.kangblog.top/api/users/student/reservations';
      this.$http.get(apiStr).then(res=>{
        console.log(res);
      })
 /*     this.$http.get('/api/student/reservation').then(res => {
        res = res.body.data
        //  console.log(res);
        for (var item in res) {
          var temp = res[item]
          if (temp.ReservationState == 0) {
            this.waitApproveList.push(temp)
          }
        }
        console.log('wait to re:' + this.alreadyApproveList)
        if (this.waitApproveList.length == 0) {
          this.showMsg = true
        } else {
          this.showList2 = true
        }
      })*/
    },
    toAlreadyApprove () {
      this.showWaitApprove = false
      this.showAlreadyApprove = true

      this.showList2 = true
      this.alreadyApproveList = []

      var apiStr ='http://www.kangblog.top/api/users/student/reservations'
      this.$http.get(apiStr).then(res=>{
        console.log(res);
      })
 /*     this.$http.get('/api/student/reservation').then(res => {
        res = res.body.data
        // console.log(res);
        for (var item in res) {
          var temp = res[item]
          if (temp.ReservationState != 0) {
            if (temp.ReservationState == 1) {
              temp.ReservationState = true
            } else {
              temp.ReservationState = false
            }
            this.alreadyApproveList.push(temp)
          }
        }
        if (this.alreadyApproveList.length == 0) {
          this.showMsg = true
        } else {
          this.showList1 = true
        }
      })*/
      console.log('already to re:' + this.alreadyApproveList)
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
    detail (classroomNum) {
      this.info.ApproverId = this.userId
      this.info.ClassroomNum = classroomNum
      for (var item in this.waitApproveList) {
        var temp = this.waitApproveList[item]
        if (temp.ClassroomNum == classroomNum) {
          this.info.StudentId = temp.StudentId
          this.info.ClassroomId = temp.ClassroomId
          this.info.Organization = temp.Organization
          this.info.ReservationInfo = temp.ReservationInfo
          this.info.Capacity = temp.Capacity

          var date = this.dateToStr(temp.Date)
          this.info.year = date.substring(0, 4)
          this.info.month = date.substring(5, 7)
          this.info.day = date.substring(8, 10)

          this.info.begin = temp.Time[0].toString()
          this.info.end = temp.Time[1].toString()
          break
        }
      }
      this.showDetail = true
    },
    closeWindow () {
      this.showDetail = false
    },
    passApprove () {
      this.info.ReservationState = 1
      for (var item in this.waitApproveList) {
        var temp = this.waitApproveList[item]
        if (temp.ClassroomNum == this.info.ClassroomNum) {
          this.waitApproveList.splice(this.waitApproveList.indexOf(temp), 1)
        }
      }
      this.showDetail = false

      // 与后台交互,更新该预订的状态
    },
    failApprove () {
      this.info.ReservationState = 2
      for (var item in this.waitApproveList) {
        var temp = this.waitApproveList[item]
        if (temp.ClassroomNum == this.info.ClassroomNum) {
          this.waitApproveList.splice(this.waitApproveList.indexOf(temp), 1)
        }
      }
      this.showDetail = false
      //与后台交互,更新该预订的状态
    },
    quit () {
      delCookie('ApproverId')
      this.$router.push('/')
    }
  }
}
</script>

<style>
 .approver{
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
