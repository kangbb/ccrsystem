<template>
  <div class="admin">
    <div class="nav">
      <div class="title">
        <img src = "../../images/search.png">
        <h1>教室预订系统</h1>
        <h3 class="username">管理 {{userId}}</h3>
      </div>
      <ul class="navigationBar">
        <li @click="classroomTag" :class="{'activeTag':showClassroom}">教室管理</li>
        <li @click="studentTag" :class="{'activeTag':showStudent}">学生管理</li>
        <li @click="approverTag" :class="{'activeTag':showApprover}">审核管理</li>
      </ul>
      <div class="exit">
        <a href="#" @click="quit">退出登录</a>
      </div>
      <div class="line"></div>
    </div>

    <div class="content">
      <div class="panel">

        <!-- 教室管理 -->
        <div class="classroom" v-show="showClassroom">
          <div class="fullClassList">
            <h1>筛选条件：</h1>
            <form id="selectForm" v-on:submit="filterClass">
              <div class="capacity select">
                <img src="../../images/user.png">
                <select v-model="capacity">
                  <option disabled="disabled" selected="selected">容量/人</option>
                  <option value='50'>50</option>
                  <option value='100'>100</option>
                  <option value='200'>200</option> 
                </select>
              </div>
              <div class="submit select"><input type="submit" value="筛选"></div>
            </form>
            <table class="tableClass" v-if="showClass">
              <tr>
                <th>序列</th>
                <th>容量</th>
                <th>教室号</th>
                <th>操作</th>
              </tr>
              <tr v-for='(items,index) in this.tempClass' v-bind:key="index">
                <td>{{index+1}}</td>
                <td>{{items.Capacity}}</td>
                <td>{{items.ClassroomNum}}</td>
                <td><button @click="reviseClass(index, items)">修改</button></td>
              </tr>
            </table>
          </div>

          <div id="reviseClassroom" v-show="classRevision" class="reviseWin">
            <form>
              <p>园区：东校园</p>
              <p>楼区：公教楼</p>
              <p>教室号：{{classNum}}</p>
              <p>教室容量：<input type="text" v-model="classCapacity"></p>
            </form>
            <button @click="classReviseSubmit">提交</button>
            <button @click="classReviseCancel">取消</button>
            <button @click="classDeleteSubmit">删除</button>
          </div>

          <button @click="displayAllClass" class="displayer">查看全部课室</button>
          <button @click="addClassroom" class="displayer" id="add">增加教室</button>
          <div id="classAdd" v-show="classroomAdd" class="reviseWin">
            <form>
              <p>园区：东校园</p>
              <p>楼区：公教楼</p>
              <p>教室号：<input type="text" v-model="addClassNum"></p>
              <p>容量：<input type="text" v-model="addCapacity"></p>
            </form>
            <button @click="addClassSubmit">提交</button>
            <button @click="closeClassSubmit">取消</button>
          </div>

        </div>

        <!-- 学生管理 -->
        <div class="student" v-show="showStudent">
          <div class="fullStudentList">
            <h1>筛选条件：</h1>
            
            <form id="selectForm" v-on:submit="filterStudent">
              <div class="grade select">
                <img src="../../images/user.png">
                <select v-model="grade">
                  <option disabled="disabled" selected="selected">年级</option>
                  <option value="14">14</option>
                  <option value="15">15</option>
                  <option value="16">16</option>
                  <option value="17">17</option> 
                </select>
              </div>
              <div class="submit select"><input type="submit" value="筛选"></div>
            </form>
            <table class="tableClass" v-if="showStu">
              <tr>
                <th>序列</th>
                <th>工号</th>
                <th>姓名</th>
                <th>操作</th>
              </tr>
              <tr v-for='(items,index) in this.tempStu' v-bind:key="index">
                <td>{{index+1}}</td>
                <td>{{items.StudentId}}</td>
                <td>{{items.StudentName}}</td>
                <td><button @click="reviseStu(index, items)">修改</button></td>
              </tr>
            </table>
          </div>

          <div id="reviseStudent" v-show="stuRevision" class="reviseWin">
            <form>
              <p>工号：{{stuId}}</p>
              <p>姓名：{{stuName}}</p>
              <p>密码：<input type="text" v-model="stuPwd"></p>
            </form>
            <button @click="stuReviseSubmit">提交</button>
            <button @click="stuReviseCancel">取消</button>
            <button @click="stuDeleteSubmit">删除</button>
          </div>
          
          <button @click="displayAllStudent" class="displayer">查看全部学生</button>
          <button @click="addStudent" class="displayer" id="add">增加学生</button>
          <div id="stuAdd" v-show="studentAdd" class="reviseWin">
            <form>
              <p>工号：<input type="text" v-model="addStuId"></p>
              <p>姓名：<input type="text" v-model="addStuName"></p>
              <p>密码：初始化默认123456</p>
            </form>
            <button @click="addStuSubmit">提交</button>
            <button @click="closeStuSubmit">取消</button>
          </div>

        </div>

        <!-- 审核管理 -->
        <div class="approver" v-show="showApprover">
          <div class="fullApproverList">
            <h1>全部审核员：</h1>
            <table class="tableClass" v-if="showApp">
              <tr>
                <th>序列</th>
                <th>工号</th>
                <th>姓名</th>
                <th>审批部门号</th>
                <th>操作</th>
              </tr>
              <tr v-for='(items,index) in this.tempApp' v-bind:key="index">
                <td>{{index+1}}</td>
                <td>{{items.ApproverId}}</td>
                <td>{{items.ApproverName}}</td>
                <td>{{items.DepartmentId}}</td>
                <td><button @click="reviseApp(index, items)">修改</button></td>
              </tr>
            </table>
          </div>

          <div id="reviseApprover" v-show="appRevision" class="reviseWin">
            <form>
              <p>工号：{{appId}}</p>
              <p>姓名：{{appName}}</p>
              <p>审批部门号：{{departmentId}}</p>
              <p>密码：<input type="text" v-model="appPwd"></p>
            </form>
            <button @click="appReviseSubmit">提交</button>
            <button @click="appReviseCancel">取消</button>
            <button @click="appDeleteSubmit">删除</button>
          </div>

          <button @click="addApprover" class="displayer" id="add">增加审核员</button>
          <div id="appAdd" v-show="approverAdd" class="reviseWin">
            <form>
              <p>工号：<input type="text" v-model="addAppId"></p>
              <p>姓名：<input type="text" v-model="addAppName"></p>
              <p>审批部门号：<input type="text" v-model="addDepId"></p>
              <p>密码：初始化默认123456</p>
            </form>
            <button @click="addAppSubmit">提交</button>
            <button @click="closeAppSubmit">取消</button>
          </div>

        </div>


      </div>
  
      <p v-show="showError">没有找到符合条件的内容</p>

    </div>
  </div>
