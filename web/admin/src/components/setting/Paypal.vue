<template>
  <div>
    <Form @submit="updateSetting()" v-slot="{ errors }">
      <div class="pb-8">
        <div class="flex items-center">
          <div class="pr-3">
            <h1>Paypal</h1>
          </div>
          <FormToggle v-model="settings.active" :disabled="Object.keys(errors).length > 0" class="pt-1" @change="active" />
        </div>
      </div>

      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="settings.client_id" :error="errors.client_id" rules="required|min:80" id="client_id" type="text" title="Client ID" ico="key" />
        </dl>

        <dl class="-my-3 mx-auto mb-0 mt-5 space-y-4 text-sm">
          <FormInput v-model.trim="settings.secret_key" :error="errors.secret_key" rules="required|min:80" id="secret_key" type="text" title="Secret key" ico="key" />
        </dl>
      </div>

      <div class="pt-8">
        <div class="flex">
          <div class="flex-none">
            <FormButton type="submit" name="Save" color="green" />
          </div>

          <div class="grow"></div>
          <div class="flex-none">
            <FormButton type="submit" name="Close" color="gray" @click="close" />
          </div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormInput, FormButton, FormToggle } from "@/components/";
import { useSystemStore } from '@/store/system';
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const settings = ref({});
const store = useSystemStore();
const props = defineProps({
  close: Function,
});

onMounted(() => {
  apiGet(`/api/_/settings/paypal`).then(res => {
    if (res.success) {
      settings.value.active = res.result.active;
      settings.value.client_id = res.result.client_id;
      settings.value.secret_key = res.result.secret_key;
    }
  });
});

const updateSetting = async () => {
  const update = {
    "paypal": {
      "client_id": settings.value.client_id,
      "secret_key": settings.value.secret_key,
      "active": settings.value.active,
    }
  };

  apiUpdate(`/api/_/settings`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const active = () => {
  const update = {
    value: settings.value.active,
  };

  apiUpdate(`/api/_/settings/paypal_active`, update).then(res => {
    if (res.success) {
      store.payments['paypal'] = settings.value.active;
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
