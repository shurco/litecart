import { createRouter, createWebHistory } from "vue-router";
import { getCookie } from "@/utils/";
import * as NProgress from "nprogress";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/install",
      name: "install",
      meta: { layouts: "BlankLayouts" },
      component: () => import("@/pages/Install.vue"),
    },
    {
      path: "/signin",
      name: "signin",
      meta: { layouts: "BlankLayouts" },
      component: () => import("@/pages/Signin.vue"),
    },
    {
      path: "/",
      name: "products",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Products.vue"),
    },
    {
      path: "/carts",
      name: "carts",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Carts.vue"),
    },
    {
      path: "/pages",
      name: "pages",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Pages.vue"), 
      children: [
        {
          path: ':page_slug',
          name: 'pagesArticle',
          component: () => import('@/pages/Pages.vue')
        },
      ],
    },
    {
      path: "/settings",
      name: "settings",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Settings.vue"),
    },

    {
      path: "/:pathMatch(.*)*",
      name: "404",
      meta: { layouts: "BlankLayouts" },
      component: () => import("@/pages/404.vue"),
    },
  ],
});

router.beforeEach((to, from, next) => {
  NProgress.start();

  let isAuthenticated = false;
  let token = getCookie("token");
  if (token) {
    isAuthenticated = true;
  }

  if (to.path === "/install") next();
  else if (!isAuthenticated && to.name !== "signin") next({ name: "signin" });
  else if (isAuthenticated && to.name == "signin") next({ name: "products" });
  else next();
});

router.afterEach(() => {
  NProgress.done();
});

export default router;
