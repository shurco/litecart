import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import VitePluginSvgSpritemap from '@spiriit/vite-plugin-svg-spritemap'
import VueDevTools from 'vite-plugin-vue-devtools';

export default defineConfig({
  //base: process.env.NODE_ENV === 'production' ? '/_/' : '/',
  base: "/_/",

  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080/",
      },
      "/uploads": {
        target: "http://localhost:8080/",
      },
    },
  },

  plugins: [
    VueDevTools(),
    vue(),
    VitePluginSvgSpritemap([
      "./src/assets/svg/*.svg",
      "./src/assets/svg/social/*.svg",
    ]),
  ],

  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
