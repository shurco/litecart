<template>
  <MainLayouts>
    <header>
      <h1>Products</h1>
      <div>
        <FormButton type="submit" name="Add" color="green" ico="arrow-right" @click="openDrawer(null, null, 'add')" />
      </div>
    </header>

    <div class="mx-auto" v-if="products.total > 0">
      <table>
        <thead>
          <tr>
            <th class="w-28"></th>
            <th>Name</th>
            <th class="w-32">URL</th>
            <th class="w-32">Price</th>
            <th class="w-24 px-4 py-2"></th>
          </tr>
        </thead>
        <tbody>
          <tr :class="{ 'opacity-30': !item.active }" v-for="(item, index) in products.products">
            <td>
              <a :href="`/uploads/${item.images[0].name}.${item.images[0].ext}`" target="_blank" v-if="item.images">
                <img style="width: 100%; max-width: 80px" :src="`/uploads/${item.images[0].name}_sm.${item.images[0].ext}`" loading="lazy" />
              </a>
              <img style="width: 100%; max-width: 80px" src="/assets/img/noimage.png" v-else />
            </td>
            <td @click="openDrawer(index, item.id, 'view')">
              <div>{{ item.name }}</div>
            </td>
            <td>
              <a :href="`/products/${item.url}`" target="_blank" v-if="item.active">{{ item.url }}</a>
              <span v-else>{{ item.url }}</span>
            </td>
            <td @click="openDrawer(index, item.id, 'view')">
              {{ costFormat(item.amount) }} {{ products.currency }}
            </td>
            <td class="px-4 py-2">
              <div class="flex">
                <div class="pr-3">
                  <SvgIcon name="pencil-square" class="h-5 w-5" @click="openDrawer(index, item.id, 'update')" />
                </div>
                <div>
                  <SvgIcon :name="item.active ? 'eye' : 'eye-slash'" class="h-5 w-5" @click="updateProductActive(index)" />
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="mx-auto" v-else>Add first product</div>

    <drawer :is-open="isDrawer.open" max-width="700px" @close="closeDrawer">
      <!-- Start view section -->
      <div v-if="isDrawer.action === 'view'">
        <div class="pb-8">
          <div class="flex items-center">
            <div class="pr-3">
              <h1>View {{ products.products[product.index].name }}</h1>
            </div>
            <div>
              <SvgIcon :name="products.products[product.index].active ? 'eye' : 'eye-slash'" class="h-5 w-5 cursor-pointer" @click="updateProductActive(product.index)" />
            </div>
          </div>
        </div>

        <div class="flow-root">
          <dl class="-my-3 divide-y divide-gray-100 text-sm">
            <DetailList name="ID">{{ product.info.id }}</DetailList>
            <DetailList name="Name">{{ product.info.name }}</DetailList>
            <DetailList name="Price">{{ product.info.amount }} {{ products.currency }}</DetailList>
            <DetailList name="URL">{{ product.info.url }}</DetailList>
            <DetailList name="Metadata">
              <div v-for="(data, index) in product.info.metadata">{{ data.key }}: {{ data.value }}</div>
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

            <DetailList name="description">{{ product.info.description }}</DetailList>
          </dl>
        </div>

        <div class="pt-8">
          <FormButton type="submit" name="Close" color="green" @click="closeDrawer" />
        </div>
      </div>
      <!-- End view section -->

      <!-- Start update section -->
      <div v-if="isDrawer.action === 'update'">
        <div class="pb-8">
          <div class="flex items-center">
            <div class="pr-3">
              <h1>Edit {{ products.products[product.index].name }}</h1>
            </div>
            <div>
              <SvgIcon :name="products.products[product.index].active ? 'eye' : 'eye-slash'" class="h-5 w-5 cursor-pointer" @click="updateProductActive(product.index)" />
            </div>
          </div>
        </div>

        <Form @submit="updateProduct" v-slot="{ errors }">
          <div class="flow-root" v-if="isDrawer.action === 'update'">
            <dl class="-my-3 text-sm mx-auto mb-0 mt-4 space-y-4">
              <FormInput v-model.trim="product.info.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="Name" ico="at-symbol" />

              <div class="flex flex-row">
                <div class="pr-3">
                  <FormInput v-model.trim="product.info.amount" :error="errors.amount" rules="required|amount" id="amount" type="text" title="Amount" ico="money" />
                </div>
                <div class="mt-3">{{ products.currency }}</div>
              </div>
              <FormInput v-model.trim="product.info.url" :error="errors.url" rules="required|alpha_num" id="url" type="text" title="Url" ico="glob-alt" />

              <hr />
              <p class="font-semibold">Metadata</p>
              <div class="flex" v-for="(data, index) in product.info.metadata" :key="index">
                <div class="grow pr-3">
                  <FormInput v-model="data.key" :id="`mtd-key-${index}`" type="text" title="" />
                </div>
                <div class="grow">
                  <FormInput v-model="data.value" :id="`mtd-value-${index}`" type="text" title="" />
                </div>
                <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteMetadataRecord(index)">
                  <SvgIcon name="trash" class="h-5 w-5" />
                </div>
              </div>
              <div class="flex">
                <div class="grow"></div>
                <div class="mt-2 flex-none">
                  <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addMetadataRecord()">
                    Add metadata record
                  </a>
                </div>
              </div>

              <hr />
              <p class="font-semibold">Attributes</p>
              <div class="flex" v-for="(value, index) in product.info.attributes" :key="index">
                <div class="grow">
                  <FormInput v-model="product.info.attributes[index]" :id="`atr-key-${index}`" type="text" title="" />
                </div>
                <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteAttributeRecord(index)">
                  <SvgIcon name="trash" class="h-5 w-5" />
                </div>
              </div>
              <div class="flex">
                <div class="grow"></div>
                <div class="mt-2 flex-none">
                  <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addAttributeRecord()">
                    Add attribute record
                  </a>
                </div>
              </div>

              <hr />
              <p class="font-semibold">Images</p>
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
              <FormUpload :productId="`${product.info.id}`" accept=".jpg,.jpeg,.png" @added="addProductImage" />

              <hr />
              <FormTextarea v-model="product.info.description" id="textarea" name="Description" />
            </dl>
          </div>

          <div class="pt-8">
            <div class="flex">
              <div class="flex-none">
                <FormButton type="submit" name="Save" color="green" class="mr-3" />
                <FormButton type="submit" name="Close" color="gray" @click="closeDrawer" />
              </div>
              <div class="grow"></div>
              <div class="mt-2 flex-none">
                <span @click="deleteProduct(product.index)" class="cursor-pointer text-red-700">Delete</span>
              </div>
            </div>
          </div>
        </Form>
      </div>
      <!-- End update section -->

      <!-- Start add section -->
      <div v-if="isDrawer.action === 'add'">
        <div class="pb-8">
          <div class="flex items-center">
            <div class="pr-3">
              <h1>Add</h1>
            </div>
          </div>
        </div>

        <Form @submit="addProduct" v-slot="{ errors }">
          <div class="flow-root" v-if="isDrawer.action === 'add'">
            <dl class="-my-3 text-sm mx-auto mb-0 mt-4 space-y-4">
              <FormInput v-model.trim="product.info.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="Name" ico="at-symbol" />
              <div class="flex flex-row">
                <div class="pr-3">
                  <FormInput v-model.trim="product.info.amount" :error="errors.amount" rules="required|amount" id="amount" type="text" title="Amount" ico="money" />
                </div>
                <div class="mt-3">{{ products.currency }}</div>
              </div>
              <FormInput v-model.trim="product.info.url" :error="errors.url" rules="required|alpha_num" id="url" type="text" title="Url" ico="glob-alt" />

              <hr />
              <p class="font-semibold">Metadata</p>
              <div class="flex" v-for="(data, index) in product.info.metadata" :key="index">
                <div class="grow pr-3">
                  <FormInput v-model="data.key" :id="`mtd-key-${index}`" type="text" title="" />
                </div>
                <div class="grow">
                  <FormInput v-model="data.value" :id="`mtd-value-${index}`" type="text" title="" />
                </div>
                <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteMetadataRecord(index)">
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
                  <FormInput v-model="product.info.attributes[index]" :id="`atr-key-${index}`" type="text" title="" />
                </div>
                <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteAttributeRecord(index)">
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
              <FormTextarea v-model="product.info.description" id="textarea" name="Description" />

            </dl>
          </div>

          <div class="pt-8">
            <div class="flex">
              <div class="flex-none">
                <FormButton type="submit" name="Add" color="green" class="mr-3" />
                <FormButton type="submit" name="Close" color="gray" @click="closeDrawer" />
              </div>
              <div class="grow"></div>
            </div>
          </div>
        </Form>
      </div>
      <!-- End update section -->
    </drawer>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { notify } from "notiwind";
