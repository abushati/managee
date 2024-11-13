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
          :location="item.location"
          :storeId="item.id"/>
          
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
}
</script>