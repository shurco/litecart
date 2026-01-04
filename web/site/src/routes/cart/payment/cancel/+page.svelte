<script lang="ts">
  import { onMount } from 'svelte'
  import { cartStore } from '$lib/stores/cart'
  import { handleNavigation } from '$lib/utils/navigation'
  import { translate } from '$lib/i18n'

  // Reactive translation function
  let t = $derived($translate)

  onMount(() => {
    cartStore.clear()
  })
</script>

<div class="flex min-h-screen items-center justify-center bg-white px-4">
  <div class="w-full max-w-2xl">
    <div class="brutal-card bg-red-300 p-12 text-center">
      <h1 class="mb-8 text-4xl font-black tracking-tighter text-black uppercase sm:text-5xl">{t('payment.cancel.title')}</h1>
      <p class="mb-8 text-lg tracking-wide text-black">
        {t('payment.cancel.message')}
      </p>
      <a
        href="/cart"
        onclick={(e) => handleNavigation(e, '/cart')}
        class="inline-block cursor-pointer border-4 border-black bg-yellow-300 px-8 py-4 text-lg font-black tracking-wider text-black uppercase transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)]"
      >
        {t('payment.cancel.returnToCart')}
      </a>
    </div>
  </div>
</div>
