<script lang="ts">
  import { onMount } from "svelte";
  import { cartStore } from "$lib/stores/cart";
  import { settingsStore } from "$lib/stores/settings";
  import { apiGet, apiPost } from "$lib/utils/api";
  import { costFormat } from "$lib/utils/costFormat";
  import { getProductImageUrl } from "$lib/utils/imageUrl";
  import {
    hasPaymentProviders,
    autoSelectProvider,
    getAvailableProviders,
  } from "$lib/utils/payment";
  import { getLocalStorage, setLocalStorage, removeLocalStorage } from "$lib/utils/browser";
  import type { PaymentMethods } from "$lib/types/models";
  import { goto } from "$app/navigation";
  import Overlay from "$lib/components/Overlay.svelte";
  import { handleNavigation } from "$lib/utils/navigation";

  let email = $state("");
  let provider = $state("");
  let payments = $state<PaymentMethods>({});
  let showOverlay = $state(false);
  let error = $state<string | undefined>(undefined);

  let cart = $derived($cartStore);
  let currency = $derived($settingsStore?.main.currency || "");

  onMount(async () => {
    email = getLocalStorage("email");
    provider = getLocalStorage("provider");

    const res = await apiGet<PaymentMethods>("/api/cart/payment");
    if (res.success && res.result) {
      payments = res.result;

      // Auto-select provider if only one is available
      const autoProvider = autoSelectProvider(payments);
      if (autoProvider) {
        provider = autoProvider;
        setLocalStorage("provider", provider);
      } else if (!hasPaymentProviders(payments)) {
        removeLocalStorage("provider");
        provider = "";
      }
    }
  });

  function showPayments(): boolean {
    return hasPaymentProviders(payments);
  }

  function showSelectPayments(): boolean {
    return getAvailableProviders(payments).length > 1;
  }

  function totalCartAmount(): string {
    const total = cart.reduce((sum, item) => sum + item.amount, 0);
    return costFormat(total);
  }

  async function checkOut(e: Event) {
    e.preventDefault();

    setLocalStorage("email", email);
    setLocalStorage("provider", provider);

    error = undefined;

    const cartData = {
      email,
      provider,
      products: cart.map((item) => ({ id: item.id, quantity: 1 })),
    };

    const res = await apiPost<{ url?: string }>("/cart/payment", cartData);
    if (res.success && res.result?.url) {
      window.location.href = res.result.url;
    } else {
      error = res.message || "Payment failed";
      showOverlay = true;
    }
  }

  function closeOverlay() {
    showOverlay = false;
    error = undefined;
  }
</script>

