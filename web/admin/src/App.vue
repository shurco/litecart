<template>
  <RouterView />

  <NotificationGroup group="bottom" position="bottom">
    <div class="fixed inset-x-0 bottom-0 flex items-start justify-end p-6 px-4 py-6 pointer-events-none" style="z-index:999;">
      <div class="w-full max-w-sm">
        <Notification v-slot="{ notifications }" enter="ease-out duration-300 transition" enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0" leave="transition ease-in duration-500" leave-from="opacity-100" leave-to="opacity-0" move="transition duration-500"
          move-delay="delay-300">
          <div v-for="notification in notifications" :key="notification.id" class="flex w-full max-w-sm mx-auto mt-4 overflow-hidden bg-white rounded-lg shadow-md">
              <div class="flex items-center justify-center w-12" :class="(notification.type === 'error' ? 'bg-red-500' : 'bg-green-500')">
                <SvgIcon name="exclamation" class="h-6 w-6 text-white fill-current" />
              </div>
              <div class="px-4 py-2 -mx-3">
                <div class="mx-3">
                  <span class="font-semibold" :class="(notification.type === 'error' ? 'text-red-500' : 'text-green-500')">{{notification.title}}</span>
                  <p class="text-sm text-gray-600">
                    {{ notification.text }}
                  </p>
                </div>
              </div>
          </div>
        </Notification>
      </div>
    </div>
  </NotificationGroup>
</template>

<script setup>
import { RouterView } from "vue-router";
import { Notification, NotificationGroup } from "notiwind";
import SvgIcon from "svg-icon";
</script>
