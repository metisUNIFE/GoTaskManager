import { createApp } from 'vue'
import './style.css'
import {createVuetify} from 'vuetify'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'
import App from './App.vue'

const vuetify = createVuetify({
    icons: {
        defaultSet: 'mdi',
    },
})

createApp(App).use(vuetify).mount('#app')
