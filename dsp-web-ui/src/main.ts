import './assets/main.css'
import 'echarts'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import VChart from 'vue-echarts'
import App from './App.vue'
import router from './router'

import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)

app.component('v-chart', VChart)

app.use(createPinia())
app.use(router)
app.use(ArcoVue)

app.mount('#app')
