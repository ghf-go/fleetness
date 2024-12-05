<template>
  <div>
    <h1>修改密码</h1>
    <el-form
      :model="postData"
      :inline="false"
      size="normal"
      label-position="top"
    >
      <el-form-item label="旧密码">
        <el-input
          v-model="postData.oldPass"
          type="password"
          placeholder="请输入旧密码"
        >
        </el-input>
      </el-form-item>
      <el-form-item label="新密码">
        <el-input
          v-model="postData.newPass"
          type="password"
          placeholder="请输入新密码"
        >
        </el-input>
      </el-form-item>
      <el-form-item label="确认密码">
        <el-input
          v-model="postData.rePass"
          type="password"
          placeholder="确认密码"
        >
        </el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="changePass">确认</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      postData: {
        oldPass: "",
        newPass: "",
        rePass: "",
      },
    };
  },
  methods: {
    async changePass() {
      if (this.postData.oldPass.length < 6) {
        this.$message({
          message: "请输入旧密码",
          type: "error",
          showClose: true,
          duration: 3000,
        });
        return;
      }
      if (this.postData.newPass.length < 6) {
        this.$message({
          message: "请输入新密码",
          type: "error",
          showClose: true,
          duration: 3000,
        });
        return;
      }
      if (this.postData.newPass != this.postData.rePass) {
        this.$message({
          message: "两次输入的密码不一致",
          type: "error",
          showClose: true,
          duration: 3000,
        });
        return;
      }
      const data = await this.$api("/account/changepasswd", this.postData);
      if (data.code == 200) {
        this.$message({
          message: "修改成功",
          type: "success",
          showClose: true,
          duration: 3000,
        });
        this.postData = {
          oldPass: "",
          newPass: "",
          rePass: "",
        };
      } else {
        this.$message({
          message: data.msg,
          type: "error",
          showClose: true,
          duration: 3000,
        });
      }
    },
  },
};
</script>

<style lang="scss" scoped></style>
