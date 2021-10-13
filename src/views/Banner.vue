<template>
  <div>
    <el-card>
      <!--新增区域-->
      <el-row :gutter="20">
        <el-col class="colLeft" :span="1">
          <div class="demo-image__preview">
            <el-tooltip class="item" effect="dark" content="在线图片查看器" placement="top-start">
              <el-image style="width: 50px; height: 40px"
                        :src="url" :preview-src-list="srcList">
              </el-image>
            </el-tooltip>
          </div>
        </el-col>
        <el-col :span="2">
          <el-button class="addBannerBtn" type="primary" @click="addDialogVisible = true">添加轮播图</el-button>
        </el-col>
      </el-row>

      <!--列表页展示区域-->
      <el-table :data="bannerlist" style="width: 100%" v-loading="loading">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="跳转链接">
                <span>{{ props.row.redirectUrl }}</span>
              </el-form-item>
              <el-form-item label="排序">
                <span>{{ props.row.orderBy }}</span>
              </el-form-item>
              <el-form-item label="URL">
                <span>{{ props.row.url }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column width="200" label="创建时间" prop="createAt"></el-table-column>
        <el-table-column width="400" label="轮播图ID" prop="bannerId"></el-table-column>
        <el-table-column width="100" label="排序" prop="orderBy"></el-table-column>
        <el-table-column width="300" label="跳转链接" prop="redirectUrl"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="showEditDialog(scope.row.id)">编辑</el-button>
            <el-button size="mini" type="danger" @click="removeBannerById(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!--添加轮播图对话框-->
      <el-dialog title="添加轮播图" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
        <!--内容主体区域-->
        <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="80px">
          <el-form-item label="url" prop="url">
            <el-input v-model="addForm.url"></el-input>
          </el-form-item>
          <el-form-item label="跳转地址" prop="redirecturl">
            <el-input v-model="addForm.redirecturl"></el-input>
          </el-form-item>
          <el-form-item label="排序" prop="orderBy">
            <el-input type="number" v-model="addForm.orderBy"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addBanner">确 定</el-button>
        </span>
      </el-dialog>

      <!--修改订单对话框-->
      <el-dialog title="修改订单" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
        <!--内容主体区域-->
        <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="85px">
          <el-form-item label="url" prop="url">
            <el-input v-model="editForm.url"></el-input>
          </el-form-item>
          <el-form-item label="跳转地址" prop="redirecturl">
            <el-input v-model="editForm.redirecturl"></el-input>
          </el-form-item>
          <el-form-item label="排序" prop="orderBy">
            <el-input type="number" v-model="editForm.orderby"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="editDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="editBanner">确 定</el-button>
        </span>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: "Banner",
    data() {
      return {
        total: 0,
        totalPage: 0,
        pagenum: 1,
        // 展示轮播图数据
        bannerlist: [],
        url: 'https://s.aigei.com/prevfiles/8a06b2f9a94047efa2f1212d29ad44d9.png?e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:XbOLF3Bab-XAWx7NoBE3MO-AIPM=',
        srcList: [],
        // 显示加载界面
        loading: true,
        // 添加窗口控制
        addDialogVisible: false,
        // 添加订单的表单数据
        addForm: {
          url: '',
          redirecturl: '',
          orderBy: 0,
          createat: '',
          updateat: ''
        },
        // 添加表单验证规则对象
        addFormRules: {
          url: [
            {required: true, message: '请输入url', trigger: 'blur'},
            {max: 255, message: 'url长度不得超过255个字符', trigger: 'blur'}
          ],
          redirecturl: [
            {required: true, message: '请输入跳转链接', trigger: 'blur'},
          ],
          orderby: [
            {required: true, message: '请输入排序', trigger: 'blur'},
          ],
        },
        // 修改窗口控制
        editDialogVisible: false,
        // 修改订单的表单数据
        editForm: {
          bannerid: '',
          url: '',
          redirecturl: '',
          orderby: 0,
          createat: '',
          updateat: ''
        },
        // 修改表单验证规则对象
        editFormRules: {
          url: [
            {required: true, message: '请输入url', trigger: 'blur'},
            {max: 255, message: 'url长度不得超过255个字符', trigger: 'blur'}
          ],
          redirecturl: [
            {required: true, message: '请输入跳转链接', trigger: 'blur'},
          ],
          orderby: [
            {required: true, message: '请输入排序', trigger: 'blur'},
          ],
        }
      }
    },
    created() {
      this.getBannerList()
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
        month = month < 10 ? '0'+month : month;
        day = day < 10 ? '0'+day : day;
        hour = hour < 10 ? '0'+hour : hour;
        min = min < 10 ? '0'+min : min;
        seconds = seconds < 10 ? '0'+seconds : seconds;

        return (year + '-'+month +'-'+day+' '+hour+':'+min+':'+seconds);
      }
    },
    methods: {
      // 获取列表页数据
      getBannerList() {
        this.$http.get('/api/banner/list').then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取列表页数据失败')
          }
          this.bannerlist = res.data.entity.data
          for(let i of this.bannerlist) {
            this.srcList.push(i.url)
          }
          this.loading = false
        })
      },
      // 删除
      removeBannerById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将永久删除该轮播图, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/banner/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('删除轮播图失败')
            }
            this.$message.success('删除轮播图成功')
            this.getBannerList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },
      // 添加窗口关闭时触发
      addDialogClosed() {
        // 清空添加对话框中输入的所有字段信息
        this.$refs.addFormRef.resetFields()
      },
      // 添加轮播图数据
      addBanner() {
        this.$refs.addFormRef.validate(valid => {
          if(!valid) return
          let newBanner = {
            url: this.addForm.url,
            redirecturl: this.addForm.redirecturl,
            orderBy: Number(this.addForm.orderBy),
            createat: this.getNowDatetime,
            updateat: this.getNowDatetime
          }
          this.$http.post('/api/banner/add', newBanner).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('添加失败')
            }
            this.$message.success('添加成功')
            this.addDialogVisible = false
            this.getBannerList()
          })
        })
      },
      // 修改窗口关闭时触发
      editDialogClosed() {
        // 清空修改对话框中输入的所有字段信息
        this.$refs.editFormRef.resetFields()
      },
      // 显示编辑框
      showEditDialog(id) {
        this.editDialogVisible = true
        this.$http.get('/api/banner/info/' + id).then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('查询指定数据失败')
          }
          console.log(res.data.entity.data);
          this.editForm.bannerid = res.data.entity.data.bannerId
          this.editForm.url = res.data.entity.data.url
          this.editForm.redirecturl = res.data.entity.data.redirectUrl
          this.editForm.orderby = res.data.entity.data.orderBy
        })
      },
      editBanner() {
        this.$refs.editFormRef.validate(valid => {
          if(!valid) return
          this.editForm.updateat = this.getNowDatetime
          this.editForm.orderby = Number(this.editForm.orderby)
          this.$http.put('/api/banner/edit', this.editForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('修改失败')
            }
            this.$message.success('修改成功')
            this.editDialogVisible = false
            this.getBannerList()
          })
        })
      }
      // 复活
      /*recoveryOrderById(id) {
        // 弹窗询问是否复活数据
        this.$confirm('此操作将恢复该轮播图, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/banner/delete/' + id).then(res => {
            if(res.status !== 200) {
              return this.$message.error('恢复轮播图失败')
            }
            this.$message.success('恢复轮播图成功')
            this.getOrderList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消恢复'
          });
        });
      }*/
    }
  }
</script>

<style scoped>
  .addBannerBtn {
    margin-left: 40px;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 100%;
  }
  .colLeft {
    margin-left: 10px;
  }
</style>
