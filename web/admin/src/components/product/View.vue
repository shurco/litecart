<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>View {{ products.products[product.index].name }}</h1>
        </div>
        <div>
          <SvgIcon :name="products.products[product.index].active ? 'eye' : 'eye-slash'" class="h-5 w-5 cursor-pointer" @click="updateActive(product.index)" />
        </div>
      </div>
    </div>

    <div class="flow-root">
      <dl class="-my-3 mt-2 divide-y divide-gray-100 text-sm">
        <DetailList name="ID">{{ product.info.id }}</DetailList>
        <DetailList name="Name">{{ product.info.name }}</DetailList>
        <DetailList name="Price">{{ product.info.amount }} {{ products.currency }}</DetailList>
        <DetailList name="Slug">{{ product.info.slug }}</DetailList>
        <DetailList name="Metadata">
          <div v-for="(data, index) in product.info.metadata">
            {{ data.key }}: {{ data.value }}
          </div>
        </DetailList>
        <DetailList name="Attributes">
          <div v-for="item in product.info.attributes">{{ item }}</div>
        </DetailList>
        <DetailList name="Created">{{ formatDate(product.info.created) }}</DetailList>
        <DetailList name="Updated" v-if="product.info.updated">{{ formatDate(product.info.updated) }}</DetailList>
        <DetailList name="Images" :grid="true" v-if="product.info.images">
          <div v-for="item in product.info.images">
            <a :href="`/uploads/${item.name}.${item.ext}`" target="_blank">
              <img style="width: 100%; max-width: 150px" :src="`/uploads/${item.name}_sm.${item.ext}`" loading="lazy" />
            </a>
          </div>
        </DetailList>

        <DetailList name="description">{{
          product.info.description
        }}</DetailList>
      </dl>
    </div>

    <div class="pt-8">
      <FormButton type="submit" name="Close" color="green" @click="close" />
    </div>
  </div>
</template>

<script setup>
import FormButton from "@/components/form/Button.vue";
import DetailList from "@/components/DetailList.vue";
import { formatDate } from "@/utils/";

import SvgIcon from "svg-icon";

const props = defineProps({
  product: {
    required: true,
  },
  products: {
    required: true,
  },
  updateActive: Function,
  close: Function,
});
</script>
