export default {
  props: {
    ico: {
      type: String,
      required: false,
      default: "",
    },
    svgClass: {
      type: String,
      required: false,
      default: "-ms-0.5 me-5.5 h-5 w-5",
    },
  },

  setup(props) { },

  template: `<span class="inline-flex items-center justify-center rounded-full px-2.5 py-0.5">
    <svg :class="(ico !== '' ? '-ms-0.5 me-1.5 '+svgClass : 'h-0 w-0' )" v-if="ico!==''">
      <use v-bind="{'xlink:href':'/_/assets/img/sprite.svg#'+ico}" />
    </svg>
    <p class="whitespace-nowrap text-xs">
      <slot />
    </p>
  </span>`
}