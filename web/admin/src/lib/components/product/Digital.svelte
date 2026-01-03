<script lang="ts">
  import { onMount } from 'svelte'
  import FormInput from '../form/Input.svelte'
  import FormButton from '../form/Button.svelte'
  import Upload from '../form/Upload.svelte'
  import SvgIcon from '../SvgIcon.svelte'
  import { loadData } from '$lib/utils/apiHelpers'
  import { apiPost, apiUpdate, apiDelete } from '$lib/utils/api'
  import { showMessage } from '$lib/utils'
  import type { Product } from '$lib/types/models'

  interface Digital {
    type: string
    files: Array<{
      id: string
      name: string
      ext: string
      orig_name?: string
    }>
    data: Array<{
      id: string
      content: string
      cart_id: string | null
    }>
  }

  interface DrawerProduct {
    product: Product
    index: number
    currency?: string
  }

  interface Props {
    drawer: DrawerProduct
    onContentUpdate?: (() => void) | undefined
    onclose?: () => void
  }

  let { drawer, onContentUpdate, onclose }: Props = $props()

  let digital = $state<Digital>({
    type: '',
    files: [],
    data: []
  })
  let loading = $state(true)

  onMount(async () => {
    await loadDigital()
  })

  async function loadDigital() {
    loading = true
    const result = await loadData<Digital>(
      `/api/_/products/${drawer.product.id}/digital`,
      'Failed to load digital content'
    )
    if (result) {
      digital = {
        type: result.type || '',
        files: result.files || [],
        data: result.data || []
      }
    }
    loading = false
  }

  function close() {
    onclose?.()
  }

  async function handleUpload(event: CustomEvent) {
    if (event.detail.success && event.detail.result) {
      digital.files = [...digital.files, event.detail.result]
      showMessage('File uploaded', 'connextSuccess')
      if (onContentUpdate) {
        onContentUpdate()
      }
    }
  }

  function isCodeSold(cartId: string | null | undefined): boolean {
    if (!cartId || cartId === null) return false
    const trimmed = String(cartId).trim()
    return trimmed !== '' && trimmed !== 'null' && trimmed !== 'undefined'
  }

  async function addDigitalData() {
    const result = await apiPost(`/api/_/products/${drawer.product.id}/digital`)
    if (result.success && result.result) {
      digital.data = [...digital.data, result.result]
      showMessage('Data added', 'connextSuccess')
      if (onContentUpdate) {
        onContentUpdate()
      }
    } else {
      showMessage(result.message || 'Failed to add data', 'connextError')
    }
  }

  async function saveData(index: number) {
    const dataItem = digital.data[index]
    // Don't allow saving if code is sold (has cart_id)
    if (!dataItem || isCodeSold(dataItem.cart_id)) return

    const update = {
      content: dataItem.content
    }
    const result = await apiUpdate(`/api/_/products/${drawer.product.id}/digital/${dataItem.id}`, update)
    if (result.success) {
      showMessage('Data saved', 'connextSuccess')
    } else {
      showMessage(result.message || 'Failed to save data', 'connextError')
    }
  }

  async function deleteDigital(type: 'file' | 'data', index: number) {
    const digitalId = type === 'file' ? digital.files[index].id : digital.data[index].id
    const result = await apiDelete(`/api/_/products/${drawer.product.id}/digital/${digitalId}`)

    if (result.success) {
      if (type === 'file') {
        digital.files = digital.files.filter((_, i) => i !== index)
      } else {
        digital.data = digital.data.filter((_, i) => i !== index)
      }
      showMessage('Deleted', 'connextSuccess')
      if (onContentUpdate) {
        onContentUpdate()
      }
    } else {
      showMessage(result.message || 'Failed to delete', 'connextError')
    }
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>Digital {digital.type}</h1>
        {#if digital.type === 'file'}
          <p class="mt-4">
            This is the product that the user purchases. Upload the files that will be sent to the buyer after payment
            to the email address provided during checkout.
          </p>
        {/if}
        {#if digital.type === 'data'}
          <p class="mt-4">
            Enter the digital product that you intend to sell. It can be a unique item, such as a license key.
          </p>
        {/if}
      </div>
    </div>
  </div>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else if digital.type === 'file'}
    <!-- File section -->
    <div class="flow-root">
      <div class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
        {#if digital.files && digital.files.length > 0}
          <div class="grid content-start">
            {#each digital.files as file, index (file.id)}
              <div class="relative mt-4 flex first:mt-0">
                <a
                  href="/secrets/{file.name}.{file.ext}"
                  target="_blank"
                  class="rounded-lg bg-gray-200 px-3 py-3"
                  rel="noopener noreferrer"
                >
                  {file.orig_name || file.name}.{file.ext}
                </a>
                <div
                  class="mt-3 ml-3 cursor-pointer"
                  role="button"
                  tabindex="0"
                  onclick={() => deleteDigital('file', index)}
                  onkeydown={(e) => {
                    if (e.key === 'Enter' || e.key === ' ') {
                      e.preventDefault()
                      deleteDigital('file', index)
                    }
                  }}
                >
                  <SvgIcon name="trash" className="h-5 w-5" stroke="currentColor" />
                </div>
              </div>
            {/each}
          </div>
        {/if}
        <Upload section="digital" productId={drawer.product.id} onadded={handleUpload} />
      </div>
    </div>
  {:else if digital.type === 'data'}
    <!-- Data section -->
    <div class="flow-root">
      <div class="mx-auto -my-3 mt-4 mb-0 space-y-4 text-sm">
        {#if digital.data && digital.data.length > 0}
          {#each digital.data as dataItem, index (dataItem.id)}
            <div class="flex">
              {#if !isCodeSold(dataItem.cart_id)}
                <!-- Not sold - editable with delete button -->
                <div class="grow">
                  <FormInput
                    id="data-{dataItem.id}"
                    type="text"
                    title=""
                    bind:value={dataItem.content}
                    onfocusout={() => saveData(index)}
                  />
                </div>
                <div
                  class="flex-none cursor-pointer pt-3 pl-3"
                  role="button"
                  tabindex="0"
                  onclick={() => deleteDigital('data', index)}
                  onkeydown={(e) => {
                    if (e.key === 'Enter' || e.key === ' ') {
                      e.preventDefault()
                      deleteDigital('data', index)
                    }
                  }}
                >
                  <SvgIcon name="trash" className="h-5 w-5" stroke="currentColor" />
                </div>
              {:else}
                <!-- Sold - read-only with badge -->
                <div class="grow">
                  <div class="flex items-center gap-2 rounded-lg bg-gray-200 px-3 py-3">
                    <span class="flex-1">{dataItem.content}</span>
                    <span
                      class="inline-flex items-center rounded-full bg-red-100 px-2.5 py-0.5 text-xs font-medium text-red-800"
                      title="This code has been sold (cart_id: {dataItem.cart_id})"
                    >
                      Sold
                    </span>
                  </div>
                </div>
              {/if}
            </div>
          {/each}
        {/if}
        <div class="flex">
          <div class="grow"></div>
          <div class="mt-2 flex-none">
            <button
              type="button"
              class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700"
              onclick={addDigitalData}
            >
              Add data
            </button>
          </div>
        </div>
      </div>
    </div>
  {:else}
    <div class="mt-4 flow-root">Select digital type</div>
  {/if}

  <div class="pt-5">
    <FormButton type="button" name="Close" color="green" onclick={close} />
  </div>
</div>
