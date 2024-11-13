<template>
  <div>
    <h1>Employees:</h1>
    <p v-if="storeDataStore.loading">Loading...</p>
    <p v-if="storeDataStore.error">{{ storeDataStore.error }}</p>
    <ul v-if="storeDataStore.storeData?.employees">
      <li v-for="item in storeDataStore.storeData.employees" :key="item.id">{{ item.name }}</li>
    </ul>
  </div>
</template>

<script setup>
import { useStoreDataStore } from '@/stores/Store';
// const store = useStoreDataStore();
const storeDataStore = useStoreDataStore();
</script>
