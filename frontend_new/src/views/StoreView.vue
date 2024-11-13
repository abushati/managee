<script setup>
import axios from 'axios'
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { onMounted } from 'vue';
import { useStoreDataStore } from '@/stores/Store';
import EmployeeComponent from '../components/employeeComponent.vue';
import FormComponent from '../components/FormComponent.vue';
import ScheduleFormComponent from '../components/ScheduleFormComponent.vue';

const route = useRoute();
const storeId = route.params.id;
const storeDataStore = useStoreDataStore();

onMounted(() => {
  storeDataStore.fetchStore(storeId);
});
</script>

<template>
  <main>
    <div>
      <EmployeeComponent/>
    </div>
    <div>
        <ScheduleFormComponent />
    </div>
    <div>
      <FormComponent />      
    </div>
  </main>
</template>


