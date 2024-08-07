import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "@/App.vue";
import router from "@/router";
import Notifications from "notiwind";
import VueTippy from "vue-tippy";

import SvgIcon from "@/components/SvgIcon.vue";
import "virtual:svg-icons-register";

import { defineRule } from "vee-validate";
import * as rules from "@vee-validate/rules";

import "@/assets/app.css";

// validate rules
Object.keys(rules)
  .filter(rule => typeof rules[rule] === "function")
  .forEach(rule => defineRule(rule, rules[rule]));
defineRule("amount", (value) => {
  if (!value || !value.length) {
    return true;
  }
  if (!/^\d+(\.\d{1,2})?$/.test(value)) {
    return "amount is not valid";
  }
  return true;
});
defineRule("slug", (value) => {
  if (!value || !value.length) {
    return true;
  }
  if (!/^[a-z0-9]+(?:-[a-z0-9]+)*$/.test(value)) {
    return "slug is not valid";
  }
  return true;
});
defineRule("confirmed", (value, [target], ctx) => {
  if (value === ctx.form[target]) {
    return true;
  }
  return "Passwords must match";
});

const pinia = createPinia();
const app = createApp(App);
app.use(pinia);
app.use(router);
app.use(Notifications);
app.use(VueTippy, {
  defaultProps: {
    theme: 'lite',
    delay: [500, null],
  },
});
app.component("SvgIcon", SvgIcon);
app.mount("#app");
