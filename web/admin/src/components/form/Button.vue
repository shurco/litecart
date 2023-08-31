<template>
  <button class="group" :class="color ? colors[color][0] + ' active:' + colors[color][1] : '', ico ? 'focus:outline-none focus:ring' : ''">
    <SvgIcon :name="ico" class="h-4 w-4 absolute -start-full transition-all group-hover:start-4" v-if="ico" />
    <span :class="ico ? 'transition-all group-hover:ms-4' : ''">
      {{ name }}
    </span>
  </button>
</template>

<script setup>
import SvgIcon from 'svg-icon'

const props = defineProps({
  color: {
    type: String,
    required: true
  },
  name: {
    type: String,
    default: 'Name'
  },
  ico: String,
});

const colors = {
  gray: ["bg-gray-600", "bg-gray-500"],
  gray_lite: ["bg-gray-400", "bg-gray-300"],
  green: ["bg-green-600", "bg-green-500"],
  yellow: ["bg-yellow-600", "bg-yellow-500"],
  red: ["bg-red-600", "bg-red-500"],
}
</script>

<style lang="scss" scoped>
button {
  @apply relative inline-flex items-center overflow-hidden rounded px-8 py-3 text-white;

  & span {
    @apply text-sm font-medium;
  }
}
</style>