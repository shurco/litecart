<template>
  <span class="badge">
    <SvgIcon :name="ico" stroke="currentColor"  :class="ico ? '-ms-0.5 me-1.5 ' + svgClass : 'h-0 w-0'" v-if="ico" />
    <p>
      <slot />
    </p>
  </span>
</template>

<script setup>
const props = defineProps({
  ico: String,
  svgClass: {
    type: String,
    required: false,
    default: "-ms-0.5 me-5.5 h-5 w-5",
  },
});
</script>

<style lang="scss" scoped>
.badge {
  @apply inline-flex items-center justify-center rounded-full px-2.5 py-0.5;

  & p {
    @apply whitespace-nowrap text-xs;
  }
}
</style>
