export default {
  data(props) {
    return {
      placeholder: "Enter " + props.id,
    }
  },

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
    type: {
      type: String,
      default: 'text'
    },
    name: {
      type: String,
      default: 'Name'
    },
    color: {
      type: String,
      default: 'indigo'
    },
    rules: String,
    ico: String,
    error: String,
  },

  template: `<div>
  <label :for="id"
    class="relative block rounded border border-gray-200 pe-10 shadow-sm text-sm focus-within:border-blue-600 focus-within:ring-1 focus-within:ring-blue-600"
    :class="error? 'border-red-500' : ''">

    <v-field :type="type" :name="id" :rules="rules" :id="id" 
      v-model="model"
      class="w-full peer border-none bg-transparent placeholder-transparent focus:border-transparent focus:outline-none focus:ring-0"
      :placeholder="placeholder" autocomplete="on"></v-field>

    <span class="pointer-events-none absolute start-2.5 top-0 -translate-y-1/2 bg-white p-0.5 text-xs text-gray-700 transition-all peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-focus:top-0 peer-focus:text-xs">
      {{name}}
    </span>

    <span class="absolute inset-y-0 end-0 grid place-content-center px-4">
      <svg class="h-5 w-5" >
        <use v-bind="{'xlink:href':'/_/assets/img/sprite.svg#'+ico}" :class="error? 'text-red-500' : 'text-gray-400'" />
      </svg>
    </span>
  </label>
  <span class="text-sm text-red-500 pl-4" v-if="error">{{ error }}</span>
</div>`
}

