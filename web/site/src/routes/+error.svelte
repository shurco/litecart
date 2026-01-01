<script lang="ts">
  import { page as pageStore } from '$app/stores'
  import { handleNavigation } from '$lib/utils/navigation'
  import { isBrowser } from '$lib/utils/browser'
  import { onDestroy } from 'svelte'

  const status = $derived($pageStore.status || 404)
  const isNotFound = $derived(status === 404)

  // Add error-page class to body
  $effect.pre(() => {
    if (isBrowser()) {
      document.body.classList.add('error-page')
    }
  })

  // Cleanup on component destroy
  onDestroy(() => {
    if (isBrowser()) {
      document.body.classList.remove('error-page')
    }
  })
</script>

<div class="flex min-h-screen items-center justify-center bg-white px-4 py-12">
  <div class="w-full max-w-3xl text-center">
    <div class="brutal-card mb-8 bg-red-300 p-12">
      <h1 class="mb-4 text-8xl font-black tracking-tighter text-black uppercase sm:text-9xl">
        {status}
      </h1>
      <p class="text-3xl font-black tracking-wider text-black uppercase">
        {isNotFound ? 'NOT FOUND' : 'ERROR'}
      </p>
    </div>

    <div class="brutal-card mb-8 p-8">
      <p class="mb-8 text-lg tracking-wide text-black">
        {isNotFound
          ? "The page you're looking for doesn't exist or has been moved."
          : 'Something went wrong. Please try again later.'}
      </p>

      <div class="flex justify-center">
        <a
          href="/"
          onclick={(e) => handleNavigation(e, '/')}
          class="inline-block cursor-pointer border-4 border-black bg-yellow-300 px-8 py-4 text-lg font-black tracking-wider text-black uppercase transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)]"
        >
          GO TO HOME
        </a>
      </div>
    </div>

    {#if isNotFound}
      <div class="brutal-card bg-white p-6">
        <p class="text-lg tracking-wide text-black">
          Try checking the URL for typos, or return to the homepage to browse our products.
        </p>
      </div>
    {/if}
  </div>
</div>
