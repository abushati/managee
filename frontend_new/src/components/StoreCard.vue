<template>
     <router-link :to="`/store/${storeId}`">
        <div class="StoreCard">
            <div>{{ storeName }}</div>
            <div>{{ location }}</div>
        </div>
    </router-link>
  </template>
  
  <script>
  
  
  export default {
    props: {
        storeId: String,
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
.StoreCard {
    display: block;
    margin: 0 auto 2rem;
    background-color: gray;
    font-size: larger;
  }

</style>
  