
export default {
  props: {
    type: {
      type: String,
      required: false,
      default: "",
    },
    svgClass: {
      type: String,
      required: false,
      default: "-ms-0.5 me-1.5 h-4 w-4",
    },
  },

  setup(props) { },

  template: `
  <span class="inline-flex items-center justify-center rounded-full px-2.5 py-0.5">
    <svg :class="(type !== '' ? '-ms-0.5 me-1.5 '+svgClass : 'h-0 w-0' )" v-if="type!==''">
      <use v-bind="{'xlink:href':'/assets/img/sprite.svg#'+type}" />
    </svg>
    <p class="whitespace-nowrap text-xs">
      <slot />
    </p>
  </span>`
}

