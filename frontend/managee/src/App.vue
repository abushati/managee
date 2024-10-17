<template>
  <div id="app">
    <h1>Vue Form Example</h1>
    <FormComponent />
  </div>
</template>

<script>
import FormComponent from './components/FormComponent.vue';

export default {
  name: 'App',
  components: {
    FormComponent,
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
