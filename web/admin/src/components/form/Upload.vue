<template>
  <div class="h-16 rounded-lg bg-gray-200 grid place-content-center cursor-pointer" @dragover="dragover" @dragleave="dragleave" @drop="drop">
    <input type="file" multiple name="fields[assetsFieldHandle][]" id="assetsFieldHandle" class="w-px h-px opacity-0 overflow-hidden absolute" @change="onChange" ref="file"
      :accept="accept" />
    <label for="assetsFieldHandle" class="block cursor-pointer">
      <SvgIcon name="plus" class="h-5 w-5" />
    </label>
  </div>
</template>

<script setup>
import { getCurrentInstance } from 'vue'
import SvgIcon from 'svg-icon'

const props = defineProps({
  accept: {
    type: String,
    default: 'image/*'
  },
  productId: String,
  status: Boolean,
})

const instance = getCurrentInstance()
const emits = defineEmits(['added'])

const onChange = () => {
  [...instance.refs.file.files].forEach((f) => {
    const formData = new FormData()
    formData.append('document', f)

    fetch(`/api/_/products/${props.productId}/image`, {
      credentials: 'include',
      method: 'POST',
      body: formData
    })
      .then((response) => response.json())
      .then((data) => {
        emits('added', data)
      })
  })
}

const dragover = (event) => {
  event.preventDefault()
  if (!event.currentTarget.classList.contains('bg-green-300')) {
    event.currentTarget.classList.remove('bg-gray-100')
    event.currentTarget.classList.add('bg-green-300')
  }
}

const dragleave = (event) => {
  event.currentTarget.classList.add('bg-gray-100')
  event.currentTarget.classList.remove('bg-green-300')
}

const drop = (event) => {
  event.preventDefault()
  instance.refs.file.files = event.dataTransfer.files
  onChange()
  event.currentTarget.classList.add('bg-gray-100')
  event.currentTarget.classList.remove('bg-green-300')
}
</script>
