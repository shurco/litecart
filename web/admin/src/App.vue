<template>
  <RouterView />
  <Alert />
</template>

<script setup>
import { RouterView } from "vue-router";
import Alert from "@/components/Alert.vue";
</script>
