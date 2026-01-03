<script lang="ts">
  import { onMount } from 'svelte'
  import { apiGet } from '$lib/utils/api'
  import type { Product } from '$lib/types/models'
  import ProductCard from '$lib/components/ProductCard.svelte'
  import Pagination from '$lib/components/Pagination.svelte'

  interface ProductsResponse {
    products: Product[]
    currency: string
    total: number
  }

  let products = $state<Product[]>([])
  let load = $state(false)
  let currentPage = $state(1)
  let limit = $state(20)
  let total = $state(0)

  async function loadProducts(page = currentPage) {
    load = false
    currentPage = page
    const res = await apiGet<ProductsResponse>(`/api/products?page=${page}&limit=${limit}`)
    if (res.success && res.result) {
      products = res.result.products || []
      total = res.result.total || 0
      load = true
    }
  }

  function handlePageChange(page: number) {
    loadProducts(page)
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  onMount(async () => {
    await loadProducts()
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
        {#each products as product, i (product.id)}
          <ProductCard {product} index={i} />
        {/each}
      </ul>

      {#if total > 0}
        <Pagination {currentPage} totalPages={Math.ceil(total / limit)} onPageChange={handlePageChange} />
      {/if}
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
