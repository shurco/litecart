<template>
  <MainLayouts>
    <header>
      <h1>Carts</h1>
    </header>

    <div class="mx-auto" v-if="carts.length > 0">
      <table>
        <thead>
          <tr>
            <th>Email</th>
            <th>Name</th>
            <th>Price</th>
            <th>Status</th>
            <th class="w-32">Created</th>
            <th class="w-32">Updated</th>
            <th class="w-12"></th>
          </tr>
        </thead>
        <tbody>
          <tr :class="{ 'bg-green-50': item.payment_status === 'paid' }" v-for="(item, index) in carts">
            <td>{{ item.email }}</td>
            <td>{{ item.name }}</td>
            <td>
              <a :href="`https://dashboard.stripe.com/payments/${item.payment_id}`" target="_blank">
                {{ costFormat(item.amount_total) }} {{ item.currency }}
              </a>
            </td>
            <td>{{ item.payment_status }}</td>
            <td>{{ formatDate(item.created) }}</td>
            <td v-if="item.updated">{{ formatDate(item.updated) }}</td>
            <td v-else></td>
            <td>
              <SvgIcon name="envelope" class="h-5 w-5 opacity-30" v-if="item.payment_status === 'cancel'" />
              <SvgIcon name="envelope" class="h-5 w-5" @click="sendEmail(item.id)" v-else />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="mx-auto" v-else>Not found carts</div>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import MainLayouts from "@/layouts/Main.vue";
import { costFormat, formatDate } from "@/utils/";
import { showMessage } from "@/utils/message";
import { apiGet, apiPost } from "@/utils/api";

import SvgIcon from "svg-icon";

const carts = ref([]);

onMounted(() => {
  listCarts();
});

const listCarts = async () => {
  apiGet(`/api/_/carts`).then(res => {
    if (res.success) {
      carts.value = res.result;
    }
  })
};

const sendEmail = async (id) => {
  apiPost(`/api/_/carts/${id}/mail`).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
