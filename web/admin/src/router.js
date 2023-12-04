import { createRouter, createWebHistory } from "vue-router";
import { getCookie } from "@/utils/";
import * as NProgress from "nprogress";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/install",
      name: "install",
      meta: { layout: "Blank" },
      component: () => import("@/pages/Install.vue"),
    },
    {
      path: "/signin",
      name: "signin",
      meta: { layout: "Blank" },
      component: () => import("@/pages/Signin.vue"),
    },
    {
      path: "/",
      name: "products",
      meta: { layout: "Main", ico: "cube" },
      component: () => import("@/pages/Products.vue"),
    },
    {
      path: "/carts",
      name: "carts",
      meta: { layout: "Main", ico: "cart" },
      component: () => import("@/pages/Carts.vue"),
    },
    {
      path: "/pages",
      name: "pages",
      meta: { layout: "Main", ico: "docs" },
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
      meta: { layout: "Main", ico: "booth", divider: true },
      redirect: to => {
        return { path: '/settings/main' }
      },
      children: [
        {
          path: 'main',
          name: 'settingsMain',
          meta: { ico: "home", title: "Main" },
          component: () => import('@/pages/settings/main.vue')
        },
        {
          path: 'password',
          name: 'settingsPassword',
          meta: { ico: "finger-print", title: "Password" },
          component: () => import('@/pages/settings/password.vue')
        },
        {
          path: 'payment',
          name: 'settingsPayment',
          meta: { ico: "money", title: "Payment" },
          component: () => import('@/pages/settings/payment.vue')
        },
        {
          path: 'webhook',
          name: 'settingsWebhook',
          meta: { ico: "webhook", title: "Webhook events" },
          component: () => import('@/pages/settings/webhook.vue')
        },
        {
          path: 'socials',
          name: 'settingsSocials',
          meta: { ico: "user-group", title: "Social" },
          component: () => import('@/pages/settings/social.vue')
        },
        {
          path: 'mail',
          name: 'settingsMail',
          meta: { ico: "at-symbol", title: "Mail" },
          component: () => import('@/pages/settings/mail.vue')
        },
      ],
    },

    {
      path: "/:pathMatch(.*)*",
      name: "404",
      meta: { layout: "Blank" },
      component: () => import("@/pages/404.vue"),
    },
  ],
});

router.beforeEach((to, from, next) => {
  NProgress.start();

  loadLayoutMiddleware(to);

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

async function loadLayoutMiddleware(route) {
  let layoutComponent;
  try {
    layoutComponent = await import(`@/layouts/${route.meta.layout}.vue`);
  } catch (e) {
    console.error('Error occurred in processing of layout: ', e);
    console.log('Mounted default layout `Blank`');
    layoutComponent = await import(`@/layouts/Blank.vue`);
  }
  route.meta.layoutComponent = layoutComponent.default;
}

export default router;
