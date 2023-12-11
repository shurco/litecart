<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1>Payment</h1>
    </header>

    <Form @submit="updateSetting" v-slot="{ errors }">
      <FormSelect v-model="payment.currency" :options="['EUR', 'USD', 'JPY', 'GBP', 'AUD', 'CAD', 'CHF', 'CNY', 'SEK']" :error="errors.currency"
        rules="required|one_of:EUR,USD,JPY,GBP,AUD,CAD,CHF,CNY,SEK" class="w-64 mt-5" id="currency" title="Currency" ico="money" />
      <div class="pt-5">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </Form>
    <hr class="mt-5" />

    <div class="mt-5">
      <h2 class="mb-5">Payment providers</h2>
      <div class="flex">
        <div class="cursor-pointer rounded p-2" @click="openDrawer('stripe')" :class="store.payments[`stripe`] ? 'bg-green-200 ' : 'bg-gray-200'">Stripe</div>
        <div class="cursor-pointer rounded p-2 ml-5" @click="openDrawer('paypal')" :class="store.payments[`paypal`] ? 'bg-green-200 ' : 'bg-gray-200'">Paypal</div>
        <div class="cursor-pointer rounded p-2 ml-5" @click="openDrawer('spectrocoin')" :class="store.payments[`spectrocoin`] ? 'bg-green-200 ' : 'bg-gray-200'">Spectrocoin
        </div>
      </div>
    </div>
  </div>

  <drawer :is-open="isDrawer.open" max-width="725px" @close="closeDrawer">
    <Stripe :close="closeDrawer" v-if="isDrawer.action === 'stripe'" />
    <Paypal :close="closeDrawer" v-if="isDrawer.action === 'paypal'" />
    <Spectrocoin :close="closeDrawer" v-if="isDrawer.action === 'spectrocoin'" />
  </drawer>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormSelect, FormButton, Drawer, Stripe, Paypal, Spectrocoin } from "@/components/";
import { showMessage } from "@/utils/message";
import { useSystemStore } from '@/store/system';
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const store = useSystemStore();
const payment = ref({});

onMounted(() => {
  apiGet(`/api/_/settings/payment`).then(res => {
    if (res.success) {
      payment.value = res.result;
    }
  });
  apiGet(`/api/cart/payment`).then(res => {
    if (res.success) {
      store.payments = res.result;
    }
  });
});

const updateSetting = async () => {
  await apiUpdate(`/api/_/settings/payment`, payment.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const isDrawer = ref({
  open: false,
  action: null,
});

const openDrawer = (action) => {
  isDrawer.value.open = true;
  isDrawer.value.action = action;
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};
</script>
