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
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
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
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

import { Form } from "vee-validate";

const props = defineProps({
  name: String,
  close: Function,
});

onMounted(() => {
  settingLetter();
});

const letter = ref({});

const settingLetter = async () => {
  apiGet(`/api/_/settings/${props.name}`).then(res => {
    if (res.success) {
      letter.value = res.result;
    }
  });
};

const updateLetter = async () => {
  apiUpdate(`/api/_/settings/${props.name}`, letter.value).then(res => {
    if (res.success) {
      showMessage(res.message);
      props.close();
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
