import { MD5 } from "crypto-js";
import axios from "axios";
//https://www.axios-http.cn/docs/req_config
//https://developer.qiniu.com/kodo/1235/vars#magicvar
//https://developer.qiniu.com/kodo/1238/go#param-uptoken
const api = {
  apiPost(name, data) {
    return axios.post(name, data, {
      headers: {
        Token: window.sessionStorage.getItem("token"),
      },
    });
  },
  fileMd5(f) {
    return new Promise((resolve, reject) => {
      const fr = new FileReader();
      fr.onload = (e) => {
        resolve(MD5(e.target.result));
      };
      fr.onerror = (e) => {
        reject(e);
      };
    });
  },
  uploadFile(f, success, onProgress) {
    const fr = new FileReader();
    fr.onload = (e) => {
      const data = this.apiPost("/upload/getToken", {
        key: MD5(e.target.result),
        file_name: f.name,
      }).then((data) => {
        if (data.code == 200) {
          if (data.data.is_exists) {
            success(data.data.url);
            return;
          }
          let formData = new FormData();
          formData.append("file", f);
          for (const k in data.data.data) {
            formData.append(k, data.data.data[k]);
          }
          axios
            .post(data.data.upload_host, formData, {
              headers: {
                "Content-Type": "multipart/form-data",
              },
              onUploadProgress: onProgress,
            })
            .then((r) => {
              api.apiPost("/upload/uploadSuccess", r);
              success(r.url);
            });
        }
      });
    };
  },
};
export default api;
