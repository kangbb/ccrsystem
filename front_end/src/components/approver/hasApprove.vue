<template>
	<div class="HasApprove">
		<h1>已审批</h1>
		<router-link to="/approver">Approver Main Page</router-link>
		<router-link to="/approver/NotApprove">待审批</router-link>

		<h2>教室预订系统</h2>
		<h3>myClassroom.com</h3>

		<h5 v-show="showMsg">没有已审批信息</h5>

		<table v-if="showList">
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
			<tr v-for="(items,index) in this.chooseList">
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
</template>

<script>
	export default {
		name: 'HasApprove',
		data() {
			return {
				chooseList: [],
				showMsg: false,
				showList: false
			}
		},
		created() {
			this.$http.get('/api/student/reservation').then(res=>{
				res = res.body.data;
		//		console.log(res);
				for (var item in res) {
					var temp = res[item];
					if (temp.ReservationState != 0) {
						if (temp.ReservationState == 1) {
							temp.ReservationState = true;
						} else {
							temp.ReservationState = false;
						}
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