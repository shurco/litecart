<script lang="ts">
  import { locale, availableLocales, type Locale } from '$lib/i18n'

  let currentLocale = $derived($locale)
  let locales = $derived($availableLocales)

  function switchLocale(newLocale: Locale) {
    locale.set(newLocale)
  }
</script>

<div class="flex items-center gap-1">
  {#each locales as loc}
    <button
      type="button"
      onclick={() => switchLocale(loc.code)}
      class="cursor-pointer border-2 border-black px-3 py-1 text-sm font-black uppercase transition-all duration-200 {currentLocale === loc.code
        ? 'bg-white text-black'
        : 'bg-transparent text-white hover:bg-white hover:text-black'}"
      aria-label="Switch to {loc.name}"
    >
      {loc.code.toUpperCase()}
    </button>
  {/each}
</div>
