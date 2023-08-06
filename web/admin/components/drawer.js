import FormButton from './form/button.js';

export default {
  components: {
    FormButton,
  },

  props: {
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
  },

  setup(props) {
    const isVisible = ref(false);
    const isTransitioning = ref(false);
    const { emit } = getCurrentInstance();

    watch(() => props.isOpen, (val) => {
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
    })

    const toggleBackgroundScrolling = (enable) => {
      const body = document.querySelector("body")
      body.style.overflow = enable ? "hidden" : null
    }

    const closeDrawer = () => {
      if (!isTransitioning.value) {
        emit("close")
      }
    }

    return { isVisible, isTransitioning, closeDrawer };
  },

  directives: {
    ClickOutside: {
      mounted: function (el, binding) {
        el.clickOutsideEvent = function (event) {
          if (!(el == event.target || el.contains(event.target))) {
            binding.value(event, el)
          }
        }
        document.addEventListener("click", el.clickOutsideEvent)
      },
      unmounted: function (el) {
        document.removeEventListener("click", el.clickOutsideEvent)
      },
    }
  },

  template: `<div :class=" isVisible ? 'visible' : 'invisible'">
    <div class="fixed inset-x-0 inset-y-0 w-full z-50 opacity-0 transition-opacity bg-black select-none" 
      style="transitionDuration: \`200 ms\`"
      :class="{ 'opacity-50': isOpen }"></div>
      
    <div id="drawer_content" class="fixed inset-y-0 h-full w-full right-0 overflow-auto flex transition-transform flex-col shadow-2xl z-[999] bg-red-200 p-6" 
      v-click-outside="closeDrawer" 
      :class=" isOpen ? 'translate-x-0' : 'translate-x-full' "
      :style="{
        maxWidth: maxWidth,
        transitionDuration: \`200 ms\`,
        backgroundColor: backgroundColor,
      }">
      <div class="pb-8">
        <slot name="header">
          <h2 class="text-2xl font-bold text-gray-900 sm:text-3xl">{{ title }}</h2>
        </slot>
      </div>
      <slot />
      <div class="pt-8">
        <slot name="footer">
          <form-button type="submit" name="Close" color="green" @click="closeDrawer" />
        </slot>
      </div>
    </div>
  </div>`
}

