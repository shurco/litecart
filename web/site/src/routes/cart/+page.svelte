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
  import type { PaymentMethods } from "$lib/types/models";
  import { goto } from "$app/navigation";
  import Overlay from "$lib/components/Overlay.svelte";

  let email = $state("");
  let provider = $state("");
  let payments = $state<PaymentMethods>({});
  let showOverlay = $state(false);
  let error = $state<string | undefined>(undefined);

  let cart = $derived($cartStore);
  let currency = $derived($settingsStore?.main.currency || "");

  onMount(async () => {
    if (typeof window !== "undefined") {
      email = localStorage.getItem("email") || "";
      provider = localStorage.getItem("provider") || "";
    }

    const res = await apiGet<PaymentMethods>("/api/cart/payment");
    if (res.success && res.result) {
      payments = res.result;

      // Auto-select provider if only one is available
      const autoProvider = autoSelectProvider(payments);
      if (autoProvider) {
        provider = autoProvider;
        localStorage.setItem("provider", provider);
      } else if (!hasPaymentProviders(payments)) {
        localStorage.removeItem("provider");
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

    if (typeof window !== "undefined") {
      localStorage.setItem("email", email);
      localStorage.setItem("provider", provider);
    }

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

<section>
  <div class="mx-auto max-w-screen-xl px-4 py-8 sm:px-6 sm:py-12 lg:px-8">
    <div class="mx-auto max-w-3xl">
      <header class="text-center">
        <h1 class="text-xl font-bold text-gray-900 sm:text-3xl">
          {cart.length > 0 ? "Your Cart" : "Your Cart is empty"}
        </h1>
      </header>

      <form onsubmit={checkOut}>
        {#if cart.length > 0}
          <div class="mt-8">
            <ul class="space-y-4">
              {#each cart as item}
                <li class="flex items-center gap-4">
                  <img
                    src={getProductImageUrl(item.image, "sm")}
                    alt={item.name}
                    class="h-16 w-16 rounded object-cover"
                  />
                  <div>
                    <a href="/products/{item.slug}" target="_blank">{item.name}</a>
                  </div>
                  <div class="flex flex-1 items-center justify-end gap-2">
                    {costFormat(item.amount)} {currency}
                    <button
                      class="text-gray-600 transition hover:text-red-600 cursor-pointer"
                      onclick={() => cartStore.remove(item.id)}
                    >
                      <span class="sr-only">Remove item</span>
                      <svg class="h-4 w-4">
                        <use href="/assets/img/sprite.svg#trash" />
                      </svg>
                    </button>
                  </div>
                </li>
              {/each}
            </ul>
            <div class="mt-8 flex justify-end border-t border-gray-100 pt-8">
              <div class="w-screen max-w-lg">
                <dl class="space-y-0.5 text-sm text-gray-700">
                  <div class="flex justify-between !text-base">
                    <dt>Total</dt>
                    <dd>{totalCartAmount()} {currency}</dd>
                  </div>
                </dl>
              </div>
            </div>

            {#if showPayments()}
              <div class="mt-8 border-t border-gray-100 pt-8">
                <div class="mx-auto max-w-xl text-center">
                  <p class="mt-4 text-gray-400">
                    To continue, you need to enter the email address to which the item
                    will be sent after payment.
                    {#if showSelectPayments()}
                      Also, choose the payment system through which the payment will be
                      made.
                    {/if}
                  </p>
                </div>
              </div>

              <div class="mt-8 border-t border-gray-100 pt-8">
                <div class="text-center">
                  <p class="mb-5 text-lg font-bold text-gray-500 sm:text-3xl"
                    >Enter email</p
                  >
                </div>
                <div class="flex place-content-center">
                  <label
                    for="email"
                    class="min-w-[50%] relative block rounded-md border-2 focus-within:border-blue-500 focus-within:ring-1 focus-within:ring-blue-500"
                    class:border-blue-500={!!email}
                    class:ring-blue-500={!!email}
                    class:ring-1={!!email}
                    class:bg-blue-100={!!email}
                    class:border-gray-200={!email}
                  >
                    <input
                      type="email"
                      bind:value={email}
                      id="email"
                      class="min-w-full peer border-none bg-transparent placeholder-transparent focus:border-transparent focus:outline-hidden focus:ring-0"
                      placeholder="Email"
                      required
                    />
                    <span
                      class="rounded pointer-events-none absolute start-2.5 top-0 -translate-y-1/2 bg-blue-500 py-0.5 px-1 text-xs text-white transition-all peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-placeholder-shown:bg-white peer-placeholder-shown:text-gray-700 peer-focus:top-0 peer-focus:text-xs"
                    >
                      Email
                    </span>
                  </label>
                </div>
              </div>

              {#if showSelectPayments()}
                <div class="mt-8 border-t border-gray-100 pt-8">
                  <div class="text-center">
                    <p class="mb-5 text-lg font-bold text-gray-500 sm:text-3xl"
                      >Select payment system</p
                    >
                  </div>
                  <div class="flex place-content-center">
                    <fieldset class="space-y-4 min-w-[50%]">
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
                            class="flex cursor-pointer items-center rounded-lg border-2 border-gray-100 bg-white p-4 hover:border-gray-200 peer-checked:border-blue-500 peer-checked:ring-1 peer-checked:bg-blue-100 peer-checked:ring-blue-500"
                          >
                            <dl class="flex flex-col">
                              <p class="text-gray-700 text-sm font-medium">Stripe</p>
                              <p class="text-gray-400 text-xs"
                                >A popular payment system that allows payments with cards
                                and<br />other widely available methods</p
                              >
                            </dl>
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
                            class="flex cursor-pointer items-center rounded-lg border-2 border-gray-100 bg-white p-4 hover:border-gray-200 peer-checked:border-blue-500 peer-checked:ring-1 peer-checked:bg-blue-100 peer-checked:ring-blue-500"
                          >
                            <dl class="flex flex-col">
                              <p class="text-gray-700 text-sm font-medium">Paypal</p>
                              <p class="text-gray-400 text-xs"
                                >A popular payment system that allows payments with cards
                                and<br />funds from a PayPal account.</p
                              >
                            </dl>
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
                            class="flex cursor-pointer items-center rounded-lg border-2 border-gray-100 bg-white p-4 hover:border-gray-200 peer-checked:border-blue-500 peer-checked:ring-1 peer-checked:bg-blue-100 peer-checked:ring-blue-500"
                          >
                            <dl class="flex flex-col">
                              <p class="text-gray-700 text-sm font-medium"
                                >Spectrocoin</p
                              >
                              <p class="text-gray-400 text-xs"
                                >Payment system allowing to pay bills with
                                cryptocurrency</p
                              >
                            </dl>
                          </label>
                        </div>
                      {/if}
                    </fieldset>
                  </div>
                </div>
              {/if}

              <div class="mt-8 flex justify-end border-t border-gray-100 pt-8">
                <div class="w-screen max-w-lg space-y-4 flex justify-end">
                  <input
                    type="submit"
                    value="Checkout"
                    disabled={!email || (showSelectPayments() && !provider)}
                    class="disabled:opacity-25 disabled:bg-gray-400 cursor-pointer block rounded bg-gray-700 px-5 py-3 text-sm text-gray-100 transition hover:bg-gray-600"
                  />
                </div>
              </div>
            {:else}
              <div class="mt-8 border-t border-gray-100 pt-8">
                <div class="mx-auto max-w-xl text-center">
                  <p class="mt-4 text-red-400">
                    To continue with the payment, the administrator must activate at least
                    one payment system.
                  </p>
                </div>
              </div>
            {/if}
          </div>
        {/if}
      </form>
    </div>
  </div>
  <Overlay show={showOverlay} {error} onClose={closeOverlay} />
</section>
