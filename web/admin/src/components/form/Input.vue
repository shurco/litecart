<template>
  <div>
    <label :for="id" :class="error ? 'border-red-500' : ''">
      <Field :type="type" :name="id" :rules="rules" :id="id" v-model="model" class="form-input field peer" :placeholder="placeholder" autocomplete="on" />
      <span v-if="title" class="peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-focus:top-0 peer-focus:text-xs peer-placeholder-shown:text-gray-400 peer-focus:text-gray-700  title">
        {{ title }}
      </span>
      <span class="ico" v-if="ico">
        <SvgIcon :name="ico" stroke="currentColor" class="h-5 w-5" :class="error ? 'text-red-500' : 'text-gray-400'" />
      </span>
    </label>
    <span class="text-sm text-red-500 pl-4" v-if="error">{{ error }}</span>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { Field } from 'vee-validate';

const props = defineProps({
  modelValue: {
    required: true
  },
  id: {
    type: String,
    default: 'name'
  },
  type: {
    type: String,
    default: 'text'
  },
  title: {
    type: String,
    default: 'Name'
  },
  color: {
    type: String,
    default: 'indigo'
  },
  rules: String,
  ico: String,
  error: String
})

const placeholder = 'Enter ' + props.id
const emits = defineEmits(['update:modelValue'])
const model = computed({
  get: () => {
    return props.modelValue
  },
  set: (val) => {
    emits('update:modelValue', val)
  }
})
</script>
