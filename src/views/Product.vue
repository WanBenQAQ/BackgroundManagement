<template>
  <div>
    <el-card>
      <!--用户添加搜索区域-->
      <el-row :gutter="20">
        <el-col :span="7">
          <!--搜索与添加区域-->
          <el-input placeholder="请输入内容" v-model="inputProductById" clearable>
            <el-button slot="append" icon="el-icon-search" @click="seachProduct"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addProductView">添加商品</el-button>
        </el-col>
      </el-row>

      <!--商品列表区域-->
      <el-table :data="productlist" style="width: 100%" v-loading="loading">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="商品ID">
                <span>{{ props.row.productId }}</span>
              </el-form-item>
              <el-form-item label="商品名称">
                <span>{{ props.row.productName }}</span>
              </el-form-item>
              <el-form-item label="商品简介">
                <span>{{ props.row.productIntro }}</span>
              </el-form-item>
              <el-form-item label="封面">
                <span>{{ props.row.productCoverImg }}</span>
              </el-form-item>
              <el-form-item label="原始价格">
                <span>{{ props.row.originalPrice }}</span>
              </el-form-item>
              <el-form-item label="销售价格">
                <span>{{ props.row.sellingPrice }}</span>
              </el-form-item>
              <el-form-item label="库存数量">
                <span>{{ props.row.stockNum }}</span>
              </el-form-item>
              <el-form-item label="库存标记">
                <span>{{ props.row.tag }}</span>
              </el-form-item>
              <el-form-item label="详细介绍">
                <span>{{ props.row.productDetailContent }}</span>
              </el-form-item>
              <el-form-item label="是否删除">
                <span>{{ props.row.isDeleted }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column width="400" label="产品ID" prop="id"></el-table-column>
        <el-table-column width="150" label="产品名称" prop="productName"></el-table-column>
        <el-table-column width="200" label="原始价格" prop="originalPrice"></el-table-column>
        <el-table-column width="200" label="销售价格" prop="sellingPrice"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="showEditDialog(scope.row.id)">编辑</el-button>
            <el-button v-if="!scope.row.isDeleted" size="mini" type="danger" @click="removeProductById(scope.row.id)">删除</el-button>
            <el-button v-else size="mini" type="info" @click="recoveryProductById(scope.row.id)">复活</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!--添加商品对话框-->
    <el-dialog title="添加商品" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
      <!--内容主体区域-->
      <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="80px">
        <el-form-item label="商品名称" prop="productname">
          <el-input v-model="addForm.productname"></el-input>
        </el-form-item>
        <el-form-item label="商品简介" prop="productintro">
          <el-input v-model="addForm.productintro"></el-input>
        </el-form-item>
        <el-form-item label="商品分类" prop="categoryid">
          <!--options用来指定数据源-->
          <!--props用来指定配置对象-->
          <el-cascader v-model="selectedKeys" :options="categorylist"
            :props="cascaderProps" @change="parentCateChanged" clearable></el-cascader>
        </el-form-item>
        <!--<el-form-item label="商品分类" prop="categoryid">
          <el-input v-model="addForm.categoryid"></el-input>
        </el-form-item>-->
        <el-form-item label="封面" prop="productcoverimg">
          <el-input v-model="addForm.productcoverimg"></el-input>
        </el-form-item>
        <el-form-item label="原始价格" prop="originalprice">
          <el-input type="number" v-model="addForm.originalprice"></el-input>
        </el-form-item>
        <el-form-item label="销售价格" prop="sellingprice">
          <el-input type="number" v-model="addForm.sellingprice"></el-input>
        </el-form-item>
        <el-form-item label="库存数量" prop="stocknum">
          <el-input type="number" v-model="addForm.stocknum"></el-input>
        </el-form-item>
        <el-form-item label="库存标记" prop="tag">
          <el-input v-model="addForm.tag"></el-input>
        </el-form-item>
        <el-form-item label="出售身份" prop="sellstatus">
          <el-input type="number" v-model="addForm.sellstatus"></el-input>
        </el-form-item>
        <el-form-item label="详细介绍" prop="productdetailcontent">
          <el-input v-model="addForm.productdetailcontent"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addProduct">确 定</el-button>
      </span>
    </el-dialog>

    <!--修改商品对话框-->
    <el-dialog title="修改商品" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
      <!--内容主体区域-->
      <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="80px">
        <el-form-item label="商品名称" prop="productname">
          <el-input v-model="editForm.productname"></el-input>
        </el-form-item>
        <el-form-item label="商品简介" prop="productintro">
          <el-input v-model="editForm.productintro"></el-input>
        </el-form-item>
        <el-form-item label="商品分类" prop="categoryid">
          <!--options用来指定数据源-->
          <!--props用来指定配置对象-->
          <el-cascader v-model="selectedKeys" :options="categorylist" :props="cascaderProps" @change="EditparentCateChanged">
          </el-cascader>
        </el-form-item>
        <!--<el-form-item label="商品分类" prop="categoryid">
          <el-input v-model="editForm.categoryid"></el-input>
        </el-form-item>-->
        <el-form-item label="封面" prop="productcoverimg">
          <el-input v-model="editForm.productcoverimg"></el-input>
        </el-form-item>
        <el-form-item label="原始价格" prop="originalprice">
          <el-input type="number" v-model="editForm.originalprice"></el-input>
        </el-form-item>
        <el-form-item label="销售价格" prop="sellingprice">
          <el-input type="number" v-model="editForm.sellingprice"></el-input>
        </el-form-item>
        <el-form-item label="库存数量" prop="stocknum">
          <el-input type="number" v-model="editForm.stocknum"></el-input>
        </el-form-item>
        <el-form-item label="库存标记" prop="tag">
          <el-input v-model="editForm.tag"></el-input>
        </el-form-item>
        <el-form-item label="出售身份" prop="sellstatus">
          <el-input type="number" v-model="editForm.sellstatus"></el-input>
        </el-form-item>
        <el-form-item label="详细介绍" prop="productdetailcontent">
          <el-input v-model="editForm.productdetailcontent"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="editProduct">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: "Product",
    data() {
      return {
        total: 0,
        totalPage: 0,
        // 存储查询到的所有分类数据
        categorylist: [],
        // 存储查询到的所有商品数据
        productlist: [],
        // 控制添加商品对话框的显示与隐藏
        addDialogVisible: false,
        // // 指定级联选择器的配置对象
        cascaderProps: {
          value: 'categoryID',
          label: 'name',
          children: 'children',
          expandTrigger: 'hover',
          checkStrictly: true
        },
        // 选中的父级分类的Id数组
        selectedKeys: [],
        // 添加商品的表单数据
        addForm: {
          productname: '',
          productintro: '',
          categoryid: '',
          productcoverimg: '',
          originalprice: 0.00,
          sellingprice: 0.00,
          stocknum: 0,
          tag: '',
          sellstatus: 0,
          productdetailcontent: ''
        },
        // 添加表单的验证规则对象
        addFormRules: {
          productname: [
            {required: true, message: '请输入商品名称', trigger: 'blur'}
          ],
          productintro: [
            {required: true, message: '请输入商品简介', trigger: 'blur'}
          ],
          productcoverimg: [
            {required: true, message: '请输入封面', trigger: 'blur'}
          ],
          originalprice: [
            {required: true, message: '请输入原始价格', trigger: 'blur'}
          ],
          stocknum: [
            {required: true, message: '请输入销售价格', trigger: 'blur'}
          ],
          sellingprice: [
            {required: true, message: '请输入库存数量', trigger: 'blur'}
          ],
          tag: [
            {required: true, message: '请输入库存标记', trigger: 'blur'}
          ],
          productdetailcontent: [
            {required: true, message: '请输入详细介绍', trigger: 'blur'}
          ]
        },
        // 控制修改商品对话框的显示与隐藏
        editDialogVisible: false,
        // 修改商品的表单数据
        editForm: {
          productid : '',
          productname: '',
          productintro: '',
          categoryid: '',
          productcoverimg: '',
          originalprice: 0.00,
          sellingprice: 0.00,
          stocknum: 0,
          tag: '',
          sellstatus: 0,
          productdetailcontent: ''
        },
        // 修改表单的验证规则对象
        editFormRules: {
          productname: [
            {required: true, message: '请输入商品名称', trigger: 'blur'}
          ],
          productintro: [
            {required: true, message: '请输入商品简介', trigger: 'blur'}
          ],
          productcoverimg: [
            {required: true, message: '请输入封面', trigger: 'blur'}
          ],
          originalprice: [
            {required: true, message: '请输入原始价格', trigger: 'blur'}
          ],
          stocknum: [
            {required: true, message: '请输入销售价格', trigger: 'blur'}
          ],
          sellingprice: [
            {required: true, message: '请输入库存数量', trigger: 'blur'}
          ],
          tag: [
            {required: true, message: '请输入库存标记', trigger: 'blur'}
          ],
          productdetailcontent: [
            {required: true, message: '请输入详细介绍', trigger: 'blur'}
          ]
        },
        inputProductById: '',
        // 显示加载界面
        loading: true
      }
    },
    created() {
      this.getproductList()
    },
    methods: {
      // 获取商品数据
      getproductList() {
        this.$http.get('/api/product/list').then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取商品数据失败')
          }
          this.total = res.data.entity.total
          this.totalPage = res.data.entity.totalPage
          this.productlist = res.data.entity.data
          console.log(this.productlist);
          this.loading = false
        })
      },
      // 删除商品信息
      removeProductById(id) {
        // 弹窗询问是否删除数据
        this.$confirm('此操作将永久删除该商品, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/product/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('删除商品失败')
            }
            this.$message.success('删除商品成功')
            this.getproductList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },
      // 根据Id恢复商品数据
      recoveryProductById(id) {
        // 弹窗询问是否恢复数据
        this.$confirm('此操作将恢复该商品, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$http.delete('/api/product/delete/' + id).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('恢复商品失败')
            }
            this.$message.success('恢复商品成功')
            this.getproductList()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消恢复'
          });
        });
      },
      // 监听添加商品对话框的关闭事件
      addDialogClosed() {
        // 清空添加商品对话框中输入的所有字段信息
        this.$refs.addFormRef.resetFields()
        this.selectedKeys = []
      },
      // 弹出添加商品窗口并且获取分类数据
      addProductView() {
        this.addDialogVisible = true

        // 获取分类数据
        this.getCategoryList()
      },
      addProduct() {
        this.$refs.addFormRef.validate(valid => {
          if(!valid) return
          // 可以发起网络请求
          let addproductForm = {
            productname: this.addForm.productname,
            productintro: this.addForm.productintro,
            categoryid: this.addForm.categoryid,
            productcoverimg: this.addForm.productcoverimg,
            originalprice: Number(this.addForm.originalprice),
            sellingprice: Number(this.addForm.sellingprice),
            stocknum: Number(this.addForm.stocknum),
            tag: this.addForm.tag,
            sellstatus: Number(this.addForm.sellstatus),
            productdetailcontent: this.addForm.productdetailcontent
          }
          this.$http.post('/api/product/add', addproductForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('用户添加失败')
            }
            this.$message.success('用户添加成功')
            this.addDialogVisible = false
            this.getproductList()
          })
        })
      },
      // 获取分类数据
      getCategoryList() {
        this.$http.get('/api/category/list').then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取分类数据失败')
          }
          let obj1 = Object.values(res.data.entity.data)
          for(let i = 0; i < obj1.length; i++) {
            let new1obj = obj1[i].children
            obj1[i].children = Object.keys(obj1[i].children)
            obj1[i].children = []
            obj1[i].children.push(new1obj)
            for(let j = 0; j < obj1[i].children.length; j++) {
              let new2obj = obj1[i].children[j].children
              obj1[i].children[j].children = Object.keys(obj1[i].children[j].children)
              obj1[i].children[j].children = []
              obj1[i].children[j].children.push(new2obj)
            }
          }
          // for(let i = 0; i < obj1.length; i++) {
          //   obj1[i].children = Object.values(obj1[i].children)
          //   for(let j = 0; j < obj1[i].children.length; j++) {
          //     obj1[i].children[j].children = Object.values(obj1[i].children[j].children)
          //   }
          // }
          this.categorylist = obj1
        })
      },
      // 监听修改商品对话框的关闭事件
      editDialogClosed() {
        this.$refs.editFormRef.resetFields()
        this.selectedKeys = []
      },
      // 弹出修改窗口并获取对应窗口的所有数据
      showEditDialog(id) {
        this.editDialogVisible = true
        // 获取分类数据
        this.getCategoryList()
        this.$http.get('/api/product/info/' + id).then(res => {
          if(res.data.entity.code !== 200) {
            return this.$message.error('获取对应产品信息失败')
          }
          this.editForm.productid = res.data.entity.data.id
          this.editForm.productname = res.data.entity.data.productName
          this.editForm.productintro = res.data.entity.data.productIntro
          this.editForm.categoryid = res.data.entity.data.categoryId
          this.editForm.productcoverimg = res.data.entity.data.productCoverImg
          this.editForm.originalprice = res.data.entity.data.originalPrice
          this.editForm.sellingprice = res.data.entity.data.sellingPrice
          this.editForm.stocknum = res.data.entity.data.stockNum
          this.editForm.tag = res.data.entity.data.tag
          this.editForm.sellstatus = res.data.entity.data.sellStatus
          this.editForm.productdetailcontent = res.data.entity.data.productDetailContent
          if(this.editForm.categoryid !== "") {
            this.$http.get('/api/category/list4backend').then(res => {
              if(res.status !== 200) {
                return this.$message.error('获取分类数据失败')
              }
              let newCategory = []
              newCategory = res.data.entity.data
              for(let i = 0; i < newCategory.length; i++) {
                for(let j in newCategory[i]) {
                  if(j === "C1CategoryID" && newCategory[i][j] === this.editForm.categoryid) {
                    this.selectedKeys = [this.editForm.categoryid]
                  } else if(j === "C2CategoryID" && newCategory[i][j] === this.editForm.categoryid) {
                    this.selectedKeys = [newCategory[i]['C1CategoryID'], this.editForm.categoryid]
                  } else if(j === "C3CategoryID" && newCategory[i][j] === this.editForm.categoryid) {
                    this.selectedKeys = [newCategory[i]['C1CategoryID'], newCategory[i]['C2CategoryID'], this.editForm.categoryid]
                  }
                }
              }
            })
          }
        })
      },
      editProduct() {
        this.$refs.editFormRef.validate(valid => {
          if(!valid) return
          let editProductForm = {
            categoryid: this.editForm.categoryid,
            originalprice: Number(this.editForm.originalprice),
            productcoverimg: this.editForm.productcoverimg,
            productdetailcontent: this.editForm.productdetailcontent,
            productid: this.editForm.productid,
            productintro: this.editForm.productintro,
            productname: this.editForm.productname,
            tag: this.editForm.tag,
            sellingprice: Number(this.editForm.sellingprice),
            sellstatus: Number(this.editForm.sellstatus),
            stocknum: Number(this.editForm.stocknum),
          }
          this.$http.put('/api/product/edit', editProductForm).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('修改用户信息失败')
            }
            this.$message.success('修改用户信息成功')
            this.editDialogVisible = false
            this.getproductList()
          })
        })
      },
      // 新增框分类选择项发生变化触发这个函数
      parentCateChanged() {
        this.addForm.categoryid = this.selectedKeys[this.selectedKeys.length - 1]
      },
      // 编辑框的分类选项发生变化触发这个函数
      EditparentCateChanged() {
        this.editForm.categoryid = this.selectedKeys[this.selectedKeys.length - 1]
      },
      seachProduct() {
        let newproductlist = []
        if(this.inputProductById !== '') {
          this.$http.get('/api/product/info/' + this.inputProductById).then(res => {
            if(res.data.entity.code !== 200) {
              return this.$message.error('查询数据失败')
            }
            this.productlist.push(res.data.entity.data)
            newproductlist.push(res.data.entity.data)
            this.productlist = newproductlist
          })
        } else {
          this.getproductList()
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
  .el-cascader {
    width: 100%;
  }
</style>
