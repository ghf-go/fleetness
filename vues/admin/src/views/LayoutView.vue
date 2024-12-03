<template>
  <el-container class="body">
    <el-aside class="left" width="200px">
      <el-menu
        :default-active="menuActive"
        :unique-opened="true"
        background-color="#545c64"
        text-color="#aaa"
        active-text-color="#FFF"
        router
      >
        <el-sub-menu v-for="(item, rk) in menus" :key="rk" :index="item.path">
          <template #title>
            <span>{{ item.name }}</span>
          </template>
          <el-menu-item
            class="item"
            v-for="sub in item.children"
            :key="`${item.path}${sub.path}`"
            :index="`${item.path}${sub.path}`"
          >
            {{ sub.name }}
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="header">
        <div class="info">
          <div class="name">{{ nickName }}</div>
          <div class="logout" @click="logout">退出登录</div>
        </div>
      </el-header>
      <el-main class="main">
        <transition name="el-zoom-in-top">
          <router-view />
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      nickName: sessionStorage.getItem("nick_name"),
      menuActive: "",
      menus: [],
    };
  },
  mounted() {
    this.menuActive = this.$route.fullPath;
    this.menus = this.$router.options.routes.filter((item) => {
      return item.meta.isMenu;
    });
    this.menus.forEach((e) => {
      e.children = e.children.filter((i) => {
        return i.meta.isMenu;
      });
      return e;
    });
    console.log(this.$route, this.$router);
  },
  methods: {
    logout() {
      sessionStorage.removeItem("nick_name");
      sessionStorage.removeItem("token");
      this.$router.replace("/login");
    },
  },
};
</script>

<style lang="scss" scoped>
.body {
  margin: 0px;
  padding: 0px;
  .left {
    background-image: linear-gradient(to bottom, #545c64, #545c64);
    min-height: 100vh;
    text-align: left;
    height: 100%;
    width: 200px;
    position: fixed;
    top: 0px;

    .item {
      text-indent: 1rem;
    }
  }
  .header {
    background-image: linear-gradient(to right, #4f4f4f, #444444);
    color: #fff;
    text-align: right;
    .info {
      text-align: right;
      .name {
      }
      .logout {
      }
    }
  }
  .main {
    text-align: left;
    margin-left: 200px;
  }
}
</style>
