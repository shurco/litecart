<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Edit {{ product.name }}</h1>
        </div>
        <div>
          <SvgIcon :name="product.active ? 'eye' : 'eye-slash'" class="h-5 w-5 cursor-pointer" @click="active" stroke="currentColor" />
        </div>
      </div>
    </div>

    <Form @submit="updateProduct" v-slot="{ errors }">
      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="product.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="Name" ico="at-symbol" />

          <div class="flex flex-row">
            <div class="pr-3">
              <FormInput v-model.trim="amount" :error="errors.amount" rules="required|amount" id="amount" type="text" title="Amount" ico="money" />
            </div>
            <div class="mt-3">{{ drawer.currency }}</div>
          </div>
          <FormInput v-model.trim="product.slug" :error="errors.slug" rules="required|slug" id="slug" type="text" title="Slug" ico="glob-alt" />

          <hr />
          <p class="font-semibold">Metadata</p>
          <div class="flex" v-for="(data, index) in product.metadata" :key="index">
            <div class="grow pr-3">
              <FormInput v-model="data.key" :id="`mtd-key-${index}`" type="text" title="Key" />
            </div>
            <div class="grow">
              <FormInput v-model="data.value" :id="`mtd-value-${index}`" type="text" title="Value" />
            </div>
            <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteMetadataRecord(index)">
              <SvgIcon name="trash" class="h-5 w-5" stroke="currentColor" />
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
          <div class="flex" v-for="(value, index) in product.attributes" :key="index">
            <div class="grow">
              <FormInput v-model="product.attributes[index]" :id="`atr-key-${index}`" type="text" title="" />
            </div>
            <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteAttributeRecord(index)">
              <SvgIcon name="trash" class="h-5 w-5" stroke="currentColor" />
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
          <div class="grid grid-cols-4 content-start gap-4" v-if="product.images !== null">
            <div v-for="(item, index) in product.images" class="relative" style="width: 100%; max-width: 150px">
              <a :href="`/uploads/${item.name}.${item.ext}`" target="_blank">
                <img :src="`/uploads/${item.name}_sm.${item.ext}`" />
              </a>
              <div class="absolute end-4 top-4 cursor-pointer bg-white p-2" @click="deleteProductImage(index)">
                <SvgIcon name="trash" class="h-5 w-5" stroke="currentColor" />
              </div>
            </div>
          </div>
          <FormUpload :productId="`${product.id}`" accept=".jpg,.jpeg,.png" section="image" @added="addProductImage" />

          <hr />
          <p class="font-semibold">Short description</p>
          <FormTextarea v-model="product.brief" id="textarea" name="Brief" />

          <hr />
          <p class="font-semibold">Description</p>
          <Editor v-model:model-value="product.description" placeholder="type description here" />
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
            <span @click="deleteProduct" class="cursor-pointer text-red-700">Delete</span>
          </div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, computed, ref } from "vue";
import { FormInput, FormButton, FormTextarea, FormUpload, Editor } from "@/components/";
import { costFormat, costStripe } from "@/utils/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate, apiDelete } from "@/utils/api";
import { Form } from "vee-validate";

const amount = ref();
const product = ref({});
const props = defineProps({
  drawer: {
    required: true,
  },
  products: {
    required: true,
  },
  updateActive: Function,
  close: Function,
});

const emits = defineEmits(["update:modelValue"]);
const products = computed({
  get: () => {
    return props.products;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

onMounted(() => {
  apiGet(`/api/_/products/${products.value.products[props.drawer.product.index].id}`,).then((res) => {
    if (res.success) {
      product.value = res.result;
      amount.value = costFormat(product.value.amount);
      if (!product.value.images) {
        product.value.images = [];
      }
    } else {
      showMessage(res.result, "connextError");
    }
  });
});

const updateProduct = async () => {
  product.value.amount = costStripe(amount.value);
  apiUpdate(`/api/_/products/${product.value.id}`, product.value).then(
    (res) => {
      if (res.success) {
        products.value.products[props.drawer.product.index] = product.value;
        showMessage(res.message);
      } else {
        showMessage(res.result, "connextError");
      }
    },
  );
};

const deleteProduct = async () => {
  apiDelete(`/api/_/products/${product.value.id}`).then((res) => {
    if (res.success) {
      products.value.products.splice(props.drawer.product.index, 1);
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

const active = async () => {
  props.updateActive(props.drawer.product.index);
  product.value.active = !product.value.active;
};

const addMetadataRecord = () => {
  const metadata = product.value.metadata || [];
  metadata.push({ key: "", value: "" });
  product.value.metadata = metadata;
};

const deleteMetadataRecord = (key) => {
  product.value.metadata.splice(key, 1);
};

const addAttributeRecord = () => {
  const attributes = product.value.attributes || [];
  attributes.push("");
  product.value.attributes = attributes;
};

const deleteAttributeRecord = (index) => {
  product.value.attributes.splice(index, 1);
};

const addProductImage = (e) => {
  if (e.success) {
    if (!product.value.images) {
      product.value.images = [];
    }
    product.value.images.push(e.result);
  }
};

const deleteProductImage = async (index) => {
  apiDelete(`/api/_/products/${product.value.id}/image/${product.value.images[index].id}`).then((res) => {
    if (res.success) {
      product.value.images.splice(index, 1);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
