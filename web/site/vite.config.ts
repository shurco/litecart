import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [sveltekit(), tailwindcss()],
  server: {
    port: 5273,
    proxy: {
      '/api': {
        target: 'http://localhost:8080/',
        changeOrigin: true
      },
      '/uploads': {
        target: 'http://localhost:8080/',
        changeOrigin: true
      },
      '^/cart/payment': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        ws: true
      }
    }
  }
})
