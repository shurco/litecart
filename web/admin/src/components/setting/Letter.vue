<template>
  <div>
    <div class="pb-8">
      <div class="flex items-center">
        <div class="pr-3">
          <h1>Update letter</h1>
        </div>
      </div>
    </div>
    <div class="flow-root">

      <div class="flow-root">
        <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
          <FormInput v-model.trim="letter.subject" id="subject" type="text" title="Subject" @focusout="updateLetter" />
        </dl>
      </div>

      <dl class="-my-3 mx-auto mb-0 space-y-4 text-sm mt-5">
        <FormTextarea v-model="letter.text" id="textarea" name="Message" :rows="15" @focusout="updateLetter" />
      </dl>
    </div>

    <div class="pt-5">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name="Close" color="gray" @click="close" />
        </div>
        <div class="grow"></div>
        <div class="flex-none">
          <FormButton type="submit" name="Test letter" color="cyan" @click="send(name)" />
        </div>
      </div>
    </div>

    <table class="mt-8 text-base">
      <tbody>
        <tr v-for="(value, key) in legend" class="cursor-default">
          <td class="w-32 font-bold">&#123;&#123;.{{ key }}&#125;&#125;</td>
          <td>{{ value }}</td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormInput, FormButton, FormTextarea } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

const letter = ref({});
const props = defineProps({
  name: String,
  legend: Object,
  send: Function,
  close: Function,
});

onMounted(() => {
  apiGet(`/api/_/settings/${props.name}`).then(res => {
    if (res.success) {
      const setting = res.result[props.name];
      letter.value.id = setting.id;
      letter.value.key = setting.key;
      letter.value.subject = JSON.parse(setting.value).subject;
      letter.value.text = JSON.parse(setting.value).text;
      letter.value.html = JSON.parse(setting.value).html;
    }
  });
});

const updateLetter = async () => {
  const value = new Object();
  value.subject = letter.value.subject;
  value.text = letter.value.text;
  value.html = letter.value.html;

  const update = {
    id: letter.value.id,
    key: letter.value.key,
    value: JSON.stringify(value),
  };

  apiUpdate(`/api/_/settings/${props.name}`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
      //props.close();
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
