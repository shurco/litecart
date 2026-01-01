<script lang="ts">
  import { onMount } from 'svelte'
  import { cartStore } from '$lib/stores/cart'
  import { settingsStore } from '$lib/stores/settings'
  import { apiGet, apiPost } from '$lib/utils/api'
  import { costFormat } from '$lib/utils/costFormat'
  import { getProductImageUrl } from '$lib/utils/imageUrl'
  import { hasPaymentProviders, autoSelectProvider, getAvailableProviders } from '$lib/utils/payment'
  import { getLocalStorage, setLocalStorage, removeLocalStorage } from '$lib/utils/browser'
  import type { PaymentMethods } from '$lib/types/models'
  import { goto } from '$app/navigation'
  import Overlay from '$lib/components/Overlay.svelte'
  import { handleNavigation } from '$lib/utils/navigation'

  let email = $state('')
  let provider = $state('')
  let payments = $state<PaymentMethods>({})
  let showOverlay = $state(false)
  let error = $state<string | undefined>(undefined)
  let isLoadingPaymentMethods = $state(false)

  let cart = $derived($cartStore)
  let currency = $derived($settingsStore?.main.currency || '')

  // Calculate total cart amount in cents
  let cartTotal = $derived(cart.reduce((sum, item) => sum + item.amount, 0))
  let isFree = $derived(cartTotal === 0)

  // Handle payment provider based on cart state
  $effect(() => {
    if (isFree) {
      // For free carts, don't auto-set provider to prevent accidental checkout
      // Provider will be set only when user explicitly clicks checkout button
      // Clear any existing provider selection when cart becomes free
      if (provider && provider !== 'dummy') {
        provider = ''
        removeLocalStorage('provider')
      }
    } else if (!isFree) {
      // If cart is no longer free, reset provider and load payment methods
      if (provider === 'dummy') {
        provider = ''
        removeLocalStorage('provider')
      }
      // Load payment methods if not already loaded and not currently loading
      if (!hasPaymentProviders(payments) && !isLoadingPaymentMethods) {
        loadPaymentMethods().catch((err) => {
          console.error('Failed to load payment methods:', err)
          error = 'Failed to load payment methods. Please refresh the page.'
          showOverlay = true
        })
      }
    }
  })

  async function loadPaymentMethods() {
    // Prevent multiple simultaneous calls
    if (isLoadingPaymentMethods) {
      return
    }

    isLoadingPaymentMethods = true
    try {
      const res = await apiGet<PaymentMethods>('/api/cart/payment')
      if (res.success && res.result) {
        payments = res.result

        // Don't auto-select provider - user must explicitly choose
        if (!hasPaymentProviders(payments)) {
          removeLocalStorage('provider')
          provider = ''
        } else {
          // Reset provider - user must choose explicitly
          provider = ''
          removeLocalStorage('provider')
        }
      } else {
        throw new Error(res.message || 'Failed to load payment methods')
      }
    } finally {
      isLoadingPaymentMethods = false
    }
  }

  onMount(async () => {
    email = getLocalStorage('email')

    // If cart is not free, load payment methods
    // $effect will also handle this, but we load here on initial mount to avoid delay
    if (!isFree && !hasPaymentProviders(payments)) {
      await loadPaymentMethods().catch((err) => {
        console.error('Failed to load payment methods on mount:', err)
        error = 'Failed to load payment methods. Please refresh the page.'
        showOverlay = true
      })
    }
    // Don't auto-set provider for free carts on mount to prevent accidental checkout
  })

  function showPayments(): boolean {
    // Hide payment systems if cart is free (will use dummy automatically)
    if (isFree) {
      return false
    }
    return hasPaymentProviders(payments)
  }

  function showSelectPayments(): boolean {
    // Never show payment selection for free carts
    if (isFree) {
      return false
    }
    // Always show payment selection for paid carts (even if only one option)
    return hasPaymentProviders(payments)
  }

  function totalCartAmount(): string {
    return costFormat(cartTotal)
  }

  async function checkOut(e: Event) {
    e.preventDefault()

    setLocalStorage('email', email)

    error = undefined

    // Ensure provider is set correctly based on cart total
    let finalProvider = provider
    if (isFree) {
      // Force dummy provider for free carts
      finalProvider = 'dummy'
      setLocalStorage('provider', finalProvider)
    } else if (!finalProvider) {
      // If cart is paid but no provider selected, show error
      error = 'Please select a payment system'
      showOverlay = true
      return
    } else {
      setLocalStorage('provider', finalProvider)
    }

    const cartData = {
      email,
      provider: finalProvider,
      products: cart.map((item) => ({ id: item.id, quantity: 1 }))
    }

    const res = await apiPost<{ url?: string }>('/cart/payment', cartData)
    if (res.success && res.result?.url) {
      window.location.href = res.result.url
    } else {
      error = res.message || 'Payment failed'
      showOverlay = true
    }
  }

  function closeOverlay() {
    showOverlay = false
    error = undefined
  }
