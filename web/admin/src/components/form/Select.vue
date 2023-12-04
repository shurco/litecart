<template>
  <div>
    <label :for="id" :class="error ? 'border-red-500' : ''">
      <Field v-slot="{ value }" v-model="model" as="select" :name="id" :id="id" :rules="rules" class="form-select field peer">
        <option value="" disabled>Please select</option>
        <option v-for="option in options" :key="option" :value="option" :selected="value && value.includes(option)">
          {{ option }}
        </option>
      </Field>

      <span class="title">{{ title }}</span>
      <span class="ico" v-if="ico">
        <SvgIcon :name="ico" stroke="currentColor" class="h-5 w-5" :class="error ? 'text-red-500' : 'text-gray-400'" />
      </span>
    </label>
    <span class="error" v-if="error">{{ error }}</span>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { Field } from "vee-validate";

const props = defineProps({
  modelValue: {
    type: String,
    default: "",
    required: true,
  },
  id: {
    type: String,
    default: "name",
  },
  title: {
    type: String,
    default: "Name",
  },
  options: {
    type: Object,
    required: true,
  },
  rules: {
    type: String,
    default: "",
  },
  color: {
    type: String,
    default: "indigo",
  },
  ico: String,
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
