<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import SvgIcon from '../SvgIcon.svelte';
  import { apiPost } from '$lib/utils/api';

  export let section: string;
  export let accept: string | undefined = undefined;
  export let productId: string | undefined = undefined;

  const dispatch = createEventDispatcher();

  let fileInput: HTMLInputElement;
  let isDragging = false;

  const onChange = async () => {
    if (!fileInput?.files) return;
    
    for (const file of fileInput.files) {
      const formData = new FormData();
      formData.append("document", file);
      const res = await apiPost(`/api/_/products/${productId}/${section}`, formData);
      dispatch('added', res);
    }
  };

  const dragover = (event: DragEvent) => {
    event.preventDefault();
    isDragging = true;
  };

  const dragleave = (event: DragEvent) => {
    event.preventDefault();
    isDragging = false;
  };

  const drop = (event: DragEvent) => {
    event.preventDefault();
    if (fileInput && event.dataTransfer?.files) {
      fileInput.files = event.dataTransfer.files;
      onChange();
    }
    isDragging = false;
  };
</script>

<div
  class="upload bg-gray-200 {isDragging ? 'bg-green-300' : ''}"
  role="button"
  tabindex="0"
  aria-label="Upload file"
  on:dragover={dragover}
  on:dragleave={dragleave}
  on:drop={drop}
  on:keydown={(e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      fileInput?.click();
    }
  }}
>
  <input
    type="file"
    multiple
    name="fields[assetsFieldHandle][]"
    id="assetsFieldHandle"
    on:change={onChange}
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
    @apply block cursor-pointer p-0 border-0 shadow-none;
  }
</style>