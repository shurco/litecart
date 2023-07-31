export default {
  data(props) { },

  components: {
    VField: Field,
  },

  emits: ['update:modelValue'],

  computed: {
    model: {
      get() {
        return this.modelValue;
      },
      set(value) {
        this.$emit('update:modelValue', value);
      },
    },
  },

  props: {
    modelValue: {
      type: String,
      default: '',
      required: true
    },
    id: {
      type: String,
      default: 'name'
    },
    name: {
      type: String,
      default: 'Name'
    },
    options: {
      type: Object,
      required: true
    },
    color: {
      type: String,
      default: 'indigo'
    },
    error: String,
  },

  template: `<div>
    <label :for="id" class="relative block rounded border border-gray-200 pe-[0.4rem] shadow-sm text-sm focus-within:border-blue-600 focus-within:ring-1 focus-within:ring-blue-600"
    :class="error? 'border-red-500' : ''">

    <v-field v-slot="{ value }" v-model="model" as="select" :name="id" :id="id" :rules="rules" class="w-full peer border-none bg-transparent placeholder-transparent focus:border-transparent focus:outline-none focus:ring-0">
      <option value="" disabled>Please select</option>
      <option v-for="option in options" :key="option" :value="option" :selected="value && value.includes(option)">{{ option }}</option>
    </v-field>

    <span class="pointer-events-none absolute start-2.5 top-0 -translate-y-1/2 bg-white p-0.5 text-xs text-gray-700 transition-all peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-focus:top-0 peer-focus:text-xs">
    {{name}}
    </span>
  </label>
  <span class="text-sm text-red-500 pl-4" v-if="error">{{ error }}</span>
</div>`
}

