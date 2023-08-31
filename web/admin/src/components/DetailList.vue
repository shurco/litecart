<template>
  <div class="detail-list">
    <dt>{{ name }}</dt>
    <dd :class="grid ? 'grid grid-cols-3 content-start gap-4' : ''">
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
  },
});
</script>

<style lang="scss" scoped>
.detail-list {
  @apply grid grid-cols-1 gap-1 py-3 sm:grid-cols-3 sm:gap-4;

  & dt {
    @apply font-medium text-gray-900;
  }

  & dd {
    @apply text-gray-700 sm:col-span-2;
  }
}
</style>
