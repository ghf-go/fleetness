<template>
  <el-upload
    class="avatar-uploader"
    ref="upload"
    :show-file-list="false"
    :auto-upload="false"
    :action="host"
    :data="params"
    :on-success="handleAvatarSuccess"
    :on-change="changeFile"
    accept="image/*"
  >
    <img v-if="modelValue" :src="modelValue" class="avatar" />
    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
  </el-upload>
</template>

<script>
export default {
  props: ["modelValue"],
  emits: ["update:modelValue"],
  data() {
    return {
      host: "",
      params: {},
    };
  },
  methods: {
    handleAvatarSuccess(e) {
      this.$emit("update:modelValue", e.url);
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
            this.$emit("update:modelValue", r.data.url);
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
  },
};
</script>

<style lang="scss" scoped>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
