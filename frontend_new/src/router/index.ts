import { createRouter, createWebHistory } from 'vue-router'
import EmployeeComponent from '../components/employeeComponent.vue'
import StoreView from '../views/StoreView.vue'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/employee',
      name: 'EmployeeComponent',
      component: EmployeeComponent,
    },
    {
      path: '/store',
      name: 'StoreComponent',
      component: StoreView
    },
    {
      path: '/',
      name: 'HomeView',
      component: HomeView
    }
  ],
})

export default router
