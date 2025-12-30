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

<section>
  <div class="max-w-screen-xl px-4 py-8 mx-auto sm:px-6 sm:py-12 lg:px-8">
    {#if load && products.length > 0}
      <ul class="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
        {#each products as product}
          <ProductCard {product} />
        {/each}
      </ul>
    {:else if load}
      <div>products not found</div>
    {/if}
  </div>
</section>
