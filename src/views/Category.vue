<template>
  <div>
    <!--新增区域-->
    <el-row :gutter="20">
      <el-button class="addCagegoryBtn" type="primary" @click="addDialogVisible = true">新增分类</el-button>
    </el-row>

    <!--列表页展示区域-->
    <el-table :data="cagegorylist" style="width: 100%" v-loading="loading">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="一级分类ID">
              <span>{{ props.row.C1CategoryID }}</span>
            </el-form-item>
            <el-form-item label="一级分类名称">
              <span>{{ props.row.C1Name }}</span>
            </el-form-item>
            <el-form-item label="一级分类描述">
              <span>{{ props.row.C1Desc }}</span>
            </el-form-item>
            <el-form-item label="一级分类排序">
              <span>{{ props.row.C1Order }}</span>
            </el-form-item>
            <el-form-item label="二级分类ID">
              <span>{{ props.row.C2CategoryID }}</span>
            </el-form-item>
            <el-form-item label="二级分类名称">
              <span>{{ props.row.C2Name }}</span>
            </el-form-item>
            <el-form-item label="二级分类排序">
              <span>{{ props.row.C2Order }}</span>
            </el-form-item>
            <el-form-item label="二级分类父分类ID">
              <span>{{ props.row.C2ParentId }}</span>
            </el-form-item>
            <el-form-item label="三级分类ID">
              <span>{{ props.row.C3CategoryID }}</span>
            </el-form-item>
            <el-form-item label="三级分类名称">
              <span>{{ props.row.C3Name }}</span>
            </el-form-item>
            <el-form-item label="三级分类排序">
              <span>{{ props.row.C3Order }}</span>
            </el-form-item>
            <el-form-item label="三级分类父分类ID">
              <span>{{ props.row.C3ParentId }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column width="300" label="一级分类" prop="C1Name"></el-table-column>
      <el-table-column width="300" label="二级分类" prop="C2Name"></el-table-column>
      <el-table-column width="300" label="三级分类" prop="C3Name"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="showEditDialog(scope.row.C3CategoryID)">编辑</el-button>
          <el-button v-if="!scope.row.isDeleted" size="mini" type="danger" @click="removeCategoryById(scope.row.C3CategoryID)">删除</el-button>
          <el-button v-else size="mini" type="info" @click="recoveryCagegoryById(scope.row.C3CategoryID)">复活</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--添加对话框-->
    <el-dialog title="添加分类" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
      <!--内容主体区域-->
      <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="110px">
        <el-form-item label="一级分类名称" prop="c1name">
          <el-input v-model="addForm.c1name"></el-input>
        </el-form-item>
        <el-form-item label="一级分类描述" prop="c1desc">
          <el-input v-model="addForm.c1desc"></el-input>
        </el-form-item>
        <el-form-item label="一级分类排序" prop="c1order">
          <el-input type="number" v-model="addForm.c1order"></el-input>
        </el-form-item>
        <el-form-item label="二级分类名称" prop="c2name">
          <el-input v-model="addForm.c2name"></el-input>
        </el-form-item>
        <el-form-item label="二级分类排序" prop="c2order">
          <el-input type="number" v-model="addForm.c2order"></el-input>
        </el-form-item>
        <el-form-item label="三级分类名称" prop="c3name">
          <el-input v-model="addForm.c3name"></el-input>
        </el-form-item>
        <el-form-item label="三级分类排序" prop="c3order">
          <el-input type="number" v-model="addForm.c3order"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addCategory">确 定</el-button>
        </span>
    </el-dialog>

    <!--修改对话框-->
    <el-dialog title="修改分类" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
      <el-steps :active="active" finish-status="success" :align-center="true">
        <el-step title="选择分类等级"></el-step>
        <el-step title="修改内容"></el-step>
      </el-steps>
      <!--内容主体区域-->
      <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="80px">
        <div v-if="this.active === 0">
          <span class="categorySelect">分类</span>
          <el-select v-model="value" placeholder="请选择" @change="changeCategory">
            <el-option v-for="item in options" :key="item.value"
                       :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </div>
        <div v-else>
          <el-form-item label="分类名称" prop="name">
            <el-input v-model="editForm.name"></el-input>
          </el-form-item>
          <el-form-item label="分类描述" prop="desc">
            <el-input v-model="editForm.desc" :disabled="isDisabled"></el-input>
          </el-form-item>
          <el-form-item label="分类排序" prop="order">
            <el-input type="number" v-model="editForm.order"></el-input>
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="editCategory">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: "Category",
    data() {
      return {
        total: 0,
        totalPage: 0,
        pagenum: 1,
        // 展示分类数据
        cagegorylist: [],
        // 显示加载界面
        loading: true,
        // 控制添加对话框的显示与隐藏
        addDialogVisible: false,
        addForm: {
          c1name: '',
          c1desc: '',
          c1order: 0,
          c2name: '',
          c2order: 0,
          c3name: '',
          c3order: 0
        },
        addFormRules: {
          c1name: [
            {required: true, message: '请输入一级分类名称', trigger: 'blur'}
          ],
          c1desc: [
            {required: true, message: '请输入一级分类描述', trigger: 'blur'}
          ],
          c1order: [
            {required: true, message: '请输入一级分类排序', trigger: 'blur'}
          ],
          c2name: [
            {required: true, message: '请输入二级分类名称', trigger: 'blur'}
          ],
          c2order: [
            {required: true, message: '请输入二级分类排序', trigger: 'blur'}
          ],
          c3name: [
            {required: true, message: '请输入三级分类名称', trigger: 'blur'}
          ],
          c3order: [
            {required: true, message: '请输入三级分类排序', trigger: 'blur'}
          ]
        },
        // 修改弹窗的显示与隐藏
        editDialogVisible: false,
        editForm: {
          categoryid: '',
          name: '',
          desc: '',
          order: 0
        },
        editFormRules: {
          name: [
            {required: true, message: '请输入分类名称', trigger: 'blur'}
          ],
          order: [
            {required: true, message: '请输入分类排序', trigger: 'blur'}
          ]
        },
        options: [{
          value: 1,
          label: '一级分类'
        }, {
          value: 2,
          label: '二级分类'
        }, {
          value: 3,
          label: '三级分类'
        }],
        // 接收分类选择项
        value: '',
        // 控制进度条
        active: 0,
        // 是否禁用编辑框内描述输入框
        isDisabled: false,
        editNewData: {}
      }
    },
    created() {
      this.getCategoryList()
    },
    methods: {
      // 获取列表页数据
      async getCategoryList() {
        const { data: res } = await this.$http.get('/api/category/list4backend')
        if(res.entity.code !== 200) {
          return this.$message.error('获取列表页数据失败')
        }
        this.cagegorylist = res.entity.data
        this.total = res.entity.total
        this.totalPage = res.entity.totalPage
        this.loading = false
      },
      // 根据Id删除指定数据
      removeCategoryById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将删除该分类, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/category/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('删除数据失败')
            }
            this.$message.success('删除数据成功')
            this.getCategoryList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },
      // 根据Id恢复对应数据
      recoveryCagegoryById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将恢复该分类, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/category/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('恢复分类失败')
            }
            this.$message.success('恢复分类成功')
            this.getCategoryList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消恢复'
          });
        });
      },
      // 关闭时触发该函数
      addDialogClosed() {
        // 清空添加对话框中输入的所有字段信息
        this.$refs.addFormRef.resetFields()
      },
      // 添加分类
      addCategory() {
        this.$refs.addFormRef.validate(valid => {
          if(!valid) return
          this.addForm.c1order = Number(this.addForm.c1order)
          this.addForm.c2order = Number(this.addForm.c2order)
          this.addForm.c3order = Number(this.addForm.c3order)
          this.$http.post('/api/category/add', this.addForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('添加失败')
            }
            this.$message.success('添加成功')
            this.addDialogVisible = false
            this.getCategoryList()
          })
        })
      },
      // 关闭时触发该函数
      editDialogClosed() {
        // 清空修改对话框中输入的所有字段信息
        this.$refs.editFormRef.resetFields()
        // 清空选项
        this.value = ''
        // 清空进度条进度
        this.active = 0
      },
      // 显示编辑弹窗
      showEditDialog(id) {
        this.$http.get('/api/category/info/' + id).then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取该条数据信息失败')
          }
          this.editNewData = Object.values(res.data.entity.data)
        })
        this.editDialogVisible = true
      },
      editCategory() {
        this.$refs.editFormRef.validate(valid => {
          if(!valid) return
          this.editForm.order = Number(this.editForm.order)
          this.$http.put('/api/category/edit', this.editForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('修改失败')
            }
            this.$message.success('修改成功')
            this.editDialogVisible = false
            this.getCategoryList()
          })
        })
      },
      changeCategory(val) {
        if(val !== '') this.active = 1
        if(val === 1) {
          this.isDisabled = false
          this.editForm.categoryid = this.editNewData[0].categoryID
          this.editForm.name = this.editNewData[0].name
          this.editForm.desc = this.editNewData[0].desc
          this.editForm.order = this.editNewData[0].order
        } else if (val === 2) {
          this.isDisabled = true
          this.editForm.categoryid = Object.values(this.editNewData[0].children)[0].categoryID
          this.editForm.name = Object.values(this.editNewData[0].children)[0].name
          this.editForm.order = Object.values(this.editNewData[0].children)[0].order
        } else if (val === 3) {
          this.isDisabled = true
          this.editForm.categoryid = Object.values(Object.values(this.editNewData[0].children)[0].children)[0].categoryID
          this.editForm.name = Object.values(Object.values(this.editNewData[0].children)[0].children)[0].name
          this.editForm.order = Object.values(Object.values(this.editNewData[0].children)[0].children)[0].order
        }
      }
    }
  }
</script>

<style scoped>
  .addCagegoryBtn {
    margin: 0 0 20px 10px;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
  .CategorySelect {
    margin-bottom: 20px;
  }
  .el-steps {
    margin-bottom: 20px;
  }
  .categorySelect {
    margin: 0 13px 0 30px;
  }
  .el-select {
    margin-bottom: 20px;
  }
</style>
