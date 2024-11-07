<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import EmployeeComponent from '../components/employeeComponent.vue';
import FormComponent from '../components/FormComponent.vue';
import ScheduleFormComponent from '../components/ScheduleFormComponent.vue';

const route = useRoute();
const storeId = ref(route.params.id);  // Capture the dynamic route param
</script>

<template>
  <main>
    <div>
      <EmployeeComponent :storeId="storeId" />
    </div>
    <div>
        <ScheduleFormComponent />
    </div>
    <div>
      <FormComponent />      
    </div>
  </main>
</template>
