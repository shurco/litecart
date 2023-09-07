import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/router";
import Notifications from 'notiwind'

import { defineRule } from 'vee-validate';
import rules from '@vee-validate/rules';

import "@/assets/app.css";

// validate rules
Object.keys(rules).forEach(rule => {
  defineRule(rule, rules[rule]);
});
defineRule("amount", (value) => {
  if (!value || !value.length) {
    return true;
  }
  if (!/^\d+(\.\d{1,2})?$/.test(value)) {
    return "amount is not valid";
  }
  return true;
});
defineRule('confirmed', (value, [target], ctx) => {
  if (value === ctx.form[target]) {
    return true;
  }
  return 'Passwords must match';
});

const app = createApp(App);
app.use(router);
app.use(Notifications);
app.mount("#app");