</template>

<script>
  import {getCookie, delCookie} from '@/common/js/cookie.js'
  export default {
    name: 'admin',
    data() {
      return {
        userId: '',
        username: '',

        capacity: '',
        showClassroom: true,
        showStudent: false,
        showApprover: false,
        showClass: false,
        fullClass: [],
        tempClass: [],
        classRevision: false,
        classId: '',
        classNum: '',
        classCapacity: '',
        classIndexer: '',
        classroomAdd: false,
        addClassNum: '',
        addCapacity: '',

        grade: '',
        showStu: false,
        fullStu: [],
        tempStu: [],
        stuRevision: false,
        stuName: '',
        stuId: '',
        stuPwd: '',
        studentAdd: false,
        addStuId: '',
        addStuName: '',
        stuIndexer: '',

        showApp: false,
        fullApp: [],
        tempApp: [],
        appIndexer: '',
        appRevision: false,
        appId: '',
        appName: '',
        appPwd: '',
        approverAdd: false,
        addAppId: '',
        addAppName: '',
        departmentId: '',
        addDepId: '',

        showError: false
      }
    },
    created() {
      let uname = getCookie('AdminId')
      console.log('uname', uname)
      this.userId = uname
      if (uname === '') {
        this.$router.push('/')
      }
      this.displayAllClass()
    },
    methods: {
      quit() {
        delCookie('AdminId')
        var apiStr = '/api/signout'
        this.$http.post(apiStr).then(res=>{
          console.log(res)
          alert('注销成功')
          this.$router.push('/')
        }).catch(err=>{
          alert('注销失败')
        })
      },
      classroomTag() {
      //  this.getElementById('shouldSel').setAttribute("selected", "selected")
        this.showClassroom = true
        this.showStudent = false
        this.showApprover = false
        this.showError = false
        this.displayAllClass()
      },
      studentTag() {
      //  this.getElementById('shouldSel').setAttribute("selected", "selected")
        this.showClassroom = false
        this.showStudent = true
        this.showApprover = false
        this.showError = false
        this.displayAllStudent()
      },
      approverTag() {
      //  this.getElementById('shouldSel').setAttribute("selected", "selected")
        this.showClassroom = false
        this.showStudent = false
        this.showApprover = true
        this.showError = false
        this.displayAllApprover()
      },
      displayAllClass() {
        var apiStr = '/api/classrooms'
        this.$http.get(apiStr).then(res=>{
          res = res.body
      //    console.log(res)
          this.fullClass = res
          this.tempClass = this.fullClass
          this.showClass = true
        }).catch(err=>{
          this.showError = true
        })
      },
      filterClass(e) {
        e.preventDefault()
        if (this.capacity == '') {
          alert('请选择容量！')
        } else {
          this.tempClass = []
          for (var items in this.fullClass) {
            var temp = this.fullClass[items]
            if (temp.Capacity == parseInt(this.capacity)) {
            //  console.log(this.tempClass)
              this.tempClass.push(temp)
            }
          }
          if (this.tempClass.length == 0) {
            this.showError = true
          }
        }
      },
      reviseClass(index, items) {
        this.classIndexer = index
        this.classId = items.ClassroomId
        this.classNum = items.ClassroomNum
        this.classCapacity = items.Capacity
        this.classRevision = true
      },
      classReviseSubmit() {
        if (this.classCapacity != '') {
          if (confirm('确认修改')) {
            var cap = parseInt(this.classCapacity)
            var apiStr = '/api/classrooms/' + this.classId
            this.$http.put(apiStr, {ClassroomCampus: '东校园', ClassroomBuilding: '公教楼', Capacity: cap, ClassroomNum: this.classNum}).then(res=>{
            //  console.log(res)
              alert('revision succeeds!')
            }).catch(err=>{
              alert('revision fails!')
            })
          }
        } else {
          alert('请输入容量！')
        }
        this.classReviseCancel()
      },
      classReviseCancel() {
        this.classRevision = false
      },
      classDeleteSubmit() {
        if (confirm('确认删除')) {
          var idx = parseInt(this.classId)
          var apiStr = '/api/classrooms/' + idx
          this.$http.delete(apiStr, {body:{Id: idx}}).then(res=>{
       //     console.log(res)
            this.tempClass.splice(parseInt(this.classIndexer), 1)
            alert('删除成功！')
          }).catch(err=>{
            alert('删除失败！')
          })
        }
        this.classReviseCancel()
      },
      addClassroom() {
        this.classroomAdd = true
      },
      addClassSubmit() {
        if (this.addClassNum != '' && this.addCapacity != '') {
          if (confirm('确认增加')) {
            var apiStr = '/api/classrooms'
            var obj = {}
            obj.ClassroomCampus = '东校园'
            obj.ClassroomBuilding = '公教楼'
            obj.ClassroomNum = this.addClassNum
            obj.Capacity = parseInt(this.addCapacity)

            this.$http.post(apiStr, obj).then(res=>{
          //    console.log(res)
              alert('成功添加！')
              this.tempClass.push(obj)
            }).catch(err=>{
              alert('添加失败！')
            })
          }
          this.addClassNum = ''
          this.addCapacity = ''
        } else {
          alert('请输入教室号和容量！')
        }
        this.closeClassSubmit()
      },
      closeClassSubmit() {
        this.classroomAdd = false
      },
      displayAllStudent() {
        var apiStr = '/api/users/students'
        this.$http.get(apiStr).then(res=>{
          res = res.body
       //   console.log(res)
          this.fullStu = res
          this.tempStu = this.fullStu
          this.showStu = true
        }).catch(err=>{
          this.showError = true
        })
      },
      filterStudent(e) {
        e.preventDefault()
        if (this.grade == '') {
          alert('请选择年级！')
        } else {
          this.tempStu = []
          for (var items in this.fullStu) {
            var temp = this.fullStu[items]
            var str = temp.StudentId.toString()
            if (str.substring(0, 2) == this.grade) {
            //  console.log(this.tempClass)
              this.tempStu.push(temp)
            }
          }
          if (this.tempStu.length == 0) {
            this.showError = true
          }
        }
      },
      reviseStu(index, items) {
        this.stuIndexer = index
        this.stuName = items.StudentName
        this.stuId = items.StudentId
        this.stuRevision = true
      },
      stuReviseSubmit() {
        if (this.stuPwd != '') {
          if (confirm('确认修改')) {
            var idx = parseInt(this.stuId)
            var apiStr = '/api/users/students/' + idx
            this.$http.put(apiStr, {StudentPwd: this.stuPwd, Id: idx}).then(res=>{
         //     console.log(res)
              alert('revision succeeds!')
            }).catch(err=>{
              alert('revision fails!')
            })
          }
        } else {
          alert('请输入密码！')
        }
        this.stuReviseCancel()
        this.stuPwd = ''
      },
      stuReviseCancel() {
        this.stuRevision = false
      },
      stuDeleteSubmit() {
        if (confirm('确认删除')) {
          var idx = parseInt(this.stuId)
          var apiStr = '/api/users/students/' + this.stuId
          this.$http.delete(apiStr, {body:{Id: idx}}).then(res=>{
      //      console.log(res)
            this.tempStu.splice(parseInt(this.stuIndexer), 1)
            alert('删除成功！')
          }).catch(err=>{
            alert('删除失败！')
          })
        }
        this.stuReviseCancel()
      },
      addStudent() {
        this.studentAdd = true
      },
      addStuSubmit() {
        if (this.addStuName != '' && this.addStuId != '') {
          if (confirm('确认增加')) {
            var apiStr = '/api/users/students'
            var obj = {}
            obj.StudentId = parseInt(this.addStuId)
          //  console.log(obj.StudentId)
            obj.StudentName = this.addStuName

            this.$http.post(apiStr, obj).then(res=>{
          //    console.log(res)
              alert('成功添加！')
              this.tempStu.push(obj)
            }).catch(err=>{
              alert('添加失败！')
            })
          }
          this.addStuName = ''
          this.addStuId = ''
        } else {
          alert('请输入工号和姓名！')
        }
        this.closeStuSubmit()
      },
      closeStuSubmit() {
        this.studentAdd = false
      },
      reviseApp(index, items) {
        this.appIndexer = index
        this.appName = items.ApproverName
        this.appId = items.ApproverId
        this.departmentId = items.DepartmentId
        this.appRevision = true
      },
      appReviseSubmit() {
        if (this.appPwd != '') {
          if (confirm('确认修改')) {
            var idx = parseInt(this.appId)
            var apiStr = '/api/users/approvers/' + idx
            this.$http.put(apiStr, {ApproverPwd: this.appPwd, Id: idx}).then(res=>{
           //   console.log(res)
              alert('revision succeeds!')
            }).catch(err=>{
              alert('revision fails!')
            })
          }
        } else {
          alert('请输入密码！')
        }
        this.appReviseCancel()
        this.appPwd = ''
      },
      appReviseCancel() {
        this.appRevision = false
      },
      appDeleteSubmit() {
        if (confirm('确认删除')) {
          var idx = parseInt(this.appId)
          var apiStr = '/api/users/approvers/' + this.appId
          console.log(idx)
          this.$http.delete(apiStr, {body:{Id: idx}}).then(res=>{
         //   console.log(res)
            this.tempApp.splice(parseInt(this.appIndexer), 1)
            alert('删除成功！')
          }).catch(err=>{
            alert('删除失败！')
          })
        }
        this.appReviseCancel()
      },
      addApprover() {
        this.approverAdd = true
      },
      addAppSubmit() {
        if (this.addAppName != '' && this.addAppId != '' && this.addDepId != '') {
          if (confirm('确认增加')) {
            var apiStr = '/api/users/approvers'
            var obj = {}
            obj.ApproverId = parseInt(this.addAppId)
          //  console.log(obj.ApproverId)
            obj.ApproverName = this.addAppName
            obj.DepartmentId = parseInt(this.addDepId)

            this.$http.post(apiStr, obj).then(res=>{
          //    console.log(res)
              alert('成功添加！')
              this.tempApp.push(obj)
            }).catch(err=>{
              alert('添加失败！')
            })
          }
          this.addDepId = ''
          this.addAppName = ''
          this.addAppId = ''
        } else {
          alert('请输入工号、姓名、审批部门号！')
        }
        this.closeAppSubmit()
      },
      closeAppSubmit() {
        this.approverAdd = false
      },
      displayAllApprover() {
        var apiStr = '/api/users/approvers'
        this.$http.get(apiStr).then(res=>{
          res = res.body
       //   console.log(res)
          this.fullApp = res
          this.tempApp = this.fullApp
          this.showApp = true
        }).catch(err=>{
          this.showError = true
        })
      }
    }
  };
