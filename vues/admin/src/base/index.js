import api from "./utils/api";
import Editor from "./widgets/Editor.vue";
import UploadImg from "././widgets/UploadImg.vue";
import UploadImgs from "././widgets/UploadImgs.vue";
import Echarts from "././widgets/Echarts.vue";
import tablerowclassName from "./utils/tablerowclassName";
import YesNo from "./widgets/YesNo.vue";
export default {
  install(app, options) {
    app.config.globalProperties.$api = api.apiPost;
    app.config.globalProperties.$uploadFile = api.uploadFile;
    app.config.globalProperties.$filemd5 = api.fileMd5;
    app.config.globalProperties.$tableRowClassName = tablerowclassName;
    app.component("Editor", Editor);
    app.component("Echarts", Echarts);
    app.component("YesNo", YesNo);
    app.component("UploadImg", UploadImg);
    app.component("UploadImgs", UploadImgs);
  },
};
