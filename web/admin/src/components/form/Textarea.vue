<template>
  <div class="textarea">
    <label :for="id" :class="error ? '!border-red-500' : ''">
      <Field as="textarea" v-model="model" :rules="rules" :id="id" :name="id" class="form-textarea text peer" :rows="rows" placeholder="Enter any additional order notes...">
        {{ $slots.default ? $slots.default()[0].children : "" }}
      </Field>
      <span class="title">{{ name }}</span>
    </label>
    <span class="text-sm text-red-500 pl-4" v-if="error">{{ error }}</span>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { Field } from 'vee-validate'

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
  rules: String,
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

<style scoped>
@reference "../../assets/app.css";

.textarea label {
  @apply p-0;
}

.textarea label .text {
  @apply resize-none;
}

.textarea label .title {
  @apply bg-zinc-50 p-0.5 peer-focus:top-0 peer-focus:text-xs;
}
</style>
