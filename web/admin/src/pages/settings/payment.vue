<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1><span class="text-gray-300">Settings</span><span class="px-3 text-gray-300">/</span>Payment</h1>
    </header>

    <div class="flex">
      <div class="cursor-pointer rounded p-2" @click="openDrawer('stripe')" :class="store.payments[`stripe`] ? 'bg-green-200 ' : 'bg-gray-200'">Stripe</div>
      <div class="cursor-pointer rounded p-2 ml-5" @click="openDrawer('paypal')" :class="store.payments[`paypal`] ? 'bg-green-200 ' : 'bg-gray-200'">Paypal</div>
      <div class="cursor-pointer rounded p-2 ml-5" @click="openDrawer('spectrocoin')" :class="store.payments[`spectrocoin`] ? 'bg-green-200 ' : 'bg-gray-200'">Spectrocoin
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
import { Drawer, Stripe, Paypal, Spectrocoin } from "@/components/";
import { useSystemStore } from '@/store/system';
import { apiGet } from "@/utils/api";

const store = useSystemStore();

const isDrawer = ref({
  open: false,
  action: null,
});

onMounted(() => {
  settingsList();
});

const settingsList = async () => {
  apiGet(`/api/cart/payment`).then(res => {
    if (res.success) {
      store.payments = res.result;
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
