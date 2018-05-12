<template>
	<div class="NotApprove">
		<h1>待审批</h1>
		<router-link to="/approver">Approver Main Page</router-link>
		<router-link to="/approver/HasApprove">已审批</router-link>

		<h2>教室预订系统</h2>
		<h3>myClassroom.com</h3>
		<h5 v-show="showMsg">没有待审批信息</h5>

		<table v-if="showList">
			<tr>
				<th>序列</th>
				<th>教室号</th>
				<th>容量</th>
				<th>操作</th>
			</tr>
			<tr v-for="(items,index) in this.chooseList">
				<td>{{index+1}}</td>
				<td>{{items.ClassroomNum}}</td>
				<td>{{items.Capacity}}</td>
				<td><button type="text" id="apply" v-on:click="detail(items.ClassroomNum)">详情</button></td>
			</tr>
		</table>

		<div id="detailWindow" v-show="showDetail">
			<button v-on:click="closeWindow">x</button>
			<form>
				<p><span>申请人：</span>{{this.info.StudentId}}</p>
				<p><span>教室号：</span>{{this.info.ClassroomNum}}</p>
				<p><span>申请日期：</span>{{this.info.year}}<span>年</span>{{this.info.month}}<span>月</span>{{this.info.day}}<span>日</span></p>
				<p><span>时间：</span>{{this.info.begin}}<span>至</span>{{this.info.end}}</p>
				<p><span>参与人数：</span>{{this.info.Capacity}}<span>人</span></p>
				<p><span>使用方隶属组织：</span>{{this.info.Organization}}</p>
				<p><span>申请教室用途：</span>{{this.info.ReservationInfo}}</p>
				<input type="button" value="通过审批" v-on:click="passApprove()">
				<input type="button" value="拒绝审批" v-on:click="failApprove()">
			</form>
		</div>

	</div>
</template>

<script>
import {getCookie, delCookie} from '@/common/js/cookie.js'
	export default {
		name: 'NotApprove',
		data() {
			return {
				userId: "",
				chooseList: [],
				showList: false,
				info: {},
				showDetail: false,
				showMsg: false
			}
		},
		methods: {
			dateToStr:function(date) {
				var year = date.Year.toString();
				var month = date.Month;
				month = this.addZero(month);
				var day = date.Day;
				day = this.addZero(day);
				var str = year + "-" + month + "-" + day;
				return str;
			},
			addZero:function(number) {
				if (number >= 0 && number <= 9) {
					number = "0" + number;
				}
				return number;
			},
			detail:function(classroomNum) {
				this.info.ApproverId = this.userId;
				this.info.ClassroomNum = classroomNum;
				for (var item in this.chooseList) {
					var temp = this.chooseList[item];
					if (temp.ClassroomNum == classroomNum) {
						this.info.StudentId = temp.StudentId;
						this.info.ClassroomId = temp.ClassroomId;
						this.info.Organization = temp.Organization;
						this.info.ReservationInfo = temp.ReservationInfo;
						this.info.Capacity = temp.Capacity;

						var date = this.dateToStr(temp.Date);
						this.info.year = date.substring(0,4);
						this.info.month = date.substring(5,7);
						this.info.day = date.substring(8,10);

						this.info.begin = temp.Time[0].toString();
						this.info.end = temp.Time[1].toString();
						break;
					}
				}
				this.showDetail = true;
			},
			closeWindow:function() {
				this.showDetail = false;
			},
			passApprove:function() {
				this.info.ReservationState = 1;
				for (var item in this.chooseList) {
					var temp = this.chooseList[item];
					if (temp.ClassroomNum == this.info.ClassroomNum) {
						this.chooseList.splice(this.chooseList.indexOf(temp),1);
					}
				}
				this.showDetail = false;
				// 写入
			},
			failApprove:function() {
				this.info.ReservationState = 2;
				// 写入
			}

		},
		created() {
			let uname = getCookie('ApproverId')
    		console.log('uname', uname)
    		this.userId = uname

    		this.$http.get('/api/student/reservation').then(res=>{
    			res = res.body.data;
			//	console.log(res);
				for (var item in res) {
					var temp = res[item];
					if (temp.ReservationState == 0) {
						this.chooseList.push(temp);
					}
				}

				if (this.chooseList.length == 0) {
					this.showMsg = true;
				} else {
					this.showList = true;
				}
    		})
		}
	}
</script>

<style>
	
</style>