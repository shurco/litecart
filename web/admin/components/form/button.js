export default {
  data() { },
  
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
  :class="color ? 'bg-'+color+'-600 active:bg-'+color+'-500' : 'bg-green-600 active:bg-'+color+'-500', ico ? 'focus:outline-none focus:ring' : ''">

  <span class="absolute -start-full transition-all group-hover:start-4" v-if="ico">
    <svg class="h-4 w-4">
      <use xlink:href="/_/assets/img/sprite.svg#arrow-right" v-if="ico==='row'" />
    </svg>
  </span>

  <span class="text-sm font-medium" :class="ico ? 'transition-all group-hover:ms-4' : ''">
    {{name}}
  </span>
</button>`
}

