<script setup>
import {onMounted, ref} from 'vue'
import api from "@/service/api.js";

const taskList = ref([]);

onMounted(async() => {
  try{
    const response = await api.getTasks.getAll();
    taskList.value = response.data;
  }catch(error){
    console.log(error);
  }
})

</script>

<template>

  <div class="task-list">
    <h1 class="titolo">Task Manager</h1>
    <div class="task-container">
      <div class="task-card" v-for="task in taskList" :key="task.id">
        <h2 class="task-title">{{task.title}}</h2>
        <span class="task-priority">Priority: {{task.priority}}</span>
        <span class="task-description">Description:
          <br><span>{{task.description}} </span>
        </span>
      </div>
    </div>

  </div>

</template>

<style scoped>
.task-list {
  align-content: flex-start;
}
.titolo {
  text-align: center;
  margin-bottom: 20px;
}
.task-container {
  display: flex;
  flex-direction: column;
}
.task-card {
  display: flex;
  flex-direction: column;
}
.task-title {
  margin-right: 10px;
}
.task-priority {
  margin-right: 10px;
}
</style>
