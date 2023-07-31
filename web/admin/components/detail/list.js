export default {
  props: {
    name: {
      type: String,
      required: true,
    },
  },

  setup(props) { },

  template: `
  <div class="grid grid-cols-1 gap-1 py-3 sm:grid-cols-3 sm:gap-4">
    <dt class="font-medium text-gray-900">{{ name }}</dt>
    <dd class="text-gray-700 sm:col-span-2">
      <slot />
    </dd>
  </div>`
}

