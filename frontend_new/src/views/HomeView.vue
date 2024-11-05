<script setup>
import EmployeeComponent from '../components/employeeComponent.vue'
import FormComponent from '../components/FormComponent.vue'
import StoreCard from '../components/StoreCard.vue'
</script>

<template>
    <main>
        <div>
          <div>Stores</div>
          <StoreCard v-for="(item, index) in stores" 
          :key="index" 
          :storeName="item.name" 
          :location="item.location"/>
        </div>
        <div>
            <EmployeeComponent />
        </div>
        <div>
            <FormComponent />      
        </div>
    </main>
  </template>

<script>
import axios from 'axios'
export default {
  name: 'HomeView',
  data() {
    return {
      stores: null,
      loading: false,
      error: null,
    }
  },
  mounted() {
    this.fetchData() // Fetch data when component is mounted
  },
  methods: {
    async fetchData() {
      this.loading = true
      try {
        const response = await axios.get('http://localhost:8080/store') // Example API
        this.stores = response.data
        //   this.employees = JSON.parse(this.data) // Populate Vue variable
      } catch (err) {
        this.error = 'Error fetching data' // Handle error
        console.log(err)
      } finally {
        this.loading = false
      }
    },
  },
}
</script>