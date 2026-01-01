<script lang="ts">
  import { onMount } from 'svelte'
  import { createEventDispatcher } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import DetailList from '../DetailList.svelte'
  import SvgIcon from '../SvgIcon.svelte'
  import { costFormat, formatDate } from '$lib/utils'
  import { loadData } from '$lib/utils/apiHelpers'
  import type { Product } from '$lib/types/models'

  interface DrawerProduct {
    product: Product
    index: number
    currency?: string
  }

  export let drawer: DrawerProduct
  export let updateActive: ((index: number) => void) | undefined

  let product: Product | null = null
  let loading = true
  let lastProductId: string | null = null

  const dispatch = createEventDispatcher()

  async function loadProduct() {
    if (!drawer?.product?.id) return

    loading = true
    const result = await loadData<Product>(`/api/_/products/${drawer.product.id}`, 'Failed to load product')
    if (result) {
      product = result
      lastProductId = drawer.product.id
    }
    loading = false
  }

  onMount(async () => {
    await loadProduct()
  })

  // Reload product when drawer.product.id changes
  $: if (drawer?.product?.id && drawer.product.id !== lastProductId) {
    loadProduct()
  }

  function close() {
    dispatch('close')
  }

  async function active() {
    if (updateActive && product) {
      await updateActive(drawer.index)
      // Update local product state reactively
      if (product) {
        product = { ...product, active: !product.active }
      }
    }
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>View {product?.name || 'Product'}</h1>
      </div>
      <div>
        {#if product}
          <SvgIcon
            name={product.active ? 'eye' : 'eye-slash'}
            className="h-5 w-5 cursor-pointer"
            on:click={active}
            stroke="currentColor"
          />
        {/if}
      </div>
    </div>
  </div>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else if product}
    <div class="flow-root">
      <dl class="-my-3 mt-2 divide-y divide-gray-100 text-sm">
        <DetailList name="ID">{product.id}</DetailList>
        <DetailList name="Name">{product.name}</DetailList>
        <DetailList name="Price">{costFormat(product.amount)} {drawer.currency || ''}</DetailList>
        <DetailList name="Slug">{product.slug}</DetailList>
        <DetailList name="Metadata">
          {#each product.metadata || [] as data}
            <div>{data.key}: {data.value}</div>
          {/each}
        </DetailList>
        <DetailList name="Attributes">
          {#each product.attributes || [] as item}
            <div>{item}</div>
          {/each}
        </DetailList>
        <DetailList name="Created">{formatDate(product.created)}</DetailList>
        {#if product.updated}
          <DetailList name="Updated">{formatDate(product.updated)}</DetailList>
        {/if}
        {#if product.images}
          <DetailList name="Images" grid={true}>
            {#each product.images as item}
              <div>
                <a href="/uploads/{item.name}.{item.ext}" target="_blank" aria-label="View full size image">
                  <img
                    style="width: 100%; max-width: 150px"
                    src="/uploads/{item.name}_sm.{item.ext}"
                    alt="{product.name} - {item.name}"
                    loading="lazy"
                  />
                </a>
              </div>
            {/each}
          </DetailList>
        {/if}
        <DetailList name="Brief (short description)">{product.brief}</DetailList>

        <div class="tiptap pt-3">{@html product.description || ''}</div>
      </dl>
    </div>
  {:else}
    <div class="py-8 text-center text-gray-500">Failed to load product</div>
  {/if}

  <div class="pt-5">
    <FormButton type="button" name="Close" color="green" on:click={close} />
  </div>
</div>
