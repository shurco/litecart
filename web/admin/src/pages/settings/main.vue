<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1><span class="text-gray-300">Settings</span><span class="px-3 text-gray-300">/</span>Main</h1>
    </header>

    <Form @submit="updateSetting" v-slot="{ errors }">
      <FormInput v-model.trim="main.site_name" :error="errors.site_name" rules="required|min:6" class="max-w-md" id="site_name" type="text" title="Site name" ico="glob-alt" />
      <FormInput v-model.trim="main.domain" :error="errors.domain" rules="required|min:4" class="max-w-md mt-5" id="domain" type="text" title="Domain" ico="glob-alt" />
      <div class="pt-5">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormInput, FormButton, FormSelect } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const main = ref({});

onMounted(() => {
  apiGet(`/api/_/settings/main`).then(res => {
    if (res.success) {
      main.value = res.result;
    }
  });
});

const updateSetting = async () => {
  await apiUpdate(`/api/_/settings/main`, main.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
