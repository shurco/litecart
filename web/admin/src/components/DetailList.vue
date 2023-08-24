<template>
  <div class="grid grid-cols-1 gap-1 py-3 sm:grid-cols-3 sm:gap-4">
    <dt class="font-medium text-gray-900">{{ name }}</dt>
    <dd class="text-gray-700 sm:col-span-2" :class="grid ? 'grid grid-cols-3 gap-4 content-start' : ''">
      <slot />
    </dd>
  </div>
</template>

<script setup>
const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  grid: {
    type: Boolean,
    default: false,
  }
});
</script>