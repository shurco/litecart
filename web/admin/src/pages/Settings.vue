<template>
  <MainLayouts>
    <div class="pb-10">
      <header>
        <h1>Settings</h1>
      </header>

      <div>
        <h2 class="mb-5">Main</h2>
        <Form @submit="updateSetting('main')" v-slot="{ errors }">
          <FormInput v-model.trim="main.site_name" :error="errors.site_name" rules="required|min:6" class="max-w-md" id="site_name" type="text" title="Site name" ico="glob-alt" />
          <div class="mt-5 flex">
            <div class="pr-3">
              <FormInput v-model.trim="main.domain" :error="errors.domain" rules="required|min:4" class="w-64" id="domain" type="text" title="Domain" ico="glob-alt" />
            </div>
            <div class="pr-3">
              <FormInput v-model.trim="main.email" :error="errors.email" rules="required|email" class="w-64" id="email" type="text" title="Email" ico="at-symbol" />
            </div>
            <div>
              <FormSelect v-model="main.currency" :options="['EUR', 'USD', 'JPY', 'GBP', 'AUD', 'CAD', 'CHF', 'CNY', 'SEK']" :error="errors.currency"
                rules="required|one_of:EUR,USD,JPY,GBP,AUD,CAD,CHF,CNY,SEK" id="currency" title="Currency" ico="money" />
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
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Password</h2>
        <Form @submit="updateSetting('password')" v-slot="{ errors }">
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
        <hr class="mt-5" />
      </div>


      <div class="mt-5">
        <h2 class="mb-5">Payment</h2>
        <div class="flex mb-7">
          <div class="cursor-pointer rounded-lg px-3 py-3" @click="openDrawer('stripe')" :class="store.payments[`stripe`] ? 'bg-green-200 ' : 'bg-gray-200'">Stripe</div>
          <div class="cursor-pointer rounded-lg px-3 py-3 ml-5" @click="openDrawer('spectrocoin')" :class="store.payments[`spectrocoin`] ? 'bg-green-200 ' : 'bg-gray-200'">Spectrocoin
          </div>
        </div>
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Webhook events</h2>
        <Form @submit="updateSetting('webhook')" v-slot="{ errors }">
          <FormInput v-model.trim="webhook.url" :error="errors.webhook_url" rules="url" class="max-w-md" id="webhook_url" type="text" title="webhook url" ico="key" />
          <div class="pt-8">
            <FormButton type="submit" name="Save" color="green" />
          </div>
        </Form>
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Socials</h2>
        <Form @submit="updateSetting('social')" v-slot="{ errors }">
          <div v-for="(value, key, index) in socialUrl" :key="index" class="mt-5 flex">
            <div class="pr-3 pt-2.5">{{ socialUrl[key] }}</div>
            <div>
              <FormInput v-model.trim="social[key]" :error="errors[`social_${key}`]" rules="alpha_num" class="w-48" :id="`social_${key}`" type="text"
                :title="key.charAt(0).toUpperCase() + key.slice(1)" :ico="key" />
            </div>
          </div>
          <div class="pt-8">
            <FormButton type="submit" name="Save" color="green" />
          </div>
        </Form>
        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">Mail letters</h2>

        <div class="flex">
          <div class="cursor-pointer rounded-lg bg-gray-200 px-3 py-3" @click="openDrawer('mail_letter_payment')">
            Letter of payment
          </div>

          <div class="cursor-pointer rounded-lg bg-gray-200 px-3 py-3 ml-5" @click="openDrawer('mail_letter_purchase')">
            Letter of purchase
          </div>
        </div>

        <hr class="mt-5" />
      </div>

      <div class="mt-5">
        <h2 class="mb-5">SMTP settings</h2>

        <div class="mb-5 flex items-center justify-between bg-red-600 px-2 py-3 text-white" v-if="!smtp.host ||
          !smtp.port ||
          !smtp.username ||
          !smtp.password
          ">
          <p class="text-sm font-medium">This section is required!</p>
        </div>

        <Form @submit="updateSetting('smtp')" v-slot="{ errors }">
          <div class="flex">
            <div class="pr-3">
              <FormInput v-model.trim="smtp.host" :error="errors.smtp_host" rules="required|min:4" class="w-64" id="smtp_host" type="text" title="SMTP host" ico="server" />
            </div>
            <div class="pr-3">
              <FormInput v-model.trim="smtp.port" :error="errors.smtp_port" rules="required|numeric" class="w-64" id="smtp_port" type="text" title="SMTP port"
                ico="arrow-left-on-rectangle" />
            </div>
            <div>
              <FormSelect v-model="smtp.encryption" :options="['None', 'SSL/TLS', 'STARTTLS']" :error="errors.smtp_encryption" rules="required" id="smtp_encryption"
                title="Encryption" ico="lock-closed" />
            </div>
          </div>
          <div class="mt-5 flex">
            <div class="pr-3">
              <FormInput v-model.trim="smtp.username" :error="errors.smtp_username" rules="required" class="w-64" id="smtp_username" type="text" title="Username" ico="user" />
            </div>
            <div>
              <FormInput v-model.trim="smtp.password" :error="errors.smtp_password" rules="required|min:6" class="w-64" id="smtp_password" type="password" title="Password"
                ico="finger-print" />
            </div>
          </div>
          <div class="flex pt-8">
            <FormButton type="submit" name="Save" color="green" class="flex-none" />
            <div class="ml-5 mt-3 flex-none">
              <span @click="sendTestLetter('smtp')" class="cursor-pointer text-red-700">Test smtp</span>
            </div>
          </div>
        </Form>
      </div>
    </div>

    <drawer :is-open="isDrawer.open" max-width="725px" @close="closeDrawer">
      <Stripe :close="closeDrawer" v-if="isDrawer.action === 'stripe'" />
      <Spectrocoin :close="closeDrawer" v-if="isDrawer.action === 'spectrocoin'" />

      <Letter :close="closeDrawer" :send="sendTestLetter" :legend="letterLegend['mail_letter_payment']" name="mail_letter_payment"
        v-if="isDrawer.action === 'mail_letter_payment'" />
      <Letter :close="closeDrawer" :send="sendTestLetter" :legend="letterLegend['mail_letter_purchase']" name="mail_letter_purchase"
        v-if="isDrawer.action === 'mail_letter_purchase'" />
    </drawer>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";

