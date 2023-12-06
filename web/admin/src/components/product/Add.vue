<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Add</h1>
        </div>
      </div>
    </div>
    <Form @submit="addProduct" v-slot="{ errors }">
      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="product.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="Name" ico="at-symbol" />
          <div class="flex flex-row">
            <div class="pr-3">
              <FormInput v-model.trim="amount" :error="errors.amount" rules="required|amount" id="amount" type="text" title="Amount" ico="money" />
            </div>
            <div class="mt-3">{{ drawer.currency }}</div>
          </div>

          <div class="flex">
            <div class="grow pr-3">
              <FormInput v-model.trim="product.slug" :error="errors.slug" rules="required|slug" id="slug" type="text" title="Slug" ico="glob-alt" />
            </div>
            <div class="grow">
              <FormSelect v-model="product.digital.type" :options="['file', 'data']" :error="errors.digital_type" rules="required" id="digital_type" title="Digital type"
                ico="cube" />
            </div>
          </div>

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
              <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addMetadataRecord">
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
              <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addAttributeRecord">
                Add attribute record
              </a>
            </div>
          </div>

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
            <FormButton type="submit" name="Add" color="green" class="mr-3" />
            <FormButton type="submit" name="Close" color="gray" @click="close" />
          </div>
          <div class="grow"></div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { FormInput, FormButton, FormSelect, FormTextarea, Editor } from "@/components/";
import { costStripe } from "@/utils/";
import { showMessage } from "@/utils/message";
import { apiPost } from "@/utils/api";
import { Form } from "vee-validate";

const amount = ref()
const product = ref({
  metadata: [],
  attributes: [],
  description: "",
  digital: {
    type: "",
  }
})
const props = defineProps({
  products: {
    required: true,
  },
  drawer: {
    required: true,
  },
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

const addMetadataRecord = async () => {
  const metadata = product.value.metadata || [];
  metadata.push({ key: "", value: "" });
  product.value.metadata = metadata;
};

const deleteMetadataRecord = async (key) => {
  product.value.metadata.splice(key, 1);
};

const addAttributeRecord = async () => {
  const attributes = product.value.attributes || [];
  attributes.push("");
  product.value.attributes = attributes;
};

const deleteAttributeRecord = async (index) => {
  product.value.attributes.splice(index, 1);
};

const addProduct = async () => {
  product.value.amount = costStripe(amount.value);
  apiPost(`/api/_/products`, product.value).then(res => {
    if (res.success) {
      products.value.products.push({
        id: res.result.id,
        name: res.result.name,
        description: res.result.description,
        amount: res.result.amount,
        slug: res.result.slug,
        created: res.result.created,
        digital: {
          type: res.result.digital.type,
        },
      });
      products.value.total++;
      props.close();
      showMessage(res.message);

      props.drawer.open = true;
      props.drawer.action = 'digital';
      props.drawer.product = {
        id: res.result.id,
        digital: res.result.digital.type,
      }
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
