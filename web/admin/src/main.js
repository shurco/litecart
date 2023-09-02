import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/router";
import Notifications from 'notiwind'

import "@/assets/app.css";

const app = createApp(App);
app.use(router);
app.use(Notifications);
app.mount("#app");