<section class="min-h-screen bg-white py-12 px-4 sm:px-6 lg:px-8">
  <div class="max-w-screen-xl mx-auto">
    <div class="mx-auto max-w-4xl">
      <!-- Header -->
      <header class="text-center mb-12">
        <h1 class="text-4xl sm:text-5xl font-black uppercase tracking-tighter text-black mb-4">
          {cart.length > 0 ? "YOUR CART" : "CART IS EMPTY"}
        </h1>
        <div class="w-32 h-1 bg-black mx-auto"></div>
      </header>

      {#if cart.length === 0}
        <div class="brutal-card p-8 mb-8 text-center">
          <p class="text-xl font-bold uppercase tracking-wide text-black mb-8">
            Your cart is empty. Add some products to continue shopping.
          </p>

          <div class="flex justify-center">
            <a
              href="/"
              onclick={(e) => handleNavigation(e, "/")}
              class="inline-block border-4 border-black bg-yellow-300 text-black px-8 py-4 font-black text-lg uppercase tracking-wider hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
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
            <h2 class="text-3xl font-black uppercase tracking-tighter text-black mb-6">
              ITEMS ({cart.length})
            </h2>
            <ul class="space-y-4">
              {#each cart as item}
                <li class="border-4 border-black bg-white p-4">
                  <div class="flex items-center gap-4">
                    <div class="border-4 border-black overflow-hidden">
                      <img
                        src={getProductImageUrl(item.image, "sm")}
                        alt={item.name}
                        class="h-20 w-20 object-cover"
                      />
                    </div>
                    <div class="flex-1">
                      <a 
                        href="/products/{item.slug}" 
                        target="_blank"
                        class="text-xl font-black uppercase tracking-tight text-black hover:underline decoration-4 decoration-yellow-300 underline-offset-4 cursor-pointer"
                      >
                        {item.name}
                      </a>
                    </div>
                    <div class="flex items-center gap-4">
                      <span class="text-2xl font-black text-black">
                        {costFormat(item.amount)} {currency}
                      </span>
                      <button
                        class="border-4 border-black bg-red-500 text-white p-2 font-black text-sm uppercase hover:shadow-[6px_6px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
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
          <div class="brutal-card p-8 mb-8 bg-yellow-300">
            <div class="flex justify-between items-center">
              <span class="text-3xl font-black uppercase tracking-tighter text-black">
                TOTAL
              </span>
              <span class="text-4xl font-black text-black">
                {totalCartAmount()} {currency}
              </span>
            </div>
          </div>

          {#if showPayments()}
            <!-- Email Input -->
            <div class="mb-8 mt-16">
              <h2 class="text-3xl font-black uppercase tracking-tighter text-black mb-6">
                ENTER EMAIL
              </h2>
              <p class="text-sm font-bold uppercase tracking-wide text-black mb-4">
                Enter the email address to which the item will be sent after payment.
                {#if showSelectPayments()}
                  Also, choose the payment system.
                {/if}
              </p>
              <label for="email" class="block">
                <input
                  type="email"
                  bind:value={email}
                  id="email"
                  required
                  class="w-full border-4 border-black bg-white px-6 py-4 font-black text-lg uppercase tracking-wider text-black focus:outline-none focus:ring-4 focus:ring-yellow-300"
                  placeholder="EMAIL@EXAMPLE.COM"
                />
              </label>
            </div>

            <!-- Payment Provider Selection -->
            {#if showSelectPayments()}
              <div class="mb-8 mt-16">
                <h2 class="text-3xl font-black uppercase tracking-tighter text-black mb-6">
                  SELECT PAYMENT SYSTEM
                </h2>
                <fieldset class="space-y-4">
                  {#if payments.stripe}
                    <div>
                      <input
                        type="radio"
                        bind:group={provider}
                        value="stripe"
                        id="stripe"
                        class="peer hidden"
                      />
                      <label
                        for="stripe"
                        class="block border-4 border-black bg-white p-6 cursor-pointer peer-checked:bg-yellow-300 peer-checked:border-yellow-300"
                      >
                        <p class="text-xl font-black uppercase tracking-tight text-black mb-2">Stripe</p>
                        <p class="text-sm font-bold text-black">
                          Popular payment system for cards and other methods
                        </p>
                      </label>
                    </div>
                  {/if}

                  {#if payments.paypal}
                    <div>
                      <input
                        type="radio"
                        bind:group={provider}
                        value="paypal"
                        id="paypal"
                        class="peer hidden"
                      />
                      <label
                        for="paypal"
                        class="block border-4 border-black bg-white p-6 cursor-pointer peer-checked:bg-yellow-300 peer-checked:border-yellow-300"
                      >
                        <p class="text-xl font-black uppercase tracking-tight text-black mb-2">PayPal</p>
                        <p class="text-sm font-bold text-black">
                          Payment system for cards and PayPal account funds
                        </p>
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
                        class="block border-4 border-black bg-white p-6 cursor-pointer peer-checked:bg-yellow-300 peer-checked:border-yellow-300"
                      >
                        <p class="text-xl font-black uppercase tracking-tight text-black mb-2">Spectrocoin</p>
                        <p class="text-sm font-bold text-black">
                          Payment system allowing payments with cryptocurrency
                        </p>
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
                disabled={!email || (showSelectPayments() && !provider)}
                class="border-4 border-black bg-green-500 text-white px-12 py-4 font-black text-xl uppercase tracking-wider transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer enabled:hover:shadow-[14px_14px_0px_0px_rgba(0,0,0,1)] enabled:hover:-translate-x-1 enabled:hover:-translate-y-1"
              >
                CHECKOUT
              </button>
            </div>
          {:else}
            <div class="brutal-card p-8 bg-red-300">
              <p class="text-xl font-black uppercase tracking-wider text-black text-center">
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
