<script lang="ts">
  import { onMount } from 'svelte'
  import { translate, locale } from '$lib/i18n'
  import { hasCookieConsent, setCookieConsent } from '$lib/utils/cookieConsent'
  import { isBrowser } from '$lib/utils/browser'

  let showBanner = $state(false)
  let t = $derived($translate)

  onMount(() => {
    if (!isBrowser()) return

    // Check if user has already given consent
    if (!hasCookieConsent()) {
      // Small delay for better UX
      setTimeout(() => {
        showBanner = true
      }, 500)
    }
  })

  function acceptCookies() {
    setCookieConsent('accepted')
    showBanner = false
  }

  function rejectCookies() {
    setCookieConsent('rejected')
    showBanner = false
  }
</script>

{#if showBanner}
  <div
    class="fixed bottom-0 left-0 right-0 z-[100] border-t-4 border-black bg-yellow-300 p-4 shadow-[0_-8px_0px_0px_rgba(0,0,0,1)]"
    role="dialog"
    aria-labelledby="cookie-consent-title"
    aria-describedby="cookie-consent-description"
  >
    <div class="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8">
      <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div class="flex-1">
          <h2 id="cookie-consent-title" class="mb-2 text-lg font-black uppercase tracking-wider text-black">
            {t('cookies.title')}
          </h2>
          <p id="cookie-consent-description" class="text-sm font-bold text-black">
            {t('cookies.description')}
            {#if t('cookies.privacyLink') !== 'cookies.privacyLink'}
              <a
                href="/privacy"
                class="underline hover:no-underline"
                onclick={(e) => {
                  e.preventDefault()
                  window.location.href = '/privacy'
                }}
              >
                {t('cookies.privacyLink')}
              </a>
            {/if}
          </p>
        </div>
        <div class="flex flex-col gap-2 sm:flex-row sm:gap-3">
          <button
            onclick={rejectCookies}
            class="border-4 border-black bg-white px-6 py-3 text-sm font-black uppercase tracking-wider text-black transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[8px_8px_0px_0px_rgba(0,0,0,1)]"
            aria-label={t('cookies.reject')}
          >
            {t('cookies.reject')}
          </button>
          <button
            onclick={acceptCookies}
            class="border-4 border-black bg-green-500 px-6 py-3 text-sm font-black uppercase tracking-wider text-white transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[8px_8px_0px_0px_rgba(0,0,0,1)]"
            aria-label={t('cookies.accept')}
          >
            {t('cookies.accept')}
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}
