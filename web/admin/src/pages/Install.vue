<template>
  <BlankLayouts>
    <div class="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-lg text-center">
        <h1 class="text-2xl font-bold sm:text-3xl">
          This is the first run! 🎉
        </h1>
        <p class="mt-4 text-gray-500">
          To get started - fill in the following fields
        </p>
      </div>

      <Form @submit="onSubmit" v-slot="{ errors }" class="mx-auto mb-0 mt-8 max-w-md space-y-4">
        <FormInput v-model="state.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email*" ico="at-symbol" />
        <FormInput v-model="state.password" :error="errors.password" id="password" type="password" rules="required|min:6" title="Password*" ico="finger-print" />
        <hr />
        <FormInput v-model="state.domain" :error="errors.domain" id="domain" type="text" rules="required" title="Domain*" ico="glob-alt" />
        <FormInput v-model="state.stripe_secret" :error="errors.stripe_secret" id="stripe_secret" type="text" rules="min:100" title="Stripe secret key" ico="key" />

        <FormButton type="submit" name="Create my cart" color="green" ico="arrow-right" />
      </Form>
    </div>
  </BlankLayouts>
</template>

<script setup>
import { ref } from "vue";

import BlankLayouts from "@/layouts/Blank.vue";
import FormInput from "@/components/form/Input.vue";
import FormButton from "@/components/form/Button.vue";
import { apiPost } from "@/utils/api";

import { Form } from "vee-validate";

const state = ref({
  email: "",
  password: "",
  domain: "",
  stripe_secret: "",
});

const onSubmit = async () => {
  apiPost(`/api/install`, state.value).then(res => {
    if (res.success) {
      window.location.href = "/_/signin/";
    }
  })
};
</script>
