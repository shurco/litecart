<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1>Webhook events</h1>
    </header>

    <Form @submit="updateSetting" v-slot="{ errors }">
      <FormInput v-model.trim="webhook.url" :error="errors.webhook_url" rules="url" class="max-w-md" id="webhook_url" type="text" title="webhook url" ico="key" />
      <div class="pt-5">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormInput, FormButton } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const webhook = ref({});

onMounted(() => {
  apiGet(`/api/_/settings/webhook`).then(res => {
    if (res.success) {
      webhook.value = res.result;
    }
  });
});

const updateSetting = async () => {
  await apiUpdate(`/api/_/settings/webhook`, webhook.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
