<script setup>
import {onMounted, ref} from 'vue'
import api from "@/service/api.js";

const popupOpen1 = ref(false);
const popupOpen2 = ref(false);
const taskList = ref([]);
const newTask = ref({
  title: "",
  description: "",
  priority: 1
});
const newPrio = ref({
  priority: 1
})

onMounted(async() => {
  try{
    const response = await api.getTasks.getAll();
    taskList.value = response.data;
  }catch(error){
    console.log(error);
  }
})

function addTask() {
  console.log("invio" , JSON.parse(JSON.stringify(newTask.value)));
  api.sendTask(newTask.value)
    .then(res => {
      console.log(res),
      newTask.value = {title: "", description: "", priority: 1};})
    .catch(err => console.log(err));
}

function updateTask(index) {
  console.log("invio" , JSON.parse(JSON.stringify(newPrio.value)));
  api.updateTask(index, newPrio.value)
    .then(res => {
      console.log(res),
      newPrio.value = {priority: 1};
    })
    .catch(err => console.log(err));
}

function removeTask(index) {
  console.log("invio " + index);
  api.deleteTask(index);
  location.reload();
}

const formatDate = (dateString) => {
  if (!dateString) return '';

  const date = new Date(dateString);

  return new Intl.DateTimeFormat('it-IT', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date);
};

</script>

<template>
    <div class="base">
      <h2 class="titolo">Task Manager</h2>
      <div class="popup-manager">
        <button class="open-btn" @click="popupOpen1 = true">New Tasks</button>
        <div v-if="popupOpen1" class="popup-overlay">
          <div class="popup-content" @click.stop>
            <button class="close-btn" @click="popupOpen1 = false">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="18" y1="6" x2="6" y2="18"></line>
                <line x1="6" y1="6" x2="18" y2="18"></line>
              </svg>
            </button>
            <div class="task-form">
              <form class="form-fill" @submit="addTask">
                <v-text-field v-model="newTask.title" label="Title" :counter="15" required></v-text-field>
                <v-text-field v-model="newTask.description" label="Description" :counter="150" required></v-text-field>
                <div class="number-input">
                  <v-radio-group class="number-input" inline v-model="newTask.priority">
                    <v-radio label="Low" color="red" :value="3"></v-radio>
                    <v-radio label="Medium" color="orange" :value="2"></v-radio>
                    <v-radio label="High" color="green" :value="1"></v-radio>
                  </v-radio-group>
                </div>
                <button type="submit" class="submit">Add Task</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="task-list">
        <table class="task-table">
          <thead>
            <tr>
              <th>Title</th>
              <th>Priority</th>
              <th>Description</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="task in taskList" :key="task.id" class="task-list-item">
              <td>{{ task.title }}</td>
              <td>{{task.priority}}</td>
              <td>{{task.description}}</td>
              <td>{{formatDate(task.created_at)}}</td>
              <td class="actions-btns">
                <button class="icon-btn edit-btn" @click="popupOpen2 = true" title="Modifica Task">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                    <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                  </svg>
                </button>
                <div v-if="popupOpen2" class="popup-overlay">
                  <div class="popup-content" @click.stop>
                    <div class="task-form">
                      <form @submit="updateTask(task.id)">
                        <div class="number-input">
                          <v-radio-group class="number-input" inline v-model="newPrio.priority">
                            <v-radio label="Low" color="red" :value="3"></v-radio>
                            <v-radio label="Medium" color="orange" :value="2"></v-radio>
                            <v-radio label="High" color="green" :value="1"></v-radio>
                          </v-radio-group>
                        </div>
                        <button type="submit" class="submit">Modify Priority</button>
                      </form>
                    </div>
                    <button class="close-btn" @click="popupOpen2 = false">Close popup</button>
                  </div>
                </div>

                <button class="icon-btn delete-btn" @click="removeTask(task.id)" title="Elimina Task">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                  </svg>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-start;
  gap: 40px;
  padding: 20px;
  width: 100%;      /* Occupa tutto lo spazio orizzontale */
  max-width: 100%;  /* Assicurati che non venga limitato */
  box-sizing: border-box;
}
.task-list {
  align-content: flex-start;
}
.titolo {
  margin: 0;
}
.task-form {
  flex: 0 0 400px
}
.number-input {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  padding-bottom: 10px;
}
.form-fill {
  display: flex;
  flex-direction: column;
}
.task-table{
  border-collapse: collapse;
  width: 100%;
}
.task-list-item{
  border-bottom: 1px solid #727272;
  border-top: 1px solid #727272;
}
.task-table tr,
.task-table th{
  padding: 12px;
  text-align: center;
}
.icon-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%; /* Cerchio perfetto */
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease-in-out;
  color: #888; /* Colore grigio di base per le icone a riposo */
}
.actions-btns{
  display: flex;
  padding: 10px;
  align-items: center;
  justify-content: center;
}
.icon-btn:hover {
  transform: scale(1.1);
  color: white;
}

.edit-btn:hover {
  background-color: #3498db; /* Blu al passaggio del mouse */
}

.delete-btn:hover {
  background-color: #ff4d4d; /* Rosso al passaggio del mouse */
}
.base{
  display: flex;
  justify-content: space-between;
  align-items: center;

  padding: 0 20px;
  margin-bottom: 20px;
  margin-top: 20px;
}
.popup-manager{
  position: relative;
}
.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.6); /* Nero al 60% di opacità */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000; /* Assicura che stia sopra a tutto il resto del sito */
}

/* La finestrella centrale */
.popup-content {
  background-color: #2a2a2a; /* Colore scuro per intonarsi al tuo tema precedente */
  color: white;
  padding: 30px;
  border-radius: 12px;
  min-width: 400px;
  max-width: 700px;
  box-shadow: 0 10px 25px rgba(0,0,0,0.5);
  text-align: center;
}

.close-btn {
  top: 15px;
  left: 15px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%; /* Cerchio perfetto */
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease-in-out;
  color: #888; /* Colore grigio di base per le icone a riposo */
}
.close-btn:hover {
  transform: scale(1.1);
  color: white;
  background-color: #ff4d4d;
}
</style>
