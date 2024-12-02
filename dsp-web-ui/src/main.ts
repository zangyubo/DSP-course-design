import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import 'echarts'
import VChart from 'vue-echarts'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.component('v-chart', VChart)

app.use(createPinia())
app.use(router)

app.mount('#app')
