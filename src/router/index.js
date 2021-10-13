import Vue from 'vue'
import VueRouter from 'vue-router'

const Login = () => import('../views/Login')
const Home = () => import('../views/Home')
const User = () => import('../views/User')
const Product = () => import('../views/Product')
const Order = () => import('../views/Order')
const Banner = () => import('../views/Banner')
const Category = () => import('../views/Category')
const Calendar = () => import('../views/Calendar')
const Map = () => import('../views/Map')

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/home',
    name: 'home',
    component: Home,
    redirect: '/user',
    children: [
      { path: '/user', component: User },
      { path: '/product', component: Product },
      { path: '/order', component: Order },
      { path: '/banner', component: Banner },
      { path: '/category', component: Category },
      { path: '/calendar', component: Calendar },
      { path: '/map', component: Map }
    ]
  }
]

const router = new VueRouter({
  routes
})

// 挂载路由导航首位
router.beforeEach((to, from, next) => {
  // to 将要访问的路径
  // from 从哪个路径跳转而来
  // next 是一个函数，表示放行
  // next() 放行 next('/login') 强制跳转
  if(to.path === '/login') return next()
  //获取token
  const tokenStr = window.sessionStorage.getItem('token')
  if(!tokenStr) return next('/login')
  next()
})

export default router
