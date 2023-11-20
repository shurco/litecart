import { ref } from "vue";
import { defineStore } from 'pinia';

export const useSystemStore = defineStore('system', () => {
  const version = ref({})
  const payments = ref({})

  return { version, payments }
})