import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import(/* webpackChunkName: "about" */ '../views/HomeView.vue')
  },
  {
    path: '/nets',
    name: 'nets',
    component: () => import(/* webpackChunkName: "about" */ '../views/NetsView.vue')
  },
  {
    path: '/hosts',
    name: 'hosts',
    component: () => import(/* webpackChunkName: "about" */ '../views/HostsView.vue')
  }
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

export default router
