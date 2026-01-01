<script lang="ts">
  import { onMount } from 'svelte'
  import { createEventDispatcher } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import DetailList from '../DetailList.svelte'
  import SvgIcon from '../SvgIcon.svelte'
  import { costFormat, formatDate, STRIPE_DASHBOARD_URL } from '$lib/utils'
  import { loadData } from '$lib/utils/apiHelpers'
  import type { CartDetail } from '$lib/types/models'

  interface DrawerCart {
    cart: {
      id: string
      email: string
      amount_total: number
      currency: string
      payment_status: 'paid' | 'pending' | 'failed'
      payment_system?: string
      payment_id?: string
      created?: string
      updated?: string
    }
  }

  export let drawer: DrawerCart

  let cart: CartDetail | null = null
  let loading = true
  let lastCartId: string | null = null

  const dispatch = createEventDispatcher()

  async function loadCart() {
    if (!drawer?.cart?.id) return

    loading = true
    const result = await loadData<CartDetail>(`/api/_/carts/${drawer.cart.id}`, 'Failed to load cart')
    if (result) {
      cart = result
      lastCartId = drawer.cart.id
    }
    loading = false
  }

  onMount(async () => {
    await loadCart()
  })

  // Reload cart when drawer.cart.id changes
  $: if (drawer?.cart?.id && drawer.cart.id !== lastCartId) {
    loadCart()
  }

  function close() {
    dispatch('close')
  }

  function getPaymentStatusColor(status: string) {
    switch (status) {
      case 'paid':
        return 'text-green-600'
      case 'pending':
        return 'text-yellow-600'
      case 'failed':
        return 'text-red-600'
      default:
        return 'text-gray-600'
    }
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>Cart Details</h1>
      </div>
    </div>
  </div>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else if cart}
    <div class="flow-root">
      <dl class="-my-3 mt-2 divide-y divide-gray-100 text-sm">
        <DetailList name="Cart ID">{cart.id}</DetailList>
        
        <DetailList name="Customer Email">
          {#if cart.email}
            <a href="mailto:{cart.email}" class="text-blue-600 hover:underline">{cart.email}</a>
          {:else}
            <span class="text-gray-400">-</span>
          {/if}
        </DetailList>

        <DetailList name="Total Amount">
          {#if !cart.amount_total || cart.amount_total === 0}
            <span class="font-bold text-green-600">free</span>
          {:else if cart.payment_id && cart.payment_system === 'stripe'}
            <a
              href="{STRIPE_DASHBOARD_URL}/{cart.payment_id}"
              target="_blank"
              class="text-blue-600 hover:underline"
            >
              {costFormat(cart.amount_total)} {cart.currency || ''}
            </a>
          {:else}
            {costFormat(cart.amount_total)} {cart.currency || ''}
          {/if}
        </DetailList>

        <DetailList name="Payment Status">
          <span class={getPaymentStatusColor(cart.payment_status || '')}>
            {cart.payment_status || '-'}
          </span>
        </DetailList>

        <DetailList name="Payment System">{cart.payment_system || '-'}</DetailList>

        {#if cart.payment_id}
          <DetailList name="Payment ID">
            {#if cart.payment_system === 'stripe'}
              <a
                href="{STRIPE_DASHBOARD_URL}/{cart.payment_id}"
                target="_blank"
                class="text-blue-600 hover:underline"
              >
                {cart.payment_id}
              </a>
            {:else}
              {cart.payment_id}
            {/if}
          </DetailList>
        {/if}

        <DetailList name="Created">{formatDate(cart.created)}</DetailList>
        
        {#if cart.updated}
          <DetailList name="Updated">{formatDate(cart.updated)}</DetailList>
        {/if}

        {#if cart.items && cart.items.length > 0}
          <DetailList name="Items" grid={false}>
            <div class="space-y-4">
              {#each cart.items as item}
                <div class="flex items-start gap-4 border-b border-gray-200 pb-4 last:border-0">
                  {#if item.image}
                    <div class="flex-shrink-0">
                      <a
                        href="/uploads/{item.image.name}.{item.image.ext}"
                        target="_blank"
                        aria-label="View full size image"
                      >
                        <img
                          class="w-20 h-20 object-cover rounded"
                          src="/uploads/{item.image.name}_sm.{item.image.ext}"
                          alt="{item.name}"
                          loading="lazy"
                        />
                      </a>
                    </div>
                  {/if}
                  <div class="flex-1 min-w-0">
                    <div class="font-medium text-gray-900">{item.name}</div>
                    <div class="text-sm text-gray-500">Slug: {item.slug}</div>
                    <div class="mt-1 text-sm text-gray-700">
                      <span class="font-medium">Price:</span> {costFormat(item.amount)} {cart.currency || ''}
                    </div>
                    <div class="mt-1 text-sm text-gray-700">
                      <span class="font-medium">Quantity:</span> {item.quantity}
                    </div>
                    <div class="mt-1 text-sm text-gray-700">
                      <span class="font-medium">Subtotal:</span> {costFormat(item.amount * item.quantity)} {cart.currency || ''}
                    </div>
                  </div>
                </div>
              {/each}
            </div>
          </DetailList>
        {:else}
          <DetailList name="Items">
            <span class="text-gray-400">No items</span>
          </DetailList>
        {/if}
      </dl>
    </div>
  {:else}
    <div class="py-8 text-center text-gray-500">Failed to load cart</div>
  {/if}

  <div class="pt-5">
    <FormButton type="button" name="Close" color="green" on:click={close} />
  </div>
</div>
