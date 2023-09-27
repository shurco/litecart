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
          <FormInput v-model="product.seo.title" :error="errors.name" id="title" type="text" title="Title" ico="glob-alt" />
          <FormInput v-model="product.seo.keywords" :error="errors.slug" id="keywords" type="text" title="Keywords" ico="glob-alt" />
          <hr />
          <FormTextarea v-model="product.seo.description" id="textarea" name="Description" />
        </dl>
      </div>

      <div class="pt-8">
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
import { onMounted, ref } from "vue";

import FormInput from "@/components/form/Input.vue";
import FormButton from "@/components/form/Button.vue";
import FormTextarea from "@/components/form/Textarea.vue";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

import { Form } from "vee-validate";

const props = defineProps({
  drawer: {
    required: true,
  },
  close: Function,
});

const product = ref({
  "seo": {
    "title": "",
    "keywords": "",
    "description": ""
  }
})

onMounted(() => {
  getProduct(props.drawer.product.id)
});

const getProduct = async (id) => {
  apiGet(`/api/_/products/${id}`).then(res => {
    if (res.success) {
      product.value = res.result;
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const updateSeo = async () => {
  apiUpdate(`/api/_/products/${product.value.id}`, product.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
