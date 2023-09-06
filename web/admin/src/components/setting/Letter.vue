<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Update letter</h1>
        </div>
      </div>
    </div>
    <Form @submit="updateLetter" v-slot="{ errors }">
      <div class="flow-root">
        <dl class="-my-3 text-sm mx-auto mb-0 mt-2 space-y-4">
          <FormTextarea v-model="letter.value" id="textarea" name="Message" :rows="15" />
        </dl>
      </div>

      <div class="pt-8">
        <div class="flex">
          <div class="flex-none">
            <FormButton type="submit" name="Save" color="green" class="mr-3" />
            <FormButton type="submit" name="Close" color="gray" @click="close" />
          </div>
          <div class="grow"></div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import FormButton from "@/components/form/Button.vue";
import FormTextarea from "@/components/form/Textarea.vue";
import { notifyMessage } from "@/utils/";

import * as NProgress from "nprogress";
import { Form } from "vee-validate";

const props = defineProps({
  name: String,
  close: Function,
})

onMounted(() => {
  settingLetter();
});

const letter = ref({})

const settingLetter = async () => {
  try {
    NProgress.start();

    const response = await fetch(`/api/_/settings/${props.name}`, {
      credentials: "include",
      method: "GET",
    });
    const { success, result } = await response.json();

    if (success) {
      letter.value = result;
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};

const updateLetter = async () => {
  try {
    NProgress.start();

    const response = await fetch(`/api/_/settings/${props.name}`, {
      credentials: "include",
      method: "PATCH",
      body: JSON.stringify(letter.value),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const { success, result, message } = await response.json();

    if (success) {
      notifyMessage("Perfect", message, "success");
      props.close();
    } else {
      notifyMessage("Error", result, "error");
    }
  } catch (error) {
    console.error(error);
  } finally {
    NProgress.done();
  }
};
</script>