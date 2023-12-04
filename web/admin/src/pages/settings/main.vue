<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1><span class="text-gray-300">Settings</span><span class="px-3 text-gray-300">/</span>Main</h1>
    </header>

    <Form @submit="updateSetting()" v-slot="{ errors }">
      <FormInput v-model.trim="main.site_name" :error="errors.site_name" rules="required|min:6" class="max-w-md" id="site_name" type="text" title="Site name" ico="glob-alt" />

      <FormSelect v-model="main.currency" :options="['EUR', 'USD', 'JPY', 'GBP', 'AUD', 'CAD', 'CHF', 'CNY', 'SEK']" :error="errors.currency"
        rules="required|one_of:EUR,USD,JPY,GBP,AUD,CAD,CHF,CNY,SEK" class="w-64 mt-5" id="currency" title="Currency" ico="money" />

      <div class="mt-5 flex">
        <div class="pr-3">
          <FormInput v-model.trim="main.domain" :error="errors.domain" rules="required|min:4" class="w-64" id="domain" type="text" title="Domain" ico="glob-alt" />
        </div>
        <div class="pr-3">
          <FormInput v-model.trim="main.email" :error="errors.email" rules="required|email" class="w-64" id="email" type="text" title="Email" ico="at-symbol" />
        </div>
      </div>
      <div class="mt-5 flex">
        <div class="pr-3">
          <FormInput v-model.trim="main.jwt.secret" :error="errors.jwt_secret" rules="required|min:30" class="w-64" id="jwt_secret" type="text" title="JWT Secret" ico="key" />
        </div>
        <div>
          <FormInput v-model.trim="main.jwt.expire_hours" :error="errors.jwt_expire_hours" rules="required|numeric" class="w-64" id="jwt_expire_hours" type="text"
            title="JWT expire hours" ico="key" />
        </div>
      </div>
      <div class="pt-8">
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

const main = ref({
  jwt: {},
});

onMounted(() => {
  settingsList();
});

const settingsList = async () => {
  apiGet(`/api/_/settings`).then(res => {
    if (res.success) {
      main.value = res.result.main;
    }
  });
};

const updateSetting = async () => {
  var update = {
    main: main.value,
  };

  apiUpdate(`/api/_/settings`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
