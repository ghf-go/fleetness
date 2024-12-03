import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import base from "./base";

createApp(App).use(ElementUI).use(base).use(router).mount("#app");
