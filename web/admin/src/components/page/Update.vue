<template>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>Page setup</h1>
      </div>
    </div>
  </div>

  <Form @submit="updatePage" v-slot="{ errors }">
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

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name="Save" color="green" class="mr-3" />
          <FormButton type="submit" name="Close" color="gray" @click="close" />
        </div>
        <div class="grow"></div>
        <div class="mt-4 flex-none">
          <span @click="deletePage" class="cursor-pointer text-red-700">Delete</span>
        </div>
      </div>
    </div>
  </Form>
</template>

<script setup>
import { computed } from "vue";
import { FormInput, FormButton, FormSelect } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiUpdate, apiDelete } from "@/utils/api";
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

const updatePage = async () => {
  const update = { ...page.value };
  apiUpdate(`/api/_/pages/${update.id}`, update).then(res => {
    if (res.success) {
      const found = pages.value.find((e) => e.id === update.id);
      found.name = update.name;
      found.slug = update.slug;
      found.position = update.position;
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const deletePage = async () => {
  const index = page.value.index;
  apiDelete(`/api/_/pages/${pages.value[index].id}`).then(res => {
    if (res.success) {
      pages.value.splice(index, 1);
    } else {
      const obj = JSON.parse(res.result);
      if (obj.code === "resource_missing") {
        console.log(obj.message);
      }
    }
    props.close();
  });
};
</script>
