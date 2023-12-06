<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1><span class="text-gray-300">Settings</span><span class="px-3 text-gray-300">/</span>Mail</h1>
    </header>

    <div>
      <h2 class="mb-5">Mail letters</h2>
      <div class="flex">
        <div class="cursor-pointer rounded bg-gray-200 p-2" @click="openDrawer('mail_letter_payment')">Letter of payment</div>
        <div class="cursor-pointer rounded bg-gray-200 p-2 ml-5" @click="openDrawer('mail_letter_purchase')">Letter of purchase</div>
      </div>
      <hr class="mt-5" />
    </div>

    <div class="mt-5">
      <h2 class="mb-5">SMTP settings</h2>
      <div class="mb-5 flex items-center justify-between bg-red-600 px-2 py-3 text-white" v-if="!smtp.host || !smtp.port || !smtp.username || !smtp.password">
        <p class="text-sm font-medium">This section is required!</p>
      </div>

      <Form @submit="updateSetting()" v-slot="{ errors }">
        <div class="flex">
          <div class="pr-3">
            <FormInput v-model.trim="smtp.host" :error="errors.smtp_host" rules="required|min:4" class="w-64" id="smtp_host" type="text" title="SMTP host" ico="server" />
          </div>
          <div class="pr-3">
            <FormInput v-model.trim="smtp.port" :error="errors.smtp_port" rules="required|numeric" class="w-64" id="smtp_port" type="text" title="SMTP port"
              ico="arrow-left-on-rectangle" />
          </div>
          <div>
            <FormSelect v-model="smtp.encryption" :options="['None', 'SSL/TLS', 'STARTTLS']" :error="errors.smtp_encryption" rules="required" id="smtp_encryption" title="Encryption"
              ico="lock-closed" />
          </div>
        </div>
        <div class="mt-5 flex">
          <div class="pr-3">
            <FormInput v-model.trim="smtp.username" :error="errors.smtp_username" rules="required" class="w-64" id="smtp_username" type="text" title="Username" ico="user" />
          </div>
          <div>
            <FormInput v-model.trim="smtp.password" :error="errors.smtp_password" rules="required|min:6" class="w-64" id="smtp_password" type="password" title="Password"
              ico="finger-print" />
          </div>
        </div>
        <div class="flex pt-8">
          <FormButton type="submit" name="Save" color="green" class="flex-none" />
          <div class="ml-5 mt-3 flex-none">
            <span @click="sendTestLetter('smtp')" class="cursor-pointer text-red-700">Test smtp</span>
          </div>
        </div>
      </Form>
    </div>
  </div>

  <drawer :is-open="isDrawer.open" max-width="725px" @close="closeDrawer">
    <Letter :close="closeDrawer" :send="sendTestLetter" :legend="letterLegend['mail_letter_payment']" name="mail_letter_payment" v-if="isDrawer.action === 'mail_letter_payment'" />
    <Letter :close="closeDrawer" :send="sendTestLetter" :legend="letterLegend['mail_letter_purchase']" name="mail_letter_purchase"
      v-if="isDrawer.action === 'mail_letter_purchase'" />
  </drawer>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { FormInput, FormButton, FormSelect, Drawer, Letter } from "@/components/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const smtp = ref({});

const isDrawer = ref({
  open: false,
  action: null,
});

const letterLegend = {
  "mail_letter_payment": {
    "Site_Name": "Site name",
    "Amount_Payment": "Amount of payment",
    "Payment_URL": "Payment link",
  },
  "mail_letter_purchase": {
    "Purchases": "Purchases",
    "Admin_Email": "Admin email",
  }
}

onMounted(() => {
  apiGet(`/api/_/settings/mail`).then(res => {
    if (res.success) {
      smtp.value = res.result;
    }
  });
});

const updateSetting = async () => {
  var update = {};
  update = smtp.value;
  update.port = Number(smtp.value.port);

  apiUpdate(`/api/_/settings/mail`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const sendTestLetter = async (letterName) => {
  apiGet(`/api/_/test/letter/${letterName}`).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const openDrawer = (action) => {
  isDrawer.value.open = true;
  isDrawer.value.action = action;
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};
</script>
