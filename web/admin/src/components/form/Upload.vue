<template>
  <div class="upload bg-gray-200" @dragover="dragover" @dragleave="dragleave" @drop="drop">
    <input type="file" multiple name="fields[assetsFieldHandle][]" id="assetsFieldHandle" @change="onChange" ref="file" :accept="accept" />
    <label for="assetsFieldHandle">
      <SvgIcon name="plus" class="h-5 w-5" />
    </label>
  </div>
</template>

<script setup>
import { getCurrentInstance } from "vue";
import { apiPost } from "@/utils/api";

import SvgIcon from "svg-icon";

const props = defineProps({
  section: {
    type: String,
    required: true,
  },
  accept: {
    type: String,
  },
  productId: String,
  status: Boolean,
});

const instance = getCurrentInstance();
const emits = defineEmits(["added"]);

const onChange = () => {
  [...instance.refs.file.files].forEach((f) => {
    const formData = new FormData();
    formData.append("document", f);
    apiPost(`/api/_/products/${props.productId}/${props.section}`, formData).then(res => {
      emits("added", res);
    })
  });
};

const dragover = (event) => {
  event.preventDefault();
  if (!event.currentTarget.classList.contains("bg-green-300")) {
    event.currentTarget.classList.remove("bg-gray-100");
    event.currentTarget.classList.add("bg-green-300");
  }
};

const dragleave = (event) => {
  event.currentTarget.classList.add("bg-gray-100");
  event.currentTarget.classList.remove("bg-green-300");
};

const drop = (event) => {
  event.preventDefault();
  instance.refs.file.files = event.dataTransfer.files;
  onChange();
  event.currentTarget.classList.add("bg-gray-100");
  event.currentTarget.classList.remove("bg-green-300");
};
</script>

<style lang="scss" scoped>
.upload {
  @apply grid h-16 cursor-pointer place-content-center rounded-lg;

  input {
    @apply absolute h-px w-px overflow-hidden opacity-0;
  }

  label {
    @apply block cursor-pointer p-0 border-0 shadow-none;
  }
}
</style>
