<script lang="ts">
  import type { Product } from "$lib/types/models";
  import { cartStore } from "$lib/stores/cart";
  import { costFormat } from "$lib/utils/costFormat";
  import { settingsStore } from "$lib/stores/settings";
  import { getFirstImageUrl } from "$lib/utils/imageUrl";
  import { toggleCartItem } from "$lib/utils/cart";
  import { handleNavigation } from "$lib/utils/navigation";

  interface Props {
    product: Product;
    index?: number;
  }

  let { product, index = 0 }: Props = $props();

  let currency = $derived($settingsStore?.main.currency || "");
  let cart = $derived($cartStore);
  let inCart = $derived(cart.some((item) => item.id === product.id));

  function handleToggleCart(e: MouseEvent) {
    e.stopPropagation();
    toggleCartItem(product, cart);
  }
</script>

<li class="h-full flex flex-col">
  <a
    href="/products/{product.slug}"
    onclick={(e) => handleNavigation(e, `/products/${product.slug}`)}
    class="block cursor-pointer flex-shrink-0"
  >
    <div class="relative overflow-hidden bg-white border-4 border-black">
      <img
        src={getFirstImageUrl(product.images, "md")}
        alt={product.name}
        class="w-full h-64 object-cover transition-transform duration-500 hover:scale-110"
        loading="lazy"
      />
      <div class="absolute top-4 right-4">
        <div class="border-4 border-black bg-yellow-300 px-3 py-1 font-black text-xs uppercase tracking-wider">
          NEW
        </div>
      </div>
    </div>
  </a>
  
  <div class="p-6 bg-white border-x-4 border-b-4 border-black flex-1 flex flex-col">
    <div class="flex items-start justify-between gap-4 mb-4 flex-1">
      <a
        href="/products/{product.slug}"
        onclick={(e) => handleNavigation(e, `/products/${product.slug}`)}
        class="flex-1 cursor-pointer"
      >
        <h3 class="text-xl font-black uppercase tracking-tight text-black mb-2 hover:underline decoration-4 decoration-yellow-300 underline-offset-4">
          {product.name}
        </h3>
      </a>
    </div>
    
    <div class="flex items-center justify-between gap-4 mt-auto">
      <div class="flex items-baseline gap-2">
        <span class="text-3xl font-black text-black tracking-tight">
          {costFormat(product.amount)}
        </span>
        <span class="text-lg font-bold text-gray-600 uppercase">{currency}</span>
      </div>
      
      <button
        onclick={handleToggleCart}
        class="border-4 border-black px-6 py-3 font-black text-sm uppercase tracking-wider hover:shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer {inCart ? 'bg-red-500 text-white' : 'bg-green-500 text-white'}"
      >
        {#if !inCart}
          <span class="flex items-center gap-2">
            <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <use href="/assets/img/sprite.svg#plus" />
            </svg>
            <span>ADD</span>
          </span>
        {:else}
          <span class="flex items-center gap-2">
            <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <use href="/assets/img/sprite.svg#minus" />
            </svg>
            <span>REMOVE</span>
          </span>
        {/if}
      </button>
    </div>
  </div>
</li>
