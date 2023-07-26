export default {
  data() {
    return {
      isVisible: ref(false),
      isTransitioning: ref(false),
    }
  },

  setup(props) {
    watch(() => props.isOpen, (val) => {
      this.isTransitioning.value = true;

      if (val) {
        const drawerContent = document.getElementsByClassName("drawer__content")[0];
        drawerContent.scrollTop = 0;

        toggleBackgroundScrolling(true);
        this.isVisible.value = true;
      } else {
        toggleBackgroundScrolling(false);
        setTimeout(() => (this.isVisible.value = false), props.speed);
      }

      setTimeout(() => (this.isTransitioning.value = false), props.speed);
    })
  },


  props: {
    title: {
      type: String,
      required: true,
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

  methods: {
    toggleBackgroundScrolling(enable) {
      const body = document.querySelector("body")
      body.style.overflow = enable ? "hidden" : null
    },

    closeDrawer() {
      if (!this.isTransitioning.value) {
        emit("close")
      }
    },
  },

  template: `<div>
  <div class="drawer" :class="{ 'is-open': isOpen, 'is-visible': isVisible }">
    <div class="drawer__overlay" :style="{ transitionDuration: '200 ms' }"></div>
    <div class="drawer__content" v-click-outside="closeDrawer" :style="{
      maxWidth: maxWidth,
      transitionDuration: '200 ms',
      backgroundColor: backgroundColor,
    }">
      <div class="pb-4">
        <h2>{{ title }}</h2>
      </div>

      <slot />
      <div class="pt-4">
        <slot name="footer">
          <button class="btn" @click="closeDrawer">Close</button>
        </slot>
      </div>
    </div>

  </div>
</div>`
}

