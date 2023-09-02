<template>
  <MainLayouts>
    <div v-if="!route.params.page_url">
      <header>
        <h1>Pages</h1>
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
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in pages" :key="item.id" @click="openPage(item.url)">
              <td>{{ item.name }}</td>
              <td>{{ item.type }}</td>
              <td>{{ item.url }}</td>
              <td>{{ formatDate(item.created) }}</td>
              <td v-if="item.updated">{{ formatDate(item.updated) }}</td>
              <td v-else></td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="mx-auto" v-else>Not found checkouts</div>
    </div>
    <div v-else>
      <header>
        <h1>{{page.name}}</h1>
      </header>

      <div class="pt-5">
        <Editor v-model="content" />

        <hr class="my-5" />
        <FormButton type="submit" name="Save" color="green" @click="updatePage" />
      </div>
    </div>

  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import MainLayouts from "@/layouts/Main.vue";
import * as NProgress from "nprogress";
import Editor from "@/components/Editor.vue";
import FormButton from "@/components/form/Button.vue";
import { formatDate } from "@/utils/";

const route = useRoute();
const router = useRouter()
const pages = ref([]);

const page = ref({});
const content = ref();

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

const updatePage = async () => {
  NProgress.start();
  page.value.content = content.value;

  try {
    const response = await fetch(`/api/_/pages/${page.value.id}`, {
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
</script>
