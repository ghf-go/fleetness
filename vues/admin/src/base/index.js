import api from "./utils/api";
export default {
  install(app, options) {
    app.config.globalProperties.$api = api.apiPost;
    app.component("Editor", import("./widgets/Editor.vue"));
  },
};
