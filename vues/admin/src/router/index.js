import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "",
    redirect: "/account/",
    meta: {
      reqauth: false,
      isMenu: false,
    },
  },
  {
    path: "/account/",
    name: "用户管理",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "el-icon-s-custom",
    },
    children: [
      {
        path: "",
        name: "用户列表",
        component: () => import("../views/account/Index.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/login",
    name: "login",
    component: () => import("../views/LoginView.vue"),
    meta: {
      reqauth: false,
      isMenu: false,
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
router.beforeEach((to, from, next) => {
  if (to.meta.reqauth && window.sessionStorage.getItem("token") == "") {
    window.localStorage.setItem("to_path", to.fullPath);
    next("/login");
  }
  next();
});

export default router;