import * as NProgress from "nprogress";
import SvgIcon from "svg-icon";
import { defineRule, Form } from "vee-validate";
import { required, numeric, alpha_num, min } from "@vee-validate/rules";
defineRule("required", required);
defineRule("min", min);
defineRule("numeric", numeric);
defineRule("alpha_num", alpha_num);
defineRule('amount', value => {
  if (!value || !value.length) {
    return true;
  }
  if (!/^\d+(\.\d{1,2})?$/.test(value)) {
    return 'amount is not valid';
  }
  return true;
});

import { costFormat, costStripe, formatDate } from "@/utils/";

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

const products = ref([]);

const product = ref({
  info: {},
  action: null,
  index: 0,
  name: null,
});

const listProducts = async () => {

  try {
    const response = await fetch("/api/_/products", {
      credentials: "include",
      method: "GET",
    });
    const data = await response.json();

    if (data.success) {
      products.value = data.result;
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const getProduct = async (id) => {
  try {
    NProgress.start();

    const response = await fetch(`/api/_/products/${id}`, {
      credentials: "include",
      method: "GET",
    });
    const data = await response.json();

    if (data.success) {
      const { info } = product.value;
      Object.assign(info, data.result);
      info.amount = costFormat(info.amount);
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const updateProduct = async () => {
  try {
    const update = { ...product.value.info };
    update.amount = costStripe(update.amount);

    NProgress.start();

    const response = await fetch(`/api/_/products/${update.id}`, {
      credentials: "include",
      method: "PATCH",
      body: JSON.stringify(update),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();

    if (data.success) {
      const found = products.value.products.find((e) => e.id === update.id);
      found.name = update.name;
      found.url = update.url;
      found.amount = update.amount;
      found.description = update.description;
      closeDrawer();
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const addProduct = async () => {
  try {
    const add = { ...product.value.info };
    add.amount = costStripe(add.amount);

    NProgress.start();

    const response = await fetch(`/api/_/products`, {
      credentials: "include",
      method: "POST",
      body: JSON.stringify(add),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();

    if (data.success) {
      products.value.products.push({
        id: data.result.id,
        name: data.result.name,
        description: data.result.description,
        amount: data.result.amount,
        url: data.result.url,
        created: data.result.created,
      });
      products.value.total++;
      closeDrawer();
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const deleteProduct = async (index) => {
  try {
    NProgress.start();

    const response = await fetch(`/api/_/products/${products.value.products[index].id}`, {
      credentials: "include",
      method: "DELETE",
    });
    const data = await response.json();

    if (data.success) {
      products.value.products.splice(index, 1);
      products.value.total--;
    } else {
      const obj = JSON.parse(data.result);
      if (obj.code === "resource_missing") {
        console.log(obj.message);
      }
    }

    closeDrawer();
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const updateProductActive = async (index) => {
  NProgress.start();

  try {
    const response = await fetch(`/api/_/products/${products.value.products[index].id}/active`, {
      credentials: "include",
      method: "PATCH",
    });
    const data = await response.json();

    if (data.success) {
      products.value.products[index].active = !products.value.products[index].active;
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const addMetadataRecord = () => {
  const metadata = product.value.info.metadata || [];
  metadata.push({ key: "", value: "" });
  product.value.info.metadata = metadata;
};

const deleteMetadataRecord = (key) => {
  product.value.info.metadata.splice(key, 1);
};

const addAttributeRecord = () => {
  const attributes = product.value.info.attributes || [];
  attributes.push("");
  product.value.info.attributes = attributes;
};

const deleteAttributeRecord = (index) => {
  product.value.info.attributes.splice(index, 1);
};

const addProductImage = (event) => {
  if (event.success) {
    product.value.info.images.push(event.result);
  }
};

const deleteProductImage = async (index, productId) => {
  try {
    NProgress.start();
    const image = product.value.info.images[index];

    const response = await fetch(`/api/_/products/${productId}/image/${image.id}`, {
      credentials: "include",
      method: "DELETE",
    });
    const data = await response.json();

    if (data.success) {
      product.value.info.images.splice(index, 1);
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
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
      product.value.info.metadata = [];
      product.value.info.attributes = [];
      product.value.info.description = "";
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