</script>

<section class="min-h-screen bg-white px-4 py-12 sm:px-6 lg:px-8">
  <div class="mx-auto max-w-screen-xl">
    <div class="mx-auto max-w-4xl">
      <!-- Header -->
      <header class="mb-12 text-center">
        <h1 class="mb-4 text-4xl font-black tracking-tighter text-black uppercase sm:text-5xl">
          {cart.length > 0 ? 'YOUR CART' : 'CART IS EMPTY'}
        </h1>
        <div class="mx-auto h-1 w-32 bg-black"></div>
      </header>

      {#if cart.length === 0}
        <div class="brutal-card mb-8 p-8 text-center">
          <p class="mb-8 text-xl font-bold tracking-wide text-black uppercase">
            Your cart is empty. Add some products to continue shopping.
          </p>

          <div class="flex justify-center">
            <a
              href="/"
              onclick={(e) => handleNavigation(e, '/')}
              class="inline-block cursor-pointer border-4 border-black bg-yellow-300 px-8 py-4 text-lg font-black tracking-wider text-black uppercase transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)]"
            >
              GO TO HOME
            </a>
          </div>
        </div>
      {/if}

      <form onsubmit={checkOut}>
        {#if cart.length > 0}
          <!-- Cart Items -->
          <div class="mb-8">
            <h2 class="mb-6 text-3xl font-black tracking-tighter text-black uppercase">
              ITEMS ({cart.length})
            </h2>
            <ul class="space-y-4">
              {#each cart as item}
                <li class="border-4 border-black bg-white p-4">
                  <div class="flex items-center gap-4">
                    <div class="overflow-hidden border-4 border-black">
                      <img src={getProductImageUrl(item.image, 'sm')} alt={item.name} class="h-20 w-20 object-cover" />
                    </div>
                    <div class="flex-1">
                      <a
                        href="/products/{item.slug}"
                        target="_blank"
                        class="cursor-pointer text-xl font-black tracking-tight text-black uppercase decoration-yellow-300 decoration-4 underline-offset-4 hover:underline"
                      >
                        {item.name}
                      </a>
                    </div>
                    <div class="flex items-center gap-4">
                      <span class="text-2xl font-black {costFormat(item.amount) === 'free' ? 'text-green-500' : 'text-black'}">
                        {costFormat(item.amount)}
                        {#if item.amount !== 0 && item.amount}
                          {currency}
                        {/if}
                      </span>
                      <button
                        type="button"
                        class="cursor-pointer border-4 border-black bg-red-500 p-2 text-sm font-black text-white uppercase transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[6px_6px_0px_0px_rgba(0,0,0,1)]"
                        onclick={() => cartStore.remove(item.id)}
                        aria-label="Remove item"
                      >
                        <svg class="h-5 w-5">
                          <use href="/assets/img/sprite.svg#trash" />
                        </svg>
                      </button>
                    </div>
                  </div>
                </li>
              {/each}
            </ul>
          </div>

          <!-- Total -->
          <div class="brutal-card mb-8 bg-yellow-300 p-8">
            <div class="flex items-center justify-between">
              <span class="text-3xl font-black tracking-tighter text-black uppercase"> TOTAL </span>
              <span class="text-4xl font-black {totalCartAmount() === 'free' ? 'text-green-500' : 'text-black'}">
                {totalCartAmount()}
                {#if cart.length > 0 && cartTotal !== 0}
                  {currency}
                {/if}
              </span>
            </div>
          </div>

          {#if isFree || showPayments()}
            <!-- Email Input -->
            <div class="mt-16 mb-8">
              <h2 class="mb-6 text-3xl font-black tracking-tighter text-black uppercase">ENTER EMAIL</h2>
              <p class="mb-4 text-sm font-bold tracking-wide text-black uppercase">
                {#if isFree}
                  Enter the email address to which the item will be sent.
                {:else}
                  Enter the email address to which the item will be sent after payment. Also, choose the payment system.
                {/if}
              </p>
              <label for="email" class="block">
                <input
                  type="email"
                  bind:value={email}
                  id="email"
                  required
                  class="w-full border-4 border-black bg-white px-6 py-4 text-lg font-black tracking-wider text-black uppercase focus:ring-4 focus:ring-yellow-300 focus:outline-none"
                  placeholder="EMAIL@EXAMPLE.COM"
                />
              </label>
            </div>

            <!-- Payment Provider Selection -->
            {#if showSelectPayments()}
              <div class="mt-16 mb-8">
                <h2 class="mb-6 text-3xl font-black tracking-tighter text-black uppercase">SELECT PAYMENT SYSTEM</h2>
                <fieldset class="space-y-4">
                  {#if payments.stripe}
                    <div>
                      <input type="radio" bind:group={provider} value="stripe" id="stripe" class="peer hidden" />
                      <label
                        for="stripe"
                        class="block cursor-pointer border-4 border-black bg-white p-6 peer-checked:border-yellow-300 peer-checked:bg-yellow-300"
                      >
                        <p class="mb-2 text-xl font-black tracking-tight text-black uppercase">Stripe</p>
                        <p class="text-sm font-bold text-black">Popular payment system for cards and other methods</p>
                      </label>
                    </div>
                  {/if}

                  {#if payments.paypal}
                    <div>
                      <input type="radio" bind:group={provider} value="paypal" id="paypal" class="peer hidden" />
                      <label
                        for="paypal"
                        class="block cursor-pointer border-4 border-black bg-white p-6 peer-checked:border-yellow-300 peer-checked:bg-yellow-300"
                      >
                        <p class="mb-2 text-xl font-black tracking-tight text-black uppercase">PayPal</p>
                        <p class="text-sm font-bold text-black">Payment system for cards and PayPal account funds</p>
                      </label>
                    </div>
                  {/if}

                  {#if payments.spectrocoin}
                    <div>
                      <input
                        type="radio"
                        bind:group={provider}
                        value="spectrocoin"
                        id="spectrocoin"
                        class="peer hidden"
                      />
                      <label
                        for="spectrocoin"
                        class="block cursor-pointer border-4 border-black bg-white p-6 peer-checked:border-yellow-300 peer-checked:bg-yellow-300"
                      >
                        <p class="mb-2 text-xl font-black tracking-tight text-black uppercase">Spectrocoin</p>
                        <p class="text-sm font-bold text-black">Payment system allowing payments with cryptocurrency</p>
                      </label>
                    </div>
                  {/if}
                </fieldset>
              </div>
            {/if}

            <!-- Checkout Button -->
            <div class="flex justify-end">
              <button
                type="submit"
                disabled={!email || (!isFree && !provider)}
                class="cursor-pointer border-4 border-black bg-green-500 px-12 py-4 text-xl font-black tracking-wider text-white uppercase transition-all duration-200 enabled:hover:-translate-x-1 enabled:hover:-translate-y-1 enabled:hover:shadow-[14px_14px_0px_0px_rgba(0,0,0,1)] disabled:cursor-not-allowed disabled:opacity-50"
              >
                {#if isFree}
                  GET FOR FREE
                {:else}
                  CHECKOUT
                {/if}
              </button>
            </div>
          {:else}
            <div class="brutal-card bg-red-300 p-8">
              <p class="text-center text-xl font-black tracking-wider text-black uppercase">
                NO PAYMENT SYSTEMS AVAILABLE. CONTACT ADMINISTRATOR.
              </p>
            </div>
          {/if}
        {/if}
      </form>
    </div>
  </div>
  <Overlay show={showOverlay} {error} onClose={closeOverlay} />
</section>
