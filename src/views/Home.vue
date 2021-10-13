<template>
  <!--头部区域-->
  <el-container class="home_container">
    <el-header>
      <div>
        <el-row :gutter="20">
          <el-col :span="20">
            <span>电商后台管理系统</span>
          </el-col>
          <el-col :span="1">
            <el-tooltip class="item" effect="dark" content="日历" placement="bottom">
              <i class="el-icon-date" @click="ToCalendar('/calendar')"></i>
            </el-tooltip>
          </el-col>
          <el-col :span="1">
            <el-tooltip class="item" effect="dark" content="定位" placement="bottom">
              <i class="el-icon-location-outline" @click="ToMap('/map')"></i>
            </el-tooltip>
          </el-col>
          <el-col :span="2">
            <el-button class="outBtn" type="info" round @click="outBtn">退出</el-button>
          </el-col>
        </el-row>

      </div>
    </el-header>
    <!--页面主体区域-->
    <el-container>
      <!--侧边栏菜单区域-->
      <el-aside :width="isCollapse ? '64px' : '200px'">
        <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
          <el-radio-button :label="false">展开</el-radio-button>
          <el-radio-button :label="true">收起</el-radio-button>
        </el-radio-group>
        <el-menu background-color="#333744" text-color="#fff" @open="handleOpen" @close="handleClose"
                 active-text-color="#ffd04b" :default-active="activePath" :collapse="isCollapse" :collapse-transition="false">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span>导航一</span>
          </template>
          <el-menu-item index="/user" @click="ToUser('/user')">
            <!--图标-->
            <i class="el-icon-user-solid"></i>
            <!--文本-->
            <span>用户列表</span>
          </el-menu-item>
          <el-menu-item index="/product" @click="ToProduct('/product')">
            <!--图标-->
            <i class="el-icon-s-cooperation"></i>
            <!--文本-->
            <span>商品列表</span>
          </el-menu-item>
          <el-menu-item index="/order" @click="ToOrder('/order')">
            <!--图标-->
            <i class="el-icon-s-order"></i>
            <!--文本-->
            <span>订单列表</span>
          </el-menu-item>
          <el-menu-item index="/banner" @click="ToBanner('/banner')">
            <!--图标-->
            <i class="el-icon-s-help"></i>
            <!--文本-->
            <span>轮播图</span>
          </el-menu-item>
          <el-menu-item index="/category" @click="ToCagegory('/category')">
            <!--图标-->
            <i class="el-icon-menu"></i>
            <!--文本-->
            <span>分类列表</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <!--详情页-->
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
  export default {
    name: "Home",
    data() {
      return {
        // 被激活的连接地址
        activePath: '/user',
        isCollapse: true
      }
    },
    created() {
      if(this.activePath === '') {
        this.activePath = '/user'
      }
      this.activePath = window.sessionStorage.getItem('activePath')
    },
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      ToUser(activePath) {
        if(this.$route.path !== '/user') {
          this.$router.push('/user')
          window.sessionStorage.setItem('activePath', activePath)
        }
      },
      ToProduct(activePath) {
        if(this.$route.path !== '/product') {
          this.$router.push('/product')
          window.sessionStorage.setItem('activePath', activePath)
        }
      },
      ToOrder(activePath) {
        if(this.$route.path !== '/order') {
          this.$router.push('/order')
          window.sessionStorage.setItem('activePath', activePath)
        }
      },
      ToBanner(activePath) {
        if(this.$route.path !== '/banner') {
          this.$router.push('/banner')
          window.sessionStorage.setItem('activePath', activePath)
        }
      },
      ToCagegory(activePath) {
        if(this.$route.path !== '/category') {
          this.$router.push('/category')
          window.sessionStorage.setItem('activePath', activePath)
        }
      },
      ToCalendar(activePath) {
        if(this.$route.path !== '/calendar') {
          this.$router.push('/calendar')
          window.sessionStorage.setItem('activePath', activePath)
        }
      },
      ToMap(activePath) {
        if(this.$route.path !== '/map') {
          this.$router.push('/map')
          window.sessionStorage.setItem('map', activePath)
        }
      },
      outBtn() {
        // 清空token
        window.sessionStorage.clear()
        this.$router.push('/login')
      }
    }
  }
</script>

<style scoped>
  .home_container {
    height: 100%;
  }
  .el-header {
    background-color: rgb(55, 61, 65);
    line-height: 60px;
    color: #fff;
    font-size: 25px;
  }
  .el-aside {
    background-color: rgb(51, 55, 68);
  }
  .el-main {
    background-color: rgb(234, 237, 241);
  }
</style>
