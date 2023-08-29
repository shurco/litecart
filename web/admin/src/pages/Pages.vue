<template>
  <MainLayouts>
    <div class="grid grid-cols-1 gap-4 pb-4 lg:grid-cols-[1fr_120px] lg:gap-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">Pages</h1>
      </div>
    </div>

    <div>
      <div>
        <nav class="flex gap-6">
          <router-link
            :to="{ name: 'pagesTerms' }"
            class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="
              route.name === 'pagesTerms'
                ? 'bg-gray-200  text-gray-700'
                : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700'
            "
            >Terms</router-link
          >
          <router-link
            :to="{ name: 'pagesPrivacy' }"
            class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="
              route.name === 'pagesPrivacy'
                ? 'bg-gray-200  text-gray-700'
                : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700'
            "
            >Privacy</router-link
          >
          <router-link
            :to="{ name: 'pagesCookies' }"
            class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="
              route.name === 'pagesCookies'
                ? 'bg-gray-200  text-gray-700'
                : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700'
            "
            >Cookies</router-link
          >
        </nav>
      </div>
    </div>

    <RouterView />
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import MainLayouts from '@/layouts/Main.vue'
// @ts-ignore
import * as NProgress from 'nprogress'

const route = useRoute()
const routeName = ref()
const pages = ref()

onMounted(async () => {
  pagesList()
  routeName.value = route.params
})

const pagesList = async () => {
  NProgress.start()

  await fetch(`/api/pages`, {
    credentials: 'include',
    method: 'GET'
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        pages.value = data.result
      }
      NProgress.done()
    })
}
</script>
