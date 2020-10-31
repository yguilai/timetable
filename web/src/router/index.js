import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'
import { Toast } from 'vant'
import getPageTitle from '@/utils/get-page-title'
import { getToken } from '@/utils/auth' // get token from cookie

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Index',
    component: () => import(/* webpackChunkName: "index" */'../views/Index/Index.vue'),
    redirect: '/timetable',
    children: [
      {
        path: 'timetable',
        name: 'Timetable',
        component: () => import(/* webpackChunkName: "index" */'../views/Timetable/Timetable.vue'),
        meta: {
          title: '课表'
        }
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import(/* webpackChunkName: "index" */'../views/Profile/Profile.vue'),
        meta: {
          title: '我的'
        }
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login/Login.vue'),
    meta: {
      title: '登录'
    }
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About/About.vue'),
    meta: {
      title: '关于'
    }
  }
]

const router = new VueRouter({
  routes
})

const whiteList = ['/login'] // no redirect whitelist

// 路由守卫
router.beforeEach(async(to, from, next) => {
  // set page title
  document.title = getPageTitle(to.meta.title)
  const hasToken = getToken()

  if (hasToken) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      const hasGetStuInfo = store.getters.stuId
      if (hasGetStuInfo) {
        next()
      } else {
        try {
          await store.dispatch('stu/getStuInfo')
          next()
        } catch (error) {
          await store.dispatch('stu/resetToken')
          Toast.fail(error || 'Has Error')
          next(`/login?redirect=${to.path}`)
        }
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
    }
  }
})

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

export default router
