<template>
  <el-upload
    v-model:file-list="fileList"
    list-type="picture-card"
    :on-preview="handlePictureCardPreview"
    :on-remove="handleRemove"
    ref="upload"
    :action="host"
    :data="params"
    :on-success="handleAvatarSuccess"
    :on-change="changeFile"
    accept="image/*"
  >
    <el-icon><Plus /></el-icon>
  </el-upload>
  <el-dialog v-model="dialogVisible">
    <img w-full :src="dialogImageUrl" alt="Preview Image" />
  </el-dialog>
</template>

<script>
export default {
  props: ["modelValue"],
  emits: ["update:modelValue"],
  data() {
    return {
      dialogVisible: false,
      dialogImageUrl: "",
      host: "",
      params: {},
    };
  },
  computed: {
    imglist() {
      if (this.modelValue instanceof Array) {
        return this.modelValue.map((item) => {
          return { name: "", url: item };
        });
      } else {
        return this.modelValue.split(",").map((item) => {
          return { name: "", url: item };
        });
      }
    },
  },
  methods: {
    getImgs(url) {
      if (this.modelValue instanceof Array) {
        let aa = this.modelValue;
        aa.push(url);
        this.$emit("update:modelValue", aa);
      } else {
        let aa = this.modelValue;
        if (aa == "") {
          aa = url;
        } else {
          aa += "," + url;
        }
        this.$emit("update:modelValue", aa);
      }
    },
    handleAvatarSuccess(e) {
      console.log("上传完成", e);
      this.$api("/upload/uploadSuccess", e);
      this.getImgs(e.url);
    },
    async changeFile(f) {
      if (f.status == "ready") {
        this.$api("/upload/getToken", {
          key: await this.$filemd5(f),
          file_name: f.name,
        }).then((r) => {
          if (r.code != 200) {
            this.$message.error("调用接口失败");
            return;
          }
          if (r.data.is_exists) {
            this.getImgs(r.data.url);
            return;
          }
          this.host = r.data.upload_host;
          this.params = r.data.data;
          setTimeout(() => {
            this.$refs.upload.submit();
          }, 500);
        });
      }
    },
    handlePictureCardPreview(item) {
      this.dialogImageUrl = item.url;
      this.dialogVisible = true;
    },
    handleRemove(f, flist) {
      console.log("删除文件", f, flist);
    },
  },
};
</script>

<style lang="scss" scoped></style>
