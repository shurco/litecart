<script lang="ts">
  import Header from '$lib/components/Header.svelte'
  import Footer from '$lib/components/Footer.svelte'
  import Overlay from '$lib/components/Overlay.svelte'
  import { settingsStore } from '$lib/stores/settings'
  import { apiGet } from '$lib/utils/api'
  import { updateSEOTags } from '$lib/utils/seo'
  import { isBrowser } from '$lib/utils/browser'
  import { page } from '$app/state'
  import { onMount } from 'svelte'

  interface Props {
    children: import('svelte').Snippet
  }

  let { children }: Props = $props()
  let showOverlay = $state(false)
  let error = $state<string | undefined>(undefined)

  // Check if this is an error page - from page status or body class
  let isErrorPageState = $state(false)

  // Check synchronously before render
  $effect.pre(() => {
    if (isBrowser()) {
      isErrorPageState =
        page.error !== null || (page.status && page.status >= 400) || document.body.classList.contains('error-page')
    }
  })

  // Also watch for changes
  $effect(() => {
    if (!isBrowser()) return

    const checkErrorPage = () => {
      isErrorPageState =
        page.error !== null || (page.status && page.status >= 400) || document.body.classList.contains('error-page')
    }

    checkErrorPage()

    // Watch for class changes on body
    const observer = new MutationObserver(checkErrorPage)
    observer.observe(document.body, { attributes: true, attributeFilter: ['class'] })

    return () => observer.disconnect()
  })

  const isErrorPage = $derived(isErrorPageState)

  onMount(async () => {
    if (!isBrowser()) return

    let cached = settingsStore.loadFromCache()
    if (!cached) {
      showOverlay = true
      const res = await apiGet('/api/settings')
      if (res.success && res.result) {
        settingsStore.set(res.result)
        settingsStore.saveToCache(res.result)

        // Update meta tags
        if (res.result.main?.site_name) {
          updateSEOTags({ title: res.result.main.site_name })
        }
      } else {
        error = res.message || 'Failed to load settings'
      }
      showOverlay = false
    } else {
      settingsStore.set(cached)
    }
  })

  function closeOverlay() {
    showOverlay = false
    error = undefined
  }
</script>

<div class="min-h-screen bg-white">
  {#if !isErrorPage}
    <header>
      <Header />
    </header>
  {/if}
  <main class="relative">
    {@render children()}
  </main>
  {#if !isErrorPage}
    <footer>
      <Footer />
    </footer>
  {/if}
  <Overlay show={showOverlay} error={error} onClose={closeOverlay} />
</div>
