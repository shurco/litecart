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
      <div class="-my-3 text-sm mx-auto mb-0 mt-2 space-y-4">
        <div class="grid content-start" v-if="digital.files !== null">
          <div v-for="(value, index) in digital.files" class="relative flex mt-4 first:mt-0">
            <a :href="`/secrets/${value.name}.${value.ext}`" target="_blank" class="bg-gray-200 px-3 py-3 rounded-lg">
              {{ value.name }}.{{ value.ext }}
            </a>
            <SvgIcon name="trash" class="h-5 w-5 mt-3 ml-3 cursor-pointer" @click="deleteDigital('file', index)" />
          </div>
        </div>
        <FormUpload :productId="`${product.info.id}`" section="digital" @added="addDigitalFile" />
      </div>
    </div>

    <!-- Data section -->
    <div class="flow-root" v-if="digital.type === 'data'">
      <div class="-my-3 text-sm mx-auto mb-0 mt-4 space-y-4">
        <div class="flex" v-for="(value, index) in digital.data" :key="index">
          <div class="grow" v-if="digital.data[index].active">
            <FormInput v-model="digital.data[index].content" :id="`${digital.data[index].id}`" type="text" title="" @focusout="saveData(index)" />
          </div>
          <div class="grow" v-else>
            <div class="bg-gray-200 px-3 py-3 rounded-lg">{{ digital.data[index].content }}</div>
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

    <div class="flow-root mt-4" v-if="!digital.type">
      Select digital type
    </div>

  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

import FormInput from "@/components/form/Input.vue";
import FormUpload from "@/components/form/Upload.vue";
import SvgIcon from 'svg-icon'

import { notify } from "notiwind";
import * as NProgress from "nprogress";

const props = defineProps({
  product: {
    required: true
  },
  close: Function,
})

const digital = ref({})

onMounted(() => {
  listDigitals(props.product.info.id);
});

const listDigitals = async (productId) => {
  try {
    NProgress.start();

    const response = await fetch(`/api/_/products/${productId}/digital`, {
      credentials: "include",
      method: "GET",
    });
    const { success, result } = await response.json();

    if (success) {
      if (result !== null) {
        digital.value.type = result.type;
        digital.value.files = result.files ? result.files : [];
        digital.value.data = result.data ? result.data : [];
      }
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const addDigitalFile = (event) => {
  if (event.success) {
    digital.value.files.push(event.result);
  }
};

const deleteDigital = async (type, index) => {
  try {
    const { id } = props.product.info;
    NProgress.start();

    var digitalID
    switch (type) {
      case "file":
        digitalID = digital.value.files[index].id;
        break;
      case "data":
        digitalID = digital.value.data[index].id;
        break;
    }

    const response = await fetch(`/api/_/products/${id}/digital/${digitalID}`, {
      credentials: "include",
      method: "DELETE",
    });
    const { success, result } = await response.json();

    if (success) {
      switch (type) {
        case "file":
          digital.value.files.splice(index, 1);
          break;
        case "data":
          digital.value.data.splice(index, 1);
          break;
      }
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const addDigitalData = async () => {
  try {
    const { id } = props.product.info;
    NProgress.start();

    const response = await fetch(`/api/_/products/${id}/digital`, {
      credentials: "include",
      method: "POST",
    });
    const { success, result } = await response.json();

    if (success) {
      digital.value.data.push(result)
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: result,
        type: "error",
      }, 4000)
    }

  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const saveData = async (index) => {
  try {
    const { id } = props.product.info;
    const update = {
      content: digital.value.data[index].content,
    }
    NProgress.start();

    const response = await fetch(`/api/_/products/${id}/digital/${digital.value.data[index].id}`, {
      credentials: "include",
      method: "PATCH",
      body: JSON.stringify(update),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const { success, result } = await response.json();

  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
}
</script>
