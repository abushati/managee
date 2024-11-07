<!-- src/components/FormComponent.vue -->
<template>
    <div>
      <h2>Input Form</h2>
      <form @submit.prevent="handleSubmit">
        <div>
            
        </div>
        <div>
            <h3>Select Days</h3>
            <div class="weekday-toggles">
            <button
                v-for="day in days"
                :key="day"
                :class="{ active: selectedDays.includes(day) }"
                @click="toggleDay(day)"
            >
                {{ day }}
            </button>
            </div>
            <p>Selected Days: {{ selectedDays.join(', ') }}</p>
        </div>
        <div>
          <label for="startTimeInput">Start Time: </label>
          <input type="time" id="startTimeInput" v-model="selectedTime" />
        </div>
        <div>
          <label for="endTimeInput">End Time: </label>
          <input type="time" id="endTimeInput" v-model="selectedTime" />
        </div>
        <button type="submit">Submit</button>
      </form>
  
      <div v-if="submitted">
        <h3>Submitted Data:</h3>
        <p>Name: {{ name }}</p>
        <p>Email: {{ email }}</p>
        <p>Age: {{ age }}</p>
      </div>
    </div>
  </template>
  

  <script>
export default {
    props: {
        storeEmployees: List
    },
    data() {
        return {
        days: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
        selectedDays: []
        };
    },
    methods: {
        toggleDay(day) {
        if (this.selectedDays.includes(day)) {
            this.selectedDays = this.selectedDays.filter(d => d !== day);
        } else {
            this.selectedDays.push(day);
        }
        }
    }
};
</script>

<style>
.weekday-toggles button {
  padding: 8px 12px;
  margin: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: pointer;
}
.weekday-toggles button.active {
  background-color: #28a745;
  color: white;
}
</style>
