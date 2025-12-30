<script lang="ts">
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import { cartStore } from "$lib/stores/cart";
  import { settingsStore } from "$lib/stores/settings";
  import { apiGet } from "$lib/utils/api";
  import { costFormat } from "$lib/utils/costFormat";
  import { getProductImageUrl } from "$lib/utils/imageUrl";
  import { goto } from "$app/navigation";

  interface CartItem {
    id: string;
    name: string;
    slug: string;
    amount: number;
    quantity: number;
    image?: { name: string; ext: string } | null;
  }

  interface CartData {
    id: string;
    email: string;
    amount_total: number;
    currency: string;
    payment_status: string;
    payment_system: string;
    items: CartItem[];
  }

  let cart = $state<CartData | null>(null);
  let loading = $state(true);
  let error = $state<string | undefined>(undefined);
  let currency = $derived($settingsStore?.main.currency || "");

  onMount(async () => {
    // Clear local cart
    cartStore.clear();

    // Get cart_id from query parameters
    const cartId = $page.url.searchParams.get("cart_id");
    if (!cartId) {
      error = "Cart ID is missing";
      loading = false;
      return;
    }

    // Load cart information from API
    const res = await apiGet<CartData>(`/api/cart/${cartId}`);
    if (res.success && res.result) {
      cart = res.result;
    } else {
      error = res.message || "Failed to load cart information";
    }
    loading = false;

    // Redirect to home page after 5 seconds
    setTimeout(() => {
      goto("/");
    }, 5000);
  });

  function totalAmount(): string {
    if (!cart) return "0";
    return costFormat(cart.amount_total);
  }
</script>

<div class="mx-auto max-w-screen-xl px-4 py-8 sm:px-6 sm:py-12 lg:px-8">
  <div class="mx-auto max-w-3xl">
    {#if loading}
      <div class="grid h-screen px-4 bg-white place-content-center">
        <p class="text-gray-600">Loading...</p>
      </div>
    {:else if error}
      <div class="grid h-screen px-4 bg-white place-content-center">
        <h1 class="text-2xl font-bold text-red-600">Error</h1>
        <p class="mt-4 text-gray-600">{error}</p>
      </div>
    {:else if cart}
      <header class="text-center">
        <h1 class="text-2xl font-bold text-green-600 sm:text-3xl">
          Payment Successful!
        </h1>
        <p class="mt-4 text-gray-600">
          Thank you for your purchase. Your order has been processed successfully.
        </p>
      </header>

      <div class="mt-8">
        <h2 class="text-xl font-bold text-gray-900 mb-4">Order Details</h2>
        <ul class="space-y-4">
          {#each cart.items as item}
            <li class="flex items-center gap-4">
              <img
                src={getProductImageUrl(item.image, "sm")}
                alt={item.name}
                class="h-16 w-16 rounded object-cover"
              />
              <div class="flex-1">
                <a href="/products/{item.slug}" class="text-gray-900 hover:text-blue-600">
                  {item.name}
                </a>
                {#if item.quantity > 1}
                  <p class="text-sm text-gray-500">Quantity: {item.quantity}</p>
                {/if}
              </div>
              <div class="text-right">
                <p class="font-semibold">{costFormat(item.amount * item.quantity)} {cart.currency}</p>
                {#if item.quantity > 1}
                  <p class="text-sm text-gray-500">{costFormat(item.amount)} each</p>
                {/if}
              </div>
            </li>
          {/each}
        </ul>

        <div class="mt-8 flex justify-end border-t border-gray-100 pt-8">
          <div class="w-screen max-w-lg">
            <dl class="space-y-0.5 text-sm text-gray-700">
              <div class="flex justify-between !text-base font-bold">
                <dt>Total</dt>
                <dd>{totalAmount()} {cart.currency}</dd>
              </div>
            </dl>
          </div>
        </div>

        <div class="mt-8 text-center">
          <p class="text-gray-600">
            Redirecting to home page in a few seconds...
          </p>
        </div>
      </div>
    {/if}
  </div>
</div>
