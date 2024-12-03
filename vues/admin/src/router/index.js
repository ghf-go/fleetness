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
        component: () => import("../views/account/users.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "changepass",
        name: "修改密码",
        component: () => import("../views/account/changepass.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/audit/",
    name: "审核",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "el-icon-s-custom",
    },
    children: [
      {
        path: "",
        name: "账号审核",
        component: () => import("../views/audit/users.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "feed",
        name: "动态审核",
        component: () => import("../views/audit/feed.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "comment",
        name: "评论审核",
        component: () => import("../views/audit/comment.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/stat/",
    name: "统计",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "el-icon-s-custom",
    },
    children: [
      {
        path: "",
        name: "账号统计",
        component: () => import("../views/stat/account.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "dot",
        name: "打点统计",
        component: () => import("../views/stat/dot.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "order",
        name: "订单统计",
        component: () => import("../views/stat/order.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "feed",
        name: "动态统计",
        component: () => import("../views/stat/feed.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "comment",
        name: "评论统计",
        component: () => import("../views/stat/comment.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/system/",
    name: "系统配置",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "Setting",
    },
    children: [
      {
        path: "",
        name: "应用配置",
        component: () => import("../views/system/app.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "appvers",
        name: "版本列表",
        component: () => import("../views/system/appvers.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "lottery",
        name: "抽奖配置",
        component: () => import("../views/system/lottery.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "friendlink",
        name: "友情链接",
        component: () => import("../views/system/friendlink.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "dot",
        name: "打点配置",
        component: () => import("../views/system/dot.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/content/",
    name: "内容管理",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "el-icon-s-custom",
    },
    children: [
      {
        path: "",
        name: "新闻列表",
        component: () => import("../views/content/news.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "feed",
        name: "动态列表",
        component: () => import("../views/content/feed.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "comment",
        name: "评论列表",
        component: () => import("../views/content/comment.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "feedback",
        name: "意见反馈",
        component: () => import("../views/content/feedback.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/mall/",
    name: "商城管理",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "el-icon-s-custom",
    },
    children: [
      {
        path: "",
        name: "商品列表",
        component: () => import("../views/mall/goods.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "category",
        name: "商品分类",
        component: () => import("../views/mall/category.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "spec",
        name: "商品规格",
        component: () => import("../views/mall/spec.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "borad",
        name: "品牌管理",
        component: () => import("../views/mall/borad.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "freight",
        name: "运费规则",
        component: () => import("../views/mall/freight.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
    ],
  },
  {
    path: "/appproject/",
    name: "项目生产",
    component: () => import("../views/LayoutView.vue"),
    meta: {
      reqauth: true,
      isMenu: true,
      icon: "el-icon-s-custom",
    },
    children: [
      {
        path: "",
        name: "模块列表",
        component: () => import("../views/appproject/modules.vue"),
        meta: {
          reqauth: true,
          isMenu: true,
        },
      },
      {
        path: "build",
        name: "生产项目",
        component: () => import("../views/appproject/build.vue"),
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
