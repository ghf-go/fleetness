import api from "./utils/api";
export default {
  install(app, options) {
    app.config.globalProperties.$api = api.apiPost;
    app.config.globalProperties.$uploadFile = api.uploadFile;
    app.config.globalProperties.$filemd5 = api.fileMd5;
    app.component("Editor", import("./widgets/Editor.vue"));
    app.component("UploadImg", import("./widgets/UploadImg.vue"));
  },
};
