<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'
  import { validators, validateFields } from '$lib/utils/validation'
  import { translate, locale, availableLocales, type Locale } from '$lib/i18n'

  // Reactive translation function
  let t = $derived($translate)
  let currentLocale = $derived($locale)
  let locales = $derived($availableLocales)

  interface MainSettings {
    site_name: string
    domain: string
    email: string
  }

  let formData = $state<MainSettings>({
    site_name: '',
    domain: '',
    email: ''
  })
  let formErrors = $state<Record<string, string>>({})
  let loading = $state(true)

  onMount(async () => {
    const loaded = await loadSettings<MainSettings>('main', formData)
    if (loaded) {
      formData = loaded
    }
    loading = false
  })

  async function handleSubmit(event: SubmitEvent) {
    event.preventDefault()
    formErrors = validateFields(formData, [
      { field: 'site_name', ...validators.minLength(6, t('settings.siteNameMinLength')) },
      { field: 'domain', ...validators.required(t('settings.domainRequired')) },
      { field: 'email', ...validators.email(t('settings.validEmailRequired')) }
    ])

    if (Object.keys(formErrors).length > 0) {
      return
    }

    await saveSettings('main', formData)
  }

  function switchLocale(newLocale: Locale) {
    locale.set(newLocale)
  }
</script>

<Main>
  <h1 class="mb-5">{t('settings.mainSettings')}</h1>

  {#if loading}
    <div class="py-8 text-center">{t('common.loading')}</div>
  {:else}
    <form onsubmit={handleSubmit} class="max-w-2xl space-y-4">
      <FormInput
        id="site_name"
        title={t('settings.siteName')}
        bind:value={formData.site_name}
        error={formErrors.site_name}
        ico="home"
      />
      <FormInput id="domain" title={t('settings.domain')} bind:value={formData.domain} error={formErrors.domain} ico="link" />
      <FormInput
        id="email"
        type="email"
        title={t('settings.email')}
        bind:value={formData.email}
        error={formErrors.email}
        ico="at-symbol"
      />
      <div class="pt-4">
        <FormButton type="submit" name={t('common.save')} color="green" />
      </div>
    </form>

    <hr class="mt-5" />

    <div class="mt-5">
      <h2 class="mb-5">{t('settings.language')}</h2>
      <div class="flex">
        {#each locales as loc, index}
          <div
            class="cursor-pointer rounded p-2 {currentLocale === loc.code ? 'bg-green-200' : 'bg-gray-200'} {index > 0 ? 'ml-5' : ''}"
            onclick={() => switchLocale(loc.code)}
            role="button"
            tabindex="0"
            onkeydown={(e) => {
              if (e.key === 'Enter' || e.key === ' ') {
                e.preventDefault()
                switchLocale(loc.code)
              }
            }}
          >
            {loc.name}
          </div>
        {/each}
      </div>
    </div>
  {/if}
</Main>
