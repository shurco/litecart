<template>
  <MainLayouts>
    <div class="grid grid-cols-1 gap-4 pb-4 lg:grid-cols-[1fr_120px] lg:gap-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">Settings</h1>
      </div>
    </div>

    <div>
      <div>
        <nav class="flex gap-6">
          <router-link :to="{ name: 'settingsMain' }" class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="(route.name === 'settingsMain' ? 'bg-gray-200  text-gray-700' : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700')">Main</router-link>
          <router-link :to="{ name: 'settingsSocials' }" class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="(route.name === 'settingsSocials' ? 'bg-gray-200  text-gray-700' : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700')">Socials</router-link>
          <router-link :to="{ name: 'settingsMail' }" class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="(route.name === 'settingsMail' ? 'bg-gray-200  text-gray-700' : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700')">Mail</router-link>
          <router-link :to="{ name: 'settingsMessages' }" class="shrink-0 rounded-lg p-2 text-sm font-medium"
            :class="(route.name === 'settingsMessages' ? 'bg-gray-200  text-gray-700' : 'text-gray-500 hover:bg-gray-50 hover:text-gray-700')">Messages</router-link>
        </nav>
      </div>
    </div>

    <RouterView />

  </MainLayouts>
</template>

<script setup>
import { useRoute } from 'vue-router';
import MainLayouts from '@/layouts/Main.vue'

const route = useRoute()
</script>
