<template>
  <div id="app">
    <h1>Vue Form Example</h1>
    <FormComponent />
    <EmployeeComponent />
  </div>
</template>

<script>
import FormComponent from './components/FormComponent.vue';
import EmployeeComponent from './components/employeeComponent.vue';

export default {
  name: 'App',
  components: {
    FormComponent,
    EmployeeComponent
  },
};
</script>

<style>
#app {
  max-width: 600px;
  margin: auto;
  padding: 20px;
}
</style>
