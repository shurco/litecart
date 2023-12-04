<template>
  <div class="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-lg text-center">
      <h1 class="text-2xl font-bold sm:text-3xl">👨‍🎨 Admin sign in</h1>
    </div>
    <Form @submit="onSubmit" v-slot="{ errors }" class="mx-auto mb-0 mt-8 max-w-md space-y-4">
      <FormInput v-model="state.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" ico="at-symbol" />
      <FormInput v-model="state.password" :error="errors.password" id="password" type="password" rules="required|min:6" title="Password" ico="finger-print" />
      <FormButton type="submit" name="Login" color="green" ico="arrow-right" />
    </Form>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { FormInput, FormButton } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiPost } from "@/utils/api";
import { Form } from "vee-validate";

const state = ref({
  email: "",
  password: "",
});

const onSubmit = async () => {
  apiPost(`/api/sign/in`, state.value).then(res => {
    if (res.success) {
      window.location.href = "/_/";
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
