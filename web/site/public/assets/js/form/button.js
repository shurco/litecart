export default {
  data(props) {
    return {
      colors: {
        green: ["bg-green-600", "bg-green-500"],
        yellow: ["bg-yellow-600", "bg-yellow-500"],
        red: ["bg-red-600", "bg-red-500"],
        blue: ["bg-blue-600", "bg-blue-500"],
      }
    }
  },

  setup(props) { },

  props: {
    color: {
      type: String,
      required: true
    },
    name: {
      type: String,
      default: 'Name'
    },
    ico: String,
  },

  template: `<button
  class="group relative inline-flex items-center overflow-hidden rounded px-8 py-3 text-white"
  :class="color ? colors[color][0]+' active:'+colors[color][1] : '', ico ? 'focus:outline-none focus:ring' : ''">

  <span class="absolute -start-full transition-all group-hover:start-4" v-if="ico">
    <svg class="h-4 w-4" >
      <use xlink:href="/assets/img/sprite.svg#arrow-right" v-if="ico==='row'" />
      <use xlink:href="/assets/img/sprite.svg#cart" v-if="ico==='cart'" />
      <use xlink:href="/assets/img/sprite.svg#trash" v-if="ico==='trash'" />
      <use xlink:href="/assets/img/sprite.svg#plus" v-if="ico==='plus'" />
    </svg>
  </span>

  <span class="text-sm font-medium" :class="ico ? 'transition-all group-hover:ms-3 group-hover:-me-3' : ''">
    {{name}}
  </span>
</button>`
}