import MainLayouts from "@/layouts/Main.vue";
import Stripe from "@/components/setting/Stripe.vue";
import Spectrocoin from "@/components/setting/Spectrocoin.vue";
import Letter from "@/components/setting/Letter.vue";
import FormInput from "@/components/form/Input.vue";
import FormButton from "@/components/form/Button.vue";
import FormSelect from "@/components/form/Select.vue";
import Drawer from "@/components/Drawer.vue";

import { useSystemStore } from '@/store/system';
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

import { Form } from "vee-validate";

const main = ref({
  jwt: {},
});
const password = ref({});
const webhook = ref({});
const social = ref({});
const smtp = ref({});

const store = useSystemStore();

const socialUrl = {
  facebook: "https://facebook.com/",
  instagram: "https://instagram.com/",
  twitter: "https://twitter.com/@",
  dribbble: "https://dribbble.com/",
  github: "https://github.com/",
};

const isDrawer = ref({
  open: false,
  action: null,
});

const letterLegend = {
  "mail_letter_payment": {
    "Site_Name": "Site name",
    "Amount_Payment": "Amount of payment",
    "Payment_URL": "Payment link",
  },
  "mail_letter_purchase": {
    "Purchases": "Purchases",
    "Admin_Email": "Admin email",
  }
}

onMounted(() => {
  settingsList();
});

const settingsList = async () => {
  apiGet(`/api/_/settings`).then(res => {
    if (res.success) {
      main.value = res.result.main;
      social.value = res.result.social;
      smtp.value = res.result.smtp;
      webhook.value = res.result.webhook;
    }
  });

  apiGet(`/api/cart/payment`).then(res => {
    if (res.success) {
      
      store.payments = res.result;
    }
  });
};

const updateSetting = async (section) => {
  var update = {};
  switch (section) {
    case "main":
      update.main = main.value;
      break;
    case "password":
      update.password = {
        old: password.value.old,
        new: password.value.new1,
      };
      break;
    case "webhook":
      update.webhook = webhook.value;
      break;
    case "social":
      update.social = social.value;
      break;
    case "smtp":
      update.smtp = smtp.value;
      update.smtp.port = Number(smtp.value.port);
      break;
    default:
      return;
  }

  apiUpdate(`/api/_/settings`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const sendTestLetter = async (letterName) => {
  apiGet(`/api/_/test/letter/${letterName}`).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const openDrawer = (action) => {
  isDrawer.value.open = true;
  isDrawer.value.action = action;
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};
</script>
