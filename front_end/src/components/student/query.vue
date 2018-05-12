<template>
	<div class="Query">
		<router-link to="/student">Student Main Page</router-link>
		<h1>教室状态查询</h1>
		<router-link to="/student/reservation">我的申请</router-link>

		<h2>教室预订系统</h2>
		<h3>myClassroom.com</h3>
		<h4>筛选条件：</h4>

		<form id="selectForm" v-on:submit="formSubmit">
			<input type="date" v-model="date" id="date">
			<div id="time">
				<select v-model="begin" id="begin">
					<option selected="selected" disabled="disabled">请选择</option>
					<option v-for="item in oneToTen">第{{item}}节</option>
				</select>
				<span>至</span>
				<select v-model="end" id="end">
					<option selected="selected" disabled="disabled">请选择</option>
					<option v-for="item in oneToTen">第{{item}}节</option>
				</select>
			</div>
			<select v-model="capacity" id="capacity">
				<option selected="selected" disabled="disabled">容量/人</option>
				<option value="50">50</option>
				<option value="100">100</option>
				<option value="200">200</option>
			</select>
			<input type="submit" value="筛选" id="submit">
		</form>

		<p v-if="showError">没有找到符合条件的教室<br/>尝试其他的筛选条件吧~</p>

		<table v-if="showList">
			<tr>
				<th>序列</th>
				<th>容量</th>
				<th>教室号</th>
				<th>操作</th>
			</tr>
			<tr v-for="(items,index) in this.chooseList">
				<td>{{index+1}}</td>
				<td>{{items.Capacity}}</td>
				<td>{{items.ClassroomNum}}</td>
				<td><button type="text" id="apply" v-on:click="applyClassroom(items.ClassroomNum)">申请</button></td>
			</tr>
		</table>

		<div id="applyWindow" v-show="showApply">
			<button v-on:click="closeWindow">x</button>
			<form v-on:submit="applySubmit">
				<p><span>教室号：</span>{{this.info.ClassroomNum}}</p>
				<p><span>申请日期：</span>{{this.info.year}}<span>年</span>{{this.info.month}}<span>月</span>{{this.info.day}}<span>日</span></p>
				<p><span>时间：</span>{{this.info.begin}}<span>至</span>{{this.info.end}}</p>
				<p><span>参与人数：</span>{{this.info.Capacity}}<span>人</span></p>
				<p>使用方隶属组织：<input type="text" placeholder="如，数据院15级软工二班" v-model="organization" id="organization"></p>
				<p>申请教室用途：<input type="text" placeholder="必须填写用途100字以内" v-model="reservationInfo" id="reservationInfo"></p>
				<input type="submit" value="提交申请">
			</form>
		</div>

	</div>
</template>

<script>
import {getCookie, delCookie} from '@/common/js/cookie.js'
	export default {
		name: 'Query',
		data() {
			return {
				userId: "",
				date: "",
				oneToTen: [1,2,3,4,5,6,7,8,9,10],
				classrooms: {},
				reservations: {},
				begin: "请选择",
				end: "请选择",
				capacity: "容量/人",
				chooseList: [],
				classroomNum: "",
				showError: false,
				showList: false,
				showApply: false,
				info: {},
				organization: "",
				reservationInfo: ""
			}
		},
		methods: {
			formSubmit:function(e) {
				e.preventDefault();
				if (this.date != "" && this.begin != "请选择" && this.end != "请选择" && this.capacity != "容量/人") {
					if (this.date < this.currentDate()) {
						alert("日期选择错误！");
					} else if (this.compareTime(this.begin) > this.compareTime(this.end)) {
						alert("时间选择错误！");
					} else {
						this.$http.get('/api/student/reservation').then(res=>{
							res = res.body.data;
							this.reservations = res;
							this.$http.get('/api/student/classroom').then(res=>{
								res = res.body.data;
								this.classrooms = res;
							//	console.log(this.reservations);
							//	console.log(this.classrooms);
								this.chooseList = this.chooseClassroom(this.reservations, this.classrooms);

								if (this.chooseList.length == 0) {
									this.showError = true;
								} else {
									this.showList = true;
								}
							})
						})
					}
				} else {
					alert("请选择日期、时间、容量！");
				}
			},
			currentDate:function() {
				var newdate = new Date();
				var year = newdate.getFullYear();
				var month = newdate.getMonth() + 1;  // 必须+1
				var day = newdate.getDate();
				month = this.addZero(month);
				day = this.addZero(day);
				var currentdate = year + "-" + month + "-" + day;
			//	console.log(currentdate);
				return currentdate;
			},
			compareTime:function(time) {
				var str = "";
				if (time[2] == "0") {
					str = time.substring(1, 3);
				} else {
					str = time[1];
				}
			//	console.log(parseInt(str));
				return parseInt(str);
			},
			chooseClassroom:function(reservations, classrooms) {
				var tempClassrooms, tempReservations;
				var chooseList=[];
				for (var item in classrooms) {
					tempClassrooms = classrooms[item];
					if (this.capacity == tempClassrooms.Capacity) {
						chooseList.push(tempClassrooms);
					}
				}

				for (var item in reservations) {
					var flag = false;
					tempReservations = reservations[item];
					if (this.capacity == tempReservations.Capacity) {
						if (this.dateToStr(tempReservations.Date) == this.date) {
							var tempTime = tempReservations.Time;
							if (!(this.compareTime(this.end) < tempTime[0] || this.compareTime(this.begin) > tempTime[1])) {
								flag = true;
							}
						}
					}
					if (flag) {
					//	console.log("0");
						for (var it in classrooms) {
							tempClassrooms = classrooms[it];
							if (tempReservations.ClassroomId == tempClassrooms.ClassroomId) {
								chooseList.splice(chooseList.indexOf(tempClassrooms), 1);
								break;
							}
						}
					}
				}
				var nodeTd = document.createElement("tr");

			//	console.log(chooseList);
				return chooseList;
			},
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
			applyClassroom:function(classroomNum) {
				this.info.StudentId = this.userId;
				for (var item in this.chooseList) {
					var temp = this.chooseList[item];
					if (temp.ClassroomNum == classroomNum) {
						this.info.ClassroomId = temp.ClassroomId;
						break;
					}
				}
				this.info.Capacity = this.capacity;
				this.info.ClassroomNum = classroomNum;
				this.info.year = this.date.substring(0,4);
				this.info.month = this.date.substring(5,7);
				this.info.day = this.date.substring(8,10);
				this.info.begin = this.begin;
				this.info.end = this.end;
				this.showApply = true;
			//	console.log(this.info);
			},
			applySubmit:function(e) {
				e.preventDefault();
				// 检查
				if (document.getElementById('organization').value == "" || document.getElementById('reservationInfo').value == "") {
					alert("请填写使用方隶属组织和申请教室用途！");
				}
				// 写入json 路径不知道要怎么搞
				this.showApply = false;
				for (var item in this.classrooms) {
					var temp = this.classrooms[item];
					if (this.info.ClassroomNum == temp.ClassroomNum) {
						this.chooseList.splice(this.chooseList.indexOf(temp), 1);
					}
				}
			},
			closeWindow:function() {
				this.showApply = false;
			}
		},
		created() {
			let uname = getCookie('StudentId')
    		console.log('uname', uname)
    		this.userId = uname
		}
	}
</script>

<style>
	
</style>