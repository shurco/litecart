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
  import FormButton from "$lib/components/FormButton.svelte";
  import { browser } from "$app/environment";

  let product = $state<Product | null>(null);
  let load = $state(false);
  let currentSlide = $state(0);

  let currency = $derived($settingsStore?.main.currency || "");
  let cart = $derived($cartStore);
  let inCart = $derived(product ? cart.some((item) => item.id === product.id) : false);

  $effect(() => {
    if ($page.params.slug) {
      loadProduct($page.params.slug);
    }
  });

  async function loadProduct(slug: string) {
    const res = await apiGet<Product>(`/api/products/${slug}`);
    if (res.success && res.result) {
      product = res.result;
      load = true;

      if (browser && product.seo) {
        updateSEOTags(product.seo);
      }
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

{#if load && product}
  <section>
    <div class="max-w-screen-xl px-4 py-8 mx-auto sm:px-6 sm:py-12 lg:px-8">
      <div class="grid grid-cols-1 gap-4 lg:grid-cols-3 lg:gap-8">
        {#if !product.images || product.images.length === 0}
          <div class="relative h-[350px] sm:h-[450px]">
            <img
              src="/assets/img/noimage.png"
              alt=""
              class="rounded-lg absolute inset-0 h-full w-full object-cover"
            />
          </div>
        {:else if product.images.length === 1}
          <div class="relative h-[350px] sm:h-[450px]">
            <img
              src={getProductImageUrl(product.images[0], "md")}
              alt={product.name}
              class="rounded-lg absolute inset-0 h-full w-full object-cover"
            />
          </div>
        {:else}
          <div class="relative overflow-hidden h-[350px] sm:h-[450px] rounded-lg">
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
              class="absolute left-0 top-1/2 bg-white p-2 rounded-r-lg cursor-pointer"
            >
              &#8592;
            </button>
            <button
              onclick={() => nextSlide(product.images!.length)}
              class="absolute right-0 top-1/2 bg-white p-2 rounded-l-lg cursor-pointer"
            >
              &#8594;
            </button>
          </div>
        {/if}

        <div class="lg:col-span-2">
          <h1 class="text-xl font-bold text-gray-900 sm:text-3xl">{product.name}</h1>
          {#if product.attributes && product.attributes.length > 0}
            <div class="mt-4">
              {#each product.attributes as attr}
                <span
                  class="mr-2 whitespace-nowrap rounded-full bg-purple-100 px-2.5 py-0.5 text-sm text-purple-700"
                >
                  {attr}
                </span>
              {/each}
            </div>
          {/if}
          {#if product.brief}
            <div class="mt-4">{product.brief}</div>
          {/if}

          <div class="flex mt-4">
            <div class="flex-none pr-8">
              {#if !inCart}
                <FormButton
                  name="Add"
                  color="green"
                  ico="plus"
                  onclick={handleToggleCart}
                />
              {:else}
                <FormButton
                  name="Remove"
                  color="red"
                  ico="trash"
                  onclick={handleToggleCart}
                />
              {/if}
            </div>
            <div class="grow relative inline-flex items-center">
              <p class="text-2xl font-black"
                >{costFormat(product.amount)} {currency}</p
              >
            </div>
          </div>
        </div>
      </div>

      {#if product.description}
        <div
          class="mt-8 prod_desc border-t border-gray-100 pt-8"
        >{@html product.description}</div>
      {/if}
    </div>
  </section>
{/if}
