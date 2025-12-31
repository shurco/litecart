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

<div class="min-h-screen bg-white py-12 px-4 sm:px-6 lg:px-8">
  <div class="max-w-screen-xl mx-auto">
    <div class="mx-auto max-w-3xl">
      {#if loading}
        <div class="brutal-card p-12 text-center">
          <div class="border-4 border-black bg-yellow-300 px-8 py-6 inline-block">
            <p class="text-2xl font-black uppercase tracking-wider text-black">
              LOADING...
            </p>
          </div>
        </div>
      {:else if error}
        <div class="brutal-card p-12 bg-red-300">
          <h1 class="text-4xl font-black uppercase tracking-tighter text-black mb-4">
            ERROR
          </h1>
          <p class="text-xl font-bold text-black">{error}</p>
        </div>
      {:else if cart}
        <div class="brutal-card p-8 sm:p-12 mb-8 bg-green-300">
          <header class="text-center">
            <h1 class="text-4xl sm:text-5xl font-black uppercase tracking-tighter text-black mb-4">
              PAYMENT SUCCESSFUL!
            </h1>
            <p class="text-xl font-bold uppercase tracking-wider text-black">
              Thank you for your purchase. Your order has been processed successfully.
            </p>
          </header>
        </div>

        <div class="brutal-card p-8 sm:p-12">
          <h2 class="text-3xl font-black uppercase tracking-tighter text-black mb-6 border-b-4 border-black pb-4">
            ORDER DETAILS
          </h2>
          <ul class="space-y-4 mb-8">
            {#each cart.items as item}
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
                      class="text-xl font-black uppercase tracking-tight text-black hover:underline decoration-4 decoration-yellow-300 underline-offset-4 cursor-pointer"
                    >
                      {item.name}
                    </a>
                    {#if item.quantity > 1}
                      <p class="text-sm font-bold text-gray-700 mt-1">Quantity: {item.quantity}</p>
                    {/if}
                  </div>
                  <div class="text-right">
                    <p class="text-2xl font-black text-black">
                      {costFormat(item.amount * item.quantity)} {cart.currency}
                    </p>
                    {#if item.quantity > 1}
                      <p class="text-sm font-bold text-gray-600">{costFormat(item.amount)} each</p>
                    {/if}
                  </div>
                </div>
              </li>
            {/each}
          </ul>

          <div class="border-t-4 border-black pt-6">
            <div class="flex justify-between items-center">
              <span class="text-3xl font-black uppercase tracking-tighter text-black">
                TOTAL
              </span>
              <span class="text-4xl font-black text-black">
                {totalAmount()} {cart.currency}
              </span>
            </div>
          </div>

          <div class="mt-8 text-center">
            <div class="border-4 border-black bg-yellow-300 px-6 py-4 inline-block">
              <p class="text-lg font-black uppercase tracking-wider text-black">
                Redirecting to home page in a few seconds...
              </p>
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
