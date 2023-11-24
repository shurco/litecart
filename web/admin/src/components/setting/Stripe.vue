<template>
  <div>
    <Form @submit="updateSetting()" v-slot="{ errors }">
      <div class="pb-8">
        <div class="flex items-center">
          <div class="pr-3">
            <h1>Stripe</h1>
          </div>
          <FormToggle v-model="settings.active" :disabled="Object.keys(errors).length > 0" class="pt-1" @change="active" />
        </div>
      </div>

      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="settings.secret_key" :error="errors.secret_key" rules="required|min:100" id="secret_key" type="text" title="Secret key" ico="key" />
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
  apiGet(`/api/_/settings/stripe`).then(res => {
    if (res.success) {
      settings.value.active = res.result.active;
      settings.value.secret_key = res.result.secret_key;
    }
  });
  console.log("load setting")
};

const updateSetting = async () => {
  const update = {
    "stripe": {
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

  apiUpdate(`/api/_/settings/stripe_active`, update).then(res => {
    if (res.success) {
      store.payments['stripe'] = settings.value.active;
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
