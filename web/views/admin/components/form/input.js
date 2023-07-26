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
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
        class="w-5 h-5 text-gray-400" :class="error? 'border-red-500' : ''" v-if="ico==='email'">
        <path stroke-linecap="round"
          d="M16.5 12a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0zm0 0c0 1.657 1.007 3 2.25 3S21 13.657 21 12a9 9 0 10-2.636 6.364M16.5 12V8.25" />
      </svg>

      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
        class="w-5 h-5 text-gray-400" :class="error? 'border-red-500' : ''" v-if="ico==='password'">
        <path stroke-linecap="round" stroke-linejoin="round"
          d="M7.864 4.243A7.5 7.5 0 0119.5 10.5c0 2.92-.556 5.709-1.568 8.268M5.742 6.364A7.465 7.465 0 004.5 10.5a7.464 7.464 0 01-1.15 3.993m1.989 3.559A11.209 11.209 0 008.25 10.5a3.75 3.75 0 117.5 0c0 .527-.021 1.049-.064 1.565M12 10.5a14.94 14.94 0 01-3.6 9.75m6.633-4.596a18.666 18.666 0 01-2.485 5.33" />
      </svg>

      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
        class="w-5 h-5 text-gray-400" :class="error? 'border-red-500' : ''" v-if="ico==='domain'">
        <path stroke-linecap="round" stroke-linejoin="round"
          d="M15.042 21.672L13.684 16.6m0 0l-2.51 2.225.569-9.47 5.227 7.917-3.286-.672zM12 2.25V4.5m5.834.166l-1.591 1.591M20.25 10.5H18M7.757 14.743l-1.59 1.59M6 10.5H3.75m4.007-4.243l-1.59-1.59" />
      </svg>

      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
        class="w-5 h-5 text-gray-400" :class="error? 'border-red-500' : ''" v-if="ico==='key'">
        <path stroke-linecap="round" stroke-linejoin="round"
          d="M15.75 5.25a3 3 0 013 3m3 0a6 6 0 01-7.029 5.912c-.563-.097-1.159.026-1.563.43L10.5 17.25H8.25v2.25H6v2.25H2.25v-2.818c0-.597.237-1.17.659-1.591l6.499-6.499c.404-.404.527-1 .43-1.563A6 6 0 1121.75 8.25z" />
      </svg>

    </span>
  </label>
  <span class="text-sm text-red-500 pl-4" v-if="error">{{ error }}</span>
</div>`
}

