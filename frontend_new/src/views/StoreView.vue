<script setup>
import axios from 'axios'
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { onMounted } from 'vue';
import { useStoreDataStore } from '@/stores/Store';
import EmployeeComponent from '../components/employeeComponent.vue';
import FormComponent from '../components/FormComponent.vue';
import ScheduleFormComponent from '../components/ScheduleFormComponent.vue';
import ForcastComponent from '../components/ForcastComponent.vue';


const route = useRoute();
const storeId = route.params.id;
const storeDataStore = useStoreDataStore();

onMounted(() => {
  storeDataStore.fetchStore(storeId);
});
</script>

<template>
  <main >
    <div id="store">
        <div>
        <EmployeeComponent/>
        </div>
        <div>
            <ScheduleFormComponent />
        </div>
        <div>
        <FormComponent />      
        </div>
    </div>
    <div>
        <ForcastComponent />
    </div>
  </main>
</template>


<style>
#store {
    display: flex;
    width: max-content;
}
</style>