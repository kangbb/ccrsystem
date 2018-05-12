<template>
	<div class="Reservation">
		<router-link to="/student">Student Main Page</router-link>
		<router-link to="/student/query">教室状态查询</router-link>
		<h1>我的申请</h1>
		<table v-if="showList">
			<tr>
				<th>序列</th>
				<th>教室号</th>
				<th>日期</th>
				<th>时间</th>
				<th>容量</th>
				<th>操作</th>
			</tr>
			<tr v-for="(items,index) in this.reservation">
				<td>{{index+1}}</td>
				<td>{{items.ClassroomNum}}</td>
				<td>{{items.Date.Year}}<span>年</span>{{items.Date.Month}}<span>月</span>{{items.Date.Day}}<span>日</span></td>
				<td><span>第</span>{{items.Time[0]}}<span>节至第</span>{{items.Time[1]}}<span>节</span></td>
				<td>{{items.Capacity}}<span>人</span></td>
				<td><button type="text" v-on:click="deleteApply(items)">取消申请</button></td>
			</tr>
		</table>
	</div>
</template>

<script>
import {getCookie, delCookie} from '@/common/js/cookie.js'
	export default {
		name: 'Reservation',
		data() {
			return {
				userId: "",
				showList: false,
				reservation: []
			}
		},
		created() {
			let uname = getCookie('StudentId')
    		console.log('uname', uname)
    		this.userId = uname

    		//  读取了存进json里面的东西
    		//  保存草稿那个不太清楚
			this.$http.get('/api/student/reservation').then(res=>{
				res = res.body.data;
			//	console.log(res);
				for (var item in res) {
					var temp = res[item];
					if (temp.StudentId == this.userId) {
						this.reservation.push(temp);
					}
				}
				if (this.reservation.length != 0) {
					this.showList = true;
				}

			})
		},
		methods: {
			deleteApply:function(items) {
				this.reservation.splice(this.reservation.indexOf(items), 1);
			}
		}
	}
</script>

<style>
	
</style>