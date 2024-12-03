import api from "./utils/api";
import Editor from "./widgets/Editor.vue";
import UploadImg from "././widgets/UploadImg.vue";
import Echarts from "././widgets/Echarts.vue";
export default {
  install(app, options) {
    app.config.globalProperties.$api = api.apiPost;
    app.config.globalProperties.$uploadFile = api.uploadFile;
    app.config.globalProperties.$filemd5 = api.fileMd5;
    app.component("Editor", Editor);
    app.component("Echarts", Echarts);
    app.component("UploadImg", UploadImg);
  },
};
