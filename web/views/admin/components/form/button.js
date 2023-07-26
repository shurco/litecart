export default {
  data() { },
  
  setup(props) { },

  props: {
    color: {
      type: String,
      default: 'indigo'
    },
    name: {
      type: String,
      default: 'Name'
    },
    ico: String,
  },

  template: `<button
  class="group relative inline-flex items-center overflow-hidden rounded px-8 py-3 text-white"
  :class="color ? 'bg-'+color+'-600 active:bg-'+color+'-500' : 'bg-green-600 active:bg-indigo-500', ico ? 'focus:outline-none focus:ring' : ''">

  <span class="absolute -start-full transition-all group-hover:start-4" v-if="ico==='row'">
    <svg class="h-5 w-5 rtl:rotate-180" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
      stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
    </svg>
  </span>

  <span class="text-sm font-medium" :class="ico ? 'transition-all group-hover:ms-4' : ''">
    {{name}}
  </span>
</button>`
}

