import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import ElementPlus from "element-plus";
import zhCn from "element-plus/es/locale/lang/zh-cn";
import "element-plus/dist/index.css";
import "dayjs/locale/zh-cn";
import base from "./base";

createApp(App)
  .use(ElementPlus, { size: "small", zIndex: 3000, locale: zhCn })
  .use(base)
  .use(router)
  .mount("#app");
