<template>
  <div>
    <label :for="id" class="relative block rounded border border-gray-200 shadow-sm text-sm focus-within:border-blue-600 focus-within:ring-1 focus-within:ring-blue-600"
      :class="error ? 'border-red-500' : ''">

      <textarea :value="model" id="id" :name="id"
        class="w-full peer resize-none border-none bg-transparent placeholder-transparent focus:border-transparent focus:outline-none focus:ring-0" rows="4"
        placeholder="Enter any additional order notes...">{{ $slots.default ? $slots.default()[0].children : '' }}</textarea>

      <span class="pointer-events-none absolute start-2.5 top-0 -translate-y-1/2 bg-white p-0.5 text-xs text-gray-700 peer-focus:top-0 peer-focus:text-xs">
        {{ name }}
      </span>

    </label>
    <span class="text-sm text-red-500 pl-4" v-if="error">{{ error }}</span>
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  modelValue: {
    required: true
  },
  id: {
    type: String,
    default: 'name'
  },
  name: {
    type: String,
    default: 'Name'
  },
  color: {
    type: String,
    default: 'indigo'
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