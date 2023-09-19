<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Edit {{ products.products[product.index].name }}</h1>
        </div>
        <div>
          <SvgIcon :name="products.products[product.index].active ? 'eye' : 'eye-slash'
            " class="h-5 w-5 cursor-pointer" @click="updateActive(product.index)" />
        </div>
      </div>
    </div>

    <Form @submit="updateProduct" v-slot="{ errors }">
      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="product.info.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="Name" ico="at-symbol" />

          <div class="flex flex-row">
            <div class="pr-3">
              <FormInput v-model.trim="product.info.amount" :error="errors.amount" rules="required|amount" id="amount" type="text" title="Amount" ico="money" />
            </div>
            <div class="mt-3">{{ products.currency }}</div>
          </div>
          <FormInput v-model.trim="product.info.slug" :error="errors.slug" rules="required|slug" id="slug" type="text" title="Slug" ico="glob-alt" />

          <hr />
          <p class="font-semibold">Metadata</p>
          <div class="flex" v-for="(data, index) in product.info.metadata" :key="index">
            <div class="grow pr-3">
              <FormInput v-model="data.key" :id="`mtd-key-${index}`" type="text" title="Key" />
            </div>
            <div class="grow">
              <FormInput v-model="data.value" :id="`mtd-value-${index}`" type="text" title="Value" />
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
          <FormUpload :productId="`${product.info.id}`" accept=".jpg,.jpeg,.png" section="image" @added="addProductImage" />

          <hr />
          <FormTextarea v-model="product.info.description" id="textarea" name="Description" />
        </dl>
      </div>

      <div class="pt-8">
        <div class="flex">
          <div class="flex-none">
            <FormButton type="submit" name="Save" color="green" class="mr-3" />
            <FormButton type="submit" name="Close" color="gray" @click="close" />
          </div>
          <div class="grow"></div>
          <div class="mt-2 flex-none">
            <span @click="deleteProduct(product.index)" class="cursor-pointer text-red-700">Delete</span>
          </div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, computed } from "vue";

import FormInput from "@/components/form/Input.vue";
import FormButton from "@/components/form/Button.vue";
import FormTextarea from "@/components/form/Textarea.vue";
import FormUpload from "@/components/form/Upload.vue";
import { costStripe } from "@/utils/";
import { showMessage } from "@/utils/message";
import { apiUpdate, apiDelete } from "@/utils/api";

import SvgIcon from "svg-icon";
import { Form } from "vee-validate";

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

const emits = defineEmits(["update:modelValue"]);

const product = computed({
  get: () => {
    return props.product;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

onMounted(() => {
  if (!product.value.info.images) {
    product.value.info.images = [];
  }
});

const products = computed({
  get: () => {
    return props.products;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

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

const updateProduct = async () => {
  const update = { ...product.value.info };
  update.amount = costStripe(update.amount);
  apiUpdate(`/api/_/products/${update.id}`, update).then(res => {
    if (res.success) {
      const found = products.value.products.find((e) => e.id === update.id);
      found.name = update.name;
      found.slug = update.slug;
      found.amount = update.amount;
      found.description = update.description;
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const deleteProduct = async (index) => {
  apiDelete(`/api/_/products/${products.value.products[index].id}`).then(res => {
    if (res.success) {
      products.value.products.splice(index, 1);
      products.value.total--;
      showMessage(res.message);
    } else {
      const obj = JSON.parse(res.result);
      if (obj.code === "resource_missing") {
        console.log(obj.message);
      }
    }
    props.close();
  });
};

const addProductImage = (event) => {
  if (event.success) {
    const update = { ...product.value.info };
    const found = products.value.products.find((e) => e.id === update.id);
    if (!found.images) {
      found.images = [];
    }
    found.images.push(event.result);

    product.value.info.images.push(event.result);
  }
};

const deleteProductImage = async (index, productId) => {
  apiDelete(`/api/_/products/${productId}/image/${product.value.info.images[index].id}`).then(res => {
    if (res.success) {
      product.value.info.images.splice(index, 1);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
