<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'
  import { translate } from '$lib/i18n'

  // Reactive translation function
  let t = $derived($translate)

  interface WebhookSettings {
    url: string
  }

  let formData = $state<WebhookSettings>({
    url: ''
  })
  let formErrors = $state<Record<string, string>>({})
  let loading = $state(true)

  onMount(async () => {
    formData = await loadSettings<WebhookSettings>('webhook', formData)
    loading = false
  })

  async function handleSubmit() {
    formErrors = {}

    if (formData.url && !/^https?:\/\/.+/.test(formData.url)) {
      formErrors.url = t('settings.validUrlExample')
      return
    }

    await saveSettings('webhook', formData)
  }
</script>

<Main>
  <h1 class="mb-5">{t('settings.webhookSettings')}</h1>

  {#if loading}
    <div class="py-8 text-center">{t('common.loading')}</div>
  {:else}
    <form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="max-w-2xl space-y-4">
      <FormInput
        id="url"
        type="url"
        title={t('settings.webhookUrl')}
        bind:value={formData.url}
        error={formErrors.url}
        ico="webhook"
        placeholder="https://example.com/webhook"
      />
      <div class="pt-4">
        <FormButton type="submit" name={t('common.save')} color="green" />
      </div>
    </form>
  {/if}
</Main>
