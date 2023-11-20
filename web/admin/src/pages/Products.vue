<template>
  <MainLayouts>
    <header>
      <h1>Products</h1>
      <div>
        <FormButton type="submit" name="Add" color="green" ico="arrow-right" @click="openDrawer(null, 'add')" />
      </div>
    </header>

    <div class="mx-auto pb-16" v-if="products.total > 0">
      <table>
        <thead>
          <tr>
            <th class="w-28"></th>
            <th>Name</th>
            <th class="w-32">Slug</th>
            <th class="w-32">Price</th>
            <th class="w-12 px-4 py-2">
              <SvgIcon name="cube" class="h-5 w-5" stroke="currentColor" />
            </th>
            <th class="w-24 px-4 py-2"></th>
          </tr>
        </thead>
        <tbody>
          <tr :class="{ 'opacity-30': !item.active }" v-for="(item, index) in products.products">
            <td>
              <a :href="`/uploads/${item.images[0].name}.${item.images[0].ext}`" target="_blank" v-if="item.images">
                <img style="width: 100%; max-width: 80px" :src="`/uploads/${item.images[0].name}_sm.${item.images[0].ext}`" loading="lazy" />
              </a>
              <img style="width: 100%; max-width: 80px" src="/assets/img/noimage.png" v-else />
            </td>
            <td @click="openDrawer(index, 'view')">
              <div>{{ item.name }}</div>
            </td>
            <td>
              <a :href="`/products/${item.slug}`" target="_blank" v-if="item.active">{{ item.slug }}</a>
              <span v-else>{{ item.slug }}</span>
            </td>
            <td @click="openDrawer(index, 'view')">
              {{ costFormat(item.amount) }} {{ products.currency }}
            </td>
            <td class="px-4 py-2">
              <SvgIcon :name="digitalTypeIco(item.digital.type)" class="h-5 w-5" @click="openDrawer(index, 'digital')" stroke="currentColor" />
            </td>
            <td class="px-4 py-2">
              <div class="flex">
                <div class="pr-3">
                  <SvgIcon name="pencil-square" class="h-5 w-5" @click="openDrawer(index, 'update')" stroke="currentColor" />
                </div>
                <div class="pr-3">
                  <SvgIcon name="rocket" class="h-5 w-5" @click="openDrawer(index, 'seo')" stroke="currentColor" />
                </div>
                <div>
                  <SvgIcon :name="item.active ? 'eye' : 'eye-slash'" class="h-5 w-5" @click="updateProductActive(index)" stroke="currentColor" />
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="mx-auto" v-else>Add first product</div>

    <drawer :is-open="isDrawer.open" max-width="710px" @close="closeDrawer">
      <ProjectView    :drawer="isDrawer" :close="closeDrawer" :updateActive="updateProductActive" v-if="isDrawer.action === 'view'" />
      <ProjectAdd     :drawer="isDrawer" :close="closeDrawer" :products="products" v-if="isDrawer.action === 'add'" />
      <ProjectUpdate  :drawer="isDrawer" :close="closeDrawer" :products="products" :updateActive="updateProductActive" v-if="isDrawer.action === 'update'" />
      <ProjectSeo     :drawer="isDrawer" :close="closeDrawer" v-if="isDrawer.action === 'seo'" />
      <ProjectDigital :drawer="isDrawer" :close="closeDrawer" v-if="isDrawer.action === 'digital'" />
    </drawer>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import MainLayouts from "@/layouts/Main.vue";
import FormButton from "@/components/form/Button.vue";
import Drawer from "@/components/Drawer.vue";
import ProjectView from "@/components/product/View.vue";
import ProjectAdd from "@/components/product/Add.vue";
import ProjectUpdate from "@/components/product/Update.vue";
import ProjectSeo from "@/components/product/Seo.vue";
import ProjectDigital from "@/components/product/Digital.vue";
import { costFormat } from "@/utils/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

onMounted(() => {
  listProducts();
});

const isDrawer = ref({
  open: false,
  action: null,
  product: {
    id: null,
    index: null,
    digital: null,
    currency: null,
  }
});

const products = ref([]);

const listProducts = async () => {
  apiGet(`/api/_/products`).then(res => {
    if (res.success) {
      products.value = res.result;
      isDrawer.value.currency = res.result.currency;
    }
  });
};

const updateProductActive = async (index) => {
  apiUpdate(`/api/_/products/${products.value.products[index].id}/active`).then(res => {
    if (res.success) {
      const name = products.value.products[index].name;
      const status = !products.value.products[index].active;
      products.value.products[index].active = status;
      if (status) {
        showMessage(`Product ${name} activated`);
      } else {
        showMessage(`Product ${name} deactivated`);
      }
    }
  })
};

const openDrawer = (index, action) => {
  isDrawer.value.open = true;
  isDrawer.value.action = action;
  if (index!==null) {
    isDrawer.value.product = {
      index: index,
      id: products.value.products[index].id,
    }
  }
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
  isDrawer.value.product = {
    index: null,
    id: null,
  }
};

const digitalTypeIco = (type) => {
  switch (type) {
    case "file":
      return "paper-clip";
    case "data":
      return "queue-list";
    default:
      return "cube-transparent";
  }
};
</script>
