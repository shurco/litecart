<template>
  <MainLayouts>
    <header>
      <h1>Checkouts</h1>
    </header>

    <div class="mx-auto" v-if="checkouts.length > 0">
      <table>
        <thead>
          <tr>
            <th>Email</th>
            <th>Name</th>
            <th>Price</th>
            <th>Status</th>
            <th class="w-32">Created</th>
            <th class="w-32">Updated</th>
          </tr>
        </thead>
        <tbody>
          <tr :class="{
            'bg-green-50': item.payment_status === 'paid',
            'bg-red-50': item.payment_status === 'cancel',
          }" v-for="(item, index) in checkouts">
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
          </tr>
        </tbody>
      </table>
    </div>
    <div class="mx-auto" v-else>Not found checkouts</div>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import MainLayouts from "@/layouts/Main.vue";
import { costFormat, formatDate } from "@/utils/";
import { apiGet } from "@/utils/api";

const checkouts = ref([]);

onMounted(() => {
  listCheckouts();
});

const listCheckouts = async () => {
  apiGet(`/api/_/checkouts`).then(res => {
    if (res.success) {
      checkouts.value = res.result;
    }
  })
};
</script>
