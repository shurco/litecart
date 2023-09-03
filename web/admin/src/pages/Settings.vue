<template>
  <MainLayouts>
    <div class="pb-10">
      <header>
        <h1>Settings</h1>
      </header>

      <div>
        <h2 class="mb-5">Main</h2>
        <Form @submit="updateSetting('main')" v-slot="{ errors }">
          <div class="flex">
            <div class="pr-3">
              <FormInput v-model.trim="main.domain" :error="errors.domain" rules="required|min:4" class="w-64" id="domain" type="text" title="Domain" ico="glob-alt" />
            </div>
            <div class="pr-3">
              <FormInput v-model.trim="main.email" :error="errors.email" rules="required|email" class="w-64" id="email" type="text" title="Email" ico="at-symbol" />
            </div>
            <div>
              <FormInput v-model.trim="main.currency" :error="errors.currency" rules="required|one_of:EUR,USD" class="w-64" id="currency" type="text" title="Currency" ico="money " />
            </div>
          </div>
          <div class="flex mt-5">
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
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Password</h2>
        <Form @submit="updateSetting('password')" v-slot="{ errors }">
          <FormInput v-model.trim="password.old" :error="errors.password_old" rules="required|min:6" class="w-96" id="password_old" type="password" title="Actual password"
            ico="finger-print" />
          <div class="flex mt-5">
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
        <hr class="mt-5" />
      </div>


      <div class="mt-5">
        <h2 class="mb-5">Stripe</h2>
        <Form @submit="updateSetting('stripe')" v-slot="{ errors }">
          <div class="flex">
            <div class="pr-3">
              <FormInput v-model.trim="stripe.secret_key" :error="errors.secret_key" rules="required|min:100" class="w-96" id="secret_key" type="text" title="Secret key" ico="key" />
            </div>
            <div>
              <FormInput v-model.trim="stripe.webhook_secret_key" :error="errors.webhook_secret_key" rules="min:100" class="w-96" id="webhook_secret_key" type="text"
                title="Webhook secret key" ico="key" />
            </div>
          </div>
          <div class="pt-8">
            <FormButton type="submit" name="Save" color="green" />
          </div>
        </Form>
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Socials</h2>
        <Form @submit="updateSetting('social')" v-slot="{ errors }">
          <div class="flex">
            <div class="pr-3 pt-2.5">
              https://facebook.com/
            </div>
            <div>
              <FormInput v-model.trim="social.facebook" :error="errors.social_facebook" rules="alpha_num" class="w-48" id="social_facebook" type="text" title="Facebook"
                ico="facebook" />
            </div>
          </div>
          <div class="flex mt-5">
            <div class="pr-3 pt-2.5">
              https://instagrammm.com/
            </div>
            <div>
              <FormInput v-model.trim="social.instagram" :error="errors.social_instagram" rules="alpha_num" class="w-48" id="social_instagram" type="text" title="Instagram"
                ico="instagram" />
            </div>
          </div>
          <div class="flex mt-5">
            <div class="pr-3 pt-2.5">
              https://twitter.com/@
            </div>
            <div>
              <FormInput v-model.trim="social.twitter" :error="errors.social_twitter" rules="alpha_num" class="w-48" id="social_twitter" type="text" title="Twitter" ico="twitter" />
            </div>
          </div>
          <div class="flex mt-5">
            <div class="pr-3 pt-2.5">
              https://dribble.com/
            </div>
            <div>
              <FormInput v-model.trim="social.dribble" :error="errors.social_dribble" rules="alpha_num" class="w-48" id="social_dribble" type="text" title="Dribble" ico="dribble" />
            </div>
          </div>
          <div class="flex mt-5">
            <div class="pr-3 pt-2.5">
              https://github.com/
            </div>
            <div>
              <FormInput v-model.trim="social.github" :error="errors.social_github" rules="alpha_num" class="w-48" id="social_github" type="text" title="Github" ico="github" />
            </div>
          </div>
          <div class="pt-8">
            <FormButton type="submit" name="Save" color="green" />
          </div>
        </Form>
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Mail</h2>
        <Form @submit="updateSetting('mail')" v-slot="{ errors }">
          <div class="flex">
            <div class="pr-3">
              <FormInput v-model.trim="mail.smtp_host" :error="errors.smtp_host" rules="required|min:4" class="w-64" id="smtp_host" type="text" title="SMTP host" ico="server" />
            </div>
            <div class="pr-3">
              <FormInput v-model.trim="mail.smtp_port" :error="errors.smtp_port" rules="required|numeric" class="w-64" id="smtp_port" type="text" title="SMTP port"
                ico="arrow-left-on-rectangle" />
            </div>
            <div>
              <FormInput v-model.trim="mail.smtp_encryption" class="w-64" id="smtp_encryption" type="text" title="SMTP encryption" ico="lock-closed" />
            </div>
          </div>
          <div class="flex mt-5">
            <div class="pr-3">
              <FormInput v-model.trim="mail.smtp_username" :error="errors.smtp_username" rules="required" class="w-64" id="smtp_username" type="text" title="Username" ico="user" />
            </div>
            <div>
              <FormInput v-model.trim="mail.smtp_password" :error="errors.smtp_password" rules="required|min:6" class="w-64" id="smtp_password" type="password" title="Password"
                ico="finger-print" />
            </div>
          </div>
          <div class="pt-8">
            <FormButton type="submit" name="Save" color="green" />
          </div>
        </Form>
      </div>

    </div>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import * as NProgress from "nprogress";
import MainLayouts from "@/layouts/Main.vue";
import { notify } from "notiwind";

import { defineRule, Form } from "vee-validate";
import { required, email, one_of, alpha_num, numeric, min } from "@vee-validate/rules";
defineRule("required", required);
defineRule("email", email);
defineRule("one_of", one_of);
defineRule("alpha_num", alpha_num);
defineRule("numeric", numeric);
defineRule("min", min);
defineRule('confirmed', (value, [target], ctx) => {
  if (value === ctx.form[target]) {
    return true;
  }
  return 'Passwords must match';
});

import FormInput from "@/components/form/Input.vue";
import FormButton from "@/components/form/Button.vue";

const main = ref({
  jwt: {},
})
const password = ref({})
const stripe = ref({})
const social = ref({})
const mail = ref({})

const route = useRoute();

onMounted(async () => {
  settingsList();
});

const settingsList = async () => {
  try {
    NProgress.start();

    const response = await fetch(`/api/_/settings`, {
      credentials: "include",
      method: "GET",
    });
    const data = await response.json();

    if (data.success) {
      main.value = data.result.main;
      stripe.value = data.result.stripe;
      social.value = data.result.social;
      mail.value = data.result.mail;
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const updateSetting = async (section) => {
  try {
    NProgress.start();
    var update = {};
    switch (section) {
      case "main":
        update.main = main.value;
        break;
      case "password":
        update.password = {
          old: password.value.old,
          new: password.value.new1
        }
        break;
      case "stripe":
        update.stripe = stripe.value;
        break;
      case "social":
        update.social = social.value;
        break;
      case "mail":
        update.mail = mail.value;
        update.mail.smtp_port = Number(mail.value.smtp_port);
        break;
      default:
        return
    }

    const response = await fetch(`/api/_/settings`, {
      credentials: "include",
      method: "PATCH",
      body: JSON.stringify(update),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();

    if (data.success) {
      notify({
        group: "bottom",
        title: "Perfect",
        text: data.message,
        type: "success",
      }, 4000)
    }else{
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
        type: "error",
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }

}
</script>
