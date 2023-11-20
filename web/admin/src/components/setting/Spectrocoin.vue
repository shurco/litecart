<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Spectrocoin</h1>
        </div>
        <FormToggle v-model="settings.active" class="pt-1" @change="active" />
      </div>
    </div>

    <Form @submit="updateSetting()" v-slot="{ errors }">
      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="settings.merchant_id" :error="errors.merchant_id" rules="required|min:36" id="merchant_id" type="text" title="Merchant ID" ico="key" />
          <FormInput v-model.trim="settings.project_id" :error="errors.project_id" rules="required|min:36" id="project_id" type="text" title="Project ID" ico="key" class="mt-5" />
          <FormTextarea v-model="settings.private_key" :error="errors.private_key" rules="required|min:1700" id="private_key" name="Private key" :rows="15" class="mt-5" />
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
import FormToggle from "@/components/form/Toggle.vue";
import FormInput from "@/components/form/Input.vue";
import FormTextarea from "@/components/form/Textarea.vue";
import FormButton from "@/components/form/Button.vue";
import { Form } from "vee-validate";

import { useSystemStore } from '@/store/system';
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

const props = defineProps({
  close: Function,
});

onMounted(() => {
  loadSettings();
});

const settings = ref({});
const store = useSystemStore();

const loadSettings = async () => {
  apiGet(`/api/_/settings/spectrocoin`).then((res) => {
    if (res.success) {
      settings.value.active = res.result.active;
      settings.value.merchant_id = res.result.merchant_id;
      settings.value.project_id = res.result.project_id;
      settings.value.private_key = res.result.private_key;
    }
  });
  console.log("load setting");
};

const updateSetting = async () => {
  const update = {
    "spectrocoin": {
      "merchant_id": settings.value.merchant_id,
      "project_id": settings.value.project_id,
      "private_key": settings.value.private_key,
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

  apiUpdate(`/api/_/settings/spectrocoin_active`, update).then(res => {
    if (res.success) {
      store.payments['spectrocoin'] = settings.value.active;
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
