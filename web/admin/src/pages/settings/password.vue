<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1><span class="text-gray-300">Settings</span><span class="px-3 text-gray-300">/</span>Password</h1>
    </header>

    <Form @submit="updateSetting()" v-slot="{ errors }">
      <FormInput v-model.trim="password.old" :error="errors.password_old" rules="required|min:6" class="w-96" id="password_old" type="password" title="Actual password"
        ico="finger-print" />
      <div class="mt-5 flex">
        <div class="pr-3">
          <FormInput v-model.trim="password.new1" :error="errors.password_new1" rules="required|min:6" class="w-96" id="password_new1" type="password" title="New password"
            ico="finger-print" />
        </div>
        <div>
          <FormInput v-model.trim="password.new2" :error="errors.password_new2" rules="required|confirmed:password_new1" class="w-96" id="password_new2" type="password"
            title="Repeat password" ico="finger-print" />
        </div>
      </div>
      <div class="pt-8">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </Form>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { FormInput, FormButton } from "@/components";
import { showMessage } from "@/utils/message";
import { apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const password = ref({});

const updateSetting = async () => {
  var update = {
    password: {
      old: password.value.old,
      new: password.value.new1,
    },
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
