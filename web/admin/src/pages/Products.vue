<template>
  <MainLayouts>
    <div class="grid grid-cols-1 gap-4 pb-4 lg:grid-cols-[1fr_120px] lg:gap-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">Products</h1>
      </div>
      <div>
        <FormButton type="submit" name="Add" color="green" ico="arrow-right" @click="openDrawer(null, null, 'add')" />
      </div>
    </div>

    <div class="mx-auto" v-if="products.total > 0">
      <table class="min-w-full divide-y-2 divide-gray-200 bg-white text-sm">
        <thead class="text-left">
          <tr>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900">
              Name
            </th>
            <th class="w-32 whitespace-nowrap px-4 py-2 font-medium text-gray-900">
              URL
            </th>
            <th class="w-32 whitespace-nowrap px-4 py-2 font-medium text-gray-900">
              Price
            </th>
            <th class="w-24 px-4 py-2"></th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr class="cursor-pointer hover:bg-gray-100 active:bg-gray-100" :class="{ 'opacity-30': !item.active }" v-for="(item, index) in products.products">
            <td class="whitespace-nowrap px-4 py-2" @click="openDrawer(index, item.id, 'view')">
              <div>
                {{ item.name }}
              </div>
            </td>
            <td class="whitespace-nowrap px-4 py-2">
              <a :href="`/products/${item.url}`" target="_blank" v-if="item.active">{{ item.url }}</a>
              <span v-else>{{ item.url }}</span>
            </td>
            <td class="whitespace-nowrap px-4 py-2" @click="openDrawer(index, item.id, 'view')">
              {{ costFormat(item.amount) }} {{ item.currency }}
            </td>
            <td class="px-4 py-2">
              <div class="flex">
                <div class="pr-3">
                  <SvgIcon name="pencil-square" class="h-5 w-5" @click="openDrawer(index, item.id, 'update')" />
                </div>
                <div>
                  <SvgIcon :name="(item.active ? 'eye' : 'eye-slash')" class="h-5 w-5" @click="updateProductActive(index)" />
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="mx-auto" v-else>Add first product</div>

    <drawer :is-open="isDrawer.open" max-width="700px" @close="closeDrawer">
      <div class="flow-root" v-if="isDrawer.action === 'view'">
        <dl class="-my-3 divide-y divide-gray-100 text-sm">
          <DetailList name="ID">{{ product.info.id }}</DetailList>
          <DetailList name="Name">{{ product.info.name }}</DetailList>
          <DetailList name="Price">{{ product.info.amount }} {{ product.info.currency }}</DetailList>
          <DetailList name="URL">{{ product.info.url }}</DetailList>

          <DetailList name="Metadata">
            <div v-for="(value, key) in product.info.metadata">
              {{ key }}: {{ value }}
            </div>
          </DetailList>

          <DetailList name="Attributes">
            <div v-for="item in product.info.attributes">{{ item }}</div>
          </DetailList>

          <DetailList name="Created">{{
            formatDate(product.info.created)
          }}</DetailList>
          <DetailList name="Updated" v-if="product.info.updated">
            {{ formatDate(product.info.updated) }}
          </DetailList>

          <DetailList name="Images" :grid="true" v-if="product.info.images !== null">
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

      <template v-slot:header v-if="isDrawer.action === 'view'">
        <div class="flex items-center">
          <div class="pr-3">
            <h2 class="text-2xl font-bold text-gray-900 sm:text-3xl">
              {{ products.products[product.index].name }}
            </h2>
          </div>
          <div>
            <SvgIcon :name="(products.products[product.index].active ? 'eye' : 'eye-slash')" class="h-5 w-5 cursor-pointer" @click="updateProductActive(product.index)" />
          </div>
        </div>
      </template>

      <div class="flow-root" v-if="isDrawer.action === 'update'">
        <dl class="-my-3 divide-y divide-gray-100 text-sm">
          <Form @submit="updateProduct" v-slot="{ errors }" class="mx-auto mb-0 mt-8 space-y-4">
            <FormInput v-model.trim="product.info.name" id="name" type="text" title="Name" ico="at-symbol" />

            <div class="flex flex-row">
              <div class="pr-3">
                <FormInput v-model="product.info.amount" id="amount" type="text" title="Amount" ico="money" />
              </div>
              <div class="mt-3">{{ product.info.currency }}</div>
            </div>
            <FormInput v-model="product.info.url" id="url" type="text" title="Url" ico="glob-alt" />

            <hr />
            <p class="font-semibold">Metadata</p>
            <div class="flex" v-for="(value, key, index) in product.info.metadata" :key="index">
              <div class="grow pr-3">
                <FormInput v-model="product.info.metadata[key]" :id="`mtd-value-${index}`" type="text" :title="`Value-${index}`" />
              </div>
              <div class="grow">
                <FormInput v-model="Object.keys(product.info.metadata)[index]" :id="`mtd-key-${index}`" type="text" :title="`Key-${index}`" />
              </div>
              <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteMetadataRecord(index, product.info.id)">
                <SvgIcon name="trash" class="h-5 w-5" />
              </div>
            </div>
            <div class="flex">
              <div class="grow"></div>
              <div class="mt-2 flex-none">
                <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addMetadataRecord">
                  Add metadata record
                </a>
              </div>
            </div>

            <hr />
            <p class="font-semibold">Attributes</p>
            <div class="flex" v-for="(value, index) in product.info.attributes" :key="index">
              <div class="grow">
                <FormInput v-model="product.info.attributes[index]" :id="`atr-key-${index}`" type="text" :title="`${index}`" />
              </div>
              <div class="flex-none cursor-pointer pl-3 pt-3">
                <SvgIcon name="trash" class="h-5 w-5" />
              </div>
            </div>
            <div class="flex">
              <div class="grow"></div>
              <div class="mt-2 flex-none">
                <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addAttributeRecord">
                  Add attribute record
                </a>
              </div>
            </div>

            <hr />
            <div class="grid grid-cols-4 content-start gap-4" v-if="product.info.images !== null">
              <div v-for="(item, index) in product.info.images" class="relative" style="width: 100%; max-width: 150px">
                <a :href="`/uploads/${item.name}.${item.ext}`" target="_blank">
                  <img :src="`/uploads/${item.name}_sm.${item.ext}`" />
                </a>
                <div class="absolute end-4 top-4 cursor-pointer bg-white p-2" @click="deleteProductImage(index, product.info.id)">
                  <SvgIcon name="trash" class="h-5 w-5" />
                </div>
              </div>
            </div>
            <FormUpload :productId="product.info.id" accept=".jpg,.jpeg,.png" @added="addProductImage" />

            <hr />
            <FormTextarea v-model="product.info.description" id="textarea" name="textarea"></FormTextarea>
          </Form>
        </dl>
      </div>

      <template v-slot:header v-if="isDrawer.action === 'update'">
        <div class="flex items-center">
          <div class="pr-3">
            <h2 class="text-2xl font-bold text-gray-900 sm:text-3xl">Edit</h2>
          </div>
        </div>
      </template>

      <template v-slot:footer v-if="isDrawer.action === 'update'">
        <div class="flex">
          <div class="flex-none">
            <FormButton type="submit" name="Save" color="green" class="mr-3" @click="updateProduct" />
            <FormButton type="submit" name="Close" color="gray" @click="closeDrawer" />
          </div>
          <div class="grow"></div>
          <div class="mt-2 flex-none">
            <span @click="deleteProduct(product.index)" class="cursor-pointer text-red-700">Delete</span>
          </div>
        </div>
      </template>

      <div class="flow-root" v-if="isDrawer.action === 'add'">
        <dl class="-my-3 divide-y divide-gray-100 text-sm">
          <FormInput v-model="product.info.name" id="name" type="text" name="Name" ico="at-symbol" />
        </dl>
      </div>

      <template v-slot:header v-if="isDrawer.action === 'add'">
        <div class="flex items-center">
          <div class="pr-3">
            <h2 class="text-2xl font-bold text-gray-900 sm:text-3xl">Add</h2>
          </div>
        </div>
      </template>
    </drawer>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";

