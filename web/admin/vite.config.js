import path from "path";

import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import VueDevTools from 'vite-plugin-vue-devtools';

export default defineConfig({
  //base: process.env.NODE_ENV === 'production' ? '/_/' : '/',

  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },

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
    createSvgIconsPlugin({
      iconDirs: [path.resolve(process.cwd(), "./src/assets/svg")],
      symbolId: "icon-[dir]-[name]",
    }),
  ],


  build: {
    rollupOptions: {
      output: {
        manualChunks(id) {
          if (id.includes('node_modules')) {
            return id.toString().split('node_modules/')[1].split('/')[0].toString();
          }
        }
      }
    }
  }
});
