<template>
  <component :is="route.meta.layoutComponent">
    <RouterView />
  </component>
  <Alert />
</template>

<script setup>
import { RouterView, useRoute } from "vue-router";
import { Alert } from "@/components/";

const route = useRoute();
</script>
