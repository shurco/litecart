<script lang="ts">
  import { onMount } from "svelte";
  import { apiGet } from "$lib/utils/api";
  import type { Product } from "$lib/types/models";
  import ProductCard from "$lib/components/ProductCard.svelte";

  let products = $state<Product[]>([]);
  let load = $state(false);

  onMount(async () => {
    const res = await apiGet<{ products: Product[] }>("/api/products");
    if (res.success && res.result) {
      products = res.result.products || [];
      load = true;
    }
  });
</script>

<section class="min-h-screen bg-white py-12 px-4 sm:px-6 lg:px-8">
  <!-- Products Section -->
  <div class="max-w-screen-xl mx-auto">
    {#if load && products.length > 0}
      <div class="mb-12 text-center">
        <h2 class="text-4xl sm:text-5xl font-black uppercase tracking-tighter mb-4 text-black">
          PRODUCTS
        </h2>
        <div class="w-32 h-1 bg-black mx-auto"></div>
      </div>
      
      <ul class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
        {#each products as product, i}
          <ProductCard {product} index={i} />
        {/each}
      </ul>
    {:else if load}
      <div class="text-center py-20">
        <div class="border-4 border-black bg-white px-8 py-6 inline-block">
          <p class="text-2xl font-black uppercase tracking-wider text-black">
            NO PRODUCTS FOUND
          </p>
        </div>
      </div>
    {:else}
      <div class="text-center py-20">
        <div class="border-4 border-black bg-yellow-300 px-8 py-6 inline-block">
          <p class="text-xl font-black uppercase tracking-wider text-black">
            LOADING...
          </p>
        </div>
      </div>
    {/if}
  </div>
</section>
