<template>
    <div>
      <h1>Data from API:</h1>
      <p v-if="loading">Loading...</p>
      <p v-if="error">{{ error }}</p>
      <ul v-if="data">
        <li v-for="item in employees" :key="item.id">{{ item.name }}</li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: 'EmployeeComponent',
    data() {
      return {
        data: null,
        employees: null, // To hold the API data
        loading: false,
        error: null,
      };
    },
    mounted() {
      this.fetchData(); // Fetch data when component is mounted
    },
    methods: {
      async fetchData() {
        this.loading = true;
        try {
          const response = await axios.get('http://localhost:8080/store/1'); // Example API
          this.data = response.data;
          this.employees = this.data.employees
        //   this.employees = JSON.parse(this.data) // Populate Vue variable
        } catch (err) {
          this.error = 'Error fetching data'; // Handle error
          console.log(err)
        } finally {
          this.loading = false;
        }
      }
    }
  };
  </script>
  