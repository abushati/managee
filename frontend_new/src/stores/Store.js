// src/stores/useEmployeeStore.js
import { defineStore } from 'pinia';
import axios from 'axios';

export const useStoreDataStore = defineStore('Store', {
  state: () => ({
    storeData: {},
    loading: false,   // Loading state
    error: null       // Error message
  }),
  actions: {
    async fetchStore(storeId) {
      this.loading = true;
      this.error = null;
      console.log(storeId)
      try {
        const storeResponse = await axios.get(`http://localhost:8080/store/${storeId}`)
        this.storeData = storeResponse.data
        console.error(this.storeData);
      } catch (error) {
        this.error = 'Failed to fetch employees';
        console.error(error);
      } finally {
        this.loading = false;
      }
    }
  },
  persist: true, // Optional: Use this if you want data to persist across reloads
});