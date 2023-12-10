<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>SEO</h1>
        </div>
      </div>
    </div>

    <Form @submit="updateSeo" v-slot="{ errors }">
      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model="page.seo.title" :error="errors.name" id="title" type="text" title="Title" ico="glob-alt" />
          <FormInput v-model="page.seo.keywords" :error="errors.slug" id="keywords" type="text" title="Keywords" ico="glob-alt" />
          <hr />
          <FormTextarea v-model="page.seo.description" id="textarea" name="Description" />
        </dl>
      </div>
      <div class="pt-5">
        <div class="flex">
          <div class="flex-none">
            <FormButton type="submit" name="Save" color="green" class="mr-3" />
            <FormButton type="submit" name="Close" color="gray" @click="close" />
          </div>
          <div class="grow"></div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { FormInput, FormButton, FormTextarea } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const props = defineProps({
  page: {
    required: true,
  },
  updateActive: Function,
  close: Function,
});

const page = computed({
  get: () => {
    return props.page;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

const updateSeo = async () => {
  const update = { ...page.value };
  apiUpdate(`/api/_/pages/${update.id}`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
