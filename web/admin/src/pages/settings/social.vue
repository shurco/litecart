<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1>Socials</h1>
    </header>

    <Form @submit="updateSetting" v-slot="{ errors }">
      <div v-for="(value, key, index) in socialUrl" :key="index" class="mt-5 flex">
        <div class="pr-3 pt-2.5" v-if="key != 'other'">{{ socialUrl[key] }}</div>
        <div class="pr-3 pt-2.5" v-else>Link to other your site</div>
        <div>
          <FormInput v-model.trim="social[key]" :error="errors[`social_${key}`]" :rules="(key == 'other' ? 'url' : 'alpha_num')" class="w-48" :id="`social_${key}`" type="text"
            :title="(key == 'other' ? `Other site` : key.charAt(0).toUpperCase() + key.slice(1))" :ico="key" />
        </div>
      </div>
      <div class="pt-5">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormInput, FormButton } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const social = ref({});
const socialUrl = {
  facebook: "https://facebook.com/",
  instagram: "https://instagram.com/",
  twitter: "https://twitter.com/@",
  dribbble: "https://dribbble.com/",
  github: "https://github.com/",
  youtube: "https://youtube.com/",
  other: ""
};

onMounted(() => {
  apiGet(`/api/_/settings/social`).then(res => {
    if (res.success) {
      social.value = res.result;
    }
  });
});

const updateSetting = async () => {
  await apiUpdate(`/api/_/settings/social`, social.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
