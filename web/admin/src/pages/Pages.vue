<template>
  <MainLayouts>
    <div v-if="!route.params.page_url">
      <header>
        <h1>Pages</h1>
        <div>
          <FormButton type="submit" name="New" color="green" ico="arrow-right" @click="openDrawer(null, 'add')" />
        </div>
      </header>

      <div class="mx-auto" v-if="pages.length > 0">
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Group</th>
              <th>Url</th>
              <th class="w-32">Created</th>
              <th class="w-32">Updated</th>
              <th class="w-24 px-4 py-2"></th>
            </tr>
          </thead>
          <tbody>
            <tr :class="{ 'opacity-30': !item.active }" v-for="(item, index) in pages" :key="item.id">
              <td @click="openPage(item.url)">{{ item.name }}</td>
              <td @click="openPage(item.url)">{{ item.type }}</td>
              <td @click="openPage(item.url)">{{ item.url }}</td>
              <td @click="openPage(item.url)">{{ formatDate(item.created) }}</td>
              <td v-if="item.updated">{{ formatDate(item.updated) }}</td>
              <td v-else></td>
              <td class="px-4 py-2">
                <div class="flex">
                  <div class="pr-3">
                    <SvgIcon name="pencil-square" class="h-5 w-5" @click="openDrawer(index, 'update')" />
                  </div>
                  <div>
                    <SvgIcon :name="item.active ? 'eye' : 'eye-slash'" class="h-5 w-5" @click="updatePageActive(index)" />
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="mx-auto" v-else>Not found pages</div>
    </div>
    <div v-else>
      <header>
        <h1>{{ page.name }}</h1>
      </header>
      <div class="pt-5">
        <Editor v-model="content" />

        <hr class="my-5" />
        <FormButton type="submit" name="Save" color="green" @click="updatePageContent" />
      </div>
    </div>

    <drawer :is-open="isDrawer.open" max-width="700px" @close="closeDrawer">
      <div v-if="isDrawer.action === 'add'">
        <div class="pb-8">
          <div class="flex items-center">
            <div class="pr-3">
              <h1>New page</h1>
            </div>
          </div>
        </div>

        <Form @submit="addPage" v-slot="{ errors }">
          <div class="flow-root">
            <dl class="-my-3 text-sm mx-auto mb-0 mt-4 space-y-4">
              <FormInput v-model.trim="page.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="name" ico="at-symbol" />
              <div class="flex">
                <div class="pr-3">
                  <FormSelect v-model="page.type" :options="typePage" :error="errors.type" rules="required" id="type" title="Type"></FormSelect>
                </div>
                <div>
                  <FormInput v-model.trim="page.url" :error="errors.url" rules="required|alpha_num" id="url" type="text" title="Url" ico="glob-alt" />
                </div>
              </div>

            </dl>
          </div>

          <div class="pt-8">
            <div class="flex">
              <div class="flex-none">
                <FormButton type="submit" name="Add" color="green" class="mr-3" />
                <FormButton type="submit" name="Close" color="gray" @click="closeDrawer" />
              </div>
              <div class="grow"></div>
            </div>
          </div>
        </Form>
      </div>

      <div v-if="isDrawer.action === 'update'">
        <div class="pb-8">
          <div class="flex items-center">
            <div class="pr-3">
              <h1>Page setup</h1>
            </div>
          </div>
        </div>

        <Form @submit="updatePage" v-slot="{ errors }">
          <div class="flow-root">
            <dl class="-my-3 text-sm mx-auto mb-0 mt-4 space-y-4">
              <FormInput v-model.trim="page.name" :error="errors.name" rules="required|min:4" id="name" type="text" title="name" ico="at-symbol" />
              <div class="flex">
                <div class="pr-3">
                  <FormSelect v-model="page.type" :options="typePage" :error="errors.type" rules="required" id="type" title="Type"></FormSelect>
                </div>
                <div>
                  <FormInput v-model.trim="page.url" :error="errors.url" rules="required|alpha_num" id="url" type="text" title="Url" ico="glob-alt" />
                </div>
              </div>

            </dl>
          </div>

          <div class="pt-8">
            <div class="flex">
              <div class="flex-none">
                <FormButton type="submit" name="Save" color="green" class="mr-3" />
                <FormButton type="submit" name="Close" color="gray" @click="closeDrawer" />
              </div>
              <div class="grow"></div>
              <div class="mt-4 flex-none">
                <span @click="deletePage" class="cursor-pointer text-red-700">Delete</span>
              </div>
            </div>
          </div>
        </Form>
      </div>
    </drawer>

  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import MainLayouts from "@/layouts/Main.vue";
