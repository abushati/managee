<!-- src/components/FormComponent.vue -->
<template>
  <div>
    <h2>Input Form</h2>
    <form @submit.prevent="handleSubmit">
      <div>
        <label for="name">Name:</label>
        <input type="text" v-model="name" id="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="email" id="email" required />
      </div>
      <div>
        <label for="age">Age:</label>
        <input
          type="number"
          v-model="age"
          @input="updateAge"
          id="age"
          required
        />
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
import { ref } from 'vue'
import axios from 'axios'

export default {
  name: 'FormComponent',
  setup() {
    const name = ref('')
    const email = ref('')
    const age = ref(null)
    const submitted = ref(false)
    const error = ref(null)

    const updateAge = event => {
      age.value = Number(event.target.value) // Update age on input
    }

    const handleSubmit = async () => {
      try {
        // Make the HTTP POST request
        const response = await axios.post('http://localhost:8080/employee', {
          name: name.value,
          email: email.value,
          age: age.value,
          //Todo: change this to use the store id from the store Pinia store
          storeid: 1,
        })

        // Handle successful response
        console.log('Response:', response.data)
        submitted.value = true
        error.value = null // Reset error
      } catch (err) {
        // Handle error
        console.error('Error:', err)
        error.value = 'Failed to submit data.'
        submitted.value = false // Reset submission status
      }
    }
    return { name, email, age, submitted, handleSubmit, updateAge }
  },
}
</script>

<style scoped>
/* Add some basic styling */
form {
  margin-bottom: 20px;
}
label {
  display: block;
  margin: 10px 0 5px;
}
input {
  padding: 8px;
  margin-bottom: 10px;
}
button {
  padding: 8px 12px;
}
</style>
