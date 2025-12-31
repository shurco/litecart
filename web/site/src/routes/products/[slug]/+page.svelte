<script lang="ts">
  import { page } from "$app/stores";
  import { apiGet } from "$lib/utils/api";
  import type { Product } from "$lib/types/models";
  import { cartStore } from "$lib/stores/cart";
  import { costFormat } from "$lib/utils/costFormat";
  import { settingsStore } from "$lib/stores/settings";
  import { getProductImageUrl } from "$lib/utils/imageUrl";
  import { toggleCartItem } from "$lib/utils/cart";
  import { updateSEOTags } from "$lib/utils/seo";
  import { isBrowser } from "$lib/utils/browser";
  import NotFoundPage from "$lib/components/NotFoundPage.svelte";

  let product = $state<Product | null>(null);
  let load = $state(false);
  let notFound = $state(false);
  let loading = $state(true);
  let currentSlide = $state(0);

  let currency = $derived($settingsStore?.main.currency || "");
  let cart = $derived($cartStore);
  let inCart = $derived(product ? cart.some((item) => item.id === product.id) : false);

  $effect(() => {
    const slug = $page.params.slug;
    if (slug) {
      // Reset state when slug changes
      product = null;
      load = false;
      notFound = false;
      currentSlide = 0;
      loadProduct(slug);
    }
  });

  async function loadProduct(slug: string) {
    const res = await apiGet<Product>(`/api/products/${slug}`);
    loading = false;
    
    if (res.success && res.result) {
      product = res.result;
      load = true;

      if (isBrowser() && product.seo) {
        updateSEOTags(product.seo);
      }
    } else {
      // Product not found
      notFound = true;
    }
  }

  function handleToggleCart() {
    if (!product) return;
    toggleCartItem(product, cart);
  }

  function nextSlide(length: number) {
    currentSlide = (currentSlide + 1) % length;
  }

  function prevSlide(length: number) {
    currentSlide = (currentSlide + length - 1) % length;
  }
</script>

{#if loading}
  <div class="min-h-screen bg-white flex items-center justify-center">
    <div class="border-4 border-black bg-yellow-300 px-8 py-6 inline-block">
      <p class="text-xl font-black uppercase tracking-wider text-black">
        LOADING...
      </p>
    </div>
  </div>
{:else if notFound}
  <NotFoundPage />
{:else if load && product}
  <section class="min-h-screen bg-white py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-screen-xl mx-auto">
      <div class="grid grid-cols-1 gap-8 lg:grid-cols-2 lg:gap-12">
        <!-- Product Images -->
        <div>
          {#if !product.images || product.images.length === 0}
            <div class="relative h-[400px] sm:h-[500px] bg-white border-4 border-black">
              <img
                src="/assets/img/noimage.png"
                alt=""
                class="absolute inset-0 h-full w-full object-cover"
              />
            </div>
          {:else if product.images.length === 1}
            <div class="relative h-[400px] sm:h-[500px] bg-white border-4 border-black overflow-hidden">
              <img
                src={getProductImageUrl(product.images[0], "md")}
                alt={product.name}
                class="absolute inset-0 h-full w-full object-cover"
              />
            </div>
          {:else}
            <div class="relative overflow-hidden h-[400px] sm:h-[500px] border-4 border-black bg-white">
              <div
                class="flex w-full h-full transition-transform duration-500 ease-in-out"
                style="transform: translateX(-{currentSlide * 100}%)"
              >
                {#each product.images as image}
                  <div class="flex-shrink-0 w-full h-full">
                    <img
                      src={getProductImageUrl(image, "md")}
                      alt={product.name}
                      class="block w-full h-full object-cover"
                    />
                  </div>
                {/each}
              </div>
              <button
                onclick={() => prevSlide(product.images!.length)}
                class="absolute left-4 top-1/2 border-4 border-black bg-yellow-300 text-black p-3 font-black text-xl hover:shadow-[6px_6px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
                aria-label="Previous image"
              >
                ←
              </button>
              <button
                onclick={() => nextSlide(product.images!.length)}
                class="absolute right-4 top-1/2 border-4 border-black bg-yellow-300 text-black p-3 font-black text-xl hover:shadow-[6px_6px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
                aria-label="Next image"
              >
                →
              </button>
            </div>
          {/if}
        </div>

        <!-- Product Info -->
        <div class="space-y-6">
          <div class="brutal-card p-8">
            <h1 class="text-4xl sm:text-5xl font-black uppercase tracking-tighter text-black mb-4">
              {product.name}
            </h1>
            
            {#if product.attributes && product.attributes.length > 0}
              <div class="flex flex-wrap gap-2 mb-6">
                {#each product.attributes as attr}
                  <span class="border-4 border-black bg-blue-300 px-4 py-2 font-black text-sm uppercase tracking-wider text-black">
                    {attr}
                  </span>
                {/each}
              </div>
            {/if}

            <div class="flex items-baseline gap-3 mb-6">
              <span class="text-5xl font-black text-black tracking-tight">
                {costFormat(product.amount)}
              </span>
              <span class="text-2xl font-bold text-gray-700 uppercase">{currency}</span>
            </div>

            {#if product.brief}
              <div class="mb-6">
                <p class="text-lg font-bold text-black leading-relaxed">
                  {product.brief}
                </p>
              </div>
            {/if}

            <button
              onclick={handleToggleCart}
              class="w-full border-4 border-black px-8 py-4 font-black text-lg uppercase tracking-wider transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)] cursor-pointer {inCart ? 'bg-red-500 text-white' : 'bg-green-500 text-white'}"
            >
              {#if !inCart}
                <span class="flex items-center justify-center gap-3">
                  <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <use href="/assets/img/sprite.svg#plus" />
                  </svg>
                  <span>ADD TO CART</span>
                </span>
              {:else}
                <span class="flex items-center justify-center gap-3">
                  <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <use href="/assets/img/sprite.svg#minus" />
                  </svg>
                  <span>REMOVE FROM CART</span>
                </span>
              {/if}
            </button>
          </div>
        </div>
      </div>

      {#if product.description}
        <div class="mt-12">
          <h2 class="text-3xl font-black uppercase tracking-tighter text-black mb-6">
            DESCRIPTION
          </h2>
          <div class="prod_desc text-lg font-bold text-black leading-relaxed">
            {@html product.description}
          </div>
        </div>
      {/if}
    </div>
  </section>
{/if}
