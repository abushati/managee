<template>
    <div class="StoreCard">
      <div>{{ storeName }}</div>
      <div>{{ location }}</div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    props: {
        storeName: String,
        location: String
    },

    name: 'StoreCard',
    data() {
      return {
        data: null,
        loading: false,
        error: null,
      }
    },
  }
  </script>

<style scoped>
.card {
    display: block;
    margin: 0 auto 2rem;
    background-color: gray;
    font-size: larger;
  }

</style>
  