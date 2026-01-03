<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte'
  import { apiPost } from '$lib/utils/api'

  interface Props {
    section: string
    accept?: string
    productId?: string
    onadded?: (res: any) => void
  }

  let { section, accept = undefined, productId = undefined, onadded }: Props = $props()

  let fileInput: HTMLInputElement | undefined = $state()
  let isDragging = $state(false)

  const onChange = async () => {
    if (!fileInput?.files) return

    for (const file of fileInput.files) {
      const formData = new FormData()
      formData.append('document', file)
      const res = await apiPost(`/api/_/products/${productId}/${section}`, formData)
      onadded?.(res)
    }
  }

  const dragover = (event: DragEvent) => {
    event.preventDefault()
    isDragging = true
  }

  const dragleave = (event: DragEvent) => {
    event.preventDefault()
    isDragging = false
  }

  const drop = (event: DragEvent) => {
    event.preventDefault()
    if (fileInput && event.dataTransfer?.files) {
      fileInput.files = event.dataTransfer.files
      onChange()
    }
    isDragging = false
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault()
      fileInput?.click()
    }
  }
</script>

<div
  class="upload bg-gray-200 {isDragging ? 'bg-green-300' : ''}"
  role="button"
  tabindex="0"
  aria-label="Upload file"
  ondragover={dragover}
  ondragleave={dragleave}
  ondrop={drop}
  onkeydown={handleKeydown}
>
  <input
    type="file"
    multiple
    name="fields[assetsFieldHandle][]"
    id="assetsFieldHandle"
    onchange={onChange}
    bind:this={fileInput}
    {accept}
  />
  <label for="assetsFieldHandle">
    <SvgIcon name="plus" className="h-5 w-5" stroke="currentColor" />
  </label>
</div>

<style>
  @reference "tailwindcss";

  :global(.upload) {
    @apply grid h-16 cursor-pointer place-content-center rounded-lg;
  }

  :global(.upload input) {
    @apply absolute h-px w-px overflow-hidden opacity-0;
  }

  :global(.upload label) {
    @apply block cursor-pointer border-0 p-0 shadow-none;
  }
</style>
