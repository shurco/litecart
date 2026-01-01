<script lang="ts">
  import { onMount } from 'svelte'
  import { apiGet } from '$lib/utils/api'
  import type { Product } from '$lib/types/models'
  import ProductCard from '$lib/components/ProductCard.svelte'

  let products = $state<Product[]>([])
  let load = $state(false)

  onMount(async () => {
    const res = await apiGet<{ products: Product[] }>('/api/products')
    if (res.success && res.result) {
      products = res.result.products || []
      load = true
    }
  })
</script>

<section class="min-h-screen bg-white px-4 py-12 sm:px-6 lg:px-8">
  <!-- Products Section -->
  <div class="mx-auto max-w-screen-xl">
    {#if load && products.length > 0}
      <div class="mb-12 text-center">
        <h2 class="mb-4 text-4xl font-black tracking-tighter text-black uppercase sm:text-5xl">PRODUCTS</h2>
        <div class="mx-auto h-1 w-32 bg-black"></div>
      </div>

      <ul class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
        {#each products as product, i}
          <ProductCard {product} index={i} />
        {/each}
      </ul>
    {:else if load}
      <div class="py-20 text-center">
        <div class="inline-block border-4 border-black bg-white px-8 py-6">
          <p class="text-2xl font-black tracking-wider text-black uppercase">NO PRODUCTS FOUND</p>
        </div>
      </div>
    {:else}
      <div class="py-20 text-center">
        <div class="inline-block border-4 border-black bg-yellow-300 px-8 py-6">
          <p class="text-xl font-black tracking-wider text-black uppercase">LOADING...</p>
        </div>
      </div>
    {/if}
  </div>
</section>
