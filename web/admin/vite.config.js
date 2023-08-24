import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'
import vitePluginVueSvgIcons from 'vite-plugin-vue-svg-icons'

export default defineConfig({
  //base: process.env.NODE_ENV === 'production' ? '/_/' : '/',
  base: '/_/',

  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080/'
      },
      '/uploads': {
        target: 'http://localhost:8080/'
      }
    }
  },

  //build: {
  //  rollupOptions: {
  //    external: ['/_/assets/img/sprite.svg']
  //  }
  //},

  plugins: [VueDevTools(), vue(), vitePluginVueSvgIcons()],

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
