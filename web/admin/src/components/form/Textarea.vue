<template>
  <div class="textarea">
    <label :for="id" :class="error ? 'border-red-500' : ''">
      <textarea v-model="model" id="id" :name="id" class="text peer" :rows="rows" placeholder="Enter any additional order notes..." style="">
        {{ $slots.default ? $slots.default()[0].children : "" }}
      </textarea>
      <span class="title">{{ name }}</span>
    </label>
    <span class="error" v-if="error">{{ error }}</span>
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  modelValue: {
    required: true,
  },
  id: {
    type: String,
    default: "name",
  },
  name: {
    type: String,
    default: "Name",
  },
  color: {
    type: String,
    default: "indigo",
  },
  rows: {
    type: Number,
    default: 4,
  },
  error: String,
});

const emits = defineEmits(["update:modelValue"]);

const model = computed({
  get: () => {
    return props.modelValue;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});
</script>

<style lang="scss" scoped>
.textarea {
  & label {
    @apply relative block rounded border border-gray-200 text-sm shadow-sm focus-within:border-blue-600 focus-within:ring-1 focus-within:ring-blue-600 p-0;

    & .text {
      @apply w-full resize-none border-none bg-transparent placeholder-transparent focus:border-transparent focus:outline-none focus:ring-0;
    }

    .title {
      @apply pointer-events-none absolute start-2.5 top-0 -translate-y-1/2 bg-zinc-50 p-0.5 text-xs text-gray-700 peer-focus:top-0 peer-focus:text-xs;
    }
  }
}
</style>
