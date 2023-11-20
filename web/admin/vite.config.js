import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vitePluginSvgsIcons from "vite-plugin-svgs-icons";

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
    vue(),
    vitePluginSvgsIcons({
      iconPrefix: 'ico',
      clearOriginFill: false,
    }),
  ],

  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
