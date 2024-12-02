import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import base from "./base";

createApp(App).use(router).use(base).mount("#app");
