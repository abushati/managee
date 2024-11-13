<template>
    <div>
      <h2>Input Form for week {{week}}, year {{year}}</h2>
      <form @submit.prevent="handleSubmit">
        <div class="employee-toggles">
          <button
                v-for="employee in employees"
                :key="employee.id"
                :class="{ active: selectedEmployeesIds.includes(employee.id) }"
                @click="toggleEmployee(employee.id)"
            >
                {{ employee.name }}
            </button>
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
          <input type="time" id="startTimeInput" v-model="selectedStartTime" @input="createBackEndTime('start')"/>
        </div>
        <div>
          <label for="endTimeInput">End Time: </label>
          <input type="time" id="endTimeInput" v-model="selectedEndTime" @input="createBackEndTime('end')"/>
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
  import { useStoreDataStore } from '@/stores/Store'; 
  import axios from 'axios';
  export default {
      data() {
        return {
        days: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
        selectedDays: [], 
        selectedEmployeesIds: [],
        year: null,
        week: null,
        selectedEndTime: null,
        selectedStartTime: null,
        selectedStartTimeBackend: null,
        selectedEndTimeBackend: null
        };
      },
      computed: {
      // Access the Pinia store data in a computed property to make it reactive
        employees() {
          const storeDataStore = useStoreDataStore();  // Access the Pinia store
          return storeDataStore.storeData.employees;   // Return the employees from the store
        },
        store() {
          const storeDataStore = useStoreDataStore();  // Access the Pinia store
          return storeDataStore.storeData.store;   // Return the employees from the store
        }
      },
      methods: {
          createBackEndTime(timeType){
            function timeInMinutesFromMidnight(timeString) {
              const [time, period] = timeString.split(" ");  // Split into time and AM/PM
              let [hours, minutes] = time.split(":").map(Number);
              if (period === "AM" && hours === 12) {
                hours = 0;
              } else if (period === "PM" && hours !== 12) {
                hours += 12;
              }
              const timeInMinutes = hours * 60 + minutes;
              return timeInMinutes;
            }
            if (timeType == 'end') {
              this.selectedEndTimeBackend = timeInMinutesFromMidnight(this.selectedEndTime)
            }
            else {
              this.selectedStartTimeBackend = timeInMinutesFromMidnight(this.selectedStartTime)
            }
          },
          toggleDay(day) {
          if (this.selectedDays.includes(day)) {
              this.selectedDays = this.selectedDays.filter(d => d !== day);
          } else {
              this.selectedDays.push(day);
          }
          },
          toggleEmployee(eId) {
            if (this.selectedEmployeesIds.includes(eId)) {
                this.selectedEmployeesIds = this.selectedDays.filter(id => id !== eId);
            } else {
                this.selectedEmployeesIds.push(eId);
            }
          },
          getYearAndWeek(date = new Date()) {
            const currentDate = new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()));
            currentDate.setUTCDate(currentDate.getUTCDate() + 4 - (currentDate.getUTCDay() || 7));
            const yearStart = new Date(Date.UTC(currentDate.getUTCFullYear(), 0, 1));
            const weekNumber = Math.ceil((((currentDate - yearStart) / 86400000) + 1) / 7);

            return {
              year: currentDate.getUTCFullYear(),
              week: weekNumber
            };
          },
          async handleSubmit(){
            console.log("submit clicked")
            var allSchedules = []
            for (const eId of this.selectedEmployeesIds) {
              var eSchedule = {'sid':this.store.id, 'eid': eId, 'week':this.week, 'year':this.year,
                              'starttime': this.selectedStartTimeBackend, 'endtime':this.selectedEndTimeBackend}
              for (const day of this.selectedDays) {
                eSchedule['day'] = this.days.indexOf(day)
                allSchedules.push(eSchedule)
                await axios.post(`http://localhost:8080/employee/${eId}/schedule`, allSchedules)
              }
            }
            
          }
      },
      mounted() {
        const { year, week } = this.getYearAndWeek();
        this.year = year;
        this.week = week;
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

.employee-toggles button {
  padding: 8px 12px;
  margin: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: pointer;
}
.employee-toggles button.active {
  background-color: #28a745;
  color: white;
}
</style>
