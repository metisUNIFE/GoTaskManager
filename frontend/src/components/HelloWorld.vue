<script setup>
import {onMounted, ref} from 'vue'
import api from "@/service/api.js";

const taskList = ref([]);
const newTask = ref({
  title: "",
  description: "",
  priority: 1
});

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

function removeTask(index) {
  api.deleteTask(index.value);
}


</script>

<template>
  <h1 class="titolo">Task Manager</h1>
  <div class="container">
    <div class="task-list">
      <div class="task-container">
        <div class="task-card" v-for="task in taskList" :key="task.id">
          <span class="task-title">Title:
            <h2>{{task.title}}</h2>
          </span>
          <span class="task-priority">Priority:
            <br><span>{{task.priority}}</span>
          </span>
          <span class="task-description">Description:
            <br><span>{{task.description}} </span>
          </span>
          <input type="button" class="task-button" value="Delete Task" @onclick="removeTask(task.id)" />
        </div>
      </div>
    </div>

    <div class="task-form">
      <form class="form-fill" @submit="addTask">
        <input type="text" placeholder="Enter Task Name" v-model="newTask.title" autocomplete="off" required/>
        <input type="text" placeholder="Enter Task Description" v-model="newTask.description" autocomplete="off" required/>
        <div class="number-input">
        <label>
          <input type="radio" :value="3" v-model.number="newTask.priority"/>
          Bassa
        </label>
        <label>
          <input type="radio" :value="2" v-model.number="newTask.priority"/>
          Media
        </label>
        <label>
          <input type="radio" :value="1" v-model.number="newTask.priority"/>
          Alta
        </label>
        </div>
        <button type="submit" class="submit">Add Task</button>
      </form>
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
  flex: 0 0 900px;
  align-content: flex-start;
}
.titolo {
  text-align: center;
  margin-bottom: 20px;
}
.task-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}
.task-card {
  display: flex;
  flex-direction: row;
  border-radius: 10px;
  background: #3a3a3a;
  min-height: 80px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}
.task-priority,
.task-title,
.task-description {
  flex: 1;              /* Ognuno prende 1/3 dello spazio disponibile (crea le colonne) */
  text-align: center;   /* Il testo dentro ogni colonna è centrato */

  /* Se vuoi essere sicuro che anche il contenuto multilinea sia centrato: */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.task-description {
  flex: 3;
  white-space: normal;
  overflow-wrap: break-word;
  word-break: break-word;
  min-width: 0;
}
.task-form {
  flex: 0 0 400px
}
.number-input {
  display: flex;
  flex-direction: row;
  justify-content: center;
}
.form-fill {
  display: flex;
  flex-direction: column;
}
</style>
