<script setup>
import {computed, onMounted, ref} from 'vue'
import api from "@/service/api.js";

const index = ref(1);
const page = ref(1);
const itemsPerPage = ref(8);
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

const header = ref([
  {
    align: 'center',
    key: 'title',
    sortable: false,
    title: 'Task',
  },
  { align: 'center', title: 'Priority', key: 'priority' },
  { align: 'center', title: 'Description', key: 'description' },
  { align: 'center', title: 'Created', key: 'created_at' },
  { align: 'center', title: 'Actions', key: 'actions' },
])

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

function clearForm() {
  newTask.value = {title: '', description: '', priority: 1};
  newPrio.value = {priority: 1};
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

const pageCount = computed(() => {
  return Math.ceil(taskList.value.length / itemsPerPage.value);
})

</script>

<template>
    <div class="base">
      <h2 class="titolo">Task Manager</h2>
      <div class="popup-manager">
        <button class="open-btn" @click="popupOpen1 = true">New Tasks</button>
        <div v-if="popupOpen1" class="popup-overlay">
          <div class="popup-content" @click.stop>
            <button class="close-btn" @click="popupOpen1 = false" v-on:click="clearForm">
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
                    <v-radio label="Low" color="green" :value="3"></v-radio>
                    <v-radio label="Medium" color="orange" :value="2"></v-radio>
                    <v-radio label="High" color="red" :value="1"></v-radio>
                  </v-radio-group>
                </div>
                <button type="submit" class="submit">Add Task</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>

  <v-data-table v-model:page="page" :headers="header" :items="taskList" :items-per-page="itemsPerPage">
    <template v-slot:item.priority="{ item }">
      <span class="priority-bubble" :class="{
                  'high' : item.priority === 1,
                  'medium' : item.priority === 2,
                  'low' : item.priority === 3,
                }">{{ item.priority }}</span>
    </template>
    <template v-slot:item.created_at="{ item }">
      {{ formatDate(item.created_at) }}
    </template>
    <template v-slot:item.actions="{ item }">
      <div class="actions-btns">

        <button class="icon-btn edit-btn" @click="popupOpen2 = true; index = item.id" title="Modifica Task">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
          </svg>
        </button>

        <button class="icon-btn delete-btn" @click="removeTask(item.id)" title="Elimina Task">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>

      </div>
    </template>
    <template v-slot:bottom>
      <div class="text-center pt-2">
        <v-pagination
            v-model="page"
            :length="pageCount"
        ></v-pagination>
      </div>
    </template>
  </v-data-table>

  <div v-if="popupOpen2" class="popup-overlay">
    <div class="popup-content" @click.stop>
      <div class="task-form">
        <form @submit="updateTask(index)">
          <div class="number-input">
            <v-radio-group class="number-input" inline v-model="newPrio.priority">
              <v-radio label="Low" color="green" :value="3"></v-radio>
              <v-radio label="Medium" color="orange" :value="2"></v-radio>
              <v-radio label="High" color="red" :value="1"></v-radio>
            </v-radio-group>
          </div>
          <button type="submit" class="submit">Modify Priority</button>
        </form>
      </div>
      <button class="close-btn" @click="popupOpen2 = false" v-on:click="clearForm">
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="18" y1="6" x2="6" y2="18"></line>
          <line x1="6" y1="6" x2="18" y2="18"></line>
        </svg>
      </button>
    </div>
  </div>
</template>

<style scoped>
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
  margin-top: 20px;
  padding-top: 10px;
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
  position: relative;
}

.close-btn {
  top: 15px;
  right: 15px;
  position: absolute;
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
.priority-bubble {
  padding: 6px 14px;
  border-radius: 20px;
  color: white;
  font-weight: 600;
  font-size: 0.9rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.3);
  display: inline-block;
  min-width: 50px;
}

.low {
  background-color: #0c8a07;
}
.medium {
  background-color: #f0ad4e;
}
.high {
  background-color: #d10000;
}
</style>
