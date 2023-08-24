<template>
  <div class="flex">
    <div class="flex-none w-48" id="menu">
      <div class="flex h-screen flex-col justify-between border-e bg-white" id="menu">
        <div class="py-6">
          <ul>
            <li>
              <router-link :to="{ name: 'products' }" class="flex items-center px-6 text-sm font-medium text-gray-900 p-4 hover:bg-gray-50 hover:text-gray-500"
                :class="(route.name === 'products' ? 'bg-gray-100' : 'bg-white')">Products</router-link>
            </li>

            <li>
              <router-link :to="{ name: 'checkouts' }" class="flex items-center px-6 text-sm font-medium text-gray-900 p-4 hover:bg-gray-50 hover:text-gray-500"
                :class="(route.name.startsWith('checkouts') ? 'bg-gray-100' : 'bg-white')">Checkouts</router-link>
            </li>

            <li>
              <router-link :to="{ name: 'pages' }" class="flex items-center px-6 text-sm font-medium text-gray-900 p-4 hover:bg-gray-50 hover:text-gray-500"
                :class="(route.name.startsWith('pages') ? 'bg-gray-100' : 'bg-white')">Pages</router-link>
            </li>

            <li>
              <router-link :to="{ name: 'settings' }" class="flex items-center px-6 text-sm font-medium text-gray-900 p-4 hover:bg-gray-50 hover:text-gray-500"
                :class="(route.name.startsWith('settings') ? 'bg-gray-100' : 'bg-white')">Settings</router-link>
            </li>
          </ul>
        </div>

        <div class="sticky inset-x-0 bottom-0 border-t border-gray-100">
          <a href="/" target="_blank" class="flex items-center px-6 text-sm font-medium text-gray-500 bg-white p-4 hover:bg-green-50 hover:text-green-500">
            Open site
          </a>

          <a href="#" class="flex items-center px-6 text-sm font-medium text-gray-500 bg-white p-4 hover:bg-red-50 hover:text-red-500" @click="signOut">
            Logout
          </a>
        </div>
      </div>
    </div>
    <div class="flex-1 mt-5 mx-5">
      <slot />
    </div>
  </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()

const signOut = async () => {
  await fetch('/api/sign/out', {
    credentials: "include",
    method: 'POST',
  })
    .then(response => response.json())
    .then(data => {
      if (data.success) {
        window.location.href = "/_/signin/";
      }
    })
};
</script>