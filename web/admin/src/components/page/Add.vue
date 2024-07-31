<template>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>New page</h1>
      </div>
    </div>
  </div>

  <Form @submit="addPage" v-slot="{ errors }">
    <div class="flow-root">
      <dl class="-my-3 mx-auto mb-0 mt-4 space-y-4 text-sm">
        <FormInput v-model.trim="page.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="name" ico="at-symbol" />
        <div class="flex">
          <div class="pr-3">
            <FormSelect v-model="page.position" :options="positionPage" :error="errors.position" rules="required" id="position" title="Position" />
          </div>
          <div>
            <FormInput v-model.trim="page.slug" :error="errors.slug" rules="required|slug" id="slug" type="text" title="Slug" ico="glob-alt" />
          </div>
        </div>
      </dl>
    </div>

    <div class="pt-5">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name="Add" color="green" class="mr-3" />
          <FormButton type="submit" name="Close" color="gray" @click="close" />
        </div>
        <div class="grow"></div>
      </div>
    </div>
  </Form>
</template>

<script setup>
import { computed } from "vue";
import { FormInput, FormButton, FormSelect } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiPost } from "@/utils/api";
import { Form } from "vee-validate";

const positionPage = ["header", "footer"];
const props = defineProps({
  page: {
    required: true,
  },
  pages: {
    required: true,
  },
  updateActive: Function,
  close: Function,
});

const emits = defineEmits(["update:modelValue"]);
const page = computed({
  get: () => {
    return props.page;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

const pages = computed({
  get: () => {
    return props.pages;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

const addPage = async () => {
  const add = { ...page.value };
  apiPost(`/api/_/pages`, add).then(res => {
    if (!Array.isArray(pages.value)) {
      pages.value = [];
    }

    if (res.success) {
      pages.value.push({
        id: res.result.id,
        name: res.result.name,
        slug: res.result.slug,
        position: res.result.position,
        created: res.result.created,
        active: res.result.active,
      });
      props.close();
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
