<template>
  <div class="login">
    <div class="form">
      <h1>欢迎回来</h1>
      <input
        v-model="reqData.login_name"
        type="text"
        placeholder="请输入账号"
      />
      <input v-model="reqData.pass" type="password" placeholder="请输入密码" />
      <input v-model="reqData.code" placeholder="验证码" />
      <button @click="login">登录</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      topath: window.sessionStorage.getItem("to_path"),
      reqData: {
        login_name: "",
        pass: "",
        code: "",
      },
    };
  },
  mounted() {
    // console.log(
    //   "asd",
    //   window.sessionStorage.getItem("nick_name"),
    //   window.sessionStorage.getItem("token")
    // );
    if (window.sessionStorage.getItem("nick_name")) {
      this.goPath();
    }
  },
  methods: {
    goPath() {
      if (this.topath) {
        this.$router.replace(this.topath);
      } else {
        this.$router.replace("/");
      }
    },
    async login() {
      if (this.reqData.login_name == "") {
        this.$message.error("请输入登录用户名");
        return;
      } else if (this.reqData.pass == "") {
        this.$message.error("请输入密码");
        return;
      }
      const data = await this.$api("/login", this.reqData);
      console.log("login ", data, data.msg);
      if (data.code == 200) {
        sessionStorage.setItem("nick_name", data.data.nick_name);
        this.goPath();
      } else {
        this.$message.error(data.msg);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.login {
  width: 100vw;
  height: 100vh;
  padding-top: 20vh;
  background-image: linear-gradient(to bottom, #666, #444);
  .form {
    width: 400px;
    margin: 0px auto;
    background-color: #333;
    text-align: center;
    padding: 1rem;
    color: #fff;
    input {
      display: block;
      width: 300px;
      margin: 0.5rem auto;
      line-height: 2.5rem;
      height: 2.5rem;
      padding-left: 0.5rem;
      border-radius: 0.5rem;
      font-size: 1rem;
      outline: none;
    }
    button {
      display: block;
      display: block;
      width: 200px;
      line-height: 3rem;
      height: 3rem;
      border-radius: 1rem;
      margin: 0.5rem auto;
      font-size: 1.2rem;
    }
  }
}
</style>
