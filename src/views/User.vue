<template>
  <div>
    <el-card>
      <!--搜索与添加区域-->
      <el-row :gutter="20">
        <el-col :span="7">
          <el-input placeholder="请输入内容" v-model="inputUserById" clearable>
            <el-button slot="append" icon="el-icon-search" @click="getUserListById"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addDialogVisible = true">添加用户</el-button>
        </el-col>
      </el-row>

      <!--用户列表区域-->
      <el-table :data="userlist" style="width: 100%" v-loading="loading">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="用户ID">
                <span>{{ props.row.id }}</span>
              </el-form-item>
              <el-form-item label="昵称">
                <span>{{ props.row.nickName }}</span>
              </el-form-item>
              <el-form-item label="手机号">
                <span>{{ props.row.mobile }}</span>
              </el-form-item>
              <el-form-item label="地址">
                <span>{{ props.row.address }}</span>
              </el-form-item>
              <el-form-item label="是否锁定">
                <span>{{ props.row.isLocked }}</span>
              </el-form-item>
              <el-form-item label="是否删除">
                <span>{{ props.row.isDeleted }}</span>
              </el-form-item>
              <el-form-item label="创建时间">
                <span>{{ props.row.createAt }}</span>
              </el-form-item>
              <el-form-item label="最近修改时间">
                <span>{{ props.row.updateAt }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column width="400" label="用户ID" prop="id"></el-table-column>
        <el-table-column width="150" label="昵称" prop="nickName"></el-table-column>
        <el-table-column width="200" label="手机号" prop="mobile"></el-table-column>
        <el-table-column width="150" label="是否锁定">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.isLocked" @change="userStateChanged(scope.row)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="showEditDialog(scope.row.id)">编辑</el-button>
            <el-button v-if="!scope.row.isDeleted" size="mini" type="danger" @click="removeUserById(scope.row.id)">删除</el-button>
            <el-button v-else size="mini" type="info" @click="recoveryUserById(scope.row.id)">复活</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!--添加用户对话框-->
    <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
      <!--内容主体区域-->
      <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="70px">
        <el-form-item label="用户名" prop="nickname">
          <el-input v-model="addForm.nickname"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="addForm.password"></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input v-model="addForm.mobile"></el-input>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="addForm.address"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="addDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="adduser">确 定</el-button>
  </span>
    </el-dialog>

    <!--修改用户的对话框-->
    <el-dialog title="修改用户" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
      <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">
        <el-form-item label="用户名" prop="nickname">
          <el-input v-model="editForm.nickname"></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input v-model="editForm.mobile"></el-input>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="editForm.address"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="edituser">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: "User",
    data() {
      // 验证手机号的规则
      var checkMobile = (rule, value, cb) => {
        //验证手机号的正则表达式
        const regMobile = /^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$/
        if(regMobile.test(value)) {
          return cb()
        }
        cb(new Error('请输入合法的手机号'))
      }
      return {
        total: 0,
        totalPage: 0,
        pagenum: 1,
        userlist: [],
        // 控制添加用户对话框的显示与隐藏
        addDialogVisible: false,
        // 添加用户的表单数据
        addForm: {
          nickname: '',
          password: '',
          address: '',
          mobile: ''
        },
        // 添加表单的验证规则对象
        addFormRules: {
          nickname: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {min: 2, max: 10, message: '用户名的长度在2~10个字符之间', trigger: 'blur'}
          ],
          password: [
             {required: true, message: '请输入密码', trigger: 'blur'},
             {min: 6, max: 15, message: '密码的长度在6~15个字符之间', trigger: 'blur'}
          ],
          address: [
            {required: true, message: '请输入邮箱', trigger: 'blur'}
          ],
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'},
            { validator: checkMobile, trigger: 'blur' }
          ]
        },
        // 控制修改用户对话框的显示与隐藏
        editDialogVisible: false,
        // 查询到的用户信息对象
        editForm: {
          id: "",
          nickname: "",
          address: "",
          mobile: "",
          createat: "",
          updateat: ""
        },
        editFormRules: {
          nickname: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {min: 2, max: 10, message: '用户名的长度在2~10个字符之间', trigger: 'blur'}
          ],
          address: [
            {required: true, message: '请输入邮箱', trigger: 'blur'}
          ],
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'},
            { validator: checkMobile, trigger: 'blur' }
          ]
        },
        // 记录搜索框输入的内容
        inputUserById: "",
        // 显示加载界面
        loading: true
      }
    },
    created() {
      this.getUserList()
    },
    computed: {
      // 获取当前日期时间并格式化
      getNowDatetime() {
        let date = new Date();  //创建日期对象
        let year = date.getFullYear(); //年
        let month = date.getMonth() + 1; //月
        let day = date.getDate(); //日
        let hour = date.getHours(); //时
        let min = date.getMinutes(); //分
        let seconds = date.getSeconds(); //秒
        //如果小于10 就拼接个0 否则不变
        month = month < 10 ? '0' + month : month;
        day = day < 10 ? '0' + day : day;
        hour = hour < 10 ? '0' + hour : hour;
        min = min < 10 ? '0' + min : min;
        seconds = seconds < 10 ? '0' + seconds : seconds;
        return (year + '-' + month + '-' + day + ' ' + hour + ':' + min + ':' + seconds);
      }
    },
    methods: {
      getUserList() {
        this.$http.get('/api/user/list').then((res) => {
          if(res.data.entity.code !== 200) {
            return this.$message.error("获取用户列表数据失败")
          } else {
            // console.log(res)
            this.total = res.data.entity.total
            this.totalPage = res.data.entity.totalPage
            this.userlist = res.data.entity.data
            this.loading = false
          }
        })
      },
      // 监听 switch 开关状态的改变
      userStateChanged(userinfo) {
        this.$http.put(`/api/user/deitlocked/${userinfo.id}`).then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('修改失败')
          }
          this.getUserList()
        })
      },
      // 监听添加用户对话框的关闭事件
      addDialogClosed() {
        // 清空添加用户对话框中输入的所有字段信息
        this.$refs.addFormRef.resetFields()
      },
      // 添加新用户
      adduser() {
        this.$refs.addFormRef.validate(valid => {
          if(!valid) return
          this.addForm.createat = this.getNowDatetime
          this.addForm.updateat = this.getNowDatetime
          // 可以发起网络请求
          this.$http.post('/api/user/add', this.addForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('用户添加失败')
            }
            this.$message.success('用户添加成功')
            this.addDialogVisible = false
            this.getUserList()
          })
        })
      },
      // 展示修改用户的对话框
      showEditDialog(id) {
        this.$http.get('/api/user/info/' + id).then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取指定用户信息失败')
          }
          // console.log(res.data.entity.data)
          this.editForm.userid = res.data.entity.data.id
          this.editForm.nickname = res.data.entity.data.nickName
          this.editForm.mobile = res.data.entity.data.mobile
          this.editForm.address = res.data.entity.data.address
          this.editDialogVisible = true
        })
      },
      // 监听修改用户对话框的关闭事件
      editDialogClosed() {
        // 清空添加用户对话框中输入的所有字段信息
        this.$refs.editFormRef.resetFields()
      },
      // 修改用户信息
      edituser() {
        this.$refs.editFormRef.validate(valid => {
          if(!valid) return
          this.editForm.updateat = this.getNowDatetime
          this.$http.put('/api/user/edit', this.editForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('修改用户信息失败')
            }
            this.$message.success('修改用户信息成功')
            this.editDialogVisible = false
            this.getUserList()
          })
        })
      },
      // 根据Id删除用户数据
      removeUserById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将永久删除该用户, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/user/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('删除用户失败')
            }
            this.$message.success('删除用户成功')
            this.getUserList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },
      // 根据Id恢复用户数据
      recoveryUserById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将恢复该用户, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/user/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('恢复用户失败')
            }
            this.$message.success('恢复用户成功')
            this.getUserList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消恢复'
          });
        });
      },
      getUserListById() {
        let searchUser = []
        if(this.inputUserById !== '') {
          this.$http.get('/api/user/info/' + this.inputUserById).then(res => {
            console.log(res);
            if(res.data.entity.code !== 200) {
              return this.$message.error('查询失败')
            }
            searchUser.push(res.data.entity.data)
            this.userlist = searchUser
          })
        } else {
          this.getUserList()
        }
      }
    }
  }
</script>

<style scoped>
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