import { notify } from "notiwind";
import * as NProgress from "nprogress";
import SvgIcon from "svg-icon";

import { defineRule, Form } from "vee-validate";
import { required, alpha_num, min } from "@vee-validate/rules";
defineRule("required", required);
defineRule("alpha_num", alpha_num);
defineRule("min", min);

import { formatDate } from "@/utils/";
import Editor from "@/components/Editor.vue";
import FormInput from "@/components/form/Input.vue";
import FormSelect from "@/components/form/Select.vue";
import FormButton from "@/components/form/Button.vue";
import Drawer from "@/components/Drawer.vue";

const route = useRoute();
const router = useRouter()
const pages = ref([]);

const page = ref({});
const content = ref();

const isDrawer = ref({
  open: false,
  action: null,
});

const typePage = ["header", "footer"];

onMounted(async () => {
  pagesList();

  if (route.params.page_url) {
    pageContent(route.params.page_url)
  }
});

const pagesList = async () => {
  NProgress.start();

  await fetch(`/api/_/pages`, {
    credentials: "include",
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        pages.value = data.result;
      }
      NProgress.done();
    });
};

const openPage = (url) => {
  router.push({ name: 'pagesArticle', params: { 'page_url': url } });
  pageContent(url)
}

const pageContent = async (url) => {
  NProgress.start();

  try {
    const response = await fetch(`/api/pages/${url}`, {
      credentials: "include",
      method: "GET",
    });
    const { status } = response;
    const data = await response.json();

    if (data.success) {
      page.value = data.result;
      content.value = data.result.content;
    }

    if (status == 404) {
      router.push({ name: '404' });
    }

  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const addPage = async () => {
  const add = { ...page.value };
  console.log(add);

  try {
    const response = await fetch(`/api/_/pages`, {
      credentials: "include",
      method: "POST",
      body: JSON.stringify(add),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();

    if (data.success) {
      pages.value.push({
        id: data.result.id,
        name: data.result.name,
        url: data.result.url,
        type: data.result.type,
        created: data.result.created,
        active: data.result.active
      });
      closeDrawer();
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
}

const updatePageContent = async () => {
  NProgress.start();
  page.value.content = content.value;

  try {
    const response = await fetch(`/api/_/pages/${page.value.id}/content`, {
      credentials: "include",
      method: "PATCH",
      body: JSON.stringify(page.value),
      headers: {
        "Content-Type": "application/json",
      },
    });
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }

};

const updatePageActive = async (index) => {
  NProgress.start();

  try {
    const response = await fetch(`/api/_/pages/${pages.value[index].id}/active`, {
      credentials: "include",
      method: "PATCH",
    });
    const data = await response.json();

    if (data.success) {
      pages.value[index].active = !pages.value[index].active;
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const updatePage = async () => {
  const update = { ...page.value };
  NProgress.start();

  try {
    const response = await fetch(`/api/_/pages/${update.id}`, {
      credentials: "include",
      method: "PATCH",
      body: JSON.stringify(update),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();

    if (data.success) {
      const found = pages.value.find((e) => e.id === update.id);
      found.name = update.name;
      found.url = update.url;
      found.type = update.type;
      closeDrawer();
    } else {
      notify({
        group: "bottom",
        title: "Error",
        text: data.result,
      }, 4000)
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
}

const deletePage = async () => {
  const index = page.value.index;
  NProgress.start();

  console.log(pages.value[index].id)

  try {
    const response = await fetch(`/api/_/pages/${pages.value[index].id}`, {
      credentials: "include",
      method: "DELETE",
    });
    const data = await response.json();

    if (data.success) {
      pages.value.splice(index, 1);
    } else {
      const obj = JSON.parse(data.result);
      if (obj.code === "resource_missing") {
        console.log(obj.message);
      }
    }

    closeDrawer();
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
}

const openDrawer = (index, action) => {
  page.value = {};
  isDrawer.value.open = true;
  isDrawer.value.action = action;

  if (action === "update") {
    page.value = {
      name: pages.value[index].name,
      url: pages.value[index].url,
      type: pages.value[index].type,
      id: pages.value[index].id,
      index: index,
    }
  }
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};

</script>
