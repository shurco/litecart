<template>
  <svg aria-hidden="true" :class="class">
    <use :xlink:href="symbolId" />
  </svg>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  prefix: {
    type: String,
    default: "sprite",
  },
  name: {
    type: String,
    required: true,
  },
  class: {
    type: String,
  },
});

const symbolId = computed(() => `/__spritemap#${props.prefix}-${props.name}`);
</script>