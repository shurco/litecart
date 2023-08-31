<template>
  <div class="pt-5">
    <Editor v-model="content" />

    <hr class="my-5" />
    <FormButton type="submit" name="Save" color="green" @click="updatePage" />
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

// @ts-ignore
import * as NProgress from "nprogress";
import Editor from "@/components/Editor.vue";
import FormButton from "@/components/form/Button.vue";

const page = ref();
const content = ref();

onMounted(() => {
  pageContent("cookies");
});

const pageContent = async (url) => {
  NProgress.start();

  await fetch(`/api/pages/${url}`, {
    credentials: "include",
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        page.value = data.result;
        content.value = data.result.content;
      }
      NProgress.done();
    });
};

const updatePage = async () => {
  NProgress.start();

  page.value.content = content.value;
  await fetch(`/api/_/pages/${page.value.id}`, {
    credentials: "include",
    method: "PATCH",
    body: JSON.stringify(page.value),
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      NProgress.done();
    });
};
</script>
