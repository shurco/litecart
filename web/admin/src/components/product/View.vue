<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>View {{ product.name }}</h1>
        </div>
        <div>
          <SvgIcon :name="product.active ? 'eye' : 'eye-slash'" class="h-5 w-5 cursor-pointer" @click="active" stroke="currentColor" />
        </div>
      </div>
    </div>

    <div class="flow-root">
      <dl class="-my-3 mt-2 divide-y divide-gray-100 text-sm">
        <DetailList name="ID">{{ product.id }}</DetailList>
        <DetailList name="Name">{{ product.name }}</DetailList>
        <DetailList name="Price">{{ costFormat(product.amount) }} {{ drawer.currency }}</DetailList>
        <DetailList name="Slug">{{ product.slug }}</DetailList>
        <DetailList name="Metadata">
          <div v-for="(data, index) in product.metadata">
            {{ data.key }}: {{ data.value }}
          </div>
        </DetailList>
        <DetailList name="Attributes">
          <div v-for="item in product.attributes">{{ item }}</div>
        </DetailList>
        <DetailList name="Created">{{ formatDate(product.created) }}</DetailList>
        <DetailList name="Updated" v-if="product.updated">{{ formatDate(product.updated) }}</DetailList>
        <DetailList name="Images" :grid="true" v-if="product.images">
          <div v-for="item in product.images">
            <a :href="`/uploads/${item.name}.${item.ext}`" target="_blank">
              <img style="width: 100%; max-width: 150px" :src="`/uploads/${item.name}_sm.${item.ext}`" loading="lazy" />
            </a>
          </div>
        </DetailList>

        <DetailList name="Brief (short description)">{{ product.brief }}</DetailList>

        <div v-html="product.description" class="pt-3 tiptap"></div>
      </dl>
    </div>

    <div class="pt-8">
      <FormButton type="submit" name="Close" color="green" @click="close" />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

import FormButton from "@/components/form/Button.vue";
import DetailList from "@/components/DetailList.vue";
import { costFormat, formatDate } from "@/utils/";
import { apiGet } from "@/utils/api";

const props = defineProps({
  drawer: {
    required: true,
  },
  updateActive: Function,
  close: Function,
});

const product = ref({})

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

const active = async () => {
  props.updateActive(props.drawer.product.index);
  product.value.active = !product.value.active;
};
</script>
