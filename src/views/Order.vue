<template>
  <div>
    <el-card>
      <!--搜索与添加区域-->
      <el-row :gutter="20">
        <el-col :span="7">
          <el-input placeholder="请输入内容" v-model="inputOrderById" clearable>
            <el-button slot="append" icon="el-icon-search" @click="seachById"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addDialogVisible = true">添加订单</el-button>
        </el-col>
      </el-row>

      <!--列表展示区域-->
      <el-table :data="orderlist" style="width: 100%" v-loading="loading">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="订单ID">
                <span>{{ props.row.orderId }}</span>
              </el-form-item>
              <el-form-item label="下单用户">
                <span>{{ props.row.nickName }}</span>
              </el-form-item>
              <el-form-item label="手机号">
                <span>{{ props.row.mobile }}</span>
              </el-form-item>
              <el-form-item label="订单总价">
                <span>{{ props.row.totalPrice }}</span>
              </el-form-item>
              <el-form-item label="支付类型">
                <span>{{ getpayTypeText(props.row.payType) }}</span>
              </el-form-item>
              <el-form-item label="支付时间">
                <span>{{ props.row.payTime }}</span>
              </el-form-item>
              <el-form-item label="额外补充">
                <span>{{ props.row.extraInfo }}</span>
              </el-form-item>
              <el-form-item label="用户地址">
                <span>{{ props.row.userAddress }}</span>
              </el-form-item>
              <el-form-item label="是否删除">
                <span>{{ props.row.isDeleted }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column width="400" label="订单ID" prop="orderId"></el-table-column>
        <el-table-column width="150" label="订单总价" prop="totalPrice"></el-table-column>
        <el-table-column width="400" label="下单地址" prop="userAddress"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="showEditDialog(scope.row.id)">编辑</el-button>
            <el-button v-if="!scope.row.isDeleted" size="mini" type="danger" @click="removeOrderById(scope.row.id)">删除</el-button>
            <el-button v-else size="mini" type="info" @click="recoveryOrderById(scope.row.id)">复活</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!--添加订单对话框-->
      <el-dialog title="添加订单" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
        <!--内容主体区域-->
        <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="85px">
          <el-form-item label="下单用户" prop="nickname">
            <el-input v-model="addForm.nickname"></el-input>
          </el-form-item>
          <el-form-item label="手机号" prop="mobile">
            <el-input v-model="addForm.mobile"></el-input>
          </el-form-item>
          <el-form-item label="订单总价" prop="totalprice">
            <el-input type="number" v-model="addForm.totalprice"></el-input>
          </el-form-item>
          <el-form-item label="支付状态" prop="paystatus">
            <el-select v-model="addForm.paystatus" clearable placeholder="请选择">
              <el-option v-for="item in optionstatus" :key="item.value"
                :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="支付类型" prop="paytype">
            <el-select v-model="addForm.paytype" clearable placeholder="请选择">
              <el-option v-for="item in optiontype" :key="item.value"
                         :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="支付时间" prop="paytime">
            <el-date-picker v-model="addForm.paytime" type="datetime"
                            placeholder="选择日期时间" value-format="yyyy-MM-dd HH:mm:ss">
            </el-date-picker>
          </el-form-item>
          <el-form-item label="额外补充" prop="extrainfo">
            <el-input v-model="addForm.extrainfo"></el-input>
          </el-form-item>
          <el-form-item label="省市区/县" prop="address1">
            <el-cascader :options="citydata" v-model="addForm.address1"></el-cascader>
          </el-form-item>
          <el-form-item label="详细地址" prop="address2">
            <el-input v-model="addForm.address2"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addOrder">确 定</el-button>
        </span>
      </el-dialog>

      <!--修改订单对话框-->
      <el-dialog title="修改订单" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
        <!--内容主体区域-->
        <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="85px">
          <el-form-item label="下单用户" prop="nickname">
            <el-input v-model="editForm.nickname"></el-input>
          </el-form-item>
          <el-form-item label="手机号" prop="mobile">
            <el-input v-model="editForm.mobile"></el-input>
          </el-form-item>
          <el-form-item label="订单总价" prop="totalprice">
            <el-input type="number" v-model="editForm.totalprice" disabled></el-input>
          </el-form-item>
          <el-form-item label="支付状态" prop="paystatus">
            <el-select v-model="editForm.paystatus" clearable placeholder="请选择">
              <el-option v-for="item in optionstatus" :key="item.value"
                         :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="支付类型" prop="paytype">
            <el-select v-model="editForm.paytype" clearable placeholder="请选择" disabled>
              <el-option v-for="item in optiontype" :key="item.value"
                         :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="支付时间" prop="paytime">
            <el-date-picker v-model="editForm.paytime" type="datetime"
                            placeholder="选择日期时间" value-format="yyyy-MM-dd HH:mm:ss" disabled>
            </el-date-picker>
          </el-form-item>
          <el-form-item label="额外补充" prop="extrainfo">
            <el-input v-model="editForm.extrainfo"></el-input>
          </el-form-item>
          <el-form-item label="详细地址" prop="useraddress">
            <el-input v-model="editForm.useraddress"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="editDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="editOrder">确 定</el-button>
        </span>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
  import citydata from "../components/citydata";
  export default {
    name: "Order",
    data() {
      // 验证手机号的规则
      let checkMobile = (rule, value, cb) => {
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
        // 展示订单数据
        orderlist: [],
        // 控制添加订单窗口的显示与隐藏
        addDialogVisible: false,
        // 添加订单的表单数据
        addForm: {
          nickname: '',
          mobile: '',
          totalprice: 0.00,
          paystatus: 0,
          paytype: 0,
          paytime: '',
          extrainfo: '',
          address1: '',
          address2: ''
        },
        // 添加表单验证规则对象
        addFormRules: {
          nickname: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {min: 2, max: 10, message: '用户名的长度在2~10个字符之间', trigger: 'blur'}
          ],
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'},
            { validator: checkMobile, trigger: 'blur' }
          ],
          totalprice: [
            {required: true, message: '请输入总金额', trigger: 'blur'},
          ],
          paystatus: [
            {required: true, message: '请选择支付状态', trigger: 'blur'},
          ],
          paytype: [
            {required: true, message: '请选择支付方式', trigger: 'blur'},
          ],
          paytime: [
            {required: true, message: '请选择支付时间', trigger: 'blur'},
          ],
          address1: [
            {required: true, message: '请选择省市区/县', trigger: 'blur'},
          ],
          address2: [
            {required: true, message: '请填写详细地址', trigger: 'blur'},
          ]
        },
        // 地址信息
        citydata,
        optionstatus: [
          { value: 0, label: '未支付' },
          { value: 1, label: '已支付' }
        ],
        optiontype: [
          { value: 0, label: '银行卡' },
          { value: 1, label: '微信' },
          { value: 2, label: '支付宝' },
        ],
        // 控制修改对话框的显示与隐藏
        editDialogVisible: false,
        // 修改订单的表单数据
        editForm: {
          orderid: '',
          nickname: '',
          mobile: '',
          totalprice: 0.00,
          paystatus: 0,
          paytype: 0,
          paytime: '',
          extrainfo: '',
          useraddress: '',
        },
        // 修改表单验证规则对象
        editFormRules: {
          nickname: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {min: 2, max: 10, message: '用户名的长度在2~10个字符之间', trigger: 'blur'}
          ],
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'},
            { validator: checkMobile, trigger: 'blur' }
          ],
          paystatus: [
            {required: true, message: '请选择支付状态', trigger: 'blur'},
          ],
          useraddress: [
            {required: true, message: '请输入地址', trigger: 'blur'},
          ]
        },
        inputOrderById: '',
        // 显示加载界面
        loading: true
      }
    },
    created() {
      this.getOrderList()
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
      getOrderList() {
        this.$http.get('/api/order/list').then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取订单列表页数据失败')
          }
          this.orderlist = res.data.entity.data
          this.loading = false
        })
      },
      // 删除
      removeOrderById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将永久删除该订单, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/order/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('删除订单失败')
            }
            this.$message.success('删除订单成功')
            this.getOrderList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },
      // 复活
      recoveryOrderById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将恢复该订单, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/order/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('恢复订单失败')
            }
            this.$message.success('恢复订单成功')
            this.getOrderList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消恢复'
          });
        });
      },
      // 监听添加订单对话框的关闭事件
      addDialogClosed() {
        // 清空添加订单对话框中输入的所有字段信息
        this.$refs.addFormRef.resetFields()
      },
      // 监听修改订单对话框的关闭时间
      editDialogClosed() {
        // 清空添加订单对话框中输入的所有字段信息
        this.$refs.editFormRef.resetFields()
      },
      // 添加订单数据
      addOrder() {
        this.$refs.addFormRef.validate(valid => {
          if(!valid) return
          let newOrderForm = {
            nickname: this.addForm.nickname,
            mobile: this.addForm.mobile,
            totalprice: Number(this.addForm.totalprice),
            paystatus: Number(this.addForm.paystatus),
            paytype: Number(this.addForm.paytype),
            paytime: this.addForm.paytime,
            extrainfo: this.addForm.extrainfo,
            useraddress: this.addForm.address1 + ',' + this.addForm.address2,
            createat: this.getNowDatetime,
            updateat: this.getNowDatetime
          }
          // 可以发起网络请求
          this.$http.post('/api/order/add', newOrderForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('添加失败')
            }
            this.$message.success('添加成功')
            this.addDialogVisible = false
            this.getOrderList()
          })
        })
      },
      // 翻译支付状态
      getpayTypeText(typeInt) {
        switch (typeInt) {
          case 0:
            return '银行卡'
          case 1:
            return '微信'
          case 2:
            return '支付宝'
        }
      },
      // 根据Id显示对应数据的编辑对话框
      showEditDialog(id) {
        this.editDialogVisible = true
        this.$http.get('/api/order/info/' + id).then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('查询订单数据失败')
          }
          this.editForm.orderid = res.data.entity.data.orderId
          this.editForm.nickname = res.data.entity.data.nickName
          this.editForm.mobile = res.data.entity.data.mobile
          this.editForm.totalprice = res.data.entity.data.totalPrice
          this.editForm.paystatus = res.data.entity.data.payStatus
          this.editForm.paytype = res.data.entity.data.payType
          this.editForm.paytime = res.data.entity.data.payTime
          this.editForm.extrainfo = res.data.entity.data.extraInfo
          this.editForm.useraddress = res.data.entity.data.userAddress
        })
      },
      // 修改订单数据
      editOrder() {
        this.$refs.editFormRef.validate(valid => {
          if(!valid) return
          let newOrderForm = {
            orderid: this.editForm.orderid,
            nickname: this.editForm.nickname,
            mobile: this.editForm.mobile,
            paystatus: Number(this.editForm.paystatus),
            extrainfo: this.editForm.extrainfo,
            useraddress: this.editForm.useraddress,
            updateat: this.getNowDatetime
          }
          // 可以发起网络请求
          this.$http.put('/api/order/edit', newOrderForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('修改失败')
            }
            this.$message.success('修改成功')
            this.editDialogVisible = false
            this.getOrderList()
          })
        })
      },
      // 根据Id查询对应订单信息
      seachById() {
        let searchOrder = []
        if(this.inputOrderById !== '') {
          this.$http.get('/api/order/info/' + this.inputOrderById).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('查询失败')
            }
            searchOrder.push(res.data.entity.data)
            this.orderlist = searchOrder
          })
        } else {
          this.getOrderList()
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
  .el-select {
    width: 100%;
  }
  .el-cascader {
    width: 100%;
  }
</style>
