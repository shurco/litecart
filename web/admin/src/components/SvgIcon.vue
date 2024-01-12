
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
    default: "icon",
  },
  name: {
    type: String,
    required: true,
  },
  class: {
    type: String,
    default: "h-6 w-6",
  },
});

const symbolId = computed(() => `#${props.prefix}-${props.name}`);
</script>