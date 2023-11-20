<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Digital {{ digital.type }}</h1>
          <p class="mt-4" v-if="digital.type === 'file'">This is the product that the user purchases. Upload the files that will be sent to the buyer after payment to the email
            address provided during checkout.</p>
          <p class="mt-4" v-if="digital.type === 'data'">Enter the digital product that you intend to sell. It can be a unique item, such as a license key.</p>
        </div>
      </div>
    </div>

    <!-- File section -->
    <div class="flow-root" v-if="digital.type === 'file'">
      <div class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
        <div class="grid content-start" v-if="digital.files !== null">
          <div v-for="(value, index) in digital.files" class="relative mt-4 flex first:mt-0">
            <a :href="`/secrets/${value.name}.${value.ext}`" target="_blank" class="rounded-lg bg-gray-200 px-3 py-3">
              {{ value.name }}.{{ value.ext }}
            </a>
            <SvgIcon name="trash" stroke="currentColor" class="ml-3 mt-3 h-5 w-5 cursor-pointer" @click="deleteDigital('file', index)" />
          </div>
        </div>
        <FormUpload :productId="`${drawer.product.id}`" section="digital" @added="addDigitalFile" />
      </div>
    </div>

    <!-- Data section -->
    <div class="flow-root" v-if="digital.type === 'data'">
      <div class="-my-3 mx-auto mb-0 mt-4 space-y-4 text-sm">
        <div class="flex" v-for="(value, index) in digital.data" :key="index">
          <div class="grow" v-if="digital.data[index].cart_id === ''">
            <FormInput v-model="digital.data[index].content" :id="`${digital.data[index].id}`" type="text" title="" @focusout="saveData(index)" />
          </div>
          <div class="grow" v-else>
            <div class="rounded-lg bg-gray-200 px-3 py-3">
              {{ digital.data[index].content }}
            </div>
          </div>
          <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteDigital('data', index)" v-if="digital.data[index].cart_id === ''">
            <SvgIcon name="trash" stroke="currentColor" class="h-5 w-5" />
          </div>
        </div>
        <div class="flex">
          <div class="grow"></div>
          <div class="mt-2 flex-none">
            <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addDigitalData">
              Add data
            </a>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-4 flow-root" v-if="!digital.type">Select digital type</div>
  </div>
</template>

<script setup>
import { onMounted, ref, computed } from "vue";

import FormInput from "@/components/form/Input.vue";
import FormUpload from "@/components/form/Upload.vue";
import { showMessage } from "@/utils/message";
import { apiGet, apiPost, apiUpdate, apiDelete } from "@/utils/api";

const props = defineProps({
  drawer: {
    required: true,
  },
  products: {
    required: true,
  },
  close: Function,
});

const digital = ref({});

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
  listDigitals();
});

const listDigitals = async () => {
  apiGet(`/api/_/products/${props.drawer.product.id}/digital`).then(res => {
    if (res.success && res.result !== null) {
      digital.value.type = res.result.type;
      digital.value.files = res.result.files ?? [];
      digital.value.data = res.result.data ?? [];
    }
  });
};

const addDigitalFile = (e) => {
  if (e.success) {
    const productToUpdate = products.value.products.find((e) => e.id === props.drawer.product.id)
    if (!productToUpdate.digital.filled) {
      productToUpdate.digital.filled = true;
    }
    digital.value.files.push(e.result);
  }
};


const addDigitalData = async () => {
  apiPost(`/api/_/products/${props.drawer.product.id}/digital`).then(res => {
    if (res.success) {
      const productToUpdate = products.value.products.find((e) => e.id === props.drawer.product.id)
      if (!productToUpdate.digital.filled) {
        productToUpdate.digital.filled = true;
      }
      digital.value.data.push(res.result);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const saveData = async (index) => {
  const update = {
    content: digital.value.data[index].content,
  };
  apiUpdate(`/api/_/products/${props.drawer.product.id}/digital/${digital.value.data[index].id}`, update)
};

const deleteDigital = async (type, index) => {
  const digitalId = type === "file" ? digital.value.files[index].id : digital.value.data[index].id;

  apiDelete(`/api/_/products/${props.drawer.product.id}/digital/${digitalId}`).then(res => {
    if (res.success) {
      const productToUpdate = products.value.products.find((e) => e.id === props.drawer.product.id);

      switch (type) {
        case "file":
          digital.value.files.splice(index, 1);
          if (digital.value.files.length === 0) {
            productToUpdate.digital.filled = false;
          }
          break;
        case "data":
          digital.value.data.splice(index, 1);
          const free_data = digital.value.data.filter((e) => e.cart_id === "").length
          if (free_data === 0) {
            productToUpdate.digital.filled = false;
          }
          break;
      }
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

</script>
