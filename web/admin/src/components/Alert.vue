<template>
  <NotificationGroup group="alerts">
    <div class="alert">
      <div class="w-full max-w-sm">
        <Notification v-slot="{ notifications }" enter="transform ease-out duration-300 transition" enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0" leave="transition ease-in duration-500" leave-from="opacity-100" leave-to="opacity-0" move="transition duration-500"
          move-delay="delay-300">
          <div v-for="notification in notifications" :key="notification.id">
            <div v-if="notification.type === 'error'" class="notification">
              <div class="ico bg-red-500">
                <SvgIcon name="exclamation" stroke="currentColor" class="h-6 w-6 text-white fill-current" />
              </div>

              <div class="message">
                <div class="mx-3">
                  <span class="font-semibold text-red-500">Error</span>
                  <p>{{ notification.text }}</p>
                </div>
              </div>
            </div>

            <div v-if="notification.type === 'info'" class="notification">
              <div class="ico bg-blue-500">
                <SvgIcon name="exclamation" stroke="currentColor" class="h-6 w-6 text-white fill-current" />
              </div>

              <div class="message">
                <div class="mx-3">
                  <span class="font-semibold text-blue-500">Info</span>
                  <p>{{ notification.text }}</p>
                </div>
              </div>
            </div>

            <div v-if="notification.type === 'success'" class="notification">
              <div class="ico bg-green-500">
                <SvgIcon name="exclamation" stroke="currentColor" class="h-6 w-6 text-white fill-current" />
              </div>

              <div class="message">
                <div class="mx-3">
                  <span class="font-semibold text-green-500">Success</span>
                  <p>{{ notification.text }}</p>
                </div>
              </div>
            </div>

            <div v-if="notification.type === 'warning'" class="notification">
              <div class="ico bg-yellow-500">
                <SvgIcon name="exclamation" stroke="currentColor" class="h-6 w-6 text-white fill-current" />
              </div>

              <div class="message">
                <div class="mx-3">
                  <span class="font-semibold text-yellow-500">Warning</span>
                  <p>{{ notification.text }}</p>
                </div>
              </div>
            </div>
          </div>
        </Notification>
      </div>
    </div>
  </NotificationGroup>
</template>

<script setup>
import { notify, Notification, NotificationGroup } from "notiwind";

const NOTIFICATION_DURATION = 4000;

const handleNotification = (eventName, notificationType) => {
  const callback = (e) => {
    notify(
      {
        group: "alerts",
        type: notificationType,
        text: e.detail,
      },
      NOTIFICATION_DURATION
    );
  };

  addEventListener(eventName, callback);
  // To remove event listener:
  // removeEventListener(eventName, callback);
};

handleNotification("connextError", "error");
handleNotification("connextSuccess", "success");
handleNotification("connextWarning", "warning");
handleNotification("connextInfo", "info");
</script>

<style lang="scss" scoped>
.alert {
  @apply fixed inset-x-0 bottom-0 flex items-start justify-end p-6 px-4 py-6 pointer-events-none;
  z-index: 999;

  & .notification {
    @apply mx-auto mt-4 flex w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-md;

    & .ico {
      @apply flex w-12 items-center justify-center;
    }

    & .message {
      @apply -mx-3 px-4 py-2;

      p {
        @apply text-sm text-gray-600;
      }
    }
  }
}
</style>