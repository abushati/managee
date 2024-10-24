import { createRouter, createWebHistory } from 'vue-router'
import EmployeeComponent from '../components/employeeComponent.vue'
import StoreView from '../views/StoreView.vue'

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
    }
    // {
    //   path: '/new_employe',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: )
    // }
  ],
})

export default router
