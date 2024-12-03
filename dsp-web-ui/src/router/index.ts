import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/addnosie',
      name: 'addnosie',
      component: () => import('../views/addnosie/add-noise.vue'),
    },
    {
      path: '/basic-analysis',
      name: 'basic-analysis',
      component: () => import('../views/basic-analysis/basic-analysis.vue'),
    },
    {
      path: '/code-decode',
      name: 'code-decode',
      component: () => import('../views/code-decode/code-decode.vue'),
    },
  ],
})

export default router
