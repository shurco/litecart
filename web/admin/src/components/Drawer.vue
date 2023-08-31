<template>
  <div :class="isVisible ? 'visible' : 'invisible'" class="drawer">
    <div class="fixed inset-x-0 inset-y-0 z-50 w-full select-none bg-black opacity-0 transition-opacity" style="transitionduration: `200 ms`" :class="{ 'opacity-50': isOpen }"></div>

    <div id="drawer_content" class="content" v-click-outside="closeDrawer" :class="isOpen ? 'translate-x-0' : 'translate-x-full'" :style="{
      maxWidth: maxWidth,
      transitionDuration: `200 ms`,
      backgroundColor: backgroundColor,
    }">
      <div class="pb-8">
        <slot name="header">
          <h1>{{ title }}</h1>
        </slot>
      </div>
      <slot />
      <div class="pt-8">
        <slot name="footer">
          <FormButton type="submit" name="Close" color="green" @click="closeDrawer" />
        </slot>
      </div>
    </div>
  </div>
</template>

<script setup>
import { getCurrentInstance, watch, ref } from "vue";
import FormButton from "@/components/form/Button.vue";

const props = defineProps({
  title: {
    type: String,
    default: "Header",
  },

  isOpen: {
    type: Boolean,
    required: false,
    default: false,
  },

  maxWidth: {
    type: String,
    required: false,
    default: "500px",
  },

  backgroundColor: {
    type: String,
    required: false,
    default: "#fafafa",
  },
});

const isVisible = ref(false);
const isTransitioning = ref(false);
const { emit } = getCurrentInstance();

watch(
  () => props.isOpen,
  (val) => {
    isTransitioning.value = true;

    if (val) {
      const drawerContent = document.getElementById("drawer_content");
      drawerContent.scrollTop = 0;

      toggleBackgroundScrolling(true);
      isVisible.value = true;
    } else {
      toggleBackgroundScrolling(false);
      setTimeout(() => (isVisible.value = false), 200);
    }

    setTimeout(() => (isTransitioning.value = false), 200);
  },
);

const toggleBackgroundScrolling = (enable) => {
  const body = document.querySelector("body");
  body.style.overflow = enable ? "hidden" : null;
};

const closeDrawer = () => {
  if (!isTransitioning.value) {
    emit("close");
  }
};

const vClickOutside = {
  mounted: (el, binding) => {
    el.clickOutsideEvent = function (event) {
      if (!(el == event.target || el.contains(event.target))) {
        binding.value(event, el);
      }
    };
    document.addEventListener("click", el.clickOutsideEvent);
  },
  unmounted: (el, binding) => {
    document.removeEventListener("click", el.clickOutsideEvent);
  },
};
</script>

<style lang="scss" scoped>
.drawer {
  & .content {
    @apply fixed inset-y-0 right-0 z-[999] flex h-full w-full flex-col overflow-auto bg-red-200 p-6 shadow-2xl transition-transform;
  }
}
</style>
