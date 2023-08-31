import { createRouter, createWebHistory } from "vue-router";
import { getCookie } from "@/utils/";

// @ts-ignore
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
      path: "/checkouts",
      name: "checkouts",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Checkouts.vue"),
    },
    {
      path: "/pages",
      name: "pages",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Pages.vue"),
      redirect: "/pages/terms",
      children: [
        /*
        {
          path: ':pageUrl',
          name: 'page',
          component: () => import('@/pages/Page.vue')
        }
        */
        {
          path: "terms",
          name: "pagesTerms",
          component: () => import("@/pages/PagesTerms.vue"),
        },
        {
          path: "privacy",
          name: "pagesPrivacy",
          component: () => import("@/pages/PagesPrivacy.vue"),
        },
        {
          path: "cookies",
          name: "pagesCookies",
          component: () => import("@/pages/PagesCookies.vue"),
        },
      ],
    },
    {
      path: "/settings",
      name: "settings",
      meta: { layouts: "MainLayouts" },
      component: () => import("@/pages/Settings.vue"),
      redirect: "/settings/main",
      children: [
        {
          path: "main",
          name: "settingsMain",
          component: () => import("@/pages/SettingsMain.vue"),
        },
        {
          path: "socials",
          name: "settingsSocials",
          component: () => import("@/pages/SettingsSocials.vue"),
        },
        {
          path: "mail",
          name: "settingsMail",
          component: () => import("@/pages/SettingsMail.vue"),
        },
        {
          path: "messages",
          name: "settingsMessages",
          component: () => import("@/pages/SettingsMessages.vue"),
        },
      ],
    },

    {
      path: "/:pathMatch(.*)",
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
