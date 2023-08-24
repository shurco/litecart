<template>
  <MainLayouts>
    <div class="grid grid-cols-1 gap-4 pb-4 lg:grid-cols-[1fr_120px] lg:gap-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">Checkouts</h1>
      </div>
    </div>


    <div class="mx-auto" v-if="checkouts.length > 0">
      <table class="min-w-full divide-y-2 divide-gray-200 bg-white text-sm">
        <thead class="text-left">
          <tr>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Email</th>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Name</th>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Price</th>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Status</th>
            <th class="whitespace-nowrap w-32 px-4 py-2 font-medium text-gray-900">Created</th>
            <th class="whitespace-nowrap w-32 px-4 py-2 font-medium text-gray-900">Updated</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr class="hover:bg-gray-100 active:bg-gray-100 cursor-pointer " :class="{ 'bg-green-50': item.payment_status === 'paid', 'bg-red-50': item.payment_status === 'cancel' }"
            v-for="item, index in checkouts">
            <td class="whitespace-nowrap px-4 py-2">{{ item.email }}</td>
            <td class="whitespace-nowrap px-4 py-2">{{ item.name }}</td>
            <td class="whitespace-nowrap px-4 py-2">
              <a :href="`https://dashboard.stripe.com/payments/${item.payment_id}`" target="_blank">
                {{ costFormat(item.amount_total) }} {{ item.currency }}
              </a>
            </td>
            <td class="whitespace-nowrap px-4 py-2">{{ item.payment_status }}</td>
            <td class="whitespace-nowrap px-4 py-2">{{ formatDate(item.created) }}</td>
            <td class="whitespace-nowrap px-4 py-2" v-if="item.updated">{{ formatDate(item.updated) }}</td>
            <td class="whitespace-nowrap px-4 py-2" v-else></td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="mx-auto" v-else>Not found checkouts</div>

  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import MainLayouts from '@/layouts/Main.vue'
import { costFormat, formatDate } from '@/utils/'

const checkouts = ref([])

onMounted(() => {
  listCheckouts()
})

const listCheckouts = async () => {
  await fetch(`/api/_/checkouts`, {
    credentials: 'include',
    method: 'GET'
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        checkouts.value = data.result
      }
    })
}
</script>