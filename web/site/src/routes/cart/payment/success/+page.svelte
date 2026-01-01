<script lang="ts">
  import { onMount } from 'svelte'
  import { page } from '$app/stores'
  import { cartStore } from '$lib/stores/cart'
  import { settingsStore } from '$lib/stores/settings'
  import { apiGet } from '$lib/utils/api'
  import { costFormat } from '$lib/utils/costFormat'
  import { getProductImageUrl } from '$lib/utils/imageUrl'
  import { goto } from '$app/navigation'

  interface CartItem {
    id: string
    name: string
    slug: string
    amount: number
    quantity: number
    image?: { name: string; ext: string } | null
  }

  interface CartData {
    id: string
    email: string
    amount_total: number
    currency: string
    payment_status: string
    payment_system: string
    items: CartItem[]
  }

  let cart = $state<CartData | null>(null)
  let loading = $state(true)
  let error = $state<string | undefined>(undefined)
  let currency = $derived($settingsStore?.main.currency || '')

  onMount(async () => {
    // Clear local cart
    cartStore.clear()

    // Get cart_id from query parameters
    const cartId = $page.url.searchParams.get('cart_id')
    if (!cartId) {
      error = 'Cart ID is missing'
      loading = false
      return
    }

    // Load cart information from API
    const res = await apiGet<CartData>(`/api/cart/${cartId}`)
    if (res.success && res.result) {
      cart = res.result
    } else {
      error = res.message || 'Failed to load cart information'
    }
    loading = false

    // Redirect to home page after 5 seconds
    setTimeout(() => {
      goto('/')
    }, 5000)
  })

  function totalAmount(): string {
    if (!cart) return '0'
    return costFormat(cart.amount_total)
  }
</script>

<div class="min-h-screen bg-white px-4 py-12 sm:px-6 lg:px-8">
  <div class="mx-auto max-w-screen-xl">
    <div class="mx-auto max-w-3xl">
      {#if loading}
        <div class="brutal-card p-12 text-center">
          <div class="inline-block border-4 border-black bg-yellow-300 px-8 py-6">
            <p class="text-2xl font-black tracking-wider text-black uppercase">LOADING...</p>
          </div>
        </div>
      {:else if error}
        <div class="brutal-card bg-red-300 p-12">
          <h1 class="mb-4 text-4xl font-black tracking-tighter text-black uppercase">ERROR</h1>
          <p class="text-xl font-bold text-black">{error}</p>
        </div>
      {:else if cart}
        <div class="brutal-card mb-8 bg-green-300 p-8 sm:p-12">
          <header class="text-center">
            <h1 class="mb-4 text-4xl font-black tracking-tighter text-black uppercase sm:text-5xl">
              PAYMENT SUCCESSFUL!
            </h1>
            <p class="text-xl font-bold tracking-wider text-black uppercase">
              Thank you for your purchase. Your order has been processed successfully.
            </p>
          </header>
        </div>

        <div class="brutal-card p-8 sm:p-12">
          <h2 class="mb-6 border-b-4 border-black pb-4 text-3xl font-black tracking-tighter text-black uppercase">
            ORDER DETAILS
          </h2>
          <ul class="mb-8 space-y-4">
            {#each cart.items as item}
              <li class="border-4 border-black bg-white p-4">
                <div class="flex items-center gap-4">
                  <div class="overflow-hidden border-4 border-black">
                    <img src={getProductImageUrl(item.image, 'sm')} alt={item.name} class="h-20 w-20 object-cover" />
                  </div>
                  <div class="flex-1">
                    <a
                      href="/products/{item.slug}"
                      class="cursor-pointer text-xl font-black tracking-tight text-black uppercase decoration-yellow-300 decoration-4 underline-offset-4 hover:underline"
                    >
                      {item.name}
                    </a>
                    {#if item.quantity > 1}
                      <p class="mt-1 text-sm font-bold text-gray-700">Quantity: {item.quantity}</p>
                    {/if}
                  </div>
                  <div class="text-right">
                    <p class="text-2xl font-black text-black">
                      {costFormat(item.amount * item.quantity)}
                      {#if item.amount !== 0 && item.amount}
                        {cart.currency}
                      {/if}
                    </p>
                    {#if item.quantity > 1}
                      <p class="text-sm font-bold text-gray-600">
                        {costFormat(item.amount)}
                        {#if item.amount !== 0 && item.amount}
                          {' each'}
                        {/if}
                      </p>
                    {/if}
                  </div>
                </div>
              </li>
            {/each}
          </ul>

          <div class="border-t-4 border-black pt-6">
            <div class="flex items-center justify-between">
              <span class="text-3xl font-black tracking-tighter text-black uppercase"> TOTAL </span>
              <span class="text-4xl font-black text-black">
                {totalAmount()}
                {#if cart && cart.amount_total !== 0}
                  {cart.currency}
                {/if}
              </span>
            </div>
          </div>

          <div class="mt-8 text-center">
            <div class="inline-block border-4 border-black bg-yellow-300 px-6 py-4">
              <p class="text-lg font-black tracking-wider text-black uppercase">
                Redirecting to home page in a few seconds...
              </p>
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
