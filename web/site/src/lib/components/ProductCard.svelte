<script lang="ts">
  import type { Product } from "$lib/types/models";
  import { cartStore } from "$lib/stores/cart";
  import { costFormat } from "$lib/utils/costFormat";
  import { settingsStore } from "$lib/stores/settings";
  import { getFirstImageUrl } from "$lib/utils/imageUrl";
  import { toggleCartItem } from "$lib/utils/cart";
  import { goto } from "$app/navigation";

  interface Props {
    product: Product;
  }

  let { product }: Props = $props();

  let currency = $derived($settingsStore?.main.currency || "");
  let cart = $derived($cartStore);
  let inCart = $derived(cart.some((item) => item.id === product.id));

  function handleToggleCart() {
    toggleCartItem(product, cart);
  }
</script>

<li>
  <a
    href="/products/{product.slug}"
    onclick={(e) => { e.preventDefault(); goto(`/products/${product.slug}`); }}
    class="block overflow-hidden group rounded-lg cursor-pointer"
  >
    <img
      src={getFirstImageUrl(product.images, "md")}
      alt={product.name}
      class="h-[150px] w-full object-cover transition duration-500 group-hover:scale-105 sm:h-[250px]"
    />
  </a>
  <div class="relative bg-white mt-2">
    <div class="flex justify-between cursor-pointer">
      <span class="tracking-wider text-gray-900"
        >{costFormat(product.amount)} {currency}</span
      >

      <button
        onclick={handleToggleCart}
        class="group relative inline-flex items-center overflow-hidden rounded px-6 py-3 text-white focus:outline-hidden focus:ring cursor-pointer"
        class:bg-green-600={!inCart}
        class:bg-red-600={inCart}
      >
        {#if !inCart}
          <span class="absolute -start-full transition-all group-hover:start-4">
            <svg class="h-4 w-4 text-white">
              <use href="/assets/img/sprite.svg#plus" />
            </svg>
          </span>
          <span class="absolute end-4 transition-all group-hover:-end-full">
            <svg class="h-4 w-4 text-white">
              <use href="/assets/img/sprite.svg#cart" />
            </svg>
          </span>
        {:else}
          <span class="absolute -start-full transition-all group-hover:start-4">
            <svg class="h-4 w-4 text-white">
              <use href="/assets/img/sprite.svg#minus" />
            </svg>
          </span>
          <span class="absolute end-4 transition-all group-hover:-end-full">
            <svg class="h-4 w-4 text-white">
              <use href="/assets/img/sprite.svg#trash" />
            </svg>
          </span>
        {/if}
      </button>
    </div>
    <a
      href="/products/{product.slug}"
      onclick={(e) => { e.preventDefault(); goto(`/products/${product.slug}`); }}
      class="block overflow-hidden group mb-5 mt-2 cursor-pointer"
    >
      <h3 class="text-sm text-gray-700 group-hover:underline group-hover:underline-offset-4"
        >{product.name}</h3
      >
    </a>
  </div>
</li>
