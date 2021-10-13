<template>
  <div class="login_container">
    <div class="login_box">
      <!--头像区域-->
      <div class="avatar_box">
        <img src="../assets/img/login.png" alt="">
      </div>
      <!--登录表单区域-->
      <el-form ref="loginFormRef" label-width="0px" :model="loginForm" :rules="loginFormRules" class="login_form">
        <!--手机号-->
        <el-form-item prop="mobile">
          <el-input v-model="loginForm.mobile" prefix-icon="el-icon-phone" placeholder="请输入手机号"></el-input>
        </el-form-item>
        <!--密码-->
        <el-form-item prop="password">
          <el-input v-model="loginForm.password" prefix-icon="el-icon-key" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <!-- 验证码 -->
        <el-form-item prop="vertify_code">
          <el-input v-model="loginForm.vertify_code" placeholder="验证码" prefix-icon="el-icon-unlock">
            <template slot="append">
              <div class="login-code" @click="refreshCode" title="看不清？点击切换">
                <vertifycode :identifyCode="identifyCode"></vertifycode>
              </div>
            </template>
          </el-input>
        </el-form-item>
        <!--按钮-->
        <el-form-item class="btns">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="info" @click="resetLoginForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
  import vertifycode from "../components/VertifyCode";
  export default {
    name: "Login",
    components: {vertifycode},
    data() {
      // 验证手机号的规则
      var checkMobile = (rule, value, cb) => {
        //验证手机号的正则表达式
        const regMobile = /^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$/
        if (regMobile.test(value)) {
          return cb()
        }
        cb(new Error('请输入合法的手机号'))
      }
      return {
        // 这是登录表单的数据绑定对象
        loginForm: {
          mobile: '',
          password: '',
          vertify_code: ''
        },
        loginFormRules: {
          // 验证用户名是否合法
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'},
            {validator: checkMobile, trigger: 'blur'}
          ],
          // 验证密码是否合法
          password: [
            {required: true, message: '请输入登录密码', trigger: 'blur'},
            {min: 6, max: 15, message: '长度在 6 到 15 个字符', trigger: 'blur'}
          ],
          // 验证码输入是否正确
          vertify_code: [
            {required: true, message: '请输入验证码', trigger: 'blur'},
            {required: true, validator: this.validateCode, change: 'blur'}
          ]
        },
        // 验证码组件
        identifyCodes: 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz0123456789',
        identifyCode: '',
      }
    },
    created() {
      this.refreshCode()
    },
    methods: {
      // 点击重置按钮，重置登录表单
      resetLoginForm() {
        this.$refs.loginFormRef.resetFields()
      },
      login() {
        this.$refs.loginFormRef.validate(valid => {
          if (!valid) return
          this.$http.post('/api/user/login', this.loginForm).then(res => {
            if (res.data.entity.code !== 200) {
              return this.$message.error('登录失败')
            }
            this.$message.success('登录成功')
            window.sessionStorage.setItem('token', res.data.entity.data.token)
            this.$router.push('/home')
          })
        })
      },
      /*// 表单【登录】
      submitLogin(formName) {
        // 登录之前进行表单验证
        this.$refs[formName].validate(async (valid) => {
          // 验证成功
          if (valid) {
            const {data: res} = await this.$http.post('/login', this.loginForm)
            if (res.meta.status === '200') {
              this.$message.success('登录成功！')
            } else {
              this.$message.error('登录失败！请检查用户名和密码输入是否正确！' + res.meta.msg)
            }
          }
        })
      },*/
      // 验证码
      randomNum(min, max) {
        return Math.floor(Math.random() * (max - min) + min)
      },
      // 生成随机验证码
      makeCode(o, l) {
        for (let i = 0; i < l; i++) {
          this.identifyCode += this.identifyCodes[
            this.randomNum(0, this.identifyCodes.length)
            ]
        }
      },
      // 验证码刷新
      refreshCode() {
        this.identifyCode = ''
        this.makeCode(this.identifyCodes, 4)
      },
      // 验证码输入校验
      validateCode(rule, value, callback) {
        if (value !== this.identifyCode) {
          callback(new Error('验证码验证错误！请输入正确的验证码！'))
        } else {
          callback()
        }
      }
    }
  }
</script>

<style lang="less" scoped>
  .login_container {
    background-image: linear-gradient(to right, #868f96 0%, #596164 100%);
    height: 100%;
  }

  .login_box {
    width: 450px;
    height: 300px;
    background-color: #ffffff;
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%); //盒子居中
  }

  .avatar_box {
    height: 130px;
    width: 130px;
    border: 1px solid #eee;
    border-radius: 50%;
    padding: 10px;
    box-shadow: 0 0 10px #ddd;
    position: absolute;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: #fff;

    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #eee;
    }
  }

  .login_form {
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
  }

  .btns {
    display: flex;
    justify-content: flex-end;
  }

  .login-code {
    cursor: pointer;
    margin: 0;
  }
</style>