// @ts-ignore
import * as NProgress from "nprogress";

import SvgIcon from 'svg-icon'
import { defineRule, Form } from "vee-validate";
import { required, email, min } from "@vee-validate/rules";
defineRule("required", required);
defineRule("email", email);
defineRule("min", min);

import { costFormat, formatDate } from "@/utils/";

import MainLayouts from "@/layouts/Main.vue";
import FormInput from "@/components/form/Input.vue";
import FormButton from "@/components/form/Button.vue";
import FormUpload from "@/components/form/Upload.vue";
import FormTextarea from "@/components/form/Textarea.vue";
import DetailList from "@/components/DetailList.vue";
import Drawer from "@/components/Drawer.vue";

onMounted(() => {
  listProducts();
});

const isDrawer = ref({
  open: false,
  action: null,
});

const products = ref({});

const product = ref({
  info: {
    id: String,
    name: String,
    amount: String,
    url: String,
    metadata: Object,
    description: String,
  },
  action: String,
  index: Number,
  name: String,
});

const listProducts = async () => {
  NProgress.start();

  await fetch("/api/_/products", {
    credentials: "include",
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        products.value = data.result;
      }
      NProgress.done();
    });
};

const getProduct = async (id) => {
  NProgress.start();

  await fetch(`/api/_/products/${id}`, {
    credentials: "include",
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        product.value.info = data.result;
        product.value.info.amount = costFormat(product.value.info.amount);
      }
      NProgress.done();
    });
};

const updateProduct = (index) => {
  console.log(index);
};

const addProduct = (index) => {
  console.log(index);
};

const deleteProduct = async (index) => {
  NProgress.start();

  await fetch(`/api/_/products/${products.value.products[index].id}`, {
    credentials: "include",
    method: "DELETE",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        products.value.products.splice(index, 1);
        products.value.total = products.value.total - 1;
        closeDrawer();
      } else {
        const obj = JSON.parse(data.result);
        if (obj.code === "resource_missing") {
          console.log(obj.message);
        }
      }
      NProgress.done();
    });
};

const updateProductActive = async (index) => {
  NProgress.start();

  await fetch(`/api/_/products/${products.value.products[index].id}/active`, {
    credentials: "include",
    method: "PATCH",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        products.value.products[index].active =
          !products.value.products[index].active;
      }
      NProgress.done();
    });
};

const addMetadataRecord = (index) => {
  console.log(index);
};

const deleteMetadataRecord = (index, productId) => {
  console.log(index);
};

const addAttributeRecord = (index) => {
  console.log(index);
};

const addProductImage = (event) => {
  if (event.success) {
    product.value.info.images.push(event.result);
  }
};

const deleteProductImage = async (index, productId) => {
  NProgress.start();
  const image = product.value.info.images[index];

  await fetch(`/api/_/products/${productId}/image/${image.id}`, {
    credentials: "include",
    method: "DELETE",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        product.value.info.images.splice(index, 1);
      }
      NProgress.done();
    });
};

const openDrawer = (index, id, action) => {
  isDrawer.value.open = true;
  isDrawer.value.action = action;
  product.value.index = index;

  switch (action) {
    case "view":
    case "update":
      getProduct(id);
      break;
    case "add":
      addProduct("add");
      break;
  }
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
  product.value.info = {};
  product.value.index = null;
};
</script>
