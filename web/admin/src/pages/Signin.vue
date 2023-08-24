<template>
  <BlankLayouts>
    <div class="py-16 sm:px-6 lg:px-8 mx-auto max-w-screen-xl px-4">
      <div class="mx-auto max-w-lg text-center">
        <h1 class="text-2xl font-bold sm:text-3xl">👨‍🎨 Admin sign in</h1>
      </div>
      <Form @submit="onSubmit" v-slot="{ errors }" class="mx-auto mb-0 mt-8 max-w-md space-y-4">
        <FormInput v-model="state.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" ico="at-symbol" />
        <FormInput v-model="state.password" :error="errors.password" id="password" type="password" rules="required|min:6" title="Password" ico="finger-print" />
        <FormButton type="submit" name="Login" color="green" ico="arrow-right" />
      </Form>
    </div>
  </BlankLayouts>
</template>

<script setup>
import { ref } from 'vue'

import { defineRule, Form } from 'vee-validate'
import { required, email, min } from '@vee-validate/rules'
defineRule('required', required)
defineRule('email', email)
defineRule('min', min)

import BlankLayouts from '@/layouts/Blank.vue'
import FormInput from '@/components/form/Input.vue'
import FormButton from '@/components/form/Button.vue'

import 'vue-toast-notification/dist/theme-default.css'
import { useToast } from 'vue-toast-notification'

const state = ref({
  email: '',
  password: ''
})

const toast = useToast()

const onSubmit = async () => {
  await fetch('/api/sign/in', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify(state.value),
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then((response) => response.json())
    .then((data) => {
      if (!data.success) {
        toast.open({
          message: data.result,
          type: 'error'
        })
      } else {
        window.location.href = '/_/'
      }
    })
}
</script>
