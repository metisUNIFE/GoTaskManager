import axios from 'axios';

const BASE_URL = '/api';

const api = axios.create({
    baseURL: BASE_URL,
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    },
    withCredentials: false,
    timeout: 5000
})

export default{
    getTasks:{
        getAll(){
            return api.get('/tasks');
        }
    },
    sendTask(task){
        return api.post('/tasks', task);
    },
    deleteTask(id){
        return api.delete(`/tasks/${id}`);
    },
    updateTask(id, priority){
        return api.patch(`/tasks/${id}`, priority);
    }

}