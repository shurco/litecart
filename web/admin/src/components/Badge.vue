<template>
  <span class="inline-flex items-center justify-center rounded-full px-2.5 py-0.5">
    <SvgIcon :name="ico" :class="(ico ? '-ms-0.5 me-1.5 ' + svgClass : 'h-0 w-0')" v-if="ico" />
    <p class="whitespace-nowrap text-xs">
      <slot />
    </p>
  </span>
</template>

<script setup>
import SvgIcon from 'svg-icon'

const props = defineProps({
  ico: String,
  svgClass: {
    type: String,
    required: false,
    default: "-ms-0.5 me-5.5 h-5 w-5",
  },
});
</script>