import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'

// Import des Font Awesome Stylesheets
import './assets/fontawesome/css/all.min.css';

const app = createApp(App)
app.use(router)
app.mount('#app')