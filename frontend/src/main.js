import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import { createApp } from 'vue';
import App from './App.vue';
import router from './router/index';
const app = createApp(App)
app.use(router)
app.use(ArcoVue);

app.mount('#app')