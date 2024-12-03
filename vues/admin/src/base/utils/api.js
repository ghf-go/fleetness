import CryptoJS from "crypto-js";
import axios from "axios";
import router from "@/router";
const instanceAxios = axios.create();
instanceAxios.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    sessionStorage.setItem("token", response.headers.token);
    if (response.data.code == 303) {
      sessionStorage.removeItem("nick_name");
      sessionStorage.removeItem("token");
      router.replace("/login");
      return response.data;
    }
    return response.data;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);
instanceAxios.interceptors.request.use(
  function (config) {
    console.log("设置header", window.sessionStorage.getItem("token"));
    // 在发送请求之前做些什么
    config.headers.Token = window.sessionStorage.getItem("token");
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

//https://www.axios-http.cn/docs/req_config
//https://developer.qiniu.com/kodo/1235/vars#magicvar
//https://developer.qiniu.com/kodo/1238/go#param-uptoken
const api = {
  apiPost(name, data) {
    return instanceAxios.post(name, data);
  },
  fileMd5(f) {
    return new Promise((resolve, reject) => {
      const fr = new FileReader();
      fr.onload = (e) => {
        console.log("fileMd5", e);
        resolve(
          CryptoJS.MD5(
            CryptoJS.lib.WordArray.create(e.target.result)
          ).toString()
        );
      };
      fr.onerror = (e) => {
        console.log("fileMd5 失败", e);
        reject(e);
      };
      fr.readAsArrayBuffer(f.raw);
    });
  },
  uploadFile(f, success, onProgress) {
    const fr = new FileReader();
    fr.onload = (e) => {
      const data = api
        .apiPost("/upload/getToken", {
          key: CryptoJS.MD5(
            CryptoJS.lib.WordArray.create(e.target.result)
          ).toString(),
          file_name: f.name,
        })
        .then((data) => {
          if (data.code == 200) {
            console.log("获取TOUKEN", data);
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
                console.log("获取到的数据", r);
                api.apiPost("/upload/uploadSuccess", r.data);
                success(r.data.url);
              });
          }
        });
    };
    fr.readAsArrayBuffer(f);
  },
};
export default api;
