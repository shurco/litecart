<template>
  <div class="pb-10">
    <header class="mb-4">
      <h1>Authentication</h1>
    </header>

    <Form @submit="updateAuth" v-slot="{ errors }">
      <FormInput v-model.trim="auth.email" :error="errors.email" rules="required|email" class="w-64" id="email" type="text" title="Email" ico="at-symbol" />
      <div class="pt-5">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </Form>
    <hr class="mt-5" />

    <div class="mt-5">
      <h2 class="mb-5">JWT parameters</h2>
      <Form @submit="updateJWT" v-slot="{ errors }">
        <div class="mt-5 flex">
          <div class="pr-3">
            <FormInput v-model.trim="jwt.secret" :error="errors.jwt_secret" rules="required|min:30" class="w-64" id="jwt_secret" type="text" title="JWT Secret" ico="key" />
          </div>
          <div>
            <FormInput v-model.trim="jwt.expire_hours" :error="errors.jwt_expire_hours" rules="required|numeric" class="w-64" id="jwt_expire_hours" type="text"
              title="JWT expire hours" ico="key" />
          </div>
        </div>
        <div class="pt-5">
          <FormButton type="submit" name="Save" color="green" />
        </div>
      </Form>
      <hr class="mt-5" />
    </div>

    <div class="mt-5">
      <h2 class="mb-5">Change password</h2>
      <Form @submit="updatePassword" v-slot="{ errors }">
        <FormInput v-model.trim="password.old" :error="errors.password_old" rules="required|min:6" class="w-96" id="password_old" type="password" title="Actual password"
          ico="finger-print" />
        <div class="mt-5 flex">
          <div class="pr-3">
            <FormInput v-model.trim="password.new" :error="errors.password_new" rules="required|min:6" class="w-96" id="password_new" type="password" title="New password"
              ico="finger-print" />
          </div>
          <div>
            <FormInput v-model.trim="password.new2" :error="errors.password_new2" rules="required|confirmed:password_new" class="w-96" id="password_new2" type="password"
              title="Repeat password" ico="finger-print" />
          </div>
        </div>
        <div class="pt-5">
          <FormButton type="submit" name="Save" color="green" />
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { FormInput, FormButton } from "@/components";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";
import { Form } from "vee-validate";

const auth = ref({});
const jwt = ref({});
const password = ref({});

onMounted(() => {
  apiGet(`/api/_/settings/auth`).then(res => {
    if (res.success) {
      auth.value = res.result;
    }
  });
  apiGet(`/api/_/settings/jwt`).then(res => {
    if (res.success) {
      jwt.value = res.result;
    }
  });
});

const updatePassword = async () => {
  var update = {
    old: password.value.old,
    new: password.value.new,
  };

  await apiUpdate(`/api/_/settings/password`, update).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const updateJWT = async () => {
  jwt.value.expire_hours = Number(jwt.value.expire_hours);
  await apiUpdate(`/api/_/settings/jwt`, jwt.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};

const updateAuth = async () => {
  await apiUpdate(`/api/_/settings/auth`, auth.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  });
};
</script>
