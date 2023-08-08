export default {
  props: {
    accept: {
      type: String,
      default: 'image/*',
    },
    productId: {
      type: String,
    },
    status: {
      type: Boolean,
    },
  },

  emits: ['added'],

  methods: {
    onChange() {
      ([...this.$refs.file.files]).forEach(f => {
        const formData = new FormData();
        formData.append('document', f);

        fetch(`/api/products/${this.productId}/image`, {
          credentials: 'include',
          method: 'POST',
          body: formData,
        })
          .then(response => response.json())
          .then(result => {
            this.$emit('added', result);
          })
      });
    },

    dragover(event) {
      event.preventDefault();
      if (!event.currentTarget.classList.contains('bg-green-300')) {
        event.currentTarget.classList.remove('bg-gray-100');
        event.currentTarget.classList.add('bg-green-300');
      }
    },

    dragleave(event) {
      event.currentTarget.classList.add('bg-gray-100');
      event.currentTarget.classList.remove('bg-green-300');
    },

    drop(event) {
      event.preventDefault();
      this.$refs.file.files = event.dataTransfer.files;
      this.onChange();
      event.currentTarget.classList.add('bg-gray-100');
      event.currentTarget.classList.remove('bg-green-300');
    },
  },

  template: `
  <div class="h-16 rounded-lg bg-gray-200 grid place-content-center cursor-pointer" @dragover="dragover" @dragleave="dragleave" @drop="drop">
  <input type="file" multiple name="fields[assetsFieldHandle][]" id="assetsFieldHandle" class="w-px h-px opacity-0 overflow-hidden absolute" @change="onChange" ref="file"
    :accept="accept" />
  <label for="assetsFieldHandle" class="block cursor-pointer">
    <svg class="h-5 w-5">
      <use xlink:href="/_/assets/img/sprite.svg#plus" />
    </svg>
  </label>
</div>`
}
