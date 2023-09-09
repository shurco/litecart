<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Digital {{ digital.type }}</h1>
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
            <SvgIcon name="trash" class="ml-3 mt-3 h-5 w-5 cursor-pointer" @click="deleteDigital('file', index)" />
          </div>
        </div>
        <FormUpload :productId="`${product.info.id}`" section="digital" @added="addDigitalFile" />
      </div>
    </div>

    <!-- Data section -->
    <div class="flow-root" v-if="digital.type === 'data'">
      <div class="-my-3 mx-auto mb-0 mt-4 space-y-4 text-sm">
        <div class="flex" v-for="(value, index) in digital.data" :key="index">
          <div class="grow" v-if="digital.data[index].active">
            <FormInput v-model="digital.data[index].content" :id="`${digital.data[index].id}`" type="text" title="" @focusout="saveData(index)" />
          </div>
          <div class="grow" v-else>
            <div class="rounded-lg bg-gray-200 px-3 py-3">
              {{ digital.data[index].content }}
            </div>
          </div>
          <div class="flex-none cursor-pointer pl-3 pt-3" @click="deleteDigital('data', index)" v-if="digital.data[index].active">
            <SvgIcon name="trash" class="h-5 w-5" />
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
import { onMounted, ref } from "vue";

import FormInput from "@/components/form/Input.vue";
import FormUpload from "@/components/form/Upload.vue";
import SvgIcon from "svg-icon";
import { showMessage } from "@/utils/message";
import { apiGet, apiPost, apiUpdate, apiDelete } from "@/utils/api";

const props = defineProps({
  product: {
    required: true,
  },
  close: Function,
});

const digital = ref({});

onMounted(() => {
  listDigitals(props.product.info.id);
});

const listDigitals = async (productId) => {
  apiGet(`/api/_/products/${productId}/digital`).then(res => {
    if (res.success && res.result !== null) {
      digital.value.type = res.result.type;
      digital.value.files = res.result.files ?? [];
      digital.value.data = res.result.data ?? [];
    }
  });
};

const addDigitalFile = (event) => {
  if (event.success) {
    digital.value.files.push(event.result);
  }
};

const deleteDigital = async (type, index) => {
  var digitalID;
  switch (type) {
    case "file":
      digitalID = digital.value.files[index].id;
      break;
    case "data":
      digitalID = digital.value.data[index].id;
      break;
  }

  apiDelete(`/api/_/products/${props.product.info.id}/digital/${digitalID}`).then(res => {
    if (res.success) {
      switch (type) {
        case "file":
          digital.value.files.splice(index, 1);
          break;
        case "data":
          digital.value.data.splice(index, 1);
          break;
      }
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const addDigitalData = async () => {
  apiPost(`/api/_/products/${props.product.info.id}/digital`).then(res => {
    if (res.success) {

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
  apiUpdate(`/api/_/products/${props.product.info.id}/digital/${digital.value.data[index].id}`, update)
};
</script>