</script>

<style>
  .admin {
    width: 100%;
  }
  .nav {
    position: relative;
    height: 150px;
    width: 1000px;
    margin: 0px auto;
  }
  div.nav .title {
    position: absolute;
    bottom: 0px;
    left: 0px;
    width: 400px;
    height: 120px;
    text-align: center;
    border: 1px solid rgba(187, 187, 187, 1);
  }
  div.nav img {
    display: block;
    float: left;
    margin-top: 10px;
    margin-left: 20px;
    height: 100px;
    width: 100px;
  }
  div.nav .title h1 {
    position: relative;
    left: 50px;
    top: 20px;
    width: 300px;
    height: 50px;
    line-height: 52px;
    color: rgba(4, 110, 97, 1);
    font-size: 30px;
    margin: 0;
  }
  div.nav .title h3 {
    position: relative;
    left: 50px;
    top: 30px;
    width: 300px;
    height: 30px;
    line-height: 26px;
    color: rgba(4, 110, 97, 1);
    font-size: 20px;
    margin: 0;
  }
  div.nav .navigationBar {
    position: absolute;
    bottom: 0px;
    left: 480px;
    margin: 0px;
    height: 50px;
  }
  .navigationBar li {
    float: left;
    display: block;
    line-height: 50px;
    height: 50px;
    color: #FFF;
    padding: 0 15px;
    font-size: 24px;
    text-align: center;
    text-decoration: none;
    background-color: #33B7A4;
    cursor: pointer;
  }
  li.activeTag {
    background-color: #046E5F;
  }
  div.nav .adminPwd {
    position: absolute;
    right: 5px;
    bottom: 30px;
  }
  div.nav .exit {
    position: absolute;
    bottom: 0px;
    right: 10px;
  }
  div.nav a {
    text-decoration: none;
    display: block;
    font-size: 16px;
    height: 20px;
    line-height: 20px;
    color: rgba(4, 110, 97, 1);
  }
  div.line {
    position: absolute;
    bottom: 0;
    width: 1000px;
    border: 1px solid rgba(187, 187, 187, 1);
  }
  .content {
    position: relative;
    width: 900px;
    margin: 0px auto;
  }
  .content h1 {
    width: 141px;
    height: 41px;
    margin: 10px;
    line-height: 41px;
    color: rgba(16, 16, 16, 1);
    font-size: 20px;
    text-align: center;
  }
  .content #selectForm {
    width: 100%;
    height: 50px;
    margin: 5px;
  }
  .select {
    float:left;
    height:30px;
    margin:10px 50px;
  }
  .select img {
    width: 20px;
    height: 20px;
    vertical-align: middle;
  }
  .submit input {
    width: 50px;
    height: 30px;
    font-size: 12px;
    background-color: #1E9205;
    color: #fff;
    border-radius: 5px;
  }
  table {
    font-size: 18px;
    color: #333;
    border-width: 1px;
    border-color: #eee;
    border-collapse: collapse;;
    margin: 20px auto;
  }
  table th {
    border-width: 1px;
    padding: 5px 10px;
    border-style: solid;
    border-color: #eee;
    background-color: #dedede;
  }
  table td {
    border-width: 1px;
    padding: 5px 10px;
    border-style: solid;
    border-color: #eee;
    background-color: #fff;
  }
  .reviseWin {
    position: absolute;
    top: 50px;
    left: 300px;
    width: 300px;
    padding: 20px 10px;
    text-align: center;
    border: 3px solid #eee;
    border-radius: 5px;
    background-color: #eee;
    box-shadow: 0 0 10px #000 inset;
  }
  .displayer {
    float: right;
    width: 100px;
    height: 30px;
    font-size: 12px;
    background-color: #1E9205;
    color: #fff;
    border-radius: 5px;
    margin-left: 10px;
    margin-top: 0px;
    display: inline-block;
  }

</style>